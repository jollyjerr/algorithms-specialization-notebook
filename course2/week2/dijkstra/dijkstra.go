package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
}

// Graph is an adjacency list representation of a graph
type Graph map[int]Node

func dijkstra(graph Graph) []int {
	fmt.Println(graph)
	return []int{2}
}

func main() {
	fmt.Println(dijkstra(loadData("./course2/week2/dijkstra/smallData.txt")))
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
		list[num] = Node{
			vertex:            num,
			edges:             convertStringSliceToEdgeSlice(rowSlice[1:]),
			shortestPathFromS: 1000000,
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
			head:   head,
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
