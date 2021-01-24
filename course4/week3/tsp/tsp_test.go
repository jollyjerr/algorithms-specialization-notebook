package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := tsp(loadData("./smallData.txt")); res != 23 {
		t.Error(res)
	}
}

func TestMiddleCase(t *testing.T) {
	if res := tsp(loadData("./medData.txt")); res != 683 {
		t.Error(res)
	}
}
