package pathfinding_test

import (
	"testing"

	"github.com/miketmoore/pathfinding"
)

func TestFindUnvisitedNeighbors(t *testing.T) {

	tests := []struct {
		getGraph                   func() *pathfinding.Graph
		unvisitedNodes             map[string]bool
		sourceNodeId               string
		expectedUnvisitedNeighbors map[string]bool
	}{
		{
			getGraph: func() *pathfinding.Graph {
				graph := pathfinding.NewGraph()
				graph.AddEdge("a", "b", 0)
				return graph
			},
			unvisitedNodes: map[string]bool{
				"b": true,
			},
			sourceNodeId: "a",
			expectedUnvisitedNeighbors: map[string]bool{
				"b": true,
			},
		},
		{
			getGraph: func() *pathfinding.Graph {
				graph := pathfinding.NewGraph()
				graph.AddEdge("a", "b", 0)
				graph.AddEdge("b", "c", 0)
				return graph
			},
			unvisitedNodes: map[string]bool{
				"b": true,
				"c": true,
			},
			sourceNodeId: "a",
			expectedUnvisitedNeighbors: map[string]bool{
				"b": true,
			},
		},
		{
			getGraph: func() *pathfinding.Graph {
				graph := pathfinding.NewGraph()
				graph.AddEdge("a", "b", 0)
				graph.AddEdge("b", "c", 0)
				graph.AddEdge("a", "c", 0)
				return graph
			},
			unvisitedNodes: map[string]bool{
				"b": true,
				"c": true,
			},
			sourceNodeId: "a",
			expectedUnvisitedNeighbors: map[string]bool{
				"b": true,
				"c": true,
			},
		},
	}

	for _, test := range tests {
		unvisitedNeighbors := pathfinding.FindUnvisitedNeighbors(
			test.getGraph(), test.sourceNodeId, test.unvisitedNodes)
		gotLength := len(unvisitedNeighbors)
		expectedLength := len(test.expectedUnvisitedNeighbors)
		if expectedLength != gotLength {
			t.Errorf("expected=%d got=%d", expectedLength, gotLength)
		}
		for key := range test.expectedUnvisitedNeighbors {
			_, ok := unvisitedNeighbors[key]
			if !ok {
				t.Errorf("unvisitedNeighbors does not contain key=%s", key)
			}
		}
	}

}
