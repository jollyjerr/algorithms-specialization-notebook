package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := tsat(buildGraph("./smallData.txt")); res != true {
		t.Error(res)
	}
}
