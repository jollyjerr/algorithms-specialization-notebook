package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if result := twoSum(loadData("smallData.txt"), 3, 10); result != 8 {
		t.Error(result)
	}
}
