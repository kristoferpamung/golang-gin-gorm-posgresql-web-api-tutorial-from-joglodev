package file_controller

import (
	"fmt"
	"joglo-dev/constanta"
	"joglo-dev/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SendStatus(ctx *gin.Context) {
	filename := ctx.MustGet("filename").(string)

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "file uploaded.",
		"file_name": filename,
	})
}

func HandleUploadFile(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(int)
	fmt.Println("userid : ", userId)

	fileHeader, _ := ctx.FormFile("file")

	if fileHeader == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "file is required.",
		})
		return
	}

	// fileType := []string{"image/jpeg"}

	// isFileValidated := utils.FileValidationByHeader(fileHeader, fileType)

	// if !isFileValidated {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"message": "file not allowed",
	// 	})

	// 	return
	// }

	fileExtension := []string{".png", "jpeg"}

	isFileValidatedByExtension := utils.FileValidationByExtension(fileHeader, fileExtension)

	if !isFileValidatedByExtension {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "file not allowed",
		})

		return
	}

	extentionFile := filepath.Ext(fileHeader.Filename)

	filename := utils.RandomFileName(extentionFile)

	isSaved := utils.SaveFile(ctx, fileHeader, filename)

	if !isSaved {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error, can't save file.",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "file uploaded.",
	})
}

func HandleRemoveFile(ctx *gin.Context) {

	filename := ctx.Param("filename")

	if filename == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "file nameis required.",
		})

		return
	}

	errRemoveFile := utils.RemoveFile(constanta.DIR_FILE + filename)

	if errRemoveFile != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errRemoveFile.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "file deleted.",
	})
}
