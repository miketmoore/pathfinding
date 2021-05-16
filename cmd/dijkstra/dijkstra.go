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

	shortestPathTree, parent, err := pathfinding.Dijkstra(graph)
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

	fmt.Println("parent path ", parent)
	for _, node := range parent {
		fmt.Printf("%s ", node.ID)
	}
	fmt.Println()

	// graph ethane {
	// 	C_0 -- H_0 [type=s];
	// 	C_0 -- H_1 [type=s];
	// 	C_0 -- H_2 [type=s];
	// 	C_0 -- C_1 [type=s];
	// 	C_1 -- H_3 [type=s];
	// 	C_1 -- H_4 [type=s];
	// 	C_1 -- H_5 [type=s];
	// }

	fmt.Println(graph.GraphVizString())

}
