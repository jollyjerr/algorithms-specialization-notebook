package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := medianMaintenance(loadData("./smallData.txt")); res != 9335 {
		t.Error(res)
	}
}

func TestOtherSimpleCase(t *testing.T) {
	if res := medianMaintenance(loadData("./otherData.txt")); res != 5174 {
		t.Error(res)
	}
}
