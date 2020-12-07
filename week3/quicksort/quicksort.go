package main

import "math/rand"

func quicksort(args []int) []int {
	if len(args) != 1 { // Only do the work if not a base case
		leftpoint, rightpoint := 0, len(args)-1
		pivot := rand.Intn(len(args))
		swapAroundPivot(args, leftpoint, rightpoint)
	}
	return args
}

func swapAroundPivot(elements []int, leftpoint int, rightpoint int) []int {
	pivotPoint := rand.Intn(rightpoint - leftpoint)
	pivot := elements[pivotPoint]
	pivotBoundary := pivotPoint + 1
	for i := pivotBoundary; i < rightpoint; i++ {
		if elements[i] < pivot && pivotBoundary > leftpoint+1 { // the only case when anything needs to happen. New e needs to be swapped && the pivot boundary has already moved
			pivotBoundary++
			elements[pivotBoundary], elements[i] = elements[i], elements[pivotBoundary]
		}
	}
	elements[pivotBoundary], elements[leftpoint] = elements[leftpoint], elements[pivotBoundary] // swap the pivot element into place
	return elements
}
