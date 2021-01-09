package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type knapsack struct {
	size       int
	numOfItems int
}

type item struct {
	value  int
	weight int
}

// Knapsack is a knap sack
func Knapsack(sack knapsack, items []item) int {
	A := buildContainer(sack.numOfItems, sack.size)
	for i := 1; i < len(A); i++ {
		currentWeight := items[i-1].weight
		currentValue := items[i-1].value

		for x := 0; x < currentWeight; x++ {
			A[i][x] = A[i-1][x]
		}

		for x := currentWeight; x < len(A[i]); x++ {
			if newVal := (A[i-1][x-currentWeight]) + currentValue; newVal > A[i-1][x] {
				A[i][x] = newVal
			} else {
				A[i][x] = A[i-1][x]
			}
		}

	}

	return A[sack.numOfItems][sack.size]
}

func buildContainer(numItems, weight int) [][]int {
	A := make([][]int, numItems+1) // +1
	for i := range A {
		A[i] = make([]int, weight+1) // +1
	}
	return A
}

func main() {
	fmt.Println(Knapsack(loadData("./course3/week4/knapsack/bigData.txt")))
}

func loadData(filepath string) (knapsack, []item) {
	items := make([]item, 0)
	sack := knapsack{}
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	row := 1
	for scanner.Scan() {
		if row == 1 {
			if size, err := strconv.Atoi(strings.Fields(scanner.Text())[0]); err == nil {
				numOfItems, err := strconv.Atoi(strings.Fields(scanner.Text())[1])
				check(err)
				sack = knapsack{
					size,
					numOfItems,
				}
			}
			row++
			continue
		}
		parseRowIntoEntry(&items, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sack, items
}

func parseRowIntoEntry(items *[]item, row string) {
	rowSlice := strings.Fields(row)
	value, err := strconv.Atoi(rowSlice[0])
	check(err)
	weight, err := strconv.Atoi(rowSlice[1])
	check(err)
	*items = append(*items, item{value, weight})
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
