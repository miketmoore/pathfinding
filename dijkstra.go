package pathfinding

import (
	"fmt"
	"math"
)

func Dijkstra(graph *Graph, sourceNodeId string) (shortestPathGraph *Graph, err error) {

	nodeDistances := map[string]float64{}
	shortestPathGraph = NewGraph()

	// find the source node
	currentNode := graph.FindNodeById(sourceNodeId)
	if currentNode == nil {
		return nil, fmt.Errorf("source node not found by id=%s", sourceNodeId)
	}

	// start building our new "shortest path graph"
	shortestPathGraph.NewNode(sourceNodeId)

	// assign tentative distance values to each node
	// source will be zero
	// all others will be infinity
	nodeDistances[sourceNodeId] = 0
	for _, node := range graph.Nodes {
		if node.ID != sourceNodeId {
			nodeDistances[node.ID] = math.Inf(1)
		}
	}

	// while shortest path set does *not* contain all nodes
	for len(shortestPathGraph.Nodes) < len(graph.Nodes) {
		// get adjacent nodes
		adjacentNodes := FindAdjacentNodes(graph, currentNode.ID)
		fmt.Printf("found %d adjacent nodes to id=%s\n", len(adjacentNodes), currentNode.ID)

		// update adjacent node distances
		for nodeId, node := range adjacentNodes {
			edge, ok := graph.FindEdge(currentNode, node)
			if ok {
				d := nodeDistances[currentNode.ID] + edge.Distance
				nodeDistances[nodeId] = d
			}
		}

		// pick node with minimum distance value
		var minNode *Node
		for nodeId, node := range graph.Nodes {
			found := shortestPathGraph.FindNodeById(nodeId)
			if found == nil {
				if minNode == nil {
					minNode = node
				} else {
					if nodeDistances[nodeId] < nodeDistances[minNode.ID] {
						minNode = node
					}
				}
			}
		}
		currentNode = minNode
		shortestPathGraph.NewNode(minNode.ID)
	}

	return shortestPathGraph, nil

}

func FindAdjacentNodes(graph *Graph, nodeId string) NodesMap {
	node := graph.FindNodeById(nodeId)
	adjacent := NodesMap{}
	for _, edge := range graph.Edges {
		if node.ID == edge.NodeA.ID {
			adjacent[edge.NodeB.ID] = edge.NodeB
		} else if node.ID == edge.NodeB.ID {
			adjacent[edge.NodeA.ID] = edge.NodeA
		}
	}
	return adjacent
}

func findNodeNotInShortestPathSet(shortestPathSet NodesMap, allNodes NodesMap) *Node {
	for nodeId, node := range allNodes {
		_, ok := shortestPathSet[nodeId]
		if !ok {
			return node
		}
	}
	return nil
}
