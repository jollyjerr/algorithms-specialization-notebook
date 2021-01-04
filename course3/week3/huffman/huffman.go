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

func huffman(elements []int) int {
	tree := Tree{}
	heap := binaryheap.NewWithIntComparator()
	for _, element := range elements {
		heap.Push(element)
	}
	count := 0
	//-----------------------------------------
	singleRound := func() {
		t1, _ := heap.Pop()
		t2, _ := heap.Pop()
		heap.Push(t1.(int) + t2.(int))

		if _, exist := tree[t1.(int)]; !exist {
			tree[t1.(int)] = count
		}
		if _, exist := tree[t2.(int)]; !exist {
			tree[t2.(int)] = count
		}

		count++

		fmt.Println(tree)
	}

	for heap.Size() > 2 {
		singleRound()
	}

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
