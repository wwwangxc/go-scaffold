package handler

import (
	"go-scaffold/internal/model"

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

func currentUser(ctx *gin.Context) (*model.Admin, error) {
	v, ok := ctx.Get("current-user")
	if !ok {
		return nil, ErrPermissionDenied
	}
	user, ok := v.(*model.Admin)
	if !ok {
		return nil, ErrPermissionDenied
	}
	return user, nil
}
