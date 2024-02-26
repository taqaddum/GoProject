package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidFileName(ctx *gin.Context) {
	fileName := ctx.Param("fileName")
	// TODO 应使用validation工具来验证文件名是否合法
	if fileName == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid file name"})
		return
	}
	ctx.Next()
}
