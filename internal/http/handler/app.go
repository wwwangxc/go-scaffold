package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWord ..
// @Tags App
// @Summary Hello Word
// @Success 200 {object} handler.Response "返回结果：{code:2000,message:"Success",data:nil}"
// @Router /api/app [get]
func HelloWord(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "hello word",
	})
}
