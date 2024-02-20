package service

import (
	"GoProject/main/mapper"
	"mime/multipart"
)

type FileService struct {
	mapApi mapper.FileMapApi
}

type FileSrvApi interface {
	Upload(file *multipart.FileHeader) error
	Download()
}

func NewFileService(fileMapper *mapper.FileMapper) *FileService {
	// TODO to be continued
	return &FileService{mapApi: fileMapper}
}

func (srv FileService) Upload(file *multipart.FileHeader) error {
	//TODO implement me
	panic("implement me")
}

func (srv FileService) Download() {
	//TODO implement me
	panic("implement me")
}
