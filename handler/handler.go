package handler

import (
	"github.com/DataWorkbench/common/constants"
	"github.com/DataWorkbench/common/utils/idgenerator"
	"gorm.io/gorm"
)

// global options in this package.
var (
	dbConn           *gorm.DB
	apiIdGenerator   *idgenerator.IDGenerator
	//reqParamIdGenerator  *idgenerator.IDGenerator
	//respParamIdGenerator *idgenerator.IDGenerator
)

type Option func()

func WithDBConn(conn *gorm.DB) Option {
	return func() {
		dbConn = conn
	}
}

func Init(opts ...Option) {
	for _, opt := range opts {
		opt()
	}
	apiIdGenerator = idgenerator.New(constants.IdPrefixDataService)
}
