package handler

import (
	"go-scaffold/internal/service"
	"go-scaffold/pkg/conf"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login ..
// @Tags 授权
// @Summary 登录
// @Description
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} handler.Response "返回结果：{code:2000,message:"Success",data:nil}"
// @Router /api/login [post]
func Login(ctx *gin.Context) {
	arg := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	if err := ctx.ShouldBindJSON(&arg); err != nil {
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
		Name:     conf.GetString("app.http.name"),
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
	issuer := conf.GetString("app.http.name")
	sessionID, _ := ctx.Cookie(issuer)
	ctx.SetCookie(issuer, "", -1, "", "", false, true)
	service.Authentication.Logout(sessionID)
	ResponseSuccess(ctx, nil)
}
