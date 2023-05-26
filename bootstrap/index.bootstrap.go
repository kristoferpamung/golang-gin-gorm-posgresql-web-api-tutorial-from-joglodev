package bootstrap

import (
	configs "joglo-dev/config"
	"joglo-dev/config/app_config"
	"joglo-dev/config/cors_config"
	"joglo-dev/database"
	"joglo-dev/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootStrapApp() {

	// LOAD .env FILE
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	// INIT CONFIG
	configs.InitConfig()

	// DATABASE CONNECTION
	database.ConnectDatabase()

	// INIT GIN ENGINE
	app := gin.Default()

	// CORS
	app.Use(cors_config.CorsConfig)

	// INIT ROUTE
	routes.InitRoute(app)

	// RUN APP
	app.Run(app_config.APP_PORT)
}
