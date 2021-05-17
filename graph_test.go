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

	nodeA := graph.FindNodeById("a")
	result0 := graph.FindEdgesForNode(nodeA)
	if len(result0) != 2 {
		t.Errorf("should have found 2 edges but found=%d", len(result0))
	}

	if result0["a~b"] == nil {
		t.Errorf("edge not found")
	}
	if result0["a~c"] == nil {
		t.Errorf("edge not found")
	}

	nodeB := graph.FindNodeById("b")
	result1 := graph.FindEdgesForNode(nodeB)
	if len(result1) != 1 {
		t.Errorf("should have found 1 edges but found=%d", len(result1))
	}

	if result0["a~b"] == nil {
		t.Errorf("edge not found")
	}

	nodeC := graph.FindNodeById("c")
	result2 := graph.FindEdgesForNode(nodeC)
	if len(result2) != 1 {
		t.Errorf("should have found 1 edges but found=%d", len(result2))
	}
}

// func TestFindEdgesForNodes(t *testing.T) {
// 	graph := pathfinding.NewGraph()

// 	graph.AddEdge("a", "b", 1)
// 	graph.AddEdge("a", "c", 2)

// 	nodeA := graph.FindNodeById("a")
// 	nodeB := graph.FindNodeById("b")
// 	nodeC := graph.FindNodeById("c")

// 	result0 := graph.FindEdgesForNodes(pathfinding.NodesMap{
// 		"a": nodeA,
// 	})

// }
