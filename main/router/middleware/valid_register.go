package middleware

import (
	"GoProject/main/enum/opstatus"
	"GoProject/main/util"
	"GoProject/main/view"
	"github.com/gin-gonic/gin"
	"log/slog"
	"strings"
)

func ValidRegister(ctx *gin.Context) {
	var reg view.RegisterReq
	if err := ctx.ShouldBind(&reg); err != nil {
		ctx.Abort()
		ctx.JSON(400, view.Fail())
		slog.Error(err.Error())
		return
	}

	reg.Username = strings.Trim(reg.Username, " ")
	if err := util.ValidateStruct(&reg); err != nil {
		ctx.Abort()
		ctx.JSON(400, view.StatusWith(opstatus.InvalidParams))
		slog.Error(err.Error())
		return
	}
	ctx.Set("register_params", reg)
}
