package util

import (
	"log/slog"
	"os"
	"path/filepath"
)

func NewFile(path string, data []byte) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0664); err != nil {
		slog.Error("创建目录失败！", err.Error())
		panic(err)
	}

	if err := os.WriteFile(path, data, 0664); err != nil {
		slog.Error("文件写入失败！", err.Error())
		panic(err)
	}
}
