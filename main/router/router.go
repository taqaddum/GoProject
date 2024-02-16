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

func (app AppRouter) Start(addr string) {
	app.userRouter()

	err := app.Engine.Run(addr)
	if err != nil {
		panic(err)
	}
}

func (app AppRouter) userRouter() {
	userGroup := app.Engine.Group("/users")
	{
		userGroup.POST("/login", middleware.ValidLogin, app.UserHandle.LoginFunc)
		userGroup.POST("/register", middleware.ValidRegister, app.UserHandle.RegisterFunc)
		//userGroup.GET("/logout", app.UserHandle.LogoutFunc)
		userGroup.GET("/:username", middleware.AuthJWT, app.UserHandle.GetDetailFunc)
	}
}

func (app AppRouter) fileRouter() {
	fileGroup := app.Engine.Group("/files")
	{
		fileGroup.POST("/upload", app.UserHandle.UploadFileFunc)
		fileGroup.GET("/download", app.UserHandle.DownloadFileFunc)
	}
}
