package main

import (
	"class-schedule/config"
	"class-schedule/db"
	"class-schedule/db/model"
	"class-schedule/router"
    "github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	db.InitDB()
	model.InitModel()
	if !config.Config.Dev {
		gin.SetMode(gin.ReleaseMode)
	}
	router.InitRouter()
	router.Router.Run(":" + config.Config.Server.Port)
}
