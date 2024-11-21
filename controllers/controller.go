package controllers

import (
	"net/http"

	"github.com/ahmadexe/prism-relay-network/data"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	network *data.PrismRelayNetwork
}

func InitController(n *data.PrismRelayNetwork) *Controller {
	return &Controller{
		network: n,
	}
}

func (c *Controller) GetNodes() {
	c.network.ViewNodes()
}

func (controller *Controller) GetRandomNode(ctx *gin.Context) {
	node := controller.network.GetNodeOnRandomIndex(ctx)
	if node != "" {
		ctx.JSON(http.StatusOK, node)
	}
}

func (controller *Controller) AddNode(ctx *gin.Context) {
	ip := ctx.Param("ip")
	controller.network.AddNode(ip)
	ctx.JSON(http.StatusOK, gin.H{"message": "Node added"})
}
