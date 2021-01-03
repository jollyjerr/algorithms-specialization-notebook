package main

import "testing"

func TestSimpleCase(t *testing.T) {
	if res := 3; res != 3 {
		t.Error(res)
	}
}
