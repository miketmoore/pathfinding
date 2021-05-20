# pathfinding

### Run tests

```
go test
```

### Generate test coverage report

```
go test -test.coverprofile=cover.out
go tool cover -html=cover.out -o coverage.html
```

## Dijkstra's Algorithm

Find the shortest path in a graph

### Run example

```
go cmd/dijkstra/dijkstra.go
```

### Example

#### Find shortest paths from source to all other nodes

```
import "github.com/miketmoore/pathfinding"

// Build graph by defining edges with a distance value for each one
graph := pathfinding.NewGraph()
graph.AddEdge("0", "1", 2)
graph.AddEdge("1", "3", 5)
graph.AddEdge("3", "2", 8)
graph.AddEdge("0", "2", 6)
graph.AddEdge("3", "5", 15)
graph.AddEdge("5", "6", 6)
graph.AddEdge("5", "4", 6)
graph.AddEdge("3", "4", 10)
graph.AddEdge("4", "6", 2)

// Run the algorithm
shortestPathGraph, nodeDistancesMap, err := pathfinding.DijkstraAllPaths(
    test.getGraph(),
    test.sourceNodeId,
)

if err != nil {
    fmt.Println(err)
    os.Exit(0)
}
```

### References

* https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Pseudocode
* https://www.freecodecamp.org/news/dijkstras-shortest-path-algorithm-visual-introduction/
* https://www.codingame.com/playgrounds/1608/shortest-paths-with-dijkstras-algorithm/dijkstras-algorithm