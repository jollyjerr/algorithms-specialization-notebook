package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := countUndirectedConnectivity(loadData("./smallData.txt")); res != 1 {
		t.Error(res)
	}
}
