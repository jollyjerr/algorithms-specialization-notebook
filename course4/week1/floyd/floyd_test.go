package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res, negative := floyd(loadData("smallData.txt")); res != -130 || negative != false {
		t.Error(res)
	}
}

func TestNegativeCycle(t *testing.T) {
	if _, negative := floyd(loadData("negCycleCheck.txt")); !negative {
		t.Error(negative)
	}
}
