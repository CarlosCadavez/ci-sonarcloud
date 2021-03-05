package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(4, 5)

	if result != 9 {
		t.Error("The result must be 9")
	}
}
