package middleware

import (
	"go-scaffold/internal/apis/rest/handler"
	"go-scaffold/internal/constant"
	"go-scaffold/internal/model"
	"go-scaffold/pkg/cache/xredis"
	"go-scaffold/pkg/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Authentication middleware
func Authentication(engine *gin.Engine) {
	engine.Use(authentication(
		"/swagger/*any",
		"/ping",
		"/api/login",
	))
}

// Authentication 身份验证中间件
func authentication(args ...string) gin.HandlerFunc {
	ignoreFn := genIgnoreFunc(args...)
	issuer := config.GetString("app.http.name")
	return func(ctx *gin.Context) {
		if ignoreFn(ctx.FullPath()) {
			return
		}
		sessionID, err := ctx.Cookie(issuer)
		if err != nil || len(sessionID) == 0 {
			handler.ResponseError(ctx, handler.ErrPermissionDenied)
			ctx.Abort()
			return
		}

		tmp := xredis.Store(constant.RedisStoreNameDB0).HGetAll(constant.RedisKeySession + sessionID)
		userID, ok := tmp["userID"]
		if !ok {
			handler.ResponseError(ctx, handler.ErrPermissionDenied)
			ctx.Abort()
			return
		}

		username, ok := tmp["username"]
		if !ok {
			handler.ResponseError(ctx, handler.ErrPermissionDenied)
			ctx.Abort()
			return
		}

		var user model.User
		id, _ := strconv.ParseUint(userID.String(), 10, 64)
		user.ID = uint(id)
		user.UserName = username.String()

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
