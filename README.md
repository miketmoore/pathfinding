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