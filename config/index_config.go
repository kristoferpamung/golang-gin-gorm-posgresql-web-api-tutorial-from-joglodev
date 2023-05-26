package configs

import (
	"joglo-dev/config/app_config"
	"joglo-dev/config/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
}
