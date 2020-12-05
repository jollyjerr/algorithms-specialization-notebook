package main

import "testing"

func TestBaseCase(t *testing.T) {
	left := [][]int{{2}}
	right := [][]int{{3}}
	if res := strassen(left, right); !areSame([][]int{{6}}, res) {
		t.Error(res)
	}
}

func TestSimpleCase(t *testing.T) {
	left := [][]int{
		{1, 2},
		{3, 4},
	}
	right := [][]int{
		{5, 6},
		{7, 8},
	}
	if res := strassen(left, right); !areSame([][]int{{6}}, res) {
		t.Error(res)
	}
}

func TestComplexCase(t *testing.T) {
	left := [][]int{
		{1, 2, 3, 4},
		{3, 4, 5, 6},
		{5, 6, 7, 8},
		{7, 8, 9, 10},
	}
	right := [][]int{
		{5, 6, 7, 8},
		{7, 8, 9, 10},
		{9, 10, 11, 12},
		{11, 12, 13, 14},
	}
	if res := strassen(left, right); !areSame([][]int{{6}}, res) {
		t.Error(res)
	}
}

func areSame(left [][]int, right [][]int) bool {
	for i := 0; i < len(left); i++ {
		for j := 0; j < len(left); j++ {
			if left[i][j] != right[i][j] {
				return false
			}
		}
	}
	return true
}
