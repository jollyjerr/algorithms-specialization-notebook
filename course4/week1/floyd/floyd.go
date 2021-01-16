package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Edge is an edge dog
type Edge struct {
	tail   int
	head   int
	length int
}

func floyd(edges []Edge) int {
	fmt.Println(edges)
	return 2
}

func main() {
	fmt.Println(floyd(loadData("./course4/week1/floyd/smallData.txt")))
}

func loadData(filepath string) []Edge {
	items := make([]Edge, 0)
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	row := 1
	for scanner.Scan() {
		if row == 1 {
			row++
			continue
		}
		parseRowIntoEntry(&items, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return items
}

func parseRowIntoEntry(items *[]Edge, row string) {
	rowSlice := strings.Fields(row)
	tail, err := strconv.Atoi(rowSlice[0])
	check(err)
	head, err := strconv.Atoi(rowSlice[1])
	check(err)
	length, err := strconv.Atoi(rowSlice[2])
	check(err)
	*items = append(*items, Edge{
		tail,
		head,
		length,
	})
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
