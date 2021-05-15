package pathfinding_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/pathfinding"
)

func TestDijkstra(t *testing.T) {

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

	edges := []*pathfinding.Edge{
		pathfinding.NewEdge(nodeA, nodeB, 1),
		pathfinding.NewEdge(nodeA, nodeC, 2),
		pathfinding.NewEdge(nodeB, nodeD, 2),
		pathfinding.NewEdge(nodeC, nodeE, 3),
		pathfinding.NewEdge(nodeD, nodeF, 6),
		pathfinding.NewEdge(nodeE, nodeF, 5),
	}

	shortestPath, err := pathfinding.Dijkstra(edges)
	if err != nil {
		fmt.Println(err)
		t.Errorf("error is unexpected")
	}

	if len(shortestPath) != 4 {
		t.Errorf("path length is unexpected: [%d]", len(shortestPath))
	}

	if shortestPath[0].ID != "a" ||
		shortestPath[1].ID != "b" ||
		shortestPath[2].ID != "d" ||
		shortestPath[3].ID != "f" {
		t.Errorf("path is invalid [%s]", pathToString(shortestPath))
	}
}

func pathToString(path []*pathfinding.Node) string {
	s := ""
	for _, node := range path {
		s = fmt.Sprintf("%s %s", s, node.ID)
	}
	return s
}
