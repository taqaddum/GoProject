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

func NewUserHandler(srv *service.UserService) *UserHandler {
	return &UserHandler{srvApi: srv}
}

type UserHandles interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	QueryDetail(ctx *gin.Context)
	QueryPage(ctx *gin.Context)
	Unsubscribe(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

func (handler UserHandler) Logout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (handler UserHandler) QueryPage(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// Login
//
//	@Description: 用户登录api，只验证密码是否正确，其他校验通过路由中间件实现
//	@receiver handler 用户控制层
//	@param ctx
func (handler UserHandler) Login(ctx *gin.Context) {
	name := ctx.GetString("username")
	passwd := ctx.GetString("password")
	token := handler.srvApi.Login(name, passwd)

	if len(token) > 0 {
		ctx.JSON(200, view.Success(gin.H{"authorization": token}))
	} else {
		ctx.JSON(400, view.StatusWith(opstatus.WrongPassword))
	}
}

// Register
//
//	@Description:
//	@receiver handler 用户控制层
//	@param ctx
func (handler UserHandler) Register(ctx *gin.Context) {
	tmp, exist := ctx.Get("register_params")
	if !exist {
		ctx.JSON(400, view.Fail())
	}

	reg, _ := tmp.(view.RegisterReq)
	err := handler.srvApi.Register(
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

func (handler UserHandler) QueryDetail(ctx *gin.Context) {
	name := ctx.Param("username")
	if name != ctx.GetString("username") {
		ctx.JSON(400, view.StatusWith(opstatus.Unauthorized))
	} else {
		result := handler.srvApi.GetDetail(name)
		ctx.JSON(200, view.Success(result))
	}
}

func (handler UserHandler) Unsubscribe(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
