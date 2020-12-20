package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

// find the sum of all medians over the course of each step of inserting data one at a time in O(logn)
func medianMaintenance(list []int) int {
	sum := 0

	maxHeap := binaryheap.NewWithIntComparator()
	inverseIntComparator := func(a, b interface{}) int {
		return -utils.IntComparator(a, b)
	}
	minHeap := binaryheap.NewWith(inverseIntComparator)

	for i, val := range list {
		if i == 0 {
			minHeap.Push(val)
			sum += val
			// fmt.Println(val)
		} else if i == 1 {
			min, _ := minHeap.Peek()
			if val > min.(int) {
				maxHeap.Push(val)
			} else {
				max, _ := minHeap.Pop()
				maxHeap.Push(max)
				minHeap.Push(val)
			}
			min, _ = minHeap.Peek()
			sum += min.(int)
			// fmt.Println(min)
		} else {
			roundWinner := val

			bigMin, _ := minHeap.Peek()
			smallMax, _ := maxHeap.Peek()

			minHeapSize := minHeap.Size()
			maxHeapSize := maxHeap.Size()

			if bigMin.(int) < val && val < smallMax.(int) {
				if minHeapSize > maxHeapSize {
					maxHeap.Push(val)
				} else {
					minHeap.Push(val)
				}
			} else if val > smallMax.(int) {
				if maxHeapSize > minHeapSize {
					smallMax, _ = maxHeap.Pop()
					minHeap.Push(smallMax)
				}
				maxHeap.Push(val)
			} else if val < bigMin.(int) {
				if minHeapSize > maxHeapSize {
					bigMin, _ = minHeap.Pop()
					maxHeap.Push(bigMin)
				}
				minHeap.Push(val)
			}

			if minHeap.Size() < maxHeap.Size() {
				winner, _ := maxHeap.Peek()
				roundWinner = winner.(int)
			} else {
				bigMin, _ = minHeap.Peek()
				roundWinner = bigMin.(int)
			}

			// fmt.Println(roundWinner, minHeap.Values(), maxHeap.Values(), i%2)

			sum += roundWinner
		}
	}

	return sum % 10000
}

func main() {
	fmt.Println(medianMaintenance(loadData("./course2/week3/medianMaintenance/data.txt")))
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
