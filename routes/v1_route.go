package routes

import (
	"joglo-dev/controllers/file_controller"
	"joglo-dev/middleware"

	"github.com/gin-gonic/gin"
)

func v1Route(routeGrup *gin.RouterGroup) {

	route := routeGrup

	// ROUTE YANG MEMBUTUHKAN MIDDLEWARE DAPAT DI GRUPKAN
	authRoute := route.Group("file", middleware.AuthMiddleware)
	authRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
	authRoute.POST("/", file_controller.HandleUploadFile)
	authRoute.POST("/middleware", middleware.AuthMiddleware, middleware.UploadFile, file_controller.SendStatus)
}
