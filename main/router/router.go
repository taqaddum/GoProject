package router

import (
	"GoProject/main/handler"
	"GoProject/main/router/middleware"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	Engine  *gin.Engine
	UserApi handler.UserHandles
	FileApi handler.FileHandles
	RPCApi  handler.RPCHandles
}

func (app AppRouter) Start(addr string) {
	app.userRouter()
	app.fileRouter()
	app.rpcRouter()

	panic(app.Engine.Run(addr))
}

func (app AppRouter) userRouter() {
	router := app.Engine.Group("/users")
	{
		router.POST("/login", middleware.ValidLogin, app.UserApi.Login)
		router.POST("/register", middleware.ValidRegister, app.UserApi.Register)
		router.POST("/logout", middleware.AuthJWT, app.UserApi.Logout)
		router.DELETE("/:id", middleware.AuthJWT, app.UserApi.Unsubscribe)
		router.GET("/:username", middleware.AuthJWT, app.UserApi.QueryDetail)
		router.GET("/", middleware.AuthJWT, app.UserApi.QueryPage)
	}
}

func (app AppRouter) fileRouter() {
	router := app.Engine.Group("/files", middleware.AuthJWT)
	{
		router.POST("/", middleware.ValidFormFile, app.FileApi.Upload)
		router.GET("/:hash", app.FileApi.Download)
		router.GET("/:id/detail", app.FileApi.QueryDetail)
		router.GET("/", middleware.ValidPage, app.FileApi.QueryPage)
		router.DELETE("/:id", app.FileApi.Remove)
		router.PATCH("/:id", middleware.ValidFileName, app.FileApi.Rename)
	}
}

func (app AppRouter) rpcRouter() {
	router := app.Engine.Group("/proto", middleware.AuthJWT)
	{
		router.POST("/predict", app.RPCApi.Predict)
	}
}
