package handler

import (
	"GoProject/main/service"
	"github.com/gin-gonic/gin"
)

type RPCHandler struct {
	srvApi service.RPCSrvApi
}

func NewRPCHandler(srv *service.RPCService) *RPCHandler {
	return &RPCHandler{srvApi: srv}
}

type RPCHandles interface {
	Predict(ctx *gin.Context)
}

func (handler RPCHandler) Predict(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
