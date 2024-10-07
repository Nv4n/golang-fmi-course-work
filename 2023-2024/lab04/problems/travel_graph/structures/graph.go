package graph

import (
	"fmt"
	"math"
)

type Path struct {
	Distance      float64
	Stops         []int
	DistanceStops []float64
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

	for i := range g.vertices {
		dist[i] = &Path{math.MaxFloat64, []int{}, []float64{}}
		visited[i] = false
	}
	dist[from].Distance = 0

	for range g.vertices {
		minDistInd := minDistance(dist, visited)
		visited[minDistInd] = true

		for v := 0; v < len(g.vertices); v++ {
			if vDist, ok := g.vertices[minDistInd][v]; !visited[v] &&
				ok &&
				dist[minDistInd].Distance != math.MaxFloat64 &&
				dist[minDistInd].Distance+vDist < dist[v].Distance {
				dist[v].Distance = dist[minDistInd].Distance + vDist
				dist[v].Stops = append(dist[v].Stops, dist[minDistInd].Stops...)
				dist[v].Stops = append(dist[v].Stops, minDistInd)
				dist[v].DistanceStops = append(dist[v].DistanceStops, dist[minDistInd].DistanceStops...)
				dist[v].DistanceStops = append(dist[v].DistanceStops, vDist)
			}
		}
	}

	for i, path := range dist {
		fmt.Printf("\nV:%d Dist:%f-->\nStops:%v\nDistances:%v\n", i, path.Distance, path.Stops, path.DistanceStops)

	}
}

func minDistance(dist map[int]*Path, visited map[int]bool) int {
	minDist := math.MaxFloat64
	var minIndex int
	for i, path := range dist {
		if !visited[i] && path.Distance <= minDist {
			minDist = path.Distance
			minIndex = i
		}
	}
	return minIndex
}
