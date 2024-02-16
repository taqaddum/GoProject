package handler

import (
	"GoProject/main/enum/opstatus"
	"GoProject/main/model"
	"GoProject/main/service"
	"GoProject/main/view"
)
import "github.com/gin-gonic/gin"

type UserHandler struct {
	srvApi service.UserSrvApi
}

type UserHandleFunc interface {
	LoginFunc(ctx *gin.Context)
	RegisterFunc(ctx *gin.Context)
	GetDetailFunc(ctx *gin.Context)
}

func NewUserHandler(srv *service.UserService) *UserHandler {
	return &UserHandler{srvApi: srv}
}

// LoginFunc
//
//	@Description: 用户登录api，只验证密码是否正确，其他校验通过路由中间件实现
//	@receiver userHandle
//	@param ctx
func (userHandle UserHandler) LoginFunc(ctx *gin.Context) {
	name := ctx.GetString("username")
	passwd := ctx.GetString("password")
	token := userHandle.srvApi.LoginByName(name, passwd)

	if len(token) > 0 {
		ctx.JSON(200, view.Success(gin.H{"authorization": token}))
	} else {
		ctx.JSON(400, view.StatusWith(opstatus.WrongPassword))
	}
}

func (userHandle UserHandler) RegisterFunc(ctx *gin.Context) {
	tmp, exist := ctx.Get("register_info")
	if !exist {
		ctx.JSON(400, view.Fail())
	}

	reg, _ := tmp.(view.RegisterReq)
	err := userHandle.srvApi.Register(
		&model.User{
			Username: reg.Username,
			Passwd:   reg.Password,
			Email:    reg.Email,
		})

	if err != nil {
		ctx.JSON(400, view.Fail())
	} else {
		ctx.JSON(200, view.Success())
	}
}

func (userHandle UserHandler) GetDetailFunc(ctx *gin.Context) {
	name := ctx.Param("username")
	if name != ctx.GetString("username") {
		ctx.JSON(400, view.StatusWith(opstatus.Unauthorized))
		return
	}
	result := userHandle.srvApi.GetDetail(name)
	ctx.JSON(200, view.Success(result))
}
