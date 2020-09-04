package handler

import (
	"encoding/json"
	"go-scaffold/internal/constant"
	"go-scaffold/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response HTTP Response
type Response struct {
	Code    constant.HTTPResponseCode `json:"code"`
	Message string                    `json:"message"`
	Data    interface{}               `json:"data"`
}

// ResponseSuccess ..
func ResponseSuccess(ctx *gin.Context, data interface{}) {
	obj := &Response{
		Code:    constant.HTTPResponseCodeSuccess,
		Message: constant.HTTPResponseCodeSuccess.String(),
		Data:    data,
	}
	log.Info(obj.String())
	ctx.JSON(http.StatusOK, obj)
}

// ResponseError ..
func ResponseError(ctx *gin.Context, err error) {
	obj := &Response{
		Code:    constant.HTTPResponseCodeServeError,
		Message: err.Error(),
		Data:    nil,
	}
	log.Info(obj.String())
	ctx.JSON(http.StatusOK, obj)
}

// ResponseWithCode ..
func ResponseWithCode(ctx *gin.Context, code constant.HTTPResponseCode) {
	obj := &Response{
		Code:    code,
		Message: code.String(),
		Data:    nil,
	}
	log.Info(obj.String())
	ctx.JSON(http.StatusOK, obj)
}

func (t *Response) String() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}
