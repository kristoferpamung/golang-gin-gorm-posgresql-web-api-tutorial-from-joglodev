package middleware

import (
	"fmt"
	"joglo-dev/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	// claimsData := ctx.MustGet("claimsData").(jwt.MapClaims)

	// fmt.Println("claimsData => email =>", claimsData["email"])

	userId := ctx.MustGet("user_id").(float64)
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

	ctx.Set("filename", filename)

	ctx.Next()
}
