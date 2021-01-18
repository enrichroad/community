package serializing

import (
	"github.com/enrichroad/community/errors"
	"github.com/enrichroad/community/pagination"
	"github.com/enrichroad/community/reflect"
)

// JsonResp json for http response
type JsonResp struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func Json(code int, message string, data interface{}) *JsonResp {
	return &JsonResp{
		ErrorCode: code,
		Message:   message,
		Data:      data,
	}
}

func JsonData(data interface{}) *JsonResp {
	return &JsonResp{
		ErrorCode: 0,
		Data:      data,
	}
}

func JsonPageData(results interface{}, page *pagination.Paging) *JsonResp {
	return JsonData(&PageResult{
		Results: results,
		Page:    page,
	})
}

func JsonCursorData(results interface{}, cursor string) *JsonResp {
	return JsonData(&CursorResult{
		Results: results,
		Cursor:  cursor,
	})
}

func JsonSuccess() *JsonResp {
	return &JsonResp{
		ErrorCode: 0,
		Data:      nil,
	}
}

func JsonError(err *errors.CodeError) *JsonResp {
	return &JsonResp{
		ErrorCode: err.Code,
		Message:   err.Message,
		Data:      err.Data,
	}
}

func JsonErrorMsg(message string) *JsonResp {
	return &JsonResp{
		ErrorCode: 0,
		Message:   message,
		Data:      nil,
	}
}

func JsonErrorCode(code int, message string) *JsonResp {
	return &JsonResp{
		ErrorCode: code,
		Message:   message,
		Data:      nil,
	}
}

func JsonErrorData(code int, message string, data interface{}) *JsonResp {
	return &JsonResp{
		ErrorCode: code,
		Message:   message,
		Data:      data,
	}
}

type RspBuilder struct {
	Data map[string]interface{}
}

func NewEmptyRspBuilder() *RspBuilder {
	return &RspBuilder{Data: make(map[string]interface{})}
}

func NewRspBuilder(obj interface{}) *RspBuilder {
	return NewRspBuilderExcludes(obj)
}

func NewRspBuilderExcludes(obj interface{}, excludes ...string) *RspBuilder {
	return &RspBuilder{Data: reflect.StructToMap(obj, excludes...)}
}

func (builder *RspBuilder) Put(key string, value interface{}) *RspBuilder {
	builder.Data[key] = value
	return builder
}

func (builder *RspBuilder) Build() map[string]interface{} {
	return builder.Data
}

func (builder *RspBuilder) JsonResult() *JsonResp {
	return JsonData(builder.Data)
}
