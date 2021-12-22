package server

import (
	"context"
	"github.com/DataWorkbench/dataservice/handler"
	datasvc "github.com/DataWorkbench/gproto/pkg/dataservicepb"
)

type DataServiceServer struct {
	datasvc.UnimplementedDataServiceServer
}

func (s *DataServiceServer) CreateDataServiceApi(ctx context.Context, req *datasvc.CreateDataSvcApiRequest) (*datasvc.CreateDataSvcApiResponse, error) {
	id, err := handler.CreateDataServiceApi(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &datasvc.CreateDataSvcApiResponse{RequestId: id} // RequestId==>API_ID
	return reply, nil
}