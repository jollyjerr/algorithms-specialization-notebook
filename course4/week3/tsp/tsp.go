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

const (
	inf = 4294967295.8574839
)

// Node is a city
type Node struct {
	X       float64
	Y       float64
	visited bool
}

func (n Node) distanceTo(node Node) float64 {
	return math.Sqrt(math.Pow((n.X-node.X), 2) + math.Pow((n.Y-node.Y), 2))
}

func tsp(elements []Node) int {
	totalDistance := 0.0

	elements[0].visited = true
	lastVisited := elements[0]

	for i := 1; i < len(elements); i++ {
		minDist := inf
		citiesOfInterest := make([]*Node, 0)

		for j := 0; j < len(elements); j++ {
			if prospect := &elements[j]; !prospect.visited {
				if distance := lastVisited.distanceTo(*prospect); distance < minDist {
					minDist = distance
					citiesOfInterest = append(citiesOfInterest, prospect)
				}
			}
		}

		for _, prospect := range citiesOfInterest { // this is to handle the weird tie rule in the assignment
			if lastVisited.distanceTo(*prospect) == minDist {
				prospect.visited = true
				lastVisited = *prospect

				totalDistance += minDist
				break
			}
		}
	}

	// go home
	totalDistance += lastVisited.distanceTo(elements[0])

	return int(math.Round(totalDistance))
}

func main() {
	fmt.Println(tsp(loadData("course4/week3/tsp/data.txt")))
}

func loadData(filepath string) []Node {
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
		parseRowIntoEntry(&nodes, &row, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nodes
}

func parseRowIntoEntry(items *[]Node, index *int, row string) {
	rowSlice := strings.Fields(row)
	X, err := strconv.ParseFloat(rowSlice[1], 64)
	check(err)
	Y, err := strconv.ParseFloat(rowSlice[2], 64)
	check(err)

	visited := false

	*items = append(*items, Node{
		X,
		Y,
		visited,
	})

	*index++
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
