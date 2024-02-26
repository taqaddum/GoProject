package service

import (
	"GoProject/main/config"
	"GoProject/main/mapper"
	"GoProject/main/util"
	"log/slog"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileService struct {
	mapApi mapper.FileMapApi
}

func NewFileService(fileMapper *mapper.FileMapper) *FileService {
	return &FileService{mapApi: fileMapper}
}

type FileSrvApi interface {
	SaveFile(username string, f *multipart.FileHeader) error
	ObtainFile(hash string) *os.File
}

func (srv FileService) SaveFile(owner string, f *multipart.FileHeader) error {
	src, err := f.Open()
	defer src.Close()
	if err != nil {
		slog.Error("文件流错误", err.Error())
	}

	dir := filepath.Join(config.GetSaveDir(), owner, time.Now().Format("20060102"))
	err = util.CopyFile(src, dir, f.Filename)
	if err != nil {
		slog.Error("文件上传异常", err.Error())
	}

	// TODO 需要计算hash，存入files表

	return err
}

func (srv FileService) ObtainFile(hash string) *os.File {
	filePath := srv.mapApi.GetPathByHash(hash) // 根据哈希值获取文件路径
	if filePath == "" {
		slog.Error("文件不存在")
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("文件打开失败", err.Error())
	}

	return file
}
