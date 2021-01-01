package main

import "testing"

func TestSmallCase(t *testing.T) {
	if res := kClusterings(loadData("./smallData.txt"), 4, 32); res != 86 {
		t.Error(res)
	}
}

func TestVerySmallCase(t *testing.T) {
	if res := kClusterings(loadData("./verySmallData.txt"), 4, 8); res != 21 {
		t.Error(res)
	}
}
