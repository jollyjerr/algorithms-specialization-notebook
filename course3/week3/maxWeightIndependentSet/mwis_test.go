package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := mwis(loadData("./smallData.txt")); res[0] != 1 {
		t.Error(res)
	}
}
