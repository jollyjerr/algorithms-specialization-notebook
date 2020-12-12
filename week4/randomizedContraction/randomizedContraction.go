package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// AdjacencyList represents a graph (according to the limits of this assignment)
type AdjacencyList map[int][]int

func randContraction(graph AdjacencyList) int {
	for len(graph) > 2 {
		// Gather keys to pick a random one
		keys := make([]int, 0, len(graph))
		for k := range graph {
			keys = append(keys, k)
		}

		// Pick a random key, random edge index, and find the other head of the edge
		randomVertex := keys[randomIntUnderThis(len(graph))]
		randomEdgeIndex := randomIntUnderThis(len(graph[randomVertex]))
		otherSideOfEdge := graph[randomVertex][randomEdgeIndex]

		// Delete the edge
		graph[randomVertex] = findAndDelete(graph[randomVertex], otherSideOfEdge)
		graph[otherSideOfEdge] = findAndDelete(graph[otherSideOfEdge], randomVertex)

		// Contract into super node
		graph[randomVertex] = append(graph[randomVertex], graph[otherSideOfEdge]...)

		// Update the other nodes to know about contraction
		for _, v := range keys {
			graph[v] = findAndReplace(graph[v], otherSideOfEdge, randomVertex)
		}

		// Remove self loops and old node
		graph[randomVertex] = findAndDelete(graph[randomVertex], randomVertex)
		delete(graph, otherSideOfEdge)
	}
	keys := make([]int, 0, 2)
	for k := range graph {
		keys = append(keys, k)
	}
	return len(graph[keys[0]])
}

func findMinCut(graph AdjacencyList) int {
	numberOfVertices := len(graph)
	bestMinimumCut := numberOfVertices * (numberOfVertices - 1) / 2 // max number of possible edges

	//approx n^2 * log(n)
	aproxNeededRotations := int((math.Pow(float64(numberOfVertices), 2))*math.Log(float64(numberOfVertices))) * 4000

	for i := 0; i < aproxNeededRotations; i++ {
		graphCopy := graph
		if roundResult := randContraction(graphCopy); roundResult < bestMinimumCut {
			bestMinimumCut = roundResult
		}
	}

	return bestMinimumCut
}

func main() {
	fmt.Println(findMinCut(loadData("./week4/randomizedContraction/data.txt")))
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
		list[num] = convertStringSliceToIntSlice(rowSlice[1:])
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

func randomIntUnderThis(max int) int {
	if max < 1 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func findAndDelete(graphVertexEdgeList []int, valueToRemove int) []int {
	response := make([]int, 0, len(graphVertexEdgeList))
	for _, v := range graphVertexEdgeList {
		if v != valueToRemove {
			response = append(response, v)
		}
	}
	return response
}

func findAndReplace(graphVertexEdgeList []int, valueToRemove, valueToReplaceWith int) []int {
	for i, v := range graphVertexEdgeList {
		if v == valueToRemove {
			graphVertexEdgeList[i] = valueToReplaceWith
		}
	}
	return graphVertexEdgeList
}
