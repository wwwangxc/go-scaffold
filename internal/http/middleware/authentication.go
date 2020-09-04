package middleware

import (
	"go-scaffold/internal/constant"
	"go-scaffold/internal/http/handler"
	"go-scaffold/internal/model"
	"go-scaffold/pkg/xredis"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UseAuthentication use authentication middleware
func UseAuthentication(engine *gin.Engine) {
	engine.Use(authentication(
		"/swagger/*any",
		"/ping",
		"/api/login",
	))
}

// Authentication 身份验证中间件
func authentication(args ...string) gin.HandlerFunc {
	ignoreFn := genIgnoreFunc(args...)
	return func(ctx *gin.Context) {
		if ignoreFn(ctx.FullPath()) {
			return
		}
		sessionID, err := ctx.Cookie("gateway-sso")
		if err != nil || len(sessionID) == 0 {
			handler.ResponseWithCode(ctx, constant.HTTPResponseCodeNotLogin)
			ctx.Abort()
			return
		}

		tmp := xredis.Cli.HGetAll(constant.RedisKeySession + sessionID)
		userID, ok := tmp["userID"]
		if !ok {
			handler.ResponseWithCode(ctx, constant.HTTPResponseCodeInvalidAuth)
			ctx.Abort()
			return
		}

		username, ok := tmp["username"]
		if !ok {
			handler.ResponseWithCode(ctx, constant.HTTPResponseCodeInvalidAuth)
			ctx.Abort()
			return
		}

		var user model.Admin
		id, _ := strconv.Atoi(userID)
		user.ID = id
		user.Username = username

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}

func genIgnoreFunc(args ...string) func(string) bool {
	ignoreMap := make(map[string]bool)
	for _, v := range args {
		ignoreMap[v] = true
	}
	return func(arg string) bool {
		return ignoreMap[arg]
	}
}
