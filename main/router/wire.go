package router

import "github.com/google/wire"

var ProviderSet = wire.NewSet(wire.Struct(new(AppRouter), "*"))
