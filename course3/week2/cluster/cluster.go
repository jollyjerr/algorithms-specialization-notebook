package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Edge is an edge
type Edge struct {
	node1 int
	node2 int
	cost  int
}

func kClusterings(edges []Edge, numberOfClustersDesired int) int {
	fmt.Println(edges)
	return 2
}

func main() {
	fmt.Println(kClusterings(loadData("./course3/week2/cluster/data.txt"), 4))
}

func loadData(filepath string) []Edge {
	data := make([]Edge, 0, 500)
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parseRowIntoEntry(&data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func parseRowIntoEntry(container *[]Edge, row string) {
	rowSlice := strings.Fields(row)
	if len(rowSlice) > 1 {
		node1, err := strconv.Atoi(rowSlice[0])
		check(err)
		node2, err := strconv.Atoi(rowSlice[1])
		check(err)
		cost, err := strconv.Atoi(rowSlice[2])
		check(err)
		*container = append(*container, Edge{
			node1,
			node2,
			cost,
		})
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
