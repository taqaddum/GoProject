package handler

import (
	"GoProject/main/service"
	"GoProject/main/view"
)
import "github.com/gin-gonic/gin"

type UserHandler struct {
	srvApi service.UserSrvApi
}

type UserHandleFunc interface {
	LoginFunc(ctx *gin.Context)
}

func NewUserHandler(srv *service.UserService) *UserHandler {
	return &UserHandler{srvApi: srv}
}

func (userHandle UserHandler) LoginFunc(ctx *gin.Context) {
	name := ctx.GetString("username")
	passwd := ctx.GetString("password")
	token := userHandle.srvApi.LoginByName(name, passwd)

	if len(token) > 0 {
		ctx.JSON(200, view.Success(gin.H{"token": token}))
	} else {
		ctx.JSON(200, view.Failed())
	}
}
