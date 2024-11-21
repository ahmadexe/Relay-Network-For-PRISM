package routes

import (
	"github.com/ahmadexe/prism-relay-network/controllers"
	"github.com/gin-gonic/gin"
)

type AppRoutes struct {
	router     *gin.Engine
	controller *controllers.Controller
}

func InitRouter(c *controllers.Controller, r *gin.Engine) *AppRoutes {
	return &AppRoutes{
		controller: c,
		router:     r,
	}
}

func (r *AppRoutes) InitRoutes() {
	router := r.router.Group("/api/v1")
	{
		router.GET("/rand/node", r.controller.GetRandomNode)
		router.POST("/add/:ip", r.controller.AddNode)
	}
}
