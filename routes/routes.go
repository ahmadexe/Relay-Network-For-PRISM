package routes

import (
	"github.com/ahmadexe/prism-relay-network/controllers"
	"github.com/gin-gonic/gin"
)

type AppRoutes struct {
	router *gin.Engine
	controller *controllers.Controller
}

func InitRouter(c *controllers.Controller, r *gin.Engine) *AppRoutes {
	return &AppRoutes{
		controller: c,
		router: r,
	}
}

func (r *AppRoutes) InitRoutes() {
	r.router.Group("/api/v1/relay") 

	// r.router.HandleFunc("/api/v1/relay", r.controller.GetRelay).Methods("GET")
	// r.router.HandleFunc("/api/v1/relay", r.controller.CreateRelay).Methods("POST")
	// r.router.HandleFunc("/api/v1/relay/{id}", r.controller.GetRelayById).Methods("GET")
	// r.router.HandleFunc("/api/v1/relay/{id}", r.controller.UpdateRelay).Methods("PUT")
	// r.router.HandleFunc("/api/v1/relay/{id}", r.controller.DeleteRelay).Methods("DELETE")
}


