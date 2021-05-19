package pathfinding

import (
	"fmt"
	"math"
)

func Dijkstra(graph *Graph, sourceNodeId string, destinationNodeId string) (shortestPathGraph *Graph, err error) {

	shortestPathGraph = NewGraph()

	// 1. Mark all nodes unvisited. Create a set of all the unvisited nodes called the unvisited set.
	unvisitedNodes := map[string]bool{}
	for _, node := range graph.Nodes {
		unvisitedNodes[node.ID] = true
	}

	// 2. Assign to every node a tentative distance value: set it to zero for our initial node and to
	// infinity for all other nodes. Set the initial node as current.[15]
	tentativeNodeDistances := map[string]float64{}
	tentativeNodeDistances[sourceNodeId] = 0
	for _, node := range graph.Nodes {
		if node.ID != sourceNodeId {
			tentativeNodeDistances[node.ID] = math.Inf(1)
		}
	}

	// 3. For the current node, consider all of its unvisited neighbours and calculate their tentative
	// distances through the current node. Compare the newly calculated tentative distance to the current
	// assigned value and assign the smaller one. For example, if the current node A is marked with a
	// distance of 6, and the edge connecting it with a neighbour B has length 2, then the distance to B
	// through A will be 6 + 2 = 8. If B was previously marked with a distance greater than 8 then change
	// it to 8. Otherwise, the current value will be kept.
	dothings(graph, sourceNodeId, destinationNodeId, unvisitedNodes, tentativeNodeDistances)

	fmt.Println(unvisitedNodes)

	return shortestPathGraph, nil

}

func dothings(
	graph *Graph,
	sourceNodeId string,
	destinationNodeId string,
	unvisitedNodes map[string]bool,
	tentativeNodeDistances map[string]float64,
) {
	// For the current node, consider all of its unvisited neighbours
	unvisitedNeighbors := FindUnvisitedNeighbors(graph, sourceNodeId, unvisitedNodes)

	// calculate their tentative distances through the current node
	for nodeId := range unvisitedNeighbors {
		CalculateTentativeDistance(graph, tentativeNodeDistances, sourceNodeId, nodeId)
	}

	// When we are done considering all of the unvisited neighbours of the current node, mark the current
	// node as visited and remove it from the unvisited set. A visited node will never be checked again.
	fmt.Printf("deleting node id=%s from unvisited set\n", sourceNodeId)
	delete(unvisitedNodes, sourceNodeId)

	// If the destination node has been marked visited (when planning a route between two specific nodes)
	// or if the smallest tentative distance among the nodes in the unvisited set is infinity (when planning
	// a complete traversal; occurs when there is no connection between the initial node and remaining
	// unvisited nodes), then stop. The algorithm has finished.

	_, ok := unvisitedNodes[destinationNodeId]

	if !ok {
		fmt.Println("destination node is no longer in the unvisited set")
		return
	}

	// Otherwise, select the unvisited node that is marked with the smallest tentative distance, set it as
	// the new "current node", and go back to step 3.
	lastId := SelectUnvisitedNodeWithSmallestTentativeDistance(unvisitedNodes, tentativeNodeDistances)

	if lastId == nil {
		fmt.Println("lastId is nil")
		return
	}

	dothings(graph, *lastId, destinationNodeId, unvisitedNodes, tentativeNodeDistances)
}

func FindUnvisitedNeighbors(
	graph *Graph,
	sourceNodeId string,
	unvisitedNodes map[string]bool,
) (unvisitedNeighbors map[string]bool) {
	unvisitedNeighbors = map[string]bool{}
	for nodeId := range unvisitedNodes {
		// is this a neighbor? yes if in edge with source node
		_, ok := graph.FindEdgeByNodeIds(sourceNodeId, nodeId)
		if ok {
			unvisitedNeighbors[nodeId] = true
		}
	}
	return unvisitedNeighbors
}

func CalculateTentativeDistance(
	graph *Graph,
	tentativeNodeDistances map[string]float64,
	u string,
	v string,
) {
	currNodeDist := tentativeNodeDistances[u]
	edge, _ := graph.FindEdgeByNodeIds(u, v)
	edgeDist := edge.Distance
	td := currNodeDist + edgeDist
	prevDist, prevDistOk := tentativeNodeDistances[v]
	if prevDistOk && prevDist > td {
		tentativeNodeDistances[v] = td
	}
}

func SelectUnvisitedNodeWithSmallestTentativeDistance(
	unvisitedNodes map[string]bool,
	tentativeNodeDistances map[string]float64,
) *string {
	var lastId *string
	var lastTentativeDistance float64
	for nodeId := range unvisitedNodes {
		if lastId == nil {
			lastId = &nodeId
			lastTentativeDistance = tentativeNodeDistances[nodeId]
		} else if tentativeNodeDistances[nodeId] < lastTentativeDistance {
			lastId = &nodeId
			lastTentativeDistance = tentativeNodeDistances[nodeId]
		}
	}
	return lastId
}
