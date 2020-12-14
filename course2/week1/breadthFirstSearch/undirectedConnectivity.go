package main

func countUndirectedConnectivity(graph AdjacencyList) int {
	count := 0
	for i := 1; i < len(graph); i++ {
		if !graph[i].explored {
			breadthFirstSearch(graph, i)
			count++
		}
	}
	return count
}
