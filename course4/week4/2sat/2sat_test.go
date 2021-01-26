package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := tsat(buildGraphs("./smallData.txt")); res != false {
		t.Error(res)
	}
}
