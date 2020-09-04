package handler

import (
	"go-scaffold/internal/constant"
	"go-scaffold/internal/model"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
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

// RouteApp route app handler
func RouteApp(engine *gin.Engine) {
	group := engine.Group("/api/app")
	{
		group.GET("", HelloWord)
	}
}

func currentUser(ctx *gin.Context) (*model.Admin, error) {
	v, ok := ctx.Get("current-user")
	if !ok {
		return nil, errors.New(constant.HTTPResponseCodeInvalidAuth.String())
	}
	user, ok := v.(*model.Admin)
	if !ok {
		return nil, errors.New(constant.HTTPResponseCodeInvalidAuth.String())
	}
	return user, nil
}
