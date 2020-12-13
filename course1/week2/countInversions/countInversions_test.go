package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if _, res := countInversions([]int{1, 3, 5, 2, 4, 6}); res != 3 {
		t.Error(res)
	}
}

func TestComplexCase(t *testing.T) {
	if _, res := countInversions([]int{1, 4, 3, 6, 78, 9, 6, 4, 3, 2, 4, 6, 7, 8, 9, 0, 2, 3, -1, 5, 4, 77, 6}); res != 131 {
		t.Error(res)
	}
}
