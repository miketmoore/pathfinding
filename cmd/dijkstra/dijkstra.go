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

	// graph.NewSourceNode("0")
	// graph.NewDestinationNode("8")

	// graph.AddEdge("0", "1", 4)
	// graph.AddEdge("1", "2", 8)
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

	shortestPathGraph, nodeDistances, err := pathfinding.DijkstraAllPaths(graph, "C")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("node distances", nodeDistances)
	fmt.Println(shortestPathGraph.GraphVizString("shortestPath"))

	// edges := graph.FindEdgesForNodes(shortestPathTree)
	// fmt.Println("---")
	// fmt.Println("edges:")
	// fmt.Println("---")
	// for edgeId, edge := range edges {
	// 	fmt.Printf("id=%s  distance=%f\n", edgeId, edge.Distance)
	// }

	// fmt.Println(graph.GraphVizString("example"))

	// // build shortest path graph
	// spGraph := pathfinding.NewGraph()

	// spGraph.NewSourceNode("0")
	// spGraph.NewDestinationNode("8")
	// for _, node := range shortestPathTree {
	// 	edgesMap := graph.FindEdgesForNode(node)
	// 	for _, edge := range edgesMap {
	// 		spGraph.AddEdge(edge.NodeA.ID, edge.NodeB.ID, edge.Distance)
	// 	}
	// }
	// fmt.Println(spGraph.GraphVizString("shortestPathGraph"))

	// for _, edge := range graph.Edges {
	// 	fmt.Printf("original weight=%.2f\n", edge.Distance)
	// }

}
