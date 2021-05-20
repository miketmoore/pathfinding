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

### References

* https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Pseudocode
* https://www.freecodecamp.org/news/dijkstras-shortest-path-algorithm-visual-introduction/
* https://www.codingame.com/playgrounds/1608/shortest-paths-with-dijkstras-algorithm/dijkstras-algorithm