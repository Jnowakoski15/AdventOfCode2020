package main

import "testing"

func Test1(t *testing.T) {
	exp := 71
	rez := solvePart1("testFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestReal2(t *testing.T) {
	exp := 2587271823407
	rez := solvePart2("mainFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}
