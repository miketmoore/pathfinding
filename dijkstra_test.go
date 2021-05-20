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

func TestCalculateTentativeDistance(t *testing.T) {
	tests := []struct {
		getGraph                                               func() *pathfinding.Graph
		tentativeNodeDistances, expectedTentativeNodeDistances map[string]float64
		nodeAId, nodeBId                                       string
	}{
		{
			getGraph: func() *pathfinding.Graph {
				graph := pathfinding.NewGraph()
				graph.AddEdge("a", "b", 2)
				return graph
			},
			tentativeNodeDistances: map[string]float64{
				"a": 10,
				"b": 13,
			},
			nodeAId: "a",
			nodeBId: "b",
			expectedTentativeNodeDistances: map[string]float64{
				"a": 10,
				"b": 12,
			},
		},
	}

	for _, test := range tests {
		pathfinding.CalculateTentativeDistance(
			test.getGraph(),
			test.tentativeNodeDistances,
			test.nodeAId,
			test.nodeBId,
		)
		expectedLength := len(test.expectedTentativeNodeDistances)
		gotLength := len(test.tentativeNodeDistances)
		if expectedLength != gotLength {
			t.Errorf("length is unexpected got=%d expected=%d", gotLength, expectedLength)
		}
		for key := range test.expectedTentativeNodeDistances {
			_, ok := test.tentativeNodeDistances[key]
			if !ok {
				t.Errorf("tentative distance for node id=%s not found", key)
			}
		}
	}

}
