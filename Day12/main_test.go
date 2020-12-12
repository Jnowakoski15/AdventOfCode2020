package main

import "testing"

func Test1(t *testing.T) {
	exp := 25
	rez := solvePart1("testFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func Test2(t *testing.T) {
	exp := 286
	rez := solvePart2("testFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}
