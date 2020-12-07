package main

import (
	"math/rand"
)

func quicksort(args []int) []int {
	if len(args) > 2 { // Only do the work if not a base case
		leftpoint, rightpoint := 0, len(args)-1
		pivot := rand.Intn(rightpoint)
		boundary := swapAroundPivot(args, pivot, leftpoint, rightpoint)

		quicksort(args[:boundary])
		quicksort(args[boundary:])
	}
	return args
}

func swapAroundPivot(elements []int, pivotIndx, leftpoint, rightpoint int) int {
	// pivot := elements[pivotIndx]
	// pivotBoundary := pivotIndx + 1
	elements[pivotIndx], elements[rightpoint] = elements[rightpoint], elements[pivotIndx] // move pivot to the right so we can iterate over the slice without the pivot

	for i := range elements {
		if elements[i] < elements[rightpoint] { // the only case when anything needs to happen
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
	return leftpoint
}
