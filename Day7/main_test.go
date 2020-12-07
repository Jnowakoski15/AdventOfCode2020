package main

import "testing"

func Test1(t *testing.T) {
	exp := 4
	rez := solvePart1("testFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func Test2(t *testing.T) {
	exp := 126
	rez := solvePart2("testFile2.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}
