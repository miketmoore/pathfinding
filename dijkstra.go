package pathfinding

import (
	"fmt"
	"math"
)

func Dijkstra(graph *Graph) (shortestPathSet NodesMap, err error) {

	nodeDistances := map[string]float64{}
	shortestPathSet = NodesMap{}

	// assign tentative distance values to each node
	// one will be zero
	// all others will be infinity
	// the node with zero will be the initial node
	var initialNode *Node
	for _, node := range graph.Nodes {
		if initialNode == nil {
			nodeDistances[node.ID] = 0
			initialNode = node
		} else {
			nodeDistances[node.ID] = math.Inf(1)
		}
	}
	fmt.Printf("initial node id=%s\n", initialNode.ID)

	return shortestPathSet, nil

}

func FindAdjacentNodes(graph *Graph, nodeId string) []*Node {
	node := graph.FindNodeById(nodeId)
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

func findNodeNotInShortestPathSet(shortestPathSet NodesMap, allNodes NodesMap) *Node {
	for nodeId, node := range allNodes {
		_, ok := shortestPathSet[nodeId]
		if !ok {
			return node
		}
	}
	return nil
}
