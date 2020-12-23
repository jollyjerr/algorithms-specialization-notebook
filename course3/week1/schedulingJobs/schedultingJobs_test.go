package main

import "testing"

func TestSimpleDifference(t *testing.T) {
	if res := scheduleJobs(loadData("./smallData.txt"), true); res != 58550 {
		t.Error(res)
	}
}

func TestSimpleRatio(t *testing.T) {
	if res := scheduleJobs(loadData("./smallData.txt"), false); res != 57262 {
		t.Error(res)
	}
}

func TestSimpleDifference2(t *testing.T) {
	if res := scheduleJobs(loadData("./smallerData.txt"), true); res != 32780 {
		t.Error(res)
	}
}

func TestSimpleRatio2(t *testing.T) {
	if res := scheduleJobs(loadData("./smallerData.txt"), false); res != 32104 {
		t.Error(res)
	}
}
