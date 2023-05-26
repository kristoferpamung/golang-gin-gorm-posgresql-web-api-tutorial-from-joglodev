package routes

import (
	"joglo-dev/config/app_config"
	"joglo-dev/controllers/book_controller"
	"joglo-dev/controllers/user_controller"
	"joglo-dev/controllers/user_controller/auth_contoller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app.Group("api")

	// ROUTE STATIC
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	// ROUTE USER
	userRoute := route.Group("user")
	userRoute.GET("/", user_controller.GetAllUser)
	userRoute.GET("/:id", user_controller.GetUserByID)
	userRoute.GET("/paginate", user_controller.GetUserPaginate)
	userRoute.POST("/", user_controller.StoreUser)
	userRoute.PATCH(":id", user_controller.UpdateUserById)
	userRoute.DELETE("/:id", user_controller.DeleteUserById)

	// ROUTE BOOK
	route.GET("/book", book_controller.GetAllBook)

	// ROUTE FILE

	v1Route(route)

	// LOGIN
	route.POST("/login", auth_contoller.Login)
}
