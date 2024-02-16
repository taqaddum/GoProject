package middleware

import (
	"GoProject/main/util"
	"GoProject/main/view"
	"github.com/gin-gonic/gin"
	"strings"
)

func ValidLogin(ctx *gin.Context) {
	name := strings.Trim(ctx.PostForm("username"), " ")
	passwd := ctx.PostForm("password")

	if util.ValidateName(name) != nil || util.ValidatePasswd(passwd) != nil {
		ctx.Abort()
		ctx.JSON(400, view.Fail())
		return
	}

	ctx.Set("username", name)
	ctx.Set("password", passwd)
}
