package service

import (
	"GoProject/main/config"
	"GoProject/main/mapper"
	"GoProject/main/util"
	"log/slog"
	"mime/multipart"
	"path/filepath"
	"time"
)

type FileService struct {
	mapApi mapper.FileMapApi
}

type FileSrvApi interface {
	SaveFile(username string, f *multipart.FileHeader) error
	SendFile()
}

func NewFileService(fileMapper *mapper.FileMapper) *FileService {
	// TODO to be continued
	return &FileService{mapApi: fileMapper}
}

func (srv FileService) SaveFile(owner string, f *multipart.FileHeader) error {
	src, err := f.Open()
	defer src.Close()
	if err != nil {
		slog.Error("文件流错误", err.Error())
		return err
	}

	dir := filepath.Join(config.GetSaveDir(), owner, time.Now().Format("20060102"))

	err = util.CopyFile(src, dir, f.Filename)
	if err != nil {
		slog.Error("文件上传异常", err.Error())
		return err
	}

	// TODO 需要计算hash，存入files表

	return nil
}

func (srv FileService) SendFile() {
	//TODO implement me
	panic("implement me")
}
