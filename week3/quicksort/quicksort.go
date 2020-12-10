package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	comparisons = 0
)

func quicksort(args []int) []int {
	if len(args) > 1 { // Only do the work if not a base case
		leftpoint, rightpoint := 0, len(args)-1

		// pivot := rand.Intn(rightpoint)
		// pivot := len(args) - 1
		pivot := findMedianOfThreePivot(args)

		boundary := swapAroundPivot(args, pivot, leftpoint, rightpoint)

		comparisons += len(args[:boundary]) - 1
		comparisons += len(args[boundary:]) - 1

		quicksort(args[:boundary])
		quicksort(args[boundary:])
	}
	return args
}

func swapAroundPivot(elements []int, pivotIndx, leftpoint, rightpoint int) int {
	pivot := elements[pivotIndx]
	// pivotBoundary := pivotIndx + 1
	// fmt.Println(elements[pivotIndx])

	// move pivot to the right so we can iterate over the slice without the pivot
	// elements[pivotIndx], elements[rightpoint] = elements[rightpoint], elements[pivotIndx]

	if pivotIndx != leftpoint {
		elements[leftpoint], elements[pivotIndx] = elements[pivotIndx], elements[leftpoint]
		pivotIndx = leftpoint
	}

	for i := range elements {
		if elements[i] < pivot { // the only case when anything needs to happen
			leftpoint++ // using the leftpoint as the pivotBoundary counter
			// fmt.Println(elements, elements[i], pivot, leftpoint)
			elements[i], elements[leftpoint] = elements[leftpoint], elements[i]
		}
	}

	// for i := leftpoint; i < rightpoint; i++ {
	// 	if elements[i] < pivot && pivotBoundary > pivotIndx+1 { // the only case when anything needs to happen. New e needs to be swapped && the pivot boundary has already moved
	// 		pivotBoundary++
	// 		elements[pivotBoundary], elements[i] = elements[i], elements[pivotBoundary]
	// 	}
	// }

	// swap the pivot element into place
	// elements[leftpoint], elements[rightpoint] = elements[rightpoint], elements[leftpoint]
	elements[leftpoint], elements[pivotIndx] = elements[pivotIndx], elements[leftpoint]

	return leftpoint + 1
}

// Assignment stuff under this comment

func findMedianOfThreePivot(args []int) int {
	var (
		first   = args[0]
		last    = args[len(args)-1]
		mid     = 0
		midIndx = 0
	)

	if len(args)%2 == 0 {
		midIndx = len(args)/2 - 1
	} else {
		midIndx = len(args) / 2
	}
	mid = args[midIndx]

	var median = first
	for _, v := range []int{last, mid} {
		if v < median {
			median = v
		}
	}

	if median == first {
		return 0
	} else if median == last {
		return len(args) - 1
	}
	return midIndx
}

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
	records := readCsvFile("./week3/quicksort/data.csv")
	quicksort(parseInput(records))
	// ans := quicksort([]int{4, 3, 2, 7, 1, -88, 44, 18, -2, 77})
	// fmt.Println(ans)
	fmt.Println(comparisons)
}
