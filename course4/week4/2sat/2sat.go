package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/stacks/arraystack"
)

// Visited keeps a record of visited nodes
type Visited map[int]bool

func tsat(graph, transposedGraph Graph) bool {
	n := len(graph)
	order := arraystack.New()
	visitedMapInv := make(Visited, n)
	visitedMap := make(Visited, n)
	leaders := make(map[int]int, n)

	for k := range transposedGraph {
		dfs1(k, order, visitedMapInv, transposedGraph)
	}

	for !order.Empty() {
		v, _ := order.Pop()
		check := dfs2(v.(int), visitedMap, graph, leaders)
		if !check {
			return false
		}
	}

	return true
}

func dfs1(v int, order *arraystack.Stack, visitedMapInv Visited, graph Graph) {
	visitedMapInv[v] = true
	for k := range graph[v] {
		if !visitedMapInv[k] {
			dfs1(k, order, visitedMapInv, graph)
		}
	}
	order.Push(v)
}

func dfs2(v int, visitedMap Visited, graph Graph, leaders map[int]int) bool {
	scc := make([]int, len(graph))
	scc = append(scc, v)
	visitedMap[v] = true
	for k := range graph[v] {
		if !visitedMap[k] {
			leaders[k] = v
			if includes(scc, mirror(k)) {
				return false
			}
			dfs2(k, visitedMap, graph, leaders)
		}
	}
	return true
}

func includes(elements []int, x int) bool {
	for _, v := range elements {
		if v == x {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(tsat(buildGraphs("course4/week4/2sat/data/data1.txt")))
}

//---------------------------------------------------------------------------/

// Graph is a graph
type Graph map[int][]int

func buildGraphs(filepath string) (Graph, Graph) {
	graph := make(Graph, 0)
	transposedGraph := make(Graph, 0)
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

		rowSlice := strings.Fields(scanner.Text())
		first, err := strconv.Atoi(rowSlice[0])
		check(err)
		second, err := strconv.Atoi(rowSlice[1])
		check(err)

		firstPrimed := mirror(first)
		secondPrimed := mirror(second)

		if _, ok := graph[firstPrimed]; ok {
			graph[firstPrimed] = append(graph[firstPrimed], second)
		} else {
			graph[firstPrimed] = []int{second}
		}

		if _, ok := graph[secondPrimed]; ok {
			graph[secondPrimed] = append(graph[secondPrimed], first)
		} else {
			graph[secondPrimed] = []int{first}
		}

		if _, ok := transposedGraph[second]; ok {
			transposedGraph[second] = append(transposedGraph[second], firstPrimed)
		} else {
			transposedGraph[second] = []int{firstPrimed}
		}

		if _, ok := transposedGraph[first]; ok {
			transposedGraph[first] = append(transposedGraph[first], secondPrimed)
		} else {
			transposedGraph[first] = []int{secondPrimed}
		}

	}

	check(scanner.Err())

	return graph, transposedGraph
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
