package pathfinding_test

import (
	"testing"

	"github.com/miketmoore/pathfinding"
)

func TestFindAdjacentNodes(t *testing.T) {

	tests := []struct {
		getGraph        func() *pathfinding.Graph
		input           string
		expectedNodeIds []string
	}{
		{
			getGraph: func() *pathfinding.Graph {
				return pathfinding.NewGraph()
			},
			input:           "a",
			expectedNodeIds: []string{},
		},
		{
			getGraph: func() *pathfinding.Graph {
				graph := pathfinding.NewGraph()
				graph.AddEdge("a", "b", 1)
				graph.AddEdge("a", "c", 1)
				graph.AddEdge("a", "d", 1)
				return graph
			},
			input:           "a",
			expectedNodeIds: []string{"b", "c", "d"},
		},
		{
			getGraph: func() *pathfinding.Graph {
				graph := pathfinding.NewGraph()
				graph.AddEdge("a", "b", 1)
				graph.AddEdge("a", "c", 1)
				graph.AddEdge("a", "d", 1)
				return graph
			},
			input:           "b",
			expectedNodeIds: []string{"a"},
		},
	}

	for _, test := range tests {
		graph := test.getGraph()
		found := pathfinding.FindAdjacentNodes(graph, test.input)
		if len(found) != len(test.expectedNodeIds) {
			t.Errorf("length is unexpected - got=%d expected=%d", len(found), len(test.expectedNodeIds))
		}
		if !ContainsInSliceNode(found, test.expectedNodeIds...) {
			t.Error("contains unexpected values")
		}
	}

}

func ContainsInSliceNode(nodes []*pathfinding.Node, expectedNodeIds ...string) bool {
	matches := 0
	for _, expectedNodeId := range expectedNodeIds {
		for _, node := range nodes {
			if node.ID == expectedNodeId {
				matches++
				continue
			}
		}
	}
	return matches == len(nodes)
}

// func TestDijkstra(t *testing.T) {

// 	tests := []struct {
// 		getGraph func() *pathfinding.Graph
// 		expected [][]string
// 	}{
// 		{
// 			expected: [][]string{
// 				{"a", "b", "d", "f"},
// 			},
// 			getGraph: func() *pathfinding.Graph {

// 				graph := pathfinding.NewGraph()

// 				nodeA := pathfinding.NewSourceNode("a")
// 				nodeB := pathfinding.NewNode("b")
// 				nodeC := pathfinding.NewNode("c")
// 				nodeD := pathfinding.NewNode("d")
// 				nodeE := pathfinding.NewNode("e")
// 				nodeF := pathfinding.NewDestinationNode("f")
// 				//  A
// 				//  | \
// 				//  |  \
// 				//  |   \
// 				//  |    \
// 				//  1     2
// 				//  |     |
// 				// [B]   [C]
// 				//  |     |
// 				//  2     3
// 				//  |     |
// 				// [D]   [E]
// 				//  |     |
// 				//  6     5
// 				//  |    /
// 				//  |   /
// 				//  |  /
// 				//  | /
// 				// [F]

// 				graph.AddEdge(nodeA, nodeB, 1)
// 				graph.AddEdge(nodeA, nodeC, 2)
// 				graph.AddEdge(nodeB, nodeD, 2)
// 				graph.AddEdge(nodeC, nodeE, 3)
// 				graph.AddEdge(nodeD, nodeF, 6)
// 				graph.AddEdge(nodeE, nodeF, 5)

// 				return graph
// 			},
// 		},
// 		// {
// 		// 	expected: [][]string{
// 		// 		{"a", "b", "d", "f"},
// 		// 		{"a", "c", "e", "f"},
// 		// 	},
// 		// 	getEdges: func() []*pathfinding.Edge {

// 		// 		nodeA := pathfinding.NewSourceNode("a")
// 		// 		nodeB := pathfinding.NewNode("b")
// 		// 		nodeC := pathfinding.NewNode("c")
// 		// 		nodeD := pathfinding.NewNode("d")
// 		// 		nodeE := pathfinding.NewNode("e")
// 		// 		nodeF := pathfinding.NewDestinationNode("f")
// 		// 		//  A
// 		// 		//  | \
// 		// 		//  |  \
// 		// 		//  |   \
// 		// 		//  |    \
// 		// 		//  1     1
// 		// 		//  |     |
// 		// 		// [B]   [C]
// 		// 		//  |     |
// 		// 		//  2     2
// 		// 		//  |     |
// 		// 		// [D]   [E]
// 		// 		//  |     |
// 		// 		//  6     6
// 		// 		//  |    /
// 		// 		//  |   /
// 		// 		//  |  /
// 		// 		//  | /
// 		// 		// [F]

// 		// 		return []*pathfinding.Edge{
// 		// 			pathfinding.NewEdge(nodeA, nodeB, 1),
// 		// 			pathfinding.NewEdge(nodeA, nodeC, 1),
// 		// 			pathfinding.NewEdge(nodeB, nodeD, 2),
// 		// 			pathfinding.NewEdge(nodeC, nodeE, 2),
// 		// 			pathfinding.NewEdge(nodeD, nodeF, 6),
// 		// 			pathfinding.NewEdge(nodeE, nodeF, 6),
// 		// 		}
// 		// 	},
// 		// },
// 		// {
// 		// 	expected: [][]string{
// 		// 		{"a", "c", "e", "f"},
// 		// 	},
// 		// 	getEdges: func() []*pathfinding.Edge {

// 		// 		nodeA := pathfinding.NewSourceNode("a")
// 		// 		nodeB := pathfinding.NewNode("b")
// 		// 		nodeC := pathfinding.NewNode("c")
// 		// 		nodeD := pathfinding.NewNode("d")
// 		// 		nodeE := pathfinding.NewNode("e")
// 		// 		nodeF := pathfinding.NewDestinationNode("f")
// 		// 		//  A
// 		// 		//  | \
// 		// 		//  |  \
// 		// 		//  |   \
// 		// 		//  |    \
// 		// 		//  1     1
// 		// 		//  |     |
// 		// 		// [B]   [C]
// 		// 		//  |     |
// 		// 		//  2     2
// 		// 		//  |     |
// 		// 		// [D]   [E]
// 		// 		//  |     |
// 		// 		//  6     4
// 		// 		//  |    /
// 		// 		//  |   /
// 		// 		//  |  /
// 		// 		//  | /
// 		// 		// [F]

// 		// 		return []*pathfinding.Edge{
// 		// 			pathfinding.NewEdge(nodeA, nodeB, 1),
// 		// 			pathfinding.NewEdge(nodeB, nodeD, 2),
// 		// 			pathfinding.NewEdge(nodeD, nodeF, 6),

// 		// 			pathfinding.NewEdge(nodeA, nodeC, 1),
// 		// 			pathfinding.NewEdge(nodeC, nodeE, 2),
// 		// 			pathfinding.NewEdge(nodeE, nodeF, 4),
// 		// 		}
// 		// 	},
// 		// },
// 	}

// 	isMatch := func(got []*pathfinding.Node, possible []string) bool {
// 		count := 0
// 		for index, id := range possible {
// 			if got[index].ID == id {
// 				count++
// 			}
// 		}
// 		return count == len(possible)
// 	}

// 	for _, test := range tests {

// 		graph := test.getGraph()
// 		shortestPathSet, shortestPath, err := pathfinding.Dijkstra(graph)
// 		if err != nil {
// 			fmt.Println(err)
// 			t.Errorf("error is unexpected")
// 		}

// 		hasMatchingPath := false
// 		for _, possiblePath := range test.expected {
// 			isMatch := false
// 			for _, id := range possiblePath {
// 				if id =
// 			}
// 			// if isMatch(shortestPathSet, possiblePath) {
// 			// 	hasMatchingPath = true
// 			// 	break
// 			// }
// 		}

// 		if !hasMatchingPath {
// 			t.Errorf("path is invalid [%s] possible expected paths: %s", pathToString(shortestPath), test.expected)
// 		}

// 	}

// }

// func pathToString(path []*pathfinding.Node) string {
// 	s := ""
// 	for _, node := range path {
// 		s = fmt.Sprintf("%s %s", s, node.ID)
// 	}
// 	return s
// }
