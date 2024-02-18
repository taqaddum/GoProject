package mapper

import (
	"GoProject/main/model"
	"xorm.io/xorm"
)

type FileMapper struct {
	db *xorm.Engine
}

type FileMapApi interface {
	Create(file *model.File) error
}

func NewFileMapper(engine *xorm.Engine) *FileMapper {
	return &FileMapper{db: engine}
}

func (f FileMapper) Create(file *model.File) error {
	//TODO implement me
	panic("implement me")
}
