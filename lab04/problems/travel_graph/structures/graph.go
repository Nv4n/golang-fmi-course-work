package graph

import (
	"fmt"
	"math"
)

type Path struct {
	Distance float64
	Stops    []int
}

type Graph struct {
	vertices map[int]map[int]float64
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int]map[int]float64),
	}
}

func (g *Graph) AddVertex(vertex int) {
	if _, exist := g.vertices[vertex]; !exist {
		g.vertices[vertex] = make(map[int]float64)
	}
}

func (g *Graph) AddEdge(from, to int, distance float64) {
	g.vertices[from][to] = distance
	g.vertices[to][from] = distance
}

func (g *Graph) FindShortestPath(from, to int) {
	dist := make(map[int]*Path)
	visited := make(map[int]bool)

	for vertex := range g.vertices {
		dist[vertex] = &Path{math.MaxFloat64, []int{}}
		visited[vertex] = false
	}
	dist[from].Distance = 0
	fmt.Println(dist)
}
