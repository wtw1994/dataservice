package handler

import (
	"context"
	"errors"
	"github.com/DataWorkbench/common/qerror"
	datasvc "github.com/DataWorkbench/gproto/pkg/dataservicepb"
	"gorm.io/gorm"
)

const (
	apiConfigTableName = "api_config"
	requestParamTableName = "api_request_parameters"
	responseParamTableName = "api_response_parameters"
)

func CreateDataServiceApi(ctx context.Context, req *datasvc.CreateDataSvcApiRequest) (id string, err error) {
	id, err = apiIdGenerator.Take()
	if err != nil {
		return
	}

	tx := dbConn.Begin().WithContext(ctx)
	if err = tx.Error; err != nil {
		return
	}
	defer func() {
		if err == nil {
			err = tx.Commit().Error
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	// check exists
	var x string
	err = tx.Table(apiConfigTableName).Select("api_id").
		Where("api_name = ? AND status != ?", req.ApiName, 3).
		Take(&x).Error
	if err == nil {
		err = qerror.ApiAlreadyExists.Format(req.ApiName)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		return
	}

	//now := time.Now().Unix()
	apiInfo := &datasvc.ApiConfig{
		ApiId:     id,
		ApiName:   req.ApiName,
		ApiPath:   req.ApiPath,
		ApiMode:   req.ApiMode,
		ApiDescription:  req.ApiDescription,
		SpaceId:       req.SpaceId,
		//Protocols: req.Protocols,
		RequestMethod: req.RequestMethod,
		ResponseType:  req.ResponseType,
		Timeout:       req.Timeout,
		VisibleRange:  req.VisibleRange,
		Datasource_Id: req.WizardDetails.WizardConnection.Datasource_Id,
		TableName:     req.WizardDetails.WizardConnection.TableName,
		//Created: now,
		//Updated: now,
	}

	if err = tx.Table(apiConfigTableName).Create(apiInfo).Error; err != nil {
		return
	}

    // create  ApiRequestParam record
	for _, r := range req.WizardDetails.RequestParams {
		rpId, err := apiIdGenerator.Take()
		if err != nil {
			return rpId,err
		}
		reqParamInfo := &datasvc.ApiRequestParams{
			ParamId:     rpId,
			ApiId:       id,
			ColumnName:   r.ColumnName,
			DefaultValue:  r.DefaultValue,
			ExampleValue:  r.ExampleValue,
			IsRequired:       r.IsRequired,
			DataType: r.DataType,
			ParamDescription:  r.ParamDescription,
			ParamName:       r.ParamName,
			ParamOperator:  r.ParamOperator,
			ParamPosition: r.ParamPosition,
			//Created: now,
			//Updated: now,
		}
		if err = tx.Table(requestParamTableName).Create(reqParamInfo).Error; err != nil {
			return rpId,err
		}
	}

	// create  ApiResponseParams record
	for _, r := range req.WizardDetails.ResponseParams {
		rpId, err := apiIdGenerator.Take()
		if err != nil {
			return rpId,err
		}
		reqParamInfo := &datasvc.ApiResponseParams{
			ParamId:     rpId,
			ApiId:       id,
			ColumnName:   r.ColumnName,
			DefaultValue:  r.DefaultValue,
			ExampleValue:  r.ExampleValue,
			DataType: r.DataType,
			ParamDescription:  r.ParamDescription,
			ParamName:       r.ParamName,
			//Created: now,
			//Updated: now,
		}
		if err = tx.Table(responseParamTableName).Create(reqParamInfo).Error; err != nil {
			return rpId,err
		}
	}
	return
}
