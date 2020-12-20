package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// number mapped to number of times you see it
type table map[int]int

func twoSum(table table, start, end int) int {

	numberOfTargetValuesThatSatisfyCondition := 0
	for target := start; target <= end; target++ {
		if target%1000 == 0 {
			fmt.Println(target)
		}
		for key := range table {
			pairMatch := -(key - target)

			if pairMatch == key {
				continue
			}

			if _, ok := table[pairMatch]; ok {
				numberOfTargetValuesThatSatisfyCondition++
				break
			}
		}
	}

	return numberOfTargetValuesThatSatisfyCondition
}

func main() {
	fmt.Println(twoSum(loadData("./course2/week4/twoSum/data.txt"), -100000, 100000))
}

func loadData(filepath string) table {
	data := table{}
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parseRowIntoEntry(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func parseRowIntoEntry(table table, row string) {
	if num, err := strconv.Atoi(row); err == nil {
		if nums, ok := table[num]; ok {
			nums++
			table[num] = nums
		} else {
			table[num] = 1
		}
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
