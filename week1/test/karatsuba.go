package main

import (
	"math"
)

func karatsuba(arg1 int64, arg2 int64) int64 {
	// Base case
	if arg1 < 10 || arg2 < 10 {
		return arg1 * arg2
	}

	// This is not bullet proof but works under the assignment constraints
	middle := findNumberOfDigits(arg1) / 2

	arg1High, arg1Low := getHighAndLowFromSplit(arg1, middle)
	arg2High, arg2Low := getHighAndLowFromSplit(arg2, middle)

	z0 := karatsuba(arg1Low, arg2Low)
	z1 := karatsuba((arg1Low + arg1High), (arg2Low + arg2High))
	z2 := karatsuba(arg1High, arg2High)

	return (z2 * int64(math.Pow(10, float64(2*middle)))) + (z1-z2-z0)*int64(math.Pow(10, float64(middle))) + z0
}

func findNumberOfDigits(number int64) int64 {
	digitCount := 0
	for number > 0 {
		digitCount++
		number = number / 10
	}
	return int64(digitCount)
}

func getHighAndLowFromSplit(num int64, digits int64) (int64, int64) {
	divisor := int64(math.Pow(10, float64(digits)))
	return num / divisor, num % divisor
}
