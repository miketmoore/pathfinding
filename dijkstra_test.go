package pathfinding_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/pathfinding"
)

func TestDijkstra(t *testing.T) {

	tests := []struct {
		getEdges func() []*pathfinding.Edge
		expected []string
	}{
		{
			expected: []string{"a", "b", "d", "f"},
			getEdges: func() []*pathfinding.Edge {

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

				return []*pathfinding.Edge{
					pathfinding.NewEdge(nodeA, nodeB, 1),
					pathfinding.NewEdge(nodeA, nodeC, 2),
					pathfinding.NewEdge(nodeB, nodeD, 2),
					pathfinding.NewEdge(nodeC, nodeE, 3),
					pathfinding.NewEdge(nodeD, nodeF, 6),
					pathfinding.NewEdge(nodeE, nodeF, 5),
				}
			},
		},
	}

	for _, test := range tests {

		shortestPath, err := pathfinding.Dijkstra(test.getEdges())
		if err != nil {
			fmt.Println(err)
			t.Errorf("error is unexpected")
		}

		gotLength := len(shortestPath)
		expectedLength := len(test.expected)

		if gotLength != expectedLength {
			t.Errorf("length is %d but is expected to be %d", gotLength, expectedLength)
		}

		for pathIndex, expectedPathValue := range test.expected {
			if shortestPath[pathIndex].ID != expectedPathValue {
				t.Errorf("path is invalid [%s]", pathToString(shortestPath))
			}
		}

	}

}

func pathToString(path []*pathfinding.Node) string {
	s := ""
	for _, node := range path {
		s = fmt.Sprintf("%s %s", s, node.ID)
	}
	return s
}
