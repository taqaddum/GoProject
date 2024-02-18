package handler

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserHandler, wire.Bind(new(UserHandleFunc), new(*UserHandler)),
	NewFileHandler, wire.Bind(new(FileHandleFunc), new(*FileHandler)),
)
