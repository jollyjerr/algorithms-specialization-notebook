package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// AdjacencyList represents a graph (according to the limits of this assignment)
type AdjacencyList map[int][]int

func randContraction(graph AdjacencyList) int {
	// for len(graph) > 2 {

	// }
	return 2 // todo count crossing edges
}

func findMinCut(graph AdjacencyList) int {
	numberOfVertices := len(graph)
	bestMinimumCut := numberOfVertices * (numberOfVertices - 1) / 2 // max number of possible edges

	//approx n^2 * log(n)
	aproxNeededRotations := int((math.Pow(float64(numberOfVertices), 2)) * math.Log(float64(numberOfVertices)))

	for i := 0; i < aproxNeededRotations; i++ {
		if roundResult := randContraction(graph); roundResult < bestMinimumCut {
			bestMinimumCut = roundResult
		}
	}

	return bestMinimumCut
}

func main() {
	fmt.Println(findMinCut(loadData("./week4/randomizedContraction/data.txt")))
}

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
