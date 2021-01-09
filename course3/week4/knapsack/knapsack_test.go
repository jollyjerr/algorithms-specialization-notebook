package main

import "testing"

func TestVerySmallCase(t *testing.T) {
	if res := Knapsack(loadData("./verySmallData.txt")); res != 147 {
		t.Error(res, 147)
	}
}

func TestSimpleCase(t *testing.T) {
	if res := Knapsack(loadData("./smallData.txt")); res != 478 {
		t.Error(res, 478)
	}
}

func TestMiddleCase(t *testing.T) {
	if res := Knapsack(loadData("./medData.txt")); res != 1618 {
		t.Error(res, 1618)
	}
}
