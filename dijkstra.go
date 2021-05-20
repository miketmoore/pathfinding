package pathfinding

import (
	"fmt"
	"math"
)

// DijkstraAllPaths finds the shortest path from the source node to all other nodes
func DijkstraAllPaths(
	graph *Graph,
	sourceNodeId string,
) (shortestPathGraph *Graph, nodeDistances map[string]float64, err error) {
	return dijkstra(graph, sourceNodeId, "", false)
}

// DijkstraDestination finds the shortest path from the source to the destination node
func DijkstraDestination(
	graph *Graph,
	sourceNodeId,
	destinationNodeId string,
) (shortestPathGraph *Graph, nodeDistances map[string]float64, err error) {
	return dijkstra(graph, sourceNodeId, destinationNodeId, true)
}

func dijkstra(
	graph *Graph,
	sourceNodeId string,
	destinationNodeId string,
	stopAtDestination bool,
) (shortestPathGraph *Graph, nodeDistances map[string]float64, err error) {

	shortestPathGraph = NewGraph()

	// 1. Mark all nodes unvisited. Create a set of all the unvisited nodes called the unvisited set.
	unvisitedNodes := map[string]bool{}
	for _, node := range graph.Nodes {
		unvisitedNodes[node.ID] = true
	}

	visitedEdges := map[string]*Edge{}

	// 2. Assign to every node a tentative distance value: set it to zero for our initial node and to
	// infinity for all other nodes. Set the initial node as current.
	tentativeNodeDistances := map[string]float64{}
	tentativeNodeDistances[sourceNodeId] = 0
	for _, node := range graph.Nodes {
		if node.ID != sourceNodeId {
			tentativeNodeDistances[node.ID] = math.Inf(1)
		}
	}

	currentNodeId := sourceNodeId

	for len(unvisitedNodes) > 0 {
		// 3. For the current node, consider all of its unvisited neighbours and calculate their tentative
		// distances through the current node. Compare the newly calculated tentative distance to the current
		// assigned value and assign the smaller one. For example, if the current node A is marked with a
		// distance of 6, and the edge connecting it with a neighbour B has length 2, then the distance to B
		// through A will be 6 + 2 = 8. If B was previously marked with a distance greater than 8 then change
		// it to 8. Otherwise, the current value will be kept.
		unvisitedNeighbors := FindUnvisitedNeighbors(graph, currentNodeId, unvisitedNodes)

		for neighborId := range unvisitedNeighbors {
			edge, edgeOk := graph.FindEdgeByNodeIds(currentNodeId, neighborId)
			if !edgeOk {
				return shortestPathGraph, tentativeNodeDistances, fmt.Errorf("neighbor id=%s found but no edge", neighborId)
			} else {
				visitedEdges[edge.Id()] = edge
				d := tentativeNodeDistances[currentNodeId] + edge.Distance
				if tentativeNodeDistances[neighborId] > d {
					tentativeNodeDistances[neighborId] = d
				}
			}
		}

		// 4. When we are done considering all of the unvisited neighbours of the current node, mark the current
		// node as visited and remove it from the unvisited set. A visited node will never be checked again.
		delete(unvisitedNodes, currentNodeId)

		// 5. If the destination node has been marked visited (when planning a route between two specific nodes)
		// or if the smallest tentative distance among the nodes in the unvisited set is infinity (when planning
		// a complete traversal; occurs when there is no connection between the initial node and remaining
		// unvisited nodes), then stop. The algorithm has finished.
		if stopAtDestination {
			_, ok := unvisitedNodes[destinationNodeId]
			if !ok {
				// the destination node has been visited, ending early
				return shortestPathGraph, tentativeNodeDistances, nil
			}
		}

		smallestTentativeDistance := math.Inf(1)
		var nodeWithSmallestTd *string
		for nodeId := range unvisitedNodes {
			if tentativeNodeDistances[nodeId] < smallestTentativeDistance {
				smallestTentativeDistance = tentativeNodeDistances[nodeId]
				copy := nodeId
				nodeWithSmallestTd = &copy
			}
		}
		if smallestTentativeDistance == math.Inf(1) {
			// the smallest tentative distance in the unvisited set is infinity, ending early
			return shortestPathGraph, tentativeNodeDistances, nil
		}

		// 6. Otherwise, select the unvisited node that is marked with the smallest tentative distance, set it
		// as the new "current node", and go back to step 3.
		currentNodeId = *nodeWithSmallestTd
	}

	return shortestPathGraph, tentativeNodeDistances, nil

}

// FindUnvisitedNeighbors returns a set of node IDs where each node is a neighbor
// of the source node ID and none of the nodes in the set have been visisted
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
