package main

import (
	"fmt"
	"os"

	"github.com/miketmoore/pathfinding"
)

func main() {

	graph := pathfinding.NewGraph()

	graph.AddEdge("0", "1", 4)
	graph.AddEdge("1", "2", 8)
	// graph.AddEdge("2", "3", 7)
	// graph.AddEdge("3", "4", 9)
	// graph.AddEdge("4", "5", 10)
	// graph.AddEdge("5", "6", 2)
	// graph.AddEdge("6", "7", 1)
	// graph.AddEdge("7", "8", 7)

	// graph.AddEdge("0", "7", 8)
	// graph.AddEdge("1", "7", 11)
	// graph.AddEdge("7", "8", 7)
	// graph.AddEdge("6", "8", 6)
	// graph.AddEdge("2", "8", 2)
	// graph.AddEdge("2", "5", 4)
	// graph.AddEdge("3", "5", 14)

	shortestPathTree, shortestPath, err := pathfinding.Dijkstra(graph)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for nodeId, node := range shortestPathTree {
		fmt.Printf("id=%s  distance=%f\n", nodeId, node.TentativeDistance)
	}

	edges := graph.FindEdgesForNodes(shortestPathTree)

	for edgeId, edge := range edges {
		fmt.Printf("id=%s  distance=%f\n", edgeId, edge.Distance)
	}

	fmt.Println("shortest path ", shortestPath)
	for _, node := range shortestPath {
		fmt.Printf("%s ", node.ID)
	}
	fmt.Println()
	fmt.Println(graph.GraphVizString())

}
