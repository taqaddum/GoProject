package handler

import (
	"GoProject/main/enum/opstatus"
	"GoProject/main/service"
	"GoProject/main/view"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
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

func (handle *FileHandler) UploadFunc(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, view.StatusWith(opstatus.InvalidParams))
		slog.Error("upload file error", err.Error())
	} else {
		err = handle.srvApi.Upload(file)
	}

	switch {
	case errors.Is(err, nil):
		ctx.JSON(200, view.Success())

	}
}

func (handle *FileHandler) DownloadFunc(ctx *gin.Context) {

}
