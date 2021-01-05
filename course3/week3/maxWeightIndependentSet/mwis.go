package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mwis(elements []int) []int {
	subProbAns := make([]int, 0)
	subProbAns = append(subProbAns, 0)
	subProbAns = append(subProbAns, elements[0])
	for i, v := range elements {
		if i > 1 {
			if newWeight := (subProbAns[i-2] + v); newWeight > subProbAns[i-1] {
				subProbAns = append(subProbAns, newWeight)
			} else {
				subProbAns = append(subProbAns, subProbAns[i-1])
			}
		}
	}

	reconstruction := make([]int, 0)
	count := len(elements)
	for count > 1 {
		if subProbAns[count-1] > subProbAns[count-2] {
			count--
		} else {
			reconstruction = append(reconstruction, elements[count])
			count -= 2
		}
	}
	reconstruction = append(reconstruction, elements[0])

	return formatAns(elements, reconstruction)
}

func formatAns(elements, reconstructed []int) []int {
	res := make([]int, 0)
	for _, v := range []int{elements[0], elements[1], elements[2], elements[3], elements[16], elements[116], elements[516], elements[996]} {
		if contains(reconstructed, v) {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
	}
	return res
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(mwis(loadData("./course3/week3/maxWeightIndependentSet/data.txt"))) // 9 19
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
