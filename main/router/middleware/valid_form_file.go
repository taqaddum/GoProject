package middleware

import (
	"GoProject/main/enum/opstatus"
	"GoProject/main/view"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func ValidFormFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.Abort()
		slog.Error("文件上传参数错误", err.Error())
		ctx.JSON(400, view.StatusWith(opstatus.InvalidParams))
		return
	}

	ctx.Set("file", file)
	ctx.Next()
}
