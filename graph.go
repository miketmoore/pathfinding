package pathfinding

import (
	"fmt"
	"math"
)

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

type Node struct {
	ID                               string
	TentativeDistance                float64
	visited, isSource, isDestination bool
}

func (g *Graph) NewNode(id string) *Node {
	node := &Node{
		ID:                id,
		visited:           false,
		TentativeDistance: math.Inf(1),
		isSource:          false,
		isDestination:     false,
	}
	g.Nodes[id] = node
	return node
}

func (g *Graph) NewSourceNode(id string) *Node {
	node := g.NewNode(id)
	node.isSource = true
	return node
}

func (g *Graph) NewDestinationNode(id string) *Node {
	node := g.NewNode(id)
	node.isDestination = true
	return node
}

type Edge struct {
	NodeA, NodeB *Node
	Distance     float64
}

func (e Edge) Id() string {
	return fmt.Sprintf(
		"%s~%s",
		e.NodeA.ID,
		e.NodeB.ID,
	)
}

func (g *Graph) FindNodeById(id string) *Node {
	for nodeId, node := range g.Nodes {
		if nodeId == id {
			return node
		}
	}
	return nil
}

func (g *Graph) AddEdge(nodeAId, nodeBId string, distance float64) {
	var nodeA *Node
	var nodeB *Node

	nodeA = g.FindNodeById(nodeAId)
	if nodeA == nil {
		nodeA = g.NewNode(nodeAId)
	}
	nodeB = g.FindNodeById(nodeBId)
	if nodeB == nil {
		nodeB = g.NewNode(nodeBId)
	}
	edge := &Edge{
		NodeA:    nodeA,
		NodeB:    nodeB,
		Distance: distance,
	}
	_, ok := g.Edges[edge.Id()]
	if !ok {
		g.Edges[edge.Id()] = edge
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

func (g *Graph) GraphVizString(name string) string {
	gvStr := fmt.Sprintf("graph %s {\n", name)
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

func (g *Graph) FindSourceNodes() NodesMap {
	nodesMap := NodesMap{}
	for _, node := range g.Nodes {
		if node.isSource {
			nodesMap[node.ID] = node
		}
	}
	return nodesMap
}
