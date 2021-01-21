package main

import (
	"bufio"
	"log"
	"os"
)

// Node is a city
type Node struct {
	index int
	X     int
	Y     int
}

// Edge is an edge
type Edge struct {
	head     int
	distance int
}

// Graph is a graph
type Graph map[int][]Edge

func loadData(filepath string) Graph {
	nodes := make([]Node, 0)
	graph := make(Graph, 0)
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
		parseRowIntoEntry(&nodes, &row, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return graph
}

func parseRowIntoEntry(items *[]Node, index *int, row string) {
	// rowSlice := strings.Fields(row)
	// X, err := strconv.Atoi(rowSlice[0])
	// check(err)
	// Y, err := strconv.Atoi(rowSlice[1])
	// check(err)

	*index++
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
