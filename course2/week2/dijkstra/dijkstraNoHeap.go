package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
	visited           bool
}

func (node *Node) edgeHeads() []int {
	edgeheads := make([]int, 0, len(node.edges))
	for _, edge := range node.edges {
		edgeheads = append(edgeheads, edge.head)
	}
	return edgeheads
}

func (n *Node) getEdgeWeight(node Node) int {
	for _, edge := range n.edges {
		if edge.head == node.vertex {
			return edge.weight
		}
	}
	return -1
}

// Graph is an adjacency list representation of a graph
type Graph map[int]Node

// get graph keys in order
func (graph Graph) keys() []int {
	keys := make([]int, len(graph))
	count := 0
	for i := range graph {
		keys[count] = i
		count++
	}
	sort.Ints(keys)
	return keys
}

func dijkstra(graph Graph) Graph {
	graph[0] = setDistance(graph[0], 0)

	visited := make(Tvisited, 0, len(graph))
	visited = append(visited, graph[0])

	for len(visited) < len(graph) {
		for _, suspect := range visited {
			for _, edge := range suspect.edges {
				if nodeToAdd := graph[edge.head]; !visited.isNodeVisited(nodeToAdd) {
					score := visited.computeGreedyScore(nodeToAdd)
					graph[nodeToAdd.vertex] = setDistance(nodeToAdd, score)
					visited = append(visited, graph[nodeToAdd.vertex])
				}
			}
		}
	}

	return graph
}

type Tvisited []Node

func (visited *Tvisited) isNodeVisited(node Node) bool {
	for _, v := range *visited {
		if v.vertex == node.vertex {
			return true
		}
	}
	return false
}

func (visited *Tvisited) computeGreedyScore(node Node) int {
	greedyScore := 10000000
	suspects := make(Tvisited, 0, len(*visited))
	for _, suspect := range *visited {
		if contains(suspect.edgeHeads(), node.vertex) {
			suspects = append(suspects, suspect)
		}
	}
	for _, suspect := range suspects {
		if weight := suspect.getEdgeWeight(node) + suspect.shortestPathFromS; weight < greedyScore {
			greedyScore = weight
		}
	}

	if greedyScore == 10000000 {
		log.Panicln(greedyScore)
	}

	return greedyScore
}

func setDistance(node Node, distance int) Node {
	node.shortestPathFromS = distance
	return node
}

func main() {
	result := dijkstra(loadData("./course2/week2/dijkstra/largerData.txt"))

	// shortest-path distances to the following ten vertices, in order: 7,37,59,82,99,115,133,165,188,197.
	for _, v := range []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197} {
		fmt.Println(result[v-1].vertex, result[v-1].shortestPathFromS)
	}
}

func loadData(filepath string) Graph {
	data := Graph{}
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parseRowIntoAdjListEntry(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func parseRowIntoAdjListEntry(list Graph, row string) {
	rowSlice := strings.Fields(row)
	if num, err := strconv.Atoi(rowSlice[0]); err == nil {
		list[num-1] = Node{ // subtracting by 1 so that arrays/adjlist can be indexed the same
			vertex:            num - 1,
			edges:             convertStringSliceToEdgeSlice(rowSlice[1:]),
			shortestPathFromS: 1000000,
			visited:           false,
		}
	}
}

func convertStringSliceToEdgeSlice(args []string) []Edge {
	response := make([]Edge, 0)
	for _, arg := range args {
		headAndWeight := strings.Split(arg, ",")
		head, err := strconv.Atoi(headAndWeight[0])
		check(err)
		weight, err := strconv.Atoi(headAndWeight[1])
		check(err)
		response = append(response, Edge{
			head:   head - 1,
			weight: weight,
		})
	}
	return response
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
