package handler

import (
	"GoProject/main/service"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	srvApi service.FileSrvApi
}

type FileHandleFunc interface {
	UploadFunc(ctx *gin.Context)
	DownloadFunc(ctx *gin.Context)
}

func NewFileHandler(srv *service.FileService) *FileHandler {
	return &FileHandler{srvApi: srv}
}
