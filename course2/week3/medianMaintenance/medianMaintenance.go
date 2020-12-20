package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

// find the sum of all medians over the course of inserting data one at a time
func medianMaintenance(list []int) int {

	tree := rbt.NewWithIntComparator()

	sum := 0

	for i, val := range list {
		tree.Put(i, val)
		fmt.Println(tree.Values()...)
		if size := tree.Size(); size%2 != 0 {
			mid, _ := tree.Get(size / 2)
			fmt.Println(mid)
			sum += mid.(int)
		}

		// sum += tree.Root.Value.(int)
	}

	return sum % 10000
}

func main() {
	fmt.Println(medianMaintenance(loadData("./course2/week3/medianMaintenance/smallData.txt")))
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
	return data
}

func parseRowIntoEntry(list *[]int, row string) {
	if num, err := strconv.Atoi(row); err == nil {
		*list = append(*list, num)
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
