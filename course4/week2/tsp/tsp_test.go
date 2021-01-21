package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := 1; res != 1 {
		t.Error(res)
	}
}
