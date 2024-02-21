package middleware

import (
	"GoProject/main/enum/httpstatus"
	"GoProject/main/view"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func ValidFormFile(ctx *gin.Context) {
	if file, err := ctx.FormFile("file"); err != nil {
		ctx.Abort()
		ctx.JSON(400, view.StatusWith(httpstatus.InvalidParams))
		slog.Error("文件上传参数错误", err.Error())
	} else {
		ctx.Set("file", file)
	}
}
