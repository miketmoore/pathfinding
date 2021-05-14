package main

import (
	"fmt"

	"github.com/miketmoore/pathfinding"
)

func main() {

	nodeA := pathfinding.NewNode("a")
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

	edges := []*pathfinding.Edge{
		pathfinding.NewEdge(nodeA, nodeB, 1),
		pathfinding.NewEdge(nodeA, nodeC, 2),
		pathfinding.NewEdge(nodeB, nodeD, 2),
		pathfinding.NewEdge(nodeC, nodeE, 3),
		pathfinding.NewEdge(nodeD, nodeF, 6),
		pathfinding.NewEdge(nodeE, nodeF, 5),
	}

	shortestPath := pathfinding.Dijkstra(edges)
	fmt.Println(shortestPath)
}
