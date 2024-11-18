package types

import (
	"fmt"
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

func FindUnvisitedMinDistance(nodes Queue) *Node {
	minDist := math.Inf(1)
	var minNode *Node
	for _, node := range nodes {
		if !node.Visited && node.Distance < minDist {
			minDist = node.Distance
			minNode = node
		}
	}
	return minNode
}

func main() {
	// Create nodes
	nodeA := &Node{ID: 0}
	nodeB := &Node{ID: 1}
	nodeC := &Node{ID: 2}
	nodeD := &Node{ID: 3}

	// Create edges
	nodeA.Neighbors = []*Edge{
		&Edge{Destination: nodeB, Weight: 4},
		&Edge{Destination: nodeC, Weight: 2},
	}
	nodeB.Neighbors = []*Edge{
		&Edge{Destination: nodeD, Weight: 5},
	}
	nodeC.Neighbors = []*Edge{
		&Edge{Destination: nodeD, Weight: 1},
	}

	// Create graph
	graph := NewGraph([]*Node{nodeA, nodeB, nodeC, nodeD})
	nodeA.Graph = graph
	nodeB.Graph = graph
	nodeC.Graph = graph
	nodeD.Graph = graph

	// Run Dijkstra's algorithm
	Dijkstra(nodeA)

	// Print results
	fmt.Println("Shortest distances:")
	for _, node := range graph.Nodes {
		fmt.Printf("Node %d: %.2f\n", node.ID, node.Distance)
	}
}
