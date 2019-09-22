package main

import "testing"

func TestVariables(t *testing.T) {
	if col[0] != [9]int{0, 9, 18, 27, 36, 45, 54, 63, 72} {
		t.Fail()
	}
	if row[0] != [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
		t.Fail()
	}
}
