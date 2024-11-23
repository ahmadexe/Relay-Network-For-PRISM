package main

import (
	"github.com/ahmadexe/prism-relay-network/configs"
	"github.com/ahmadexe/prism-relay-network/controllers"
	"github.com/ahmadexe/prism-relay-network/data"
	"github.com/ahmadexe/prism-relay-network/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs := configs.InitConfigs()
	gin.SetMode(configs.Mode)

	network := data.InitNetwork()
	controller := controllers.InitController(network)

	router := gin.Default()
	networkRouter := routes.InitRouter(controller, router)

	networkRouter.InitRoutes()

	go network.PingNodes() 

	router.Run(":" + configs.Port)
}
