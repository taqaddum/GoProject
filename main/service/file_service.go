package service

import "GoProject/main/mapper"

type FileService struct {
	mapApi mapper.FileMapApi
}

type FileSrvApi interface {
	Upload()
	Download()
}

func NewFileService(fileMapper *mapper.FileMapper) *FileService {
	// TODO to be continued
	return &FileService{mapApi: fileMapper}
}

func (f FileService) Upload() {
	//TODO implement me
	panic("implement me")
}

func (f FileService) Download() {
	//TODO implement me
	panic("implement me")
}
