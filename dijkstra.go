package pathfinding

import (
	"fmt"
	"math"
)

type Node struct {
	ID                               string
	TentativeDistance                float64
	visited, isSource, isDestination bool
}

func NewNode(id string) *Node {
	return &Node{
		ID:                id,
		visited:           false,
		TentativeDistance: math.Inf(1),
		isSource:          false,
		isDestination:     false,
	}
}

func NewSourceNode(id string) *Node {
	node := NewNode(id)
	node.isSource = true
	return node
}

func NewDestinationNode(id string) *Node {
	node := NewNode(id)
	node.isDestination = true
	return node
}

type Edge struct {
	NodeA, NodeB *Node
	Distance     float64
}

func (e Edge) ID() string {
	return fmt.Sprintf(
		"%s~%s",
		e.NodeA.ID,
		e.NodeB.ID,
	)
}

func NewEdge(nodeA, nodeB *Node, distance float64) *Edge {
	return &Edge{
		NodeA:    nodeA,
		NodeB:    nodeB,
		Distance: distance,
	}
}

// func buildEdgesMapFromSlice(edges []*Edge) EdgesMap {
// 	edgesMap := EdgesMap{}
// 	for _, edge := range edges {
// 		edgeKey := buildEdgeKeyFromNodes(edge.nodeA, edge.nodeB)
// 		edgesMap[edgeKey] = edge
// 	}
// 	return edgesMap
// }

type Graph struct {
	Nodes NodesMap
	Edges EdgesMap
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: NodesMap{},
		Edges: EdgesMap{},
	}
}

func (g *Graph) AddNode(node *Node) {
	_, ok := g.Nodes[node.ID]
	if !ok {
		g.Nodes[node.ID] = node
	}
}

func (g *Graph) AddEdge(nodeA, nodeB *Node, distance float64) {
	g.AddNode(nodeA)
	g.AddNode(nodeB)
	edge := NewEdge(nodeA, nodeB, distance)
	_, ok := g.Edges[edge.ID()]
	if !ok {
		g.Edges[edge.ID()] = edge
	}
}

func (g *Graph) FindEdge(u, v *Node) (*Edge, bool) {
	for _, edge := range g.Edges {
		if (edge.NodeA.ID == u.ID && edge.NodeB.ID == v.ID) ||
			(edge.NodeA.ID == v.ID && edge.NodeB.ID == u.ID) {
			return edge, true
		}
	}
	return nil, false
}

func (g *Graph) FindEdgesForNode(node *Node) EdgesMap {
	edgesMap := EdgesMap{}

	for edgeId, edge := range g.Edges {
		if edge.NodeA.ID == node.ID || edge.NodeB.ID == node.ID {
			edgesMap[edgeId] = edge
		}
	}

	return edgesMap
}

func (g *Graph) FindEdgesForNodes(nodesMap NodesMap) EdgesMap {
	edgesMap := EdgesMap{}

	for _, node := range nodesMap {
		edges := g.FindEdgesForNode(node)
		for edgeId, edge := range edges {
			_, ok := edgesMap[edgeId]
			if !ok {
				edgesMap[edgeId] = edge
			}
		}
	}

	return edgesMap
}

func (g *Graph) GraphVizString() string {
	gvStr := "graph shortestPath {\n"
	for _, edge := range g.Edges {
		gvStr = fmt.Sprintf("%s  %s -- %s [label=%.2f]\n", gvStr, edge.NodeA.ID, edge.NodeB.ID, edge.Distance)
	}

	for _, node := range g.Nodes {
		if node.isSource {
			gvStr = fmt.Sprintf("%s  %s [shape=diamond];\n", gvStr, node.ID)
		} else if node.isDestination {
			gvStr = fmt.Sprintf("%s  %s [shape=square];\n", gvStr, node.ID)
		}
	}

	gvStr = fmt.Sprintf("%s}", gvStr)
	return gvStr
}

func Dijkstra(graph *Graph) (shortestPathSet NodesMap, parent []*Node, err error) {

	// 1) Create a set sptSet (shortest path tree set) that keeps track of vertices included in shortest path tree, i.e.,
	// whose minimum distance from source is calculated and finalized. Initially, this set is empty.

	// Shortest path tree set
	// Keys are node IDs
	// Values are nodes
	// These nodes have a minimum distance from the source that is calculated and finalized
	shortestPathSet = NodesMap{}

	// 2) Assign a distance value to all vertices in the input graph. Initialize all distance values as INFINITE.
	// Assign distance value as 0 for the source vertex so that it is picked first.

	// This is already done during instantiation of each Node

	// 	3) While sptSet doesn’t include all vertices
	// ….a) Pick a vertex u which is not there in sptSet and has minimum distance value.
	// ….b) Include u to sptSet.
	// ….c) Update distance value of all adjacent vertices of u. To update the distance values, iterate through all adjacent vertices.
	// For every adjacent vertex v, if sum of distance value of u (from source) and weight of edge u-v, is less than the distance value of v,
	// then update the distance value of v.

	findNodeNotInShortestPathSet := func(shortestPathSet NodesMap, allNodes NodesMap) *Node {

		for nodeId, node := range allNodes {
			_, ok := shortestPathSet[nodeId]
			if !ok {
				return node
			}
		}
		return nil
	}

	findAdjacentNodes := func(node *Node) []*Node {
		adjacent := []*Node{}
		for _, edge := range graph.Edges {
			if node.ID == edge.NodeA.ID {
				adjacent = append(adjacent, edge.NodeB)
			} else if node.ID == edge.NodeB.ID {
				adjacent = append(adjacent, edge.NodeA)
			}
		}
		return adjacent
	}

	parent = []*Node{}

	// While the shortest path set does not contain all nodes
	for len(shortestPathSet) != len(graph.Nodes) {
		// Pick a vertex u which is not there in sptSet and has minimum distance value.
		u := findNodeNotInShortestPathSet(shortestPathSet, graph.Nodes)
		if u == nil {
			// TODO
			return shortestPathSet, parent, fmt.Errorf("node not found")
		}
		// Include u to sptSet.
		shortestPathSet[u.ID] = u

		// Update distance value of all adjacent vertices of u. To update the distance values, iterate through all adjacent vertices.
		adjacent := findAdjacentNodes(u)
		for _, v := range adjacent {
			// TODO update distance
			// get distance from edge between u and v
			edge, ok := graph.FindEdge(u, v)
			if !ok {
				// TODO
				return shortestPathSet, parent, fmt.Errorf("edge not found")
			}
			d := edge.Distance

			// For every adjacent vertex v, if sum of distance value of u (from source) and weight of edge u-v, is less than the distance value of v,
			// then update the distance value of v.
			// TODO
			v.TentativeDistance = d
		}

		// Pick the vertex with minimum distance value and not already included in SPT (not in sptSET).
		min := math.Inf(1)
		var last *Node
		for _, v := range adjacent {
			if v.TentativeDistance < min {
				last = v
			}
		}

		if last == nil {
			// TODO
			return shortestPathSet, parent, fmt.Errorf("last is nil")
		} else {
			parent = append(parent, last)
			shortestPathSet[last.ID] = last

			// Update the distance values of adjacent vertices of last
			// TODO ... just continue loop

		}

	}

	return shortestPathSet, parent, nil

}

type NodesMap map[string]*Node
type EdgesMap map[string]*Edge
