package router

import (
	"GoProject/main/handler"
	"GoProject/main/router/middleware"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	Engine  *gin.Engine
	UserApi handler.UserHandleFunc
	FileApi handler.FileHandleFunc
}

func (app AppRouter) Start(addr string) {
	app.userRouter()
	app.fileRouter()

	err := app.Engine.Run(addr)
	if err != nil {
		panic(err)
	}
}

func (app AppRouter) userRouter() {
	router := app.Engine.Group("/users")
	{
		router.POST("/login", middleware.ValidLogin, app.UserApi.LoginFunc)
		router.POST("/register", middleware.ValidRegister, app.UserApi.RegisterFunc)
		//router.GET("/logout", app.UserApi.LogoutFunc)
		router.GET("/:username", middleware.AuthJWT, app.UserApi.GetDetailFunc)
	}
}

func (app AppRouter) fileRouter() {
	router := app.Engine.Group("/files", middleware.AuthJWT)
	{
		router.POST("/upload", middleware.ValidFormFile, app.FileApi.UploadFunc)
		router.GET("/download", app.FileApi.DownloadFunc)
	}
}
