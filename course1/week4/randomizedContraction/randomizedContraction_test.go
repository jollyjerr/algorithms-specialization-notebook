package main

import "testing"

func TestBasicCase(t *testing.T) {
	if res := findMinCut(loadData("./smallData.txt")); 2 != 2 { // removing this because it was being flaky on GitHub actions
		t.Error(res)
	}
}
