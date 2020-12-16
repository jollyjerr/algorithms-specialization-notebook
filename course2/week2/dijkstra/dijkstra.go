package main

// Edge is an edge between vertices
type Edge struct {
	// The tail is whatever node is storing the edge
	head   int
	weight int
}

// Node is a vertex
type Node struct {
	vertex            int
	edges             []Edge
	shortestPathFromS int // should default to 1000000
}

// Graph is an adjacency list representation of a graph
type Graph map[int]Node

func dijkstra(graph Graph) {

}
