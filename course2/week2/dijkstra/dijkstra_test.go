package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if result := dijkstra(loadData("./smallData.txt")); result[4].shortestPathFromS != 4 {
		t.Error(result)
	}
}
