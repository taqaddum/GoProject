package router

import (
	"GoProject/main/handler"
	"GoProject/main/router/middleware"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	Engine     *gin.Engine
	UserHandle handler.UserHandleFunc
}

func (app AppRouter) Start() {
	app.userGroup()

	err := app.Engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func (app AppRouter) userGroup() {
	userGroup := app.Engine.Group("/users")
	{
		userGroup.POST("/login", middleware.ValidLoginParams, app.UserHandle.LoginFunc)
	}
}
