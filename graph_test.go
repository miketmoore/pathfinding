package pathfinding_test

import (
	"math"
	"testing"

	"github.com/miketmoore/pathfinding"
)

func TestGraph(t *testing.T) {
	graph := pathfinding.NewGraph()
	if len(graph.Edges) != 0 {
		t.Errorf("error")
	}
	if len(graph.Nodes) != 0 {
		t.Errorf("error")
	}
}

func TestNewNode(t *testing.T) {
	graph := pathfinding.NewGraph()

	node := graph.NewNode("a")

	if node == nil {
		t.Errorf("node is nil but should not be")
	}

	if node.ID != "a" {
		t.Errorf("node ID is unexpected")
	}

	if node.TentativeDistance != math.Inf(1) {
		t.Errorf("node TentativeDistance is unexpected")
	}

	found, ok := graph.Nodes["a"]
	if found == nil {
		t.Errorf("node not found but should exist")
	}
	if ok == false {
		t.Errorf("node was found so ok should be true")
	}
}

// func TestNewEdge(t *testing.T) {
// 	graph := pathfinding.NewGraph()

// 	nodeA := graph.NewNode("a")
// 	nodeB := graph.NewNode("b")

// 	edge := graph.NewEdge(nodeA, nodeB, 123)

// 	if edge == nil {
// 		t.Errorf("edge is nil but should not be")
// 	}

// 	if edge.NodeA == nil {
// 		t.Errorf("NodeA is nil but should not be")
// 	}

// 	if edge.NodeB == nil {
// 		t.Errorf("NodeB is nil but should not be")
// 	}

// 	if edge.Distance != 123 {
// 		t.Errorf("edge distance is unexpected")
// 	}
// }

// func TestEdgeId(t *testing.T) {
// 	graph := pathfinding.NewGraph()

// 	nodeA := graph.NewNode("a")
// 	nodeB := graph.NewNode("b")

// 	edge := graph.NewEdge(nodeA, nodeB, 123)

// 	if edge.Id() != "a~b" {
// 		t.Errorf("edge id is unexpected")
// 	}
// }

func TestFindNodebyIdNoneFound(t *testing.T) {
	graph := pathfinding.NewGraph()

	node := graph.FindNodeById("a")
	if node != nil {
		t.Errorf("node found but should not be found")
	}
}

func TestFindNodebyIdIsFound(t *testing.T) {
	graph := pathfinding.NewGraph()

	graph.NewNode("a")

	node := graph.FindNodeById("a")
	if node == nil {
		t.Errorf("node not found but should be found")
	}
}

func TestAddEdge(t *testing.T) {
	graph := pathfinding.NewGraph()

	graph.AddEdge("a", "b", 234)

	nodeA := graph.Nodes["a"]
	if nodeA == nil || nodeA.ID != "a" {
		t.Errorf("node a not found but should be found")
	}

	nodeB := graph.Nodes["b"]
	if nodeB == nil || nodeB.ID != "b" {
		t.Errorf("node b not found but should be found")
	}

	edge := graph.Edges["a~b"]

	if edge == nil {
		t.Errorf("edge not found but should be found")
	}

}

func TestFindEdgeNotFound(t *testing.T) {
	graph := pathfinding.NewGraph()

	nodeA := graph.Nodes["a"]
	nodeB := graph.Nodes["b"]

	edge, ok := graph.FindEdge(nodeA, nodeB)
	if ok {
		t.Errorf("ok should be false")
	}
	if edge != nil {
		t.Errorf("edge is not nil but should be")
	}
}

func TestFindEdgeIsFound(t *testing.T) {
	graph := pathfinding.NewGraph()
	graph.AddEdge("a", "b", 234)

	nodeA := graph.Nodes["a"]
	nodeB := graph.Nodes["b"]

	edge, ok := graph.FindEdge(nodeA, nodeB)
	if !ok {
		t.Errorf("ok should be true")
	}
	if edge == nil || edge.Id() != "a~b" {
		t.Errorf("edge not found but should be found")
	}
}

func TestFindEdgesForNodeInputIsNil(t *testing.T) {
	graph := pathfinding.NewGraph()

	edgesMap := graph.FindEdgesForNode(nil)

	if len(edgesMap) != 0 {
		t.Errorf("map length should be zero")
	}
}

func TestFindEdgesForNodeNoneFound(t *testing.T) {
	graph := pathfinding.NewGraph()

	node := graph.NewNode("a")

	edgesMap := graph.FindEdgesForNode(node)

	if len(edgesMap) != 0 {
		t.Errorf("map length should be zero")
	}
}

func TestFindEdgesForNodeSomeFound(t *testing.T) {
	graph := pathfinding.NewGraph()

	graph.AddEdge("a", "b", 1)
	graph.AddEdge("a", "c", 2)

	assert := func(nodeId string, expectedEdgeIds []string) {
		input := pathfinding.NodesMap{
			nodeId: graph.FindNodeById(nodeId),
		}
		got := graph.FindEdgesForNodes(input)
		if len(got) != len(expectedEdgeIds) {
			t.Errorf("response length is unexpected")
		}
		for _, edgeId := range expectedEdgeIds {
			_, ok := got[edgeId]
			if !ok {
				t.Error("edge is not in results")
			}
		}
	}

	assert("a", []string{"a~b", "a~c"})
	assert("b", []string{"a~b"})
	assert("c", []string{"a~c"})

}

func TestFindEdgesForNodes(t *testing.T) {
	graph := pathfinding.NewGraph()

	graph.AddEdge("a", "b", 1)
	graph.AddEdge("a", "c", 2)
	graph.AddEdge("d", "e", 3)

	assert := func(nodeIds []string, expectedEdgeIds []string) {
		input := pathfinding.NodesMap{}
		for _, id := range nodeIds {
			input[id] = graph.FindNodeById(id)
		}
		got := graph.FindEdgesForNodes(input)
		if len(got) != len(expectedEdgeIds) {
			t.Errorf("response length is unexpected")
		}
		for _, edgeId := range expectedEdgeIds {
			_, ok := got[edgeId]
			if !ok {
				t.Error("edge is not in results")
			}
		}
	}

	assert([]string{"a"}, []string{"a~b", "a~c"})
	assert([]string{"a", "b"}, []string{"a~b", "a~c"})
	assert([]string{"a", "b", "c"}, []string{"a~b", "a~c"})
	assert([]string{"d"}, []string{"d~e"})
	assert([]string{"e"}, []string{"d~e"})

}
