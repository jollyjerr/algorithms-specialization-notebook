package main

import "testing"

func TestBaseCase(t *testing.T) {
	if res := mergeSort([]int{1}); !equal(res, []int{1}) {
		t.Error()
	}
}

func TestSimpleCase(t *testing.T) {
	if res := mergeSort([]int{4, 6, 5}); !equal(res, []int{4, 5, 6}) {
		t.Error()
	}
}

func TestEdgeCase(t *testing.T) {
	if res := mergeSort([]int{3, 0, 4, -1, 5}); !equal(res, []int{-1, 0, 3, 4, 5}) {
		t.Error()
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
