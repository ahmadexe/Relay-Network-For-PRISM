package data

import "fmt"

type PrismRelayNetwork struct {
    nodes []string
}

func InitNetwork() (*PrismRelayNetwork) {
	return &PrismRelayNetwork{nodes: []string{}}
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

func (network *PrismRelayNetwork) ViewNode() {
	fmt.Println(network.nodes)
}

func (network *PrismRelayNetwork) GetNode() ([]string) {
	return network.nodes
}
