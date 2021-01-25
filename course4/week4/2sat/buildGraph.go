package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func buildGraph(filepath string) []Node {
	nodes := make([]Node, 0)
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
		parseRowIntoEntry(&nodes, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nodes
}

func parseRowIntoEntry(items *[]Node, row string) {
	rowSlice := strings.Fields(row)
	first, err := strconv.Atoi(rowSlice[0])
	check(err)
	second, err := strconv.Atoi(rowSlice[1])
	check(err)

	firstPrimed := mirror(first)
	secondPrimed := mirror(second)

	node1 := Node{
		vertex:   firstPrimed,
		edgesTo:  []int{second},
		explored: false,
	}

	node2 := Node{
		vertex:   secondPrimed,
		edgesTo:  []int{first},
		explored: false,
	}

	*items = append(*items, node1)
	*items = append(*items, node2)
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
