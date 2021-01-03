package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func huffman(elements []int) int {
	fmt.Println(elements)
	return 2
}

func main() {
	fmt.Println(huffman(loadData("./course3/week3/huffman/smallData.txt")))
}

func loadData(filepath string) []int {
	data := make([]int, 0, 500)
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
