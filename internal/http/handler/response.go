package handler

import (
	"encoding/json"
	"fmt"
	"go-scaffold/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response HTTP Response
type HTTPResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseSuccess response http success
func ResponseSuccess(ctx *gin.Context, data interface{}) {
	resp := &HTTPResponse{
		Code: CodeSuccess,
		Data: data,
	}
	log.Info(fmt.Sprintf("%s %s %s", ctx.FullPath(), ctx.Request.Method, resp.String()))
	ctx.JSON(http.StatusOK, resp)
}

// ResponseError response http error
func ResponseError(ctx *gin.Context, code *HTTPResponseCode) {
	resp := &HTTPResponse{
		Code:    code.Code,
		Message: code.Message,
	}
	log.Fatal(fmt.Sprintf("%s %s %s", ctx.FullPath(), ctx.Request.Method, resp.String()))
	ctx.JSON(http.StatusOK, resp)
}

func (t *HTTPResponse) String() string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}

// HTTPResponseCode HTTP response code
type HTTPResponseCode struct {
	Code    int
	Message string
}

// New will modify the message of ResponseCode
func (t *HTTPResponseCode) New(message string) *HTTPResponseCode {
	return &HTTPResponseCode{
		Code:    t.Code,
		Message: message,
	}
}

// Newf will modify the message of ResponseCode with parameters
func (t *HTTPResponseCode) Newf(message string, args interface{}) *HTTPResponseCode {
	return &HTTPResponseCode{
		Code:    t.Code,
		Message: fmt.Sprintf(message, args),
	}
}

// Error return error message
func (t *HTTPResponseCode) Error() string {
	return t.Message
}

const (
	CodePermissionDenied = 1001
	CodeInvalidParams    = 1002

	CodeSuccess = 2000

	CodeInternal  = 5000
	CodeServeBusy = 5001
)

var (
	ErrPermissionDenied = &HTTPResponseCode{Code: CodePermissionDenied, Message: "permission denied"}
	ErrInvalidParams    = &HTTPResponseCode{Code: CodeInvalidParams, Message: "invalid parameters"}

	ErrInternal  = &HTTPResponseCode{Code: CodeInternal, Message: "internal error"}
	ErrServeBusy = &HTTPResponseCode{Code: CodeServeBusy, Message: "serve busy"}
)
