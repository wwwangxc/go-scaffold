package handler

import (
	"go-scaffold/internal/constant"
	"go-scaffold/internal/model"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidAuth = errors.New(constant.HTTPResponseCodeInvalidAuth.String())
)

// RoutePing ..
func RoutePing(engine *gin.Engine) {
	engine.GET("/ping", func(ctx *gin.Context) {
		ResponseSuccess(ctx, "pong!")
	})
}

// RouteAuthentication route authentication handler
func RouteAuthentication(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	{
		apiGroup.POST("/login", Login)
		apiGroup.GET("/logout", Logout)
	}
}

func currentUser(ctx *gin.Context) (*model.Admin, error) {
	v, ok := ctx.Get("current-user")
	if !ok {
		return nil, ErrInvalidAuth
	}
	user, ok := v.(*model.Admin)
	if !ok {
		return nil, ErrInvalidAuth
	}
	return user, nil
}
