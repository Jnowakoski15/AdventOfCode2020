package main

import "testing"

func Test1(t *testing.T) {
	exp := 820
	rez := solvePart1("testFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}
