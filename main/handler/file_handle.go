package handler

import (
	"GoProject/main/service"
	"GoProject/main/view"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"time"
)

type FileHandler struct {
	srvApi service.FileSrvApi
}

func NewFileHandler(srv *service.FileService) *FileHandler {
	return &FileHandler{srvApi: srv}
}

type FileHandles interface {
	Upload(ctx *gin.Context)
	Download(ctx *gin.Context)
	QueryDetail(ctx *gin.Context)
	QueryPage(ctx *gin.Context)
	Remove(ctx *gin.Context)
	Rename(ctx *gin.Context)
}

func (handle *FileHandler) QueryDetail(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (handle *FileHandler) QueryPage(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (handle *FileHandler) Remove(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (handle *FileHandler) Rename(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (handle *FileHandler) Upload(ctx *gin.Context) {
	tmp, _ := ctx.Get("file")
	file := tmp.(*multipart.FileHeader)

	err := handle.srvApi.SaveFile("admin", file)
	if err != nil {
		ctx.JSON(500, view.Fail())
		slog.Error(err.Error())
		return
	}

	ctx.JSON(200, view.Success())
}

func (handle *FileHandler) Download(ctx *gin.Context) {
	hash := ctx.Param("hash") // 获取哈希值
	file := handle.srvApi.ObtainFile(hash)
	defer file.Close()

	// 设置响应头
	fileInfo, err := file.Stat()
	if err != nil {
		ctx.JSON(500, view.Fail())
		return
	}

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name()))
	ctx.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// 开始下载时间
	startTime := time.Now()
	// 复制文件内容到响应体
	if _, err := io.Copy(ctx.Writer, file); err != nil {
		ctx.JSON(500, view.Fail())
		return
	}
	// 计算下载完成时间
	completionTime := time.Since(startTime)

	// 在响应体中返回下载完成时间
	ctx.JSON(http.StatusOK, view.Success(gin.H{"completion_time": completionTime.String()}))
}
