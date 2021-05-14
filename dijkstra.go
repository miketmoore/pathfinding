package pathfinding

import (
	"fmt"
	"math"
)

type Node struct {
	ID                string
	visited           bool
	tentativeDistance float64
	isDestination     bool
}

func NewNode(id string) *Node {
	return &Node{
		ID:                id,
		visited:           false,
		tentativeDistance: math.Inf(1),
		isDestination:     false,
	}
}

func NewDestinationNode(id string) *Node {
	node := NewNode(id)
	node.isDestination = true
	return node
}

type Edge struct {
	nodeA, nodeB *Node
	distance     float64
}

func NewEdge(nodeA, nodeB *Node, distance float64) *Edge {
	return &Edge{
		nodeA:    nodeA,
		nodeB:    nodeB,
		distance: distance,
	}
}

func buildUnvisitedNodesMapFromEdges(edges []*Edge) (initial *Node, nodesMap NodesMap) {
	unvisited := NodesMap{}
	for _, edge := range edges {
		if initial == nil {
			initial = edge.nodeA
		}
		unvisited[edge.nodeA.ID] = edge.nodeA
		unvisited[edge.nodeB.ID] = edge.nodeB
	}
	return initial, unvisited
}

func buildEdgesMapFromSlice(edges []*Edge) EdgesMap {
	edgesMap := EdgesMap{}
	for _, edge := range edges {
		edgeKey := buildEdgeKeyFromNodes(edge.nodeA, edge.nodeB)
		edgesMap[edgeKey] = edge
	}
	return edgesMap
}

func Dijkstra(edges []*Edge) NodesMap {
	initial, unvisited := buildUnvisitedNodesMapFromEdges(edges)
	edgesMap := buildEdgesMapFromSlice(edges)
	initial.tentativeDistance = 0

	path := NodesMap{initial.ID: initial}

	traverse(initial, unvisited, edgesMap, path)

	return path
}

// Let the node at which we are starting be called the initial node. Let the distance of node Y be the distance from the initial node to Y.
// Dijkstra's algorithm will assign some initial distance values and will try to improve them step by step.

// Mark all nodes unvisited. Create a set of all the unvisited nodes called the unvisited set.
// Assign to every node a tentative distance value: set it to zero for our initial node and to infinity for all other nodes. Set the initial node as current.[15]

// For the current node, consider all of its unvisited neighbours and calculate their tentative distances through the current node.
// Compare the newly calculated tentative distance to the current assigned value and assign the smaller one.
// For example, if the current node A is marked with a distance of 6, and the edge connecting it with a neighbour B has length 2,
// then the distance to B through A will be 6 + 2 = 8. If B was previously marked with a distance greater than 8 then change it to 8.
// Otherwise, the current value will be kept.

// When we are done considering all of the unvisited neighbours of the current node, mark the current node as visited and remove it from the unvisited set.
// A visited node will never be checked again.

// If the destination node has been marked visited (when planning a route between two specific nodes) or if the smallest tentative distance among the nodes
// in the unvisited set is infinity (when planning a complete traversal; occurs when there is no connection between the initial node and remaining unvisited nodes), then stop.
// The algorithm has finished.

// Otherwise, select the unvisited node that is marked with the smallest tentative distance, set it as the new "current node", and go back to step 3.
// When planning a route, it is actually not necessary to wait until the destination node is "visited" as above: the algorithm can stop once the destination node has the smallest tentative distance among all "unvisited" nodes (and thus could be selected as the next "current").

type NodesMap map[string]*Node
type EdgesMap map[string]*Edge

func findPossibleNeighborNodes(currentNode *Node, edges EdgesMap) NodesMap {
	possibleNeighbors := NodesMap{}

	for _, edge := range edges {
		if edge.nodeA == currentNode {
			possibleNeighbors[edge.nodeB.ID] = edge.nodeB
		} else if edge.nodeB == currentNode {
			possibleNeighbors[edge.nodeA.ID] = edge.nodeA
		}
	}

	return possibleNeighbors
}

// For the current node, consider all of its unvisited neighbours and calculate their tentative distances through the current node.
// Compare the newly calculated tentative distance to the current assigned value and assign the smaller one.
// For example, if the current node A is marked with a distance of 6, and the edge connecting it with a neighbour B has length 2,
// then the distance to B through A will be 6 + 2 = 8. If B was previously marked with a distance greater than 8 then change it to 8.
// Otherwise, the current value will be kept.
func traverse(currentNode *Node, unvisitedNodes NodesMap, edges EdgesMap, path NodesMap) {

	possibleNeighbors := findPossibleNeighborNodes(currentNode, edges)

	actualNeighbors := []*Node{}

	for _, possibleNeighbor := range possibleNeighbors {
		if !possibleNeighbor.visited {
			edgeKey := buildEdgeKeyFromNodes(currentNode, possibleNeighbor)
			edge, edgeOk := edges[edgeKey]
			if edgeOk {

				// this is an actual neighbor node
				actualNeighbors = append(actualNeighbors, possibleNeighbor)

				d := currentNode.tentativeDistance + edge.distance

				if possibleNeighbor.tentativeDistance > d {
					possibleNeighbor.tentativeDistance = d
				}
			}
		}
	}

	currentNode.visited = true

	if currentNode.isDestination {
		return
	}

	delete(unvisitedNodes, currentNode.ID)

	if isFinishEarlyCase(unvisitedNodes) {
		return
	}

	var nextNode *Node
	for _, node := range actualNeighbors {
		if nextNode == nil {
			nextNode = node
		} else if node.tentativeDistance < nextNode.tentativeDistance {
			nextNode = node
		}
	}

	if nextNode == nil {
		fmt.Println("no next node found")
		return
	}

	path[nextNode.ID] = nextNode

	if nextNode != nil {
		traverse(nextNode, unvisitedNodes, edges, path)
	}

}

func isFinishEarlyCase(unvisitedNodes NodesMap) bool {
	// is destination node in unvisitedNodes?
	// does it have the smallest tentative distance?
	// if true for both, we are finished
	smallestTentativeDistance := math.Inf(1)
	var destinationNode *Node
	for _, node := range unvisitedNodes {
		if node.tentativeDistance < smallestTentativeDistance {
			smallestTentativeDistance = node.tentativeDistance
		}
		if node.isDestination {
			destinationNode = node
		}
	}
	if destinationNode != nil && destinationNode.tentativeDistance < smallestTentativeDistance {
		fmt.Println("destination node is unvisited and has the smallest tentative distance among all unvisited nodes, so we can return early")
		return true
	}

	return false
}

func buildEdgeKeyFromNodes(nodeA, nodeB *Node) string {
	return fmt.Sprintf(
		"%s~%s",
		nodeA.ID,
		nodeB.ID,
	)
}

// func buildNodesAndEdges(grid *mazegen.Grid) (initialNode *Node, nodesMap NodesMap, edgesMap EdgesMap) {
// 	nodesMap = NodesMap{}
// 	edgesMap = EdgesMap{}

// 	cacheNode := func(node *Node) {
// 		key := buildNodeKey(node)
// 		_, ok := nodesMap[key]
// 		if !ok {
// 			nodesMap[key] = node
// 		}
// 	}

// 	cacheEdge := func(nodeA, nodeB *Node) {
// 		cacheNode(nodeA)
// 		cacheNode(nodeB)

// 		edge := newEdge(nodeA, nodeB, 1)
// 		// key := buildEdgeKey(edge)
// 		key := buildEdgeKeyFromNodes(nodeA, nodeB)
// 		_, ok := edgesMap[key]
// 		if !ok {
// 			edgesMap[key] = edge
// 		}
// 	}

// 	for y, row := range grid.Cells {
// 		for x, cell := range row {

// 			thisNode := newNode(y, x)

// 			if initialNode == nil {
// 				initialNode = thisNode
// 			}

// 			// build edges
// 			if cell.Walls[mazegen.North] {
// 				// build edge from this to north node
// 				if y > 0 {
// 					cacheEdge(thisNode, newNode(y-1, x))
// 				}
// 			}
// 			if cell.Walls[mazegen.East] {
// 				// build edge from this node to east node
// 				maxX := len(row) - 1
// 				if x < maxX {
// 					cacheEdge(thisNode, newNode(y, x+1))
// 				}
// 			}
// 			if cell.Walls[mazegen.South] {
// 				// build edge from this node to south node
// 				maxY := grid.Rows
// 				if y < maxY {
// 					cacheEdge(thisNode, newNode(y+1, x))
// 				}
// 			}
// 			if cell.Walls[mazegen.West] {
// 				// build edge from this node to west node
// 				if x > 0 {
// 					cacheEdge(thisNode, newNode(y, x-1))
// 				}
// 			}
// 		}
// 	}

// 	return initialNode, nodesMap, edgesMap
// }
