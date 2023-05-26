package utils

import (
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {

	rand.Seed(time.Now().UnixMilli())

	b := make([]byte, n)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func RandomFileName(extensionFile string, prefix ...string) string {
	currentPrefix := "file"

	if len(prefix) > 0 {
		if prefix[0] != "" {
			currentPrefix = prefix[0]
		}
	}

	currentTime := time.Now().Format("20060102")

	filename := fmt.Sprintf("%s-%s-%s%s", currentPrefix, currentTime, RandomString(5), extensionFile)

	return filename
}

func FileValidationByHeader(fileHeader *multipart.FileHeader, fileType []string) bool {
	contentType := fileHeader.Header.Get("Content-Type")

	log.Println("content-type ", contentType)

	result := false

	for _, typeFile := range fileType {
		if contentType == typeFile {
			result = true
			break
		}
	}

	return result
}

func FileValidationByExtension(fileHeader *multipart.FileHeader, fileExtension []string) bool {
	extension := filepath.Ext(fileHeader.Filename)

	log.Println("extension ", extension)

	result := false

	for _, typeFile := range fileExtension {
		if extension == typeFile {
			result = true
			break
		}
	}

	return result
}

func SaveFile(ctx *gin.Context, fileHeader *multipart.FileHeader, filename string) bool {
	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", filename))

	if errUpload != nil {
		log.Println("Can't save file.")
		return false
	} else {
		return true
	}
}

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)

	if err != nil {
		log.Println("Failed to remove file")
		return err
	}

	return nil
}
