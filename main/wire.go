//go:build wireinject
// +build wireinject

package main

import (
	"GoProject/main/component"
	"GoProject/main/handler"
	"GoProject/main/mapper"
	"GoProject/main/router"
	"GoProject/main/service"
	"github.com/google/wire"
)

func NewApp() *router.AppRouter {
	wire.Build(
		component.NewPostgres, component.NewRouter,
		mapper.ProviderSet, service.ProviderSet,
		handler.ProviderSet, router.ProviderSet)
	return nil
}
