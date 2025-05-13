package main

import (
	guitar_config "server/guitar/config"
	guitar_core "server/guitar/core"
	guitar_route "server/guitar/route"

	"server/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func main(){
	database.ConnectDb()

	database.DB.AutoMigrate(&guitar_core.Guitar{})

	guitar := guitar_config.InitGuitarFeature()

	r := gin.Default()

	guitar_route.GuitarRoutes(r, guitar.Controller)

	r.Run(":8080")
}