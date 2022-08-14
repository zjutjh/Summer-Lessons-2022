package main

import (
	"day5/qa_system/app/midwares"
	"day5/qa_system/config/config"
	"day5/qa_system/config/database"
	"day5/qa_system/config/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.Init()

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)

	router.Init(r)

	err := r.Run(":" + config.Config.GetString("router.post"))
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}
