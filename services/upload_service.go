package services

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveImage(file multipart.File, fileName string) error {

	err := os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return err
	}

	dstPath := filepath.Join("uploads", fileName)

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil
}
