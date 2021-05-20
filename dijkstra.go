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
	// infinity for all other nodes. Set the initial node as current.
	tentativeNodeDistances := map[string]float64{}
	tentativeNodeDistances[sourceNodeId] = 0
	for _, node := range graph.Nodes {
		if node.ID != sourceNodeId {
			tentativeNodeDistances[node.ID] = math.Inf(1)
		}
	}

	debug := func() {
		fmt.Println("unvisited nodes", unvisitedNodes)
		fmt.Println("tentative distances cache", tentativeNodeDistances)
	}

	currentNodeId := sourceNodeId

	for len(unvisitedNodes) > 0 {
		fmt.Printf("LOOP currentNodeId=%s\n", currentNodeId)
		// 3. For the current node, consider all of its unvisited neighbours and calculate their tentative
		// distances through the current node. Compare the newly calculated tentative distance to the current
		// assigned value and assign the smaller one. For example, if the current node A is marked with a
		// distance of 6, and the edge connecting it with a neighbour B has length 2, then the distance to B
		// through A will be 6 + 2 = 8. If B was previously marked with a distance greater than 8 then change
		// it to 8. Otherwise, the current value will be kept.
		unvisitedNeighbors := FindUnvisitedNeighbors(graph, currentNodeId, unvisitedNodes)
		fmt.Printf(" unvisited neighbors of current node id=%s\n", currentNodeId)
		fmt.Println(" ", unvisitedNeighbors)

		for neighborId := range unvisitedNeighbors {
			fmt.Printf(" finding edge current=%s neighbor=%s\n", currentNodeId, neighborId)
			edge, edgeOk := graph.FindEdgeByNodeIds(currentNodeId, neighborId)
			if !edgeOk {
				fmt.Printf(" neighbor id=%s found but no edge\n", neighborId)
				// return shortestPathGraph, fmt.Errorf("neighbor id=%s found but no edge", neighborId)
			} else {
				fmt.Printf(" neighbor id=%s and edge found\n", neighborId)
				d := tentativeNodeDistances[currentNodeId] + edge.Distance
				fmt.Printf(" tentative distance from current=%s to neighbor=%s %.2f\n", currentNodeId, neighborId, d)
				if tentativeNodeDistances[neighborId] > d {
					fmt.Printf(" existing tentative distance %.2f is greater than new %.2f\n", tentativeNodeDistances[neighborId], d)
					tentativeNodeDistances[neighborId] = d
				}
			}
			fmt.Println("---")
		}

		// 4. When we are done considering all of the unvisited neighbours of the current node, mark the current
		// node as visited and remove it from the unvisited set. A visited node will never be checked again.
		delete(unvisitedNodes, currentNodeId)

		// 5. If the destination node has been marked visited (when planning a route between two specific nodes)
		// or if the smallest tentative distance among the nodes in the unvisited set is infinity (when planning
		// a complete traversal; occurs when there is no connection between the initial node and remaining
		// unvisited nodes), then stop. The algorithm has finished.
		_, ok := unvisitedNodes[destinationNodeId]
		if !ok {
			fmt.Println("the destination node has been visited, ending early")
			debug()
			return shortestPathGraph, nil
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
			fmt.Println("the smallest tentative distance in the unvisited set is infinity, ending early")
			debug()
			return shortestPathGraph, nil
		}

		// 6. Otherwise, select the unvisited node that is marked with the smallest tentative distance, set it
		// as the new "current node", and go back to step 3.
		currentNodeId = *nodeWithSmallestTd
	}

	debug()

	return shortestPathGraph, nil

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
			fmt.Printf(">>> caching for source=%s unvisited neighbor id=%s\n", sourceNodeId, nodeId)
			unvisitedNeighbors[nodeId] = true
		}
	}
	return unvisitedNeighbors
}

// CalculateTentativeDistance updates the tentative node distance map for the
// nodes u and v
func CalculateTentativeDistance(
	graph *Graph,
	tentativeNodeDistances map[string]float64,
	nodeAId string,
	nodeBId string,
) {
	nodeADist := tentativeNodeDistances[nodeAId]
	edge, _ := graph.FindEdgeByNodeIds(nodeAId, nodeBId)
	edgeDist := edge.Distance
	td := nodeADist + edgeDist
	prevDist, prevDistOk := tentativeNodeDistances[nodeBId]
	if prevDistOk && prevDist > td {
		tentativeNodeDistances[nodeBId] = td
	}
}

// SelectUnvisitedNodeWithSmallestTentativeDistance returns a node ID if
// there are still unvisited nodes. The node ID returned is the one with
// the minimum tentative distance among the remaining unvisited nodes.
// If no node is found, a nil pointer is returned.
func SelectUnvisitedNodeWithSmallestTentativeDistance(
	unvisitedNodes map[string]bool,
	tentativeNodeDistances map[string]float64,
) (string, bool) {
	var lastId *string
	lastDistance := math.Inf(1)

	for nodeId := range unvisitedNodes {
		if lastId == nil {
			// Need to copy string otherwise value pointed by lastId
			// variable will change on every iteration of the loop
			copy := nodeId
			lastId = &copy
			lastDistance = tentativeNodeDistances[nodeId]
		} else {
			td, tdOk := tentativeNodeDistances[nodeId]
			if tdOk {
				if td < lastDistance {
					lastDistance = td
					lastId = &nodeId
				}
			}
		}
	}

	if lastId == nil {
		return "", false
	}
	return *lastId, true
}
