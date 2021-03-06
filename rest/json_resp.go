package rest

import (
	"github.com/enrichroad/community/errors"
	"github.com/enrichroad/community/pagination"
	"github.com/enrichroad/community/reflect"
)

type Resp struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
}

func Json(code int, message string, data interface{}, success bool) *Resp {
	return &Resp{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   success,
	}
}

func JsonData(data interface{}) *Resp {
	return &Resp{
		ErrorCode: 0,
		Data:      data,
		Success:   true,
	}
}

func JsonPageData(results interface{}, page *pagination.Paging) *Resp {
	return JsonData(&PageResult{
		Results: results,
		Page:    page,
	})
}

func JsonCursorData(results interface{}, cursor string) *Resp {
	return JsonData(&CursorResult{
		Results: results,
		Cursor:  cursor,
	})
}

func JsonSuccess() *Resp {
	return &Resp{
		ErrorCode: 0,
		Data:      nil,
		Success:   true,
	}
}

func JsonError(err *errors.CodeError) *Resp {
	return &Resp{
		ErrorCode: err.Code,
		Message:   err.Message,
		Data:      err.Data,
		Success:   false,
	}
}

func JsonErrorMsg(message string) *Resp {
	return &Resp{
		ErrorCode: 0,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}

func JsonErrorCode(code int, message string) *Resp {
	return &Resp{
		ErrorCode: code,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}

func JsonErrorData(code int, message string, data interface{}) *Resp {
	return &Resp{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   false,
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

func (builder *RspBuilder) JsonResult() *Resp {
	return JsonData(builder.Data)
}
