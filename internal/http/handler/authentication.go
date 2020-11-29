package handler

import (
	"go-scaffold/internal/http/dto"
	"go-scaffold/internal/service"
	"go-scaffold/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login ..
// @Tags 授权
// @Summary 登录
// @Description 登陆成功后，签发一个15分钟效期的cookie，用于后端鉴权
// @Param arg body dto.Login true "用户登陆信息"
// @Success 200 {object} handler.Response "返回结果：{code:2000,message:"Success",data:nil}"
// @Router /api/login [post]
func Login(ctx *gin.Context) {
	arg := &dto.Login{}
	if err := ctx.ShouldBindJSON(arg); err != nil {
		ResponseError(ctx, err)
		return
	}
	sessionID, err := service.Authentication.Login(
		arg.Username, arg.Password)
	if err != nil {
		ResponseError(ctx, err)
		return
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     config.GetString("app.http.name"),
		Value:    sessionID,
		HttpOnly: true,
		MaxAge:   15 * 60,
	})
	ResponseSuccess(ctx, nil)
}

// Logout ..
// @Tags 授权
// @Summary 登出
// @Success 200 {object} handler.Response "返回结果：{code:2000,message:"Success",data:nil}"
// @Router /api/logout [get]
func Logout(ctx *gin.Context) {
	issuer := config.GetString("app.http.name")
	sessionID, _ := ctx.Cookie(issuer)
	ctx.SetCookie(issuer, "", -1, "", "", false, true)
	service.Authentication.Logout(sessionID)
	ResponseSuccess(ctx, nil)
}
