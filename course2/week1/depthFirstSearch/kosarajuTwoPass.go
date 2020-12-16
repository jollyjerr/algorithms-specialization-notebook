package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Node is a graph vertex (directed in this assignment)
type Node struct {
	vertex   int
	edgesTo  []int
	explored bool
	leader   int
}

// Graph is an adjacency list representation of a graph
type Graph map[int]Node

func kosarajuTwoPass(graph Graph) []int {
	transposedGraph := transposeGraph(graph)
	order := dfsLoop(transposedGraph, 1)
	graph = switchNodeNamesWithIndexInOrder(graph, order)
	leaders := dfsLoop(graph, 2)
	return leaders
}

type bookKeeping struct {
	finishingTimes        []int
	countOfNodesProcessed int
	leader                int
}

func dfsLoop(graph Graph, pass int) []int {
	books := bookKeeping{
		finishingTimes:        make([]int, len(graph)),
		countOfNodesProcessed: 0,
		leader:                0,
	}

	for i := len(graph); i >= 1; i-- {
		if !graph[i].explored {
			books.leader = i
			depthFirstSearch(graph, i, &books)
		}
	}

	if pass == 1 {
		return books.finishingTimes
	}

	leaders := make([]int, len(graph))
	for i, key := range graph.keys() {
		leaders[i] = graph[key].leader
	}
	return leaders
}

func depthFirstSearch(graph Graph, nodeIndex int, books *bookKeeping) {
	graph[nodeIndex] = markLeaderAndExplored(graph, nodeIndex, books.leader)
	for _, nestedVal := range graph[nodeIndex].edgesTo {
		if !graph[nestedVal].explored {
			depthFirstSearch(graph, nestedVal, books)
		}
	}
	books.countOfNodesProcessed++
	books.finishingTimes[books.countOfNodesProcessed-1] = nodeIndex
}

func main() {
	fmt.Println(findTopFiveSCCs(countNodesInSCCs(kosarajuTwoPass(loadData("./course2/week1/depthFirstSearch/data.txt")))))
}

func transposeGraph(graph Graph) Graph { // TODO look for a better solution
	res := Graph{}
	for _, key := range graph.keys() {
		for _, nValue := range graph[key].edgesTo {
			if val, ok := res[nValue]; ok {
				val.edgesTo = append(val.edgesTo, key)
				res[nValue] = val
			} else {
				res[nValue] = Node{
					vertex:   nValue,
					edgesTo:  []int{key},
					explored: false,
					leader:   0,
				}
			}
		}
	}
	for index := range graph {
		if val, ok := res[index]; !ok {
			res[index] = Node{
				vertex:   val.vertex,
				edgesTo:  []int{},
				explored: false,
				leader:   0,
			}
		}
	}
	return res
}

func markLeaderAndExplored(graph Graph, index, leader int) Node {
	shallow := graph[index]
	shallow.explored = true
	shallow.leader = leader
	return shallow
}

func switchNodeNamesWithIndexInOrder(graph Graph, order []int) Graph {
	full := len(order)
	res := Graph{}
	for i, value := range order {

		if i%50000 == 0 {
			fmt.Println(i, "out of", full)
		}

		node := graph[value]
		node.edgesTo = renameEdgesTo(node.edgesTo, order)
		res[i+1] = graph[value]
	}
	return res
}

func renameEdgesTo(edges []int, order []int) []int {
	for i, value := range edges {
		edges[i] = indexOf(value, order) + 1
	}
	return edges
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func (graph Graph) keys() []int {
	keys := make([]int, len(graph))
	count := 0
	for i := range graph {
		keys[count] = i
		count++
	}
	sort.Ints(keys)
	return keys
}

func countNodesInSCCs(elements []int) map[int]int {
	res := make(map[int]int, 0)
	for _, key := range elements {
		if val, ok := res[key]; ok {
			res[key] = val + 1
		} else {
			res[key] = 1
		}
	}
	return res
}

func findTopFiveSCCs(elements map[int]int) []int {
	res := make([]int, len(elements))

	count := 0
	for _, value := range elements {
		res[count] = value
		count++
	}

	sort.Ints(res)

	return res
}

//------------------------------------------------------------------------------

func loadData(filepath string) Graph {
	data := Graph{}
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

func parseRowIntoAdjListEntry(graph Graph, row string) {
	rowSlice := strings.Fields(row)
	if num, err := strconv.Atoi(rowSlice[0]); err == nil {
		if val, ok := graph[num]; ok {
			val.edgesTo = append(val.edgesTo, toString(rowSlice[1]))
			graph[num] = val
		} else {
			graph[num] = Node{
				vertex:   num,
				edgesTo:  append(make([]int, 0), toString(rowSlice[1])),
				explored: false,
				leader:   0,
			}
		}
	}
}

func toString(arg string) int {
	num, err := strconv.Atoi(arg)
	if err != nil {
		log.Panicln(err)
	}
	return num
}
