package main

// Sort an array of distinct elements
// Array length may be odd or even

func mergeSort(args []int) []int {
	lengthOfArgs := len(args)

	// Base case
	if lengthOfArgs <= 1 {
		return args
	}

	// find a middle point
	breakpoint := lengthOfArgs / 2

	return merge(mergeSort(args[:breakpoint]), mergeSort(args[breakpoint:]))
}

func merge(left, right []int) []int {

	var (
		size = len(left) + len(right)
		i    = 0
		j    = 0
		ret  = make([]int, size, size)
	)

	for k := 0; k < size; k++ {
		// handle odd number of elements
		if i > len(left)-1 && j <= len(right)-1 {
			ret[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			ret[k] = left[i]
			i++
			// compare and sort by lowest element
		} else if left[i] < right[j] {
			ret[k] = left[i]
			i++
		} else {
			ret[k] = right[j]
			j++
		}
	}
	return ret
}
