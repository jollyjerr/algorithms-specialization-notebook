package main

import "testing"

func TestBasicCase(t *testing.T) {
	if res := rSelect([]int{1, 2, 3, 4, 5, 6, 7, 9, 8}, 4); res != 5 {
		t.Error(res)
	}
}

func TestAdvancedCase(t *testing.T) {
	if res := rSelect([]int{-4, 3, 5, -7, 2, 9}, 3); res != 3 {
		t.Error(res)
	}
}
