package main

// map can be used to represent a graph using the fact that the value of a map
// could be a map or struct. In this following example, a string key is mapped to a
// map containing strings, thus bearing a graph representation
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
