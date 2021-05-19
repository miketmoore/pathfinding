package pathfinding

import (
	"fmt"
	"math"
)

func Dijkstra(graph *Graph) (shortestPathSet NodesMap, err error) {

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

	// While the shortest path set does not contain all nodes
	for len(shortestPathSet) != len(graph.Nodes) {
		// Pick a vertex u which is not there in sptSet and has minimum distance value.
		u := findNodeNotInShortestPathSet(shortestPathSet, graph.Nodes)

		if u == nil {
			// TODO
			return shortestPathSet, fmt.Errorf("node not found")
		}
		// Include u to sptSet.
		shortestPathSet[u.ID] = u

		// Update distance value of all adjacent vertices of u. To update the distance values, iterate through all adjacent vertices.
		adjacent := FindAdjacentNodes(graph, u.ID)
		for _, v := range adjacent {
			// TODO update distance
			// get distance from edge between u and v
			edge, ok := graph.FindEdge(u, v)
			if !ok {
				// TODO
				return shortestPathSet, fmt.Errorf("edge not found")
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
			return shortestPathSet, fmt.Errorf("last is nil")
		} else {
			shortestPathSet[last.ID] = last

			// Update the distance values of adjacent vertices of last
			// TODO ... just continue loop

		}

	}

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
