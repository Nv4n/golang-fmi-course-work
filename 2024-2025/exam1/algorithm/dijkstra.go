package algorithm

import (
	"exam1/types"
	"math"
)

var queue types.Queue

func Dijkstra(graph *types.Graph, source *types.Node, target *types.Node) map[*types.Node]*types.Node {
	prev := make(map[*types.Node]*types.Node)

	for _, n := range graph.Nodes {
		if n == source {
			n.Distance = 0
		} else {
			n.Distance = math.Inf(1)
		}
		prev[n] = nil
		queue.Push(n)
	}

	for queue.Len() > 0 {
		current := queue.Pop()
		if current == target {
			break
		}
		if current.Visited {
			continue
		}
		current.Visited = true

		for _, edge := range current.Neighbors {
			newDistance := current.Distance + edge.Weight
			if !edge.Destination.Visited && newDistance < edge.Destination.Distance {
				edge.Destination.Distance = newDistance
				//queue.Push(edge.Destination)
				prev[edge.Destination] = current
			}
		}
	}
	return prev
}

//function Dijkstra(Graph, source):
//
//    for each vertex v in Graph.Vertices:
//        dist[v] ← INFINITY
//        prev[v] ← UNDEFINED
//        add v to Q
//    dist[source] ← 0
//
//    while Q is not empty:
//         u ← vertex in Q with minimum dist[u]
//         remove u from Q
//
//         for each neighbor v of u still in Q:
//             alt ← dist[u] + Graph.Edges(u, v)
//             if alt < dist[v]:
//                 dist[v] ← alt
//                 prev[v] ← u
//
//     return dist[], prev[]
