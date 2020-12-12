package main

import (
	"fmt"
	"math/rand"
)

func rSelect(elements []int, statistic int) int { // todo how do we treat even/odd lengths? Which way is perfered?
	if len(elements) == 1 { // base case
		return elements[0]
	}

	leftpoint, rightpoint := 0, len(elements)-1

	pivot := rand.Intn(rightpoint)

	boundary := swapAroundPivot(elements, pivot, leftpoint, rightpoint)

	if boundary == statistic { // really lucky
		return elements[boundary]
	} else if boundary > statistic {
		if split := boundary - 1; split > 0 {
			return rSelect(elements[:split], statistic)
		}
	}
	return rSelect(elements[boundary:], statistic-boundary)
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

func main() {
	test := []int{1, 2, 3, 4, 5}
	res := rSelect(test, 3)
	fmt.Println(res)
	nextTry := rSelect([]int{1, 2, 3, 4, 5, 6, 7, 9, 8}, 5)
	fmt.Println(nextTry)
}
