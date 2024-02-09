// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"GoProject/main/component"
	"GoProject/main/handler"
	"GoProject/main/mapper"
	"GoProject/main/router"
	"GoProject/main/service"
)

// Injectors from wire.go:

func NewApp() *router.AppRouter {
	engine := component.NewRouter()
	xormEngine := component.NewPostgres()
	userMapper := mapper.NewUserMapper(xormEngine)
	userService := service.NewUserService(userMapper)
	userHandler := handler.NewUserHandler(userService)
	appRouter := &router.AppRouter{
		Engine:     engine,
		UserHandle: userHandler,
	}
	return appRouter
}
