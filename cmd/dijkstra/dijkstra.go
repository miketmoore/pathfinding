package main

import (
	"fmt"
	"os"

	"github.com/miketmoore/pathfinding"
)

func main() {

	graph := pathfinding.NewGraph()

	graph.NewSourceNode("A")
	// graph.NewDestinationNode("")
	graph.AddEdge("A", "B", 3)
	graph.AddEdge("B", "E", 1)
	graph.AddEdge("E", "D", 7)
	graph.AddEdge("B", "D", 5)
	graph.AddEdge("C", "D", 2)
	graph.AddEdge("B", "C", 7)
	graph.AddEdge("A", "C", 1)

	shortestPathGraph, nodeDistances, err := pathfinding.DijkstraAllPaths(graph, "C")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("node distances", nodeDistances)
	fmt.Println(shortestPathGraph.GraphVizString("shortestPath"))

}
