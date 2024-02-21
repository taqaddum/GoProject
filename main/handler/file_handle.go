package handler

import (
	"GoProject/main/service"
	"GoProject/main/view"
	"github.com/gin-gonic/gin"
	"log/slog"
	"mime/multipart"
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
	tmp, has := ctx.Get("file")
	file := tmp.(*multipart.FileHeader)
	if !has {
		slog.Error("上传文件不存在")
		ctx.JSON(400, view.Fail())
		return
	}

	err := handle.srvApi.SaveFile("admin", file)
	if err != nil {
		ctx.JSON(500, view.Fail())
		slog.Error(err.Error())
		return
	}

	ctx.JSON(200, view.Success())
}

func (handle *FileHandler) DownloadFunc(ctx *gin.Context) {
}
