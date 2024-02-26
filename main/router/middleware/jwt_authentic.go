package middleware

import (
	"GoProject/main/enum/opstatus"
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
		ctx.JSON(400, view.StatusWith(opstatus.UnLogin))
		return
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		ctx.Abort()
		ctx.JSON(401, view.StatusWith(opstatus.TokenExpired))
		return
	}

	ctx.Set("id", claims.ID)
	ctx.Set("username", claims.Username)
	ctx.Set("role", claims.Authority)
	ctx.Next()
}
