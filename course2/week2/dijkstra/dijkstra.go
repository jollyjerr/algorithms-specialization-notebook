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

// Edge is an edge between vertices
type Edge struct {
	// The tail is whatever node is storing the edge
	head   int
	weight int
}

// Node is a vertex
type Node struct {
	vertex            int
	edges             []Edge
	shortestPathFromS int // should default to 1000000
	visited           bool
}

// Graph is an adjacency list representation of a graph
type Graph map[int]Node

// get graph keys in order
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

func dijkstra(graph Graph) []int {
	heap := Heap{
		size:      len(graph),
		array:     make([]Node, 0, len(graph)),
		positions: make([]int, 0, len(graph)),
	}

	for key := range graph.keys() {
		heap.addNode(graph[key], key)
	}

	heap.decreaseKey(0, 0)

	for !heap.isEmpty() {
		node := heap.extractMin()
		for _, edge := range node.edges {
			if heap.isNodeInMinHeap(edge.head) {
				heap.decreaseKey(edge.head, edge.weight)
			}
		}
	}

	return heap.positions
}

func main() {
	fmt.Println(dijkstra(loadData("./course2/week2/dijkstra/smallData.txt")))
}

//************************************************************************************ Heap

// Heap is a priority queue and api for dijkstra's loop
type Heap struct {
	size      int
	array     []Node
	positions []int
}

func (heap *Heap) addNode(node Node, position int) {
	heap.array = append(heap.array, node)
	heap.positions = append(heap.positions, position)
}

func (heap *Heap) swap(i1, i2 int) {
	heap.positions[i1], heap.positions[i2] = i2, i1
	heap.array[i1], heap.array[i2] = heap.array[i2], heap.array[i1]
}

func (heap *Heap) minHeapifyDown(index int) {
	parent := index
	left, right := left(index), right(index)

	if heap.array[left].shortestPathFromS < heap.array[parent].shortestPathFromS {
		parent = left
	}
	if heap.array[right].shortestPathFromS < heap.array[parent].shortestPathFromS {
		parent = right
	}

	if parent != index {
		heap.swap(parent, index)
		heap.minHeapifyDown(parent)
	}
}

func (heap *Heap) minHeapifyUp(index int) {
	for heap.array[parent(index)].shortestPathFromS > heap.array[index].shortestPathFromS {
		heap.swap(index, parent(index))
		index = parent(index)
	}
}

func (heap *Heap) decreaseKey(vertex, distance int) {
	index := heap.positions[vertex]
	heap.array[index].shortestPathFromS = distance
	heap.minHeapifyUp(index)
}

func (heap *Heap) extractMin() Node {
	if heap.isEmpty() {
		return Node{}
	}

	root := heap.array[0]
	lastNode := heap.array[len(heap.array)-1]

	// move the last node into root position
	heap.array[0] = heap.array[len(heap.array)-1]
	heap.array[len(heap.array)-1] = Node{}

	// update positions mapping
	heap.positions[lastNode.vertex] = 0
	heap.positions[root.vertex] = heap.size - 1

	heap.size--
	heap.minHeapifyDown(0)

	return root
}

// check if heap is empty
func (heap *Heap) isEmpty() bool {
	return heap.size == 0
}

// Check if vertex has been pulled out of heap and discovered
func (heap *Heap) isNodeInMinHeap(vertex int) bool {
	fmt.Println(vertex)
	return heap.positions[vertex] < heap.size
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
func right(i int) int {
	return 2*i + 2
}

//************************************************************************************ Assignment stuff

func loadData(filepath string) Graph {
	data := Graph{}
	f, err := os.Open(filepath)
	check(err)
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

func parseRowIntoAdjListEntry(list Graph, row string) {
	rowSlice := strings.Fields(row)
	if num, err := strconv.Atoi(rowSlice[0]); err == nil {
		list[num-1] = Node{ // subtracting by 1 so that arrays/adjlist can be indexed the same
			vertex:            num - 1,
			edges:             convertStringSliceToEdgeSlice(rowSlice[1:]),
			shortestPathFromS: 1000000,
			visited:           false,
		}
	}
}

func convertStringSliceToEdgeSlice(args []string) []Edge {
	response := make([]Edge, 0)
	for _, arg := range args {
		headAndWeight := strings.Split(arg, ",")
		head, err := strconv.Atoi(headAndWeight[0])
		check(err)
		weight, err := strconv.Atoi(headAndWeight[1])
		check(err)
		response = append(response, Edge{
			head:   head - 1,
			weight: weight,
		})
	}
	return response
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
