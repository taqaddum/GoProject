package handler

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserHandler, wire.Bind(new(UserHandles), new(*UserHandler)),
	NewFileHandler, wire.Bind(new(FileHandles), new(*FileHandler)),
	NewRPCHandler, wire.Bind(new(RPCHandles), new(*RPCHandler)),
)
