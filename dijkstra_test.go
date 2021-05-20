package pathfinding_test

import (
	"fmt"
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

func TestDijkstraAllPaths(t *testing.T) {

	tests := []struct {
		getGraph              func() *pathfinding.Graph
		expectedNodeDistances map[string]float64
	}{
		{
			getGraph: func() *pathfinding.Graph {

				graph := pathfinding.NewGraph()

				graph.NewSourceNode("A")
				graph.AddEdge("A", "B", 3)
				graph.AddEdge("B", "E", 1)
				graph.AddEdge("E", "D", 7)
				graph.AddEdge("B", "D", 5)
				graph.AddEdge("C", "D", 2)
				graph.AddEdge("B", "C", 7)
				graph.AddEdge("A", "C", 1)

				return graph
			},
			expectedNodeDistances: map[string]float64{
				"A": 1,
				"B": 4,
				"C": 0,
				"D": 2,
				"E": 5,
			},
		},
	}

	for _, test := range tests {

		_, nodeDistances, err := pathfinding.DijkstraAllPaths(test.getGraph(), "C")
		if err != nil {
			fmt.Println(err)
			t.Error("an unexpected error was returned")
		}

		for nodeId, distance := range test.expectedNodeDistances {
			_, ok := nodeDistances[nodeId]
			if !ok {
				t.Errorf("node=%s is not in the map but should be", nodeId)
			}
			if nodeDistances[nodeId] != distance {
				t.Errorf("node=%s distance got=%f expected=%f", nodeId, nodeDistances[nodeId], distance)
			}
		}
	}

}
