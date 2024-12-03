package data

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
	"net/http"
)

type PrismRelayNetwork struct {
    nodes []string
}

func InitNetwork() (*PrismRelayNetwork) {
	return &PrismRelayNetwork{nodes: []string{"Ahmad", "Rizky", "Fikri"}}
}

func (network *PrismRelayNetwork) AddNode(ip string) {
    network.nodes = append(network.nodes, ip)
}

func (network *PrismRelayNetwork) RemoveNode(ip string) {
	for i, v := range network.nodes {
		if v == ip {
			network.nodes = append(network.nodes[:i], network.nodes[i + 1:]...)
		}
	}
}

func (network *PrismRelayNetwork) ViewNodes() {
	fmt.Println(network.nodes)
}

func (network *PrismRelayNetwork) GetNodes() ([]string) {
	return network.nodes
}

func (network *PrismRelayNetwork) GetNodeOnRandomIndex(ctx *gin.Context) (string) {
	if len(network.nodes) == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "No nodes available"})
		return ""
	}	

	return network.nodes[rand.Intn(len(network.nodes))]
}
func (network *PrismRelayNetwork) PingNodes() {
	copiedList := make([]string, len(network.nodes))
	copy(copiedList, network.nodes)

	for _, node := range copiedList {
		fmt.Println("Pinging node: " + node)
		_, err := http.Get("http://" + node + "/is_alive")
		if err != nil {
			network.RemoveNode(node)
		}
	}
}