package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := kosarajuTwoPass(loadData("./smallerData.txt")); !equal(res, []int{6, 4, 4, 4, 6, 6, 9, 9, 9}) {
		t.Error(res)
	}
}

func TestBiggerCase(t *testing.T) {
	if res := kosarajuTwoPass(loadData("./smallData.txt")); !equal(res, []int{1, 4, 4, 4, 6, 6, 12, 12, 12, 12, 12, 12}) {
		t.Error(res)
	}
}

func equal(res, solution []int) bool {
	for i, v := range solution {
		if res[i] != v {
			return false
		}
	}
	return true
}
