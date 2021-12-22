package server

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DataWorkbench/common/gormwrap"
	"github.com/DataWorkbench/common/gtrace"
	"github.com/DataWorkbench/common/utils/buildinfo"
	"github.com/DataWorkbench/glog"
	"gorm.io/gorm"

	"github.com/DataWorkbench/common/grpcwrap"
	"github.com/DataWorkbench/common/metrics"
	datasvc "github.com/DataWorkbench/gproto/pkg/dataservicepb"

	"github.com/DataWorkbench/dataservice/config"
	"github.com/DataWorkbench/dataservice/handler"
)

// Start for start the http server
func Start() (err error) {
	fmt.Printf("%s pid=%d program_build_info: %s\n",
		time.Now().Format(time.RFC3339Nano), os.Getpid(), buildinfo.JSONString)

	var cfg *config.Config

	cfg, err = config.Load()
	if err != nil {
		return
	}

	// init parent logger
	lp := glog.NewDefault().WithLevel(glog.Level(cfg.LogLevel))
	ctx := glog.WithContext(context.Background(), lp)

	var (
		db           *gorm.DB
		rpcServer    *grpcwrap.Server
		metricServer *metrics.Server
		tracer       gtrace.Tracer
		tracerCloser io.Closer
	)

	defer func() {
		rpcServer.GracefulStop()
		_ = metricServer.Shutdown(ctx)
		if tracerCloser != nil {
			_ = tracerCloser.Close()
		}
		_ = lp.Close()
	}()

	tracer, tracerCloser, err = gtrace.New(cfg.Tracer)
	if err != nil {
		return
	}
	ctx = gtrace.ContextWithTracer(ctx, tracer)

	// init gorm.DB
	db, err = gormwrap.NewMySQLConn(ctx, cfg.MySQL)
	if err != nil {
		return
	}

	// init prometheus server
	metricServer, err = metrics.NewServer(ctx, cfg.MetricsServer)
	if err != nil {
		return err
	}

	// init grpc.Server
	grpcwrap.SetLogger(lp, cfg.GRPCLog)
	rpcServer, err = grpcwrap.NewServer(ctx, cfg.GRPCServer)
	if err != nil {
		return
	}

	handler.Init(handler.WithDBConn(db))
	rpcServer.Register(func(s *grpcwrap.GServer) {
		datasvc.RegisterDataServiceServer(s, &DataServiceServer{})
	})

	// handle signal
	sigGroup := []os.Signal{syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM}
	sigChan := make(chan os.Signal, len(sigGroup))
	signal.Notify(sigChan, sigGroup...)

	blockChan := make(chan struct{})

	// run grpc server
	go func() {
		_ = rpcServer.ListenAndServe()
		blockChan <- struct{}{}
	}()

	go func() {
		// Ignore metrics server error.
		_ = metricServer.ListenAndServe()
	}()

	go func() {
		sig := <-sigChan
		lp.Info().String("receive system signal", sig.String()).Fire()
		blockChan <- struct{}{}
	}()

	<-blockChan
	return
}
