package handler

import (
	"GoProject/main/enum/httpstatus"
	"GoProject/main/service"
	"GoProject/main/view"
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
		ctx.JSON(400, view.StatusWith(httpstatus.InvalidParams))
		slog.Error("文件上传参数错误", err.Error())
		return
	}

	err = handle.srvApi.SaveFile("admin", file)
	if err != nil {
		ctx.JSON(400, view.StatusWith(httpstatus.UploadFileError))
		return
	}

	ctx.JSON(200, view.Success())
}

func (handle *FileHandler) DownloadFunc(ctx *gin.Context) {

}
