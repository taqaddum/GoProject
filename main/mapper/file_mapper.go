package mapper

import (
	"GoProject/main/model"
	"errors"
	"log/slog"
	"xorm.io/xorm"
)

type FileMapper struct {
	db *xorm.Engine
}

func NewFileMapper(engine *xorm.Engine) *FileMapper {
	return &FileMapper{db: engine}
}

type FileMapApi interface {
	Create(file *model.File) error
	GetPathByHash(hash string) string
}

func (mapper FileMapper) Create(file *model.File) error {
	if mapper.IsExist(file.Hash) {
		return errors.New("文件已存在")
	}

	_, err := mapper.db.Insert(file)
	if err != nil {
		slog.Error(err.Error())
	}

	return err
}

func (mapper FileMapper) GetPathByHash(hash string) string {
	file := &model.File{Hash: hash}
	if _, err := mapper.db.Get(file); err != nil {
		slog.Error(err.Error())
		return ""
	}
	return file.Path
}

func (mapper FileMapper) IsExist(hash string) bool {
	file := &model.File{Hash: hash}
	if _, err := mapper.db.Get(file); err != nil {
		return false
	}
	return true
}
