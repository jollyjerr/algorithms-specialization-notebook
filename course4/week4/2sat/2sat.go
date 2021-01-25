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

func tsat(graph, transposedGraph [][]int) bool {
	visitedRecord := make([]bool, 0, len(graph))

	return false
}

func main() {
	fmt.Println(tsat(buildGraphs("course4/week4/2sat/smallData.txt")))
}

//---------------------------------------------------------------------------/

// [][]int represents a graph. [[tail, head],[tail, head]...n]
func buildGraphs(filepath string) ([][]int, [][]int) {
	graph := make([][]int, 0)
	transposedGraph := make([][]int, 0)
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	row := -1
	for scanner.Scan() {
		if row == -1 {
			row++
			continue
		}
		parseRowIntoEntry(&graph, &transposedGraph, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return graph, transposedGraph
}

func parseRowIntoEntry(graph, transposedGraph *[][]int, row string) {
	rowSlice := strings.Fields(row)
	first, err := strconv.Atoi(rowSlice[0])
	check(err)
	second, err := strconv.Atoi(rowSlice[1])
	check(err)

	firstPrimed := mirror(first)
	secondPrimed := mirror(second)

	node1 := []int{firstPrimed, second}
	node2 := []int{secondPrimed, first}

	node1T := []int{second, firstPrimed}
	node2T := []int{first, secondPrimed}

	*graph = append(*graph, node1)
	*graph = append(*graph, node2)

	*transposedGraph = append(*transposedGraph, node1T)
	*transposedGraph = append(*transposedGraph, node2T)
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func mirror(number int) int {
	if number > 0 {
		return -number
	} else if number < 0 {
		return int(math.Abs(float64(number)))
	}
	return 0
}
