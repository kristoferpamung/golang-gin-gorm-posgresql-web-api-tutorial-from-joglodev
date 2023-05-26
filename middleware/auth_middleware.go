package middleware

import (
	"fmt"
	"joglo-dev/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid token.",
		})

		return
	}

	token := strings.Replace(bearerToken, "Bearer ", "", -1)

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})

		return
	}

	claimsData, err := utils.DecodeToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	ctx.Set("claimsData", claimsData)
	ctx.Set("user_id", claimsData["id"])
	ctx.Set("user_name", claimsData["name"])
	ctx.Set("email", claimsData["email"])

	ctx.Next()
}

func TokenMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	fmt.Println(token)

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})

		return
	}

	if token != "123" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "token not valid.",
		})

		return
	}

	ctx.Next()
}
