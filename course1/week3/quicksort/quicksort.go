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
	if len(args) == 1 { // base case
		return args
	}

	comparisons += len(args) - 1
	leftpoint, rightpoint := 0, len(args)-1

	pivot := rand.Intn(rightpoint)
	// pivot := leftpoint
	// pivot := len(args) - 1
	// pivot := findMedianOfThreePivot(args)

	boundary := swapAroundPivot(args, pivot, leftpoint, rightpoint)

	if split := boundary - 1; split > 0 {
		quicksort(args[:split])
	}
	if len(args[boundary:]) > 1 {
		quicksort(args[boundary:])
	}

	return args
}

func swapAroundPivot(elements []int, pivotIndx, leftpoint, rightpoint int) int {

	if pivotIndx != leftpoint {
		elements[leftpoint], elements[pivotIndx] = elements[pivotIndx], elements[leftpoint]
		pivotIndx = leftpoint
	}

	pivot := elements[pivotIndx]
	i := leftpoint + 1

	for j := range elements {
		if j > 0 {
			if elements[j] < pivot { // the only case when anything needs to happen
				elements[j], elements[i] = elements[i], elements[j]
				i++
			}
		}
	}

	// swap the pivot element into place
	elements[pivotIndx], elements[i-1] = elements[i-1], elements[pivotIndx]

	return i
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
		midIndx = (len(args) / 2) - 1
	} else {
		midIndx = len(args) / 2
	}
	mid = args[midIndx]

	var median = first // I am sure there is a better way to do this but it is too late for this lol
	if (mid > first && mid < last) || (mid > last && mid < first) {
		median = mid
	} else if (last > mid && last < first) || (last > first && last < mid) {
		median = last
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
