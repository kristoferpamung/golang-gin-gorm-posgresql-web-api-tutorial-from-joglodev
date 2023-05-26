package book_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get All Book Function",
	})
}
