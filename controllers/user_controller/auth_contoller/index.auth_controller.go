package auth_contoller

import (
	"joglo-dev/database"
	"joglo-dev/models"
	"joglo-dev/requests"
	"joglo-dev/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context) {

	loginReq := new(requests.LoginRequest)

	err := ctx.ShouldBindJSON(&loginReq)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := new(models.User)

	errUser := database.DB.Table("users").Where("email = ?", loginReq.Email).Find(&user).Error

	if errUser != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "credential not valid.",
		})
		return
	}

	if user.Email == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "credential not valid.",
		})
		return
	}

	// Check Password
	if loginReq.Password != "12345" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "credential not valid.",
		})
		return
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed generate token.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login succesfully",
		"token":   token,
	})
}
