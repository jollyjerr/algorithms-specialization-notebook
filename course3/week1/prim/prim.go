package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/trees/binaryheap"
)

type edge struct {
	node1                 int
	node2                 int
	cost                  int
	ifHeldByTailTheHeadIs int
}

type node struct {
	vertex int
	edges  []edge
	// heap key
	key       int
	edgeToKey edge
}

type graph map[int]node

type frontier []node

func (f frontier) contains(vertex int) bool {
	for _, node := range f {
		if node.vertex == vertex {
			return true
		}
	}
	return false
}

func (f frontier) calculateHeapKey(node node) (int, edge) {
	lowestCost := 1 << 31 // max int32
	edgeToKey := edge{}
	for _, edge := range node.edges {
		if f.contains(edge.ifHeldByTailTheHeadIs) && edge.cost < lowestCost {
			lowestCost = edge.cost
			edgeToKey = edge
		}
	}
	return lowestCost, edgeToKey
}

func (f *frontier) push(graph graph, vertex int) {
	*f = append(*f, graph[vertex])
	delete(graph, vertex)
}

func prim(graph graph) int {
	T := make([]edge, 0)
	X := make(frontier, 0, len(graph))
	X.push(graph, 1)

	heap := generateHeap(X, graph)

	for heap.Size() > 0 {
		popped, _ := heap.Pop()
		X.push(graph, popped.(node).vertex)
		T = append(T, popped.(node).edgeToKey)
		heap = generateHeap(X, graph) // this heap does not support deletions so I am just re-creating it each time :/
	}

	return sumEdges(T)
}

func generateHeap(X frontier, graph graph) *binaryheap.Heap {
	heap := binaryheap.NewWith(byKey)
	for _, node := range graph {
		node.key, node.edgeToKey = X.calculateHeapKey(node)
		heap.Push(node)
	}
	return heap
}

// Custom comparator (sort by keys)
func byKey(a, b interface{}) int {
	// Type assertion
	c1 := a.(node)
	c2 := b.(node)
	switch {
	case c1.key > c2.key:
		return 1
	case c1.key < c2.key:
		return -1
	default:
		return 0
	}
}

func sumEdges(edges []edge) int {
	sum := 0
	for _, v := range edges {
		sum += v.cost
	}
	return sum
}

func main() {
	fmt.Println(prim(loadData("./course3/week1/prim/data.txt")))
}

func loadData(filepath string) graph {
	data := graph{}
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parseRowIntoEntry(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func parseRowIntoEntry(container graph, row string) {
	rowSlice := strings.Fields(row)
	if len(rowSlice) > 2 {
		node1, err := strconv.Atoi(rowSlice[0])
		check(err)
		node2, err := strconv.Atoi(rowSlice[1])
		check(err)
		cost, err := strconv.Atoi(rowSlice[2])
		check(err)

		newEdge := edge{
			node1:                 node1,
			node2:                 node2,
			cost:                  cost,
			ifHeldByTailTheHeadIs: node2,
		}

		if v, exists := container[node1]; exists {
			v.edges = append(v.edges, newEdge)
			container[node1] = v
		} else {
			container[node1] = node{
				vertex: node1,
				edges:  []edge{newEdge},
			}
		}

		newEdge.ifHeldByTailTheHeadIs = node1

		if v, exists := container[node2]; exists {
			v.edges = append(v.edges, newEdge)
			container[node2] = v
		} else {
			container[node2] = node{
				vertex: node2,
				edges:  []edge{newEdge},
			}
		}
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
