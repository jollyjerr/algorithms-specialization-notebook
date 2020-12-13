package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// returns (a sorted version of the input array, the number of inversions)
func countInversions(args []int) ([]int, int) {
	// Base Case
	lengthOfArgs := len(args)
	if lengthOfArgs == 1 {
		return args, 0
	}

	firstHalfSorted, firstHalfCount := countInversions(args[:lengthOfArgs/2])
	secondHalfSorted, secondHalfCount := countInversions(args[lengthOfArgs/2:])

	fullSorted, splitCount := mergeAndCountSplitInv(firstHalfSorted, secondHalfSorted)

	return fullSorted, (firstHalfCount + secondHalfCount + splitCount)
}

func mergeAndCountSplitInv(left []int, right []int) ([]int, int) {
	var (
		size     = len(left) + len(right)
		ret      = make([]int, size, size)
		i        = 0 // left pointer to traverse
		j        = 0 // right pointer to traverse
		invCount = 0
	)

	for k := 0; k < size; k++ {
		// compare and sort by lowest element
		if i > len(left)-1 && j <= len(right)-1 {
			// left array depleted
			ret[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			// right array depleted
			ret[k] = left[i]
			i++
		} else if left[i] < right[j] {
			ret[k] = left[i]
			i++
		} else {
			ret[k] = right[j]
			j++
			// Add split inversions here based on claim in notes
			invCount += len(left[i:])
		}
	}

	return ret, invCount
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
	records := readCsvFile("./week2/countInversions/data.csv")
	_, answer := countInversions(parseInput(records))
	fmt.Println(answer)
}
