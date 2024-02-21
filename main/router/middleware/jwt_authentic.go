package middleware

import (
	"GoProject/main/enum/httpstatus"
	"GoProject/main/util"
	"GoProject/main/view"
	"github.com/gin-gonic/gin"
	"time"
)

func AuthJWT(ctx *gin.Context) {
	token := ctx.GetHeader("authorization")
	claims, err := util.ParseToken(token)
	if err != nil {
		ctx.Abort()
		ctx.JSON(400, view.StatusWith(httpstatus.UnLogin))
		return
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		ctx.Abort()
		ctx.JSON(401, view.StatusWith(httpstatus.TokenExpired))
		return
	}
}
