package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Node is a graph vertex
type Node struct {
	vertex   int
	explored bool
	edges    []int
}

// AdjacencyList is a representation of a graph
type AdjacencyList map[int]Node

func breadthFirstSearch(graph AdjacencyList, startVertex int) {
	queue := append(make([]Node, 0, len(graph)), exploredEdition(graph, startVertex))

	for len(queue) > 0 {
		V := queue[0]

		for _, edge := range V.edges {
			if suspect := graph[edge]; !suspect.explored {
				suspect.explored = true
				graph[edge] = suspect
				queue = append(queue, suspect)
			}
		}

		queue = safelyDequeue(queue)
	}
}

func main() {
	fmt.Println(countUndirectedConnectivity(loadData("./course2/week1/breadthFirstSearch/smallData.txt")))
}

func safelyDequeue(queue []Node) []Node {
	queue[0] = Node{}
	return queue[1:]
}

func exploredEdition(graph AdjacencyList, index int) Node {
	shallow := graph[index]
	shallow.explored = true
	graph[index] = shallow
	return shallow
}

// utils
func loadData(filepath string) AdjacencyList {
	data := AdjacencyList{}
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
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

func parseRowIntoAdjListEntry(list AdjacencyList, row string) {
	rowSlice := strings.Fields(row)
	if num, err := strconv.Atoi(rowSlice[0]); err == nil {
		list[num] = Node{
			vertex:   num,
			edges:    convertStringSliceToIntSlice(rowSlice[1:]),
			explored: false,
		}
	}
}

func convertStringSliceToIntSlice(args []string) []int {
	response := make([]int, 0)
	for _, arg := range args {
		if num, err := strconv.Atoi(arg); err == nil {
			response = append(response, num)
		}
	}
	return response
}
