package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func ValidLoginParams(ctx *gin.Context) {
	name := strings.Trim(ctx.PostForm("username"), " ")
	passwd := ctx.PostForm("password")

	ctx.Set("username", name)
	ctx.Set("password", passwd)
}
