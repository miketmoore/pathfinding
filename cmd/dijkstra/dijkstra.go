package main

import (
	"fmt"
	"os"

	"github.com/miketmoore/pathfinding"
)

func main() {

	nodeA := pathfinding.NewSourceNode("a")
	nodeB := pathfinding.NewNode("b")
	nodeC := pathfinding.NewNode("c")
	nodeD := pathfinding.NewNode("d")
	nodeE := pathfinding.NewNode("e")
	nodeF := pathfinding.NewDestinationNode("f")
	//  A
	//  | \
	//  |  \
	//  |   \
	//  |    \
	//  1     2
	//  |     |
	// [B]   [C]
	//  |     |
	//  2     3
	//  |     |
	// [D]   [E]
	//  |     |
	//  6     5
	//  |    /
	//  |   /
	//  |  /
	//  | /
	// [F]

	// shortest path is A > B > D > F (distance=9)
	// longest path is A > C > E > F (distance=10)

	graph := pathfinding.NewGraph()

	// nodeX := pathfinding.NewNode("x")
	// graph.AddNode(nodeX)

	graph.AddEdge(nodeA, nodeB, 1)
	graph.AddEdge(nodeA, nodeC, 2)
	graph.AddEdge(nodeB, nodeD, 2)
	graph.AddEdge(nodeC, nodeE, 3)
	graph.AddEdge(nodeD, nodeF, 6)
	graph.AddEdge(nodeE, nodeF, 5)

	shortestPath, err := pathfinding.Dijkstra(graph)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for nodeId := range shortestPath {
		fmt.Printf("id=%s\n", nodeId)
	}

	edges := graph.FindEdgesForNodes(shortestPath)

	for edgeId := range edges {
		fmt.Printf("id=%s\n", edgeId)
	}

}
