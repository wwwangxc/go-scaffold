package handler

import (
	"go-scaffold/internal/service"
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
	sessionID, err := service.Authentication.Login(
		ctx.PostForm("username"), ctx.PostForm("password"))
	if err != nil {
		ResponseError(ctx, err)
		return
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "gateway-sso",
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
	sessionID, _ := ctx.Cookie("gateway-sso")
	ctx.SetCookie("gateway-sso", "", -1, "", "", false, true)
	service.Authentication.Logout(sessionID)
	ResponseSuccess(ctx, nil)
}
