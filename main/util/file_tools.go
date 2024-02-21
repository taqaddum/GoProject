package util

import (
	"io"
	"log/slog"
	"mime/multipart"
	"os"
)

func CopyFile(src multipart.File, dir string, filename string) error {
	if err := os.MkdirAll(dir, 0664); err != nil {
		slog.Error("创建目录失败！", err.Error())
		return err
	}

	dst, err := os.Create(dir + "/" + filename)
	defer dst.Close()
	if err != nil {
		slog.Error("创建文件失败！", err.Error())
		return err
	}

	if _, err := io.Copy(dst, src); err != nil {
		slog.Error("复制文件失败！", err.Error())
		return err
	}

	return nil
}
