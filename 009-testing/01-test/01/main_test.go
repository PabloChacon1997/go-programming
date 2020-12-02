package main

import "testing"

func TestMiSuma(t *testing.T) {
	v := miSuma(2,4)
	if v != 6 {
		t.Error("Expected: ", 6, "Got: ", miSuma(2,4))
	}
}
