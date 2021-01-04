package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/emirpasic/gods/trees/binaryheap"
)

// Tree is a tree sorta
type Tree map[int]int

// Element is an element
type Element struct {
	Nodes []int
	Cost  int
}

func byCost(a, b interface{}) int {
	c1 := a.(Element)
	c2 := b.(Element)

	switch {
	case c1.Cost > c2.Cost:
		return 1
	case c1.Cost < c2.Cost:
		return -1
	default:
		return 0
	}
}

func huffman(elements []int) int {
	tree := Tree{}
	heap := binaryheap.NewWith(byCost)
	for _, element := range elements {
		heap.Push(Element{
			Nodes: []int{element},
			Cost:  element,
		})
	}
	count := 0
	//-----------------------------------------
	singleRound := func() {
		t1, _ := heap.Pop()
		t2, _ := heap.Pop()
		nodes1 := t1.(Element).Nodes
		nodes1 = append(nodes1, t2.(Element).Nodes...)
		heap.Push(Element{
			Nodes: nodes1,
			Cost:  t1.(Element).Cost + t2.(Element).Cost,
		})

		if len(t1.(Element).Nodes) == 1 {
			tree[t1.(Element).Nodes[0]] = count
		}
		if len(t2.(Element).Nodes) == 1 {
			tree[t2.(Element).Nodes[0]] = count
		}

		fmt.Println(t1, t2)

		count++
	}

	// Build the tree
	for heap.Size() > 2 {
		singleRound()
	}
	// Reconstruct codes
	t1, _ := heap.Pop()
	t2, _ := heap.Pop()

	fmt.Println(t1, t2)

	return 2
}

func main() {
	fmt.Println(huffman(loadData("./course3/week3/huffman/smallData.txt")))
}

func loadData(filepath string) []int {
	data := make([]int, 0)
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
	return data[1:]
}

func parseRowIntoEntry(container *[]int, row string) {
	if res, err := strconv.Atoi(row); err == nil {
		*container = append(*container, res)
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
