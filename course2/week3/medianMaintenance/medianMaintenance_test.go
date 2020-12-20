package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := 2; res != 2 {
		t.Error(res)
	}
}
