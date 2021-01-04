package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if min, max := huffman(loadData("./smallData.txt")); min != 2 || max != 5 {
		t.Error(min, max)
	}
}

func TestMidCase(t *testing.T) {
	if min, max := huffman(loadData("./midData.txt")); min != 5 || max != 11 {
		t.Error(min, max)
	}
}
