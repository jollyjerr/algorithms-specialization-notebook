package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type job struct {
	weight     int
	length     int
	difference int
	ratio      float64
	completion int
}

func scheduleJobs(elements []job, sortbyDif bool) int {
	if sortbyDif {
		sort.Sort(byDifference(elements))
	} else {
		sort.Sort(byRatio(elements))
	}
	return sum(calcCompletionTimes(elements))
}

func calcCompletionTimes(elements []job) []job {
	currentTime := 0
	for i, v := range elements {
		currentTime += v.length
		v.completion = currentTime
		elements[i] = v
	}
	return elements
}

func sum(elements []job) int {
	sum := 0
	for _, v := range elements {
		sum += (v.completion * v.weight)
	}
	return sum
}

// byDifference is a sort api
type byDifference []job

func (bd byDifference) Len() int { return len(bd) }
func (bd byDifference) Less(i, j int) bool {
	if bd[i].difference == bd[j].difference {
		return bd[i].weight > bd[j].weight
	}
	return bd[i].difference > bd[j].difference
}
func (bd byDifference) Swap(i, j int) { bd[i], bd[j] = bd[j], bd[i] }

// byRatio is a sort api
type byRatio []job

func (br byRatio) Len() int           { return len(br) }
func (br byRatio) Less(i, j int) bool { return br[i].ratio > br[j].ratio }
func (br byRatio) Swap(i, j int)      { br[i], br[j] = br[j], br[i] }

func main() {
	fmt.Println(scheduleJobs(loadData("./course3/week1/schedulingJobs/data.txt"), false)) // 1. 69119377652, 2. 67311454237
}

func loadData(filepath string) []job {
	data := make([]job, 0, 10000)
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

func parseRowIntoEntry(container *[]job, row string) {
	rowSlice := strings.Fields(row)
	if len(rowSlice) > 1 {
		weight, err := strconv.Atoi(rowSlice[0])
		check(err)
		length, err := strconv.Atoi(rowSlice[1])
		check(err)
		*container = append(*container, job{
			weight:     weight,
			length:     length,
			difference: weight - length,
			ratio:      float64(weight) / float64(length),
		})
	}
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
