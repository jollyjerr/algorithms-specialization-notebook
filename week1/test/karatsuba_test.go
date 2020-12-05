package main

import (
	"testing"
)

func TestSimpleCase(t *testing.T) {
	if res := karatsuba(11, 11); res != 121 {
		t.Error()
	}
}

func TestInterestingCase(t *testing.T) {
	if res := karatsuba(345498, 238471293746); res != 82391355046655508 {
		t.Error()
	}
}
