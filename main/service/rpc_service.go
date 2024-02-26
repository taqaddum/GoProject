package service

import (
	"GoProject/main/config"
	"log/slog"
	"net/rpc"
)

type RPCService struct {
}

func NewRPCService() *RPCService {
	return &RPCService{}
}

type RPCSrvApi interface {
	Predict(hash []string) string
}

func (srv RPCService) Predict(hash []string) string {
	client, err := rpc.Dial("tcp", config.GetServerAddr())
	defer client.Close()

	var reply string
	err = client.Call("Yolo.Predict", hash, &reply)
	if err != nil {
		slog.Error(err.Error())
		return ""
	}

	return reply
}
