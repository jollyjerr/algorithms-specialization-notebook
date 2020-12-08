package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var (
	comparisons = 0
)

func quicksort(args []int) []int {
	if len(args) > 2 { // Only do the work if not a base case
		leftpoint, rightpoint := 0, len(args)-1
		pivot := rand.Intn(rightpoint)
		// pivot := 0
		boundary := swapAroundPivot(args, pivot, leftpoint, rightpoint)

		comparisons += len(args[:boundary]) - 1
		comparisons += len(args[boundary:]) - 1

		quicksort(args[:boundary])
		quicksort(args[boundary:])
	}
	return args
}

func swapAroundPivot(elements []int, pivotIndx, leftpoint, rightpoint int) int {
	// pivot := elements[pivotIndx]
	// pivotBoundary := pivotIndx + 1
	// fmt.Println(elements[pivotIndx])
	elements[pivotIndx], elements[rightpoint] = elements[rightpoint], elements[pivotIndx] // move pivot to the right so we can iterate over the slice without the pivot

	for i := range elements {
		if elements[i] < elements[rightpoint] { // the only case when anything needs to happen
			fmt.Println("swap")
			elements[i], elements[leftpoint] = elements[leftpoint], elements[i]
			leftpoint++ // using the leftpoint as the pivotBoundary counter
		}
	}

	// for i := leftpoint; i < rightpoint; i++ {
	// 	if elements[i] < pivot && pivotBoundary > pivotIndx+1 { // the only case when anything needs to happen. New e needs to be swapped && the pivot boundary has already moved
	// 		pivotBoundary++
	// 		elements[pivotBoundary], elements[i] = elements[i], elements[pivotBoundary]
	// 	}
	// }

	elements[leftpoint], elements[rightpoint] = elements[rightpoint], elements[leftpoint] // swap the pivot element into place
	fmt.Println(elements)
	return leftpoint
}

// Assignment stuff under this comment

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func parseInput(csvdata [][]string) []int {
	res := make([]int, 0)
	for i := 0; i < len(csvdata); i++ {
		num := csvdata[i][0]
		newint, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		res = append(res, newint)
	}
	return res
}

func main() {
	// records := readCsvFile("./week3/quicksort/data.csv")
	// quicksort(parseInput(records))
	quicksort([]int{3, 2, 1})
	fmt.Println(comparisons)
}
