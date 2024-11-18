package types

import (
	"math"
)

type Graph struct {
	Nodes []*Node
}

type Node struct {
	ID        int
	Distance  float64
	Visited   bool
	Neighbors []*Edge
}

type Edge struct {
	Destination *Node
	Weight      float64
}

func NewGraph(nodes []*Node) *Graph {
	return &Graph{Nodes: nodes}
}

func InitializeDistances(graph Graph) {
	for _, node := range graph.Nodes {
		node.Distance = math.Inf(1)
		node.Visited = false
	}
}
