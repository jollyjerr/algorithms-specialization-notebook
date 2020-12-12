package main

import "testing"

func TestBasicCase(t *testing.T) {
	if res := findMinCut(loadData("./smallData.txt")); res != 2 {
		t.Error(res)
	}
}
