package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := prim(loadData("./smallData.txt")); res != 7 {
		t.Error(res)
	}
}
