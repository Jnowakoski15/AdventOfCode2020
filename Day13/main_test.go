package main

import "testing"

func Test1(t *testing.T) {
	exp := 295
	rez := solvePart1("testFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func Test2(t *testing.T) {
	exp := 1068781
	rez := solvePart2("testFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestSmall2(t *testing.T) {
	exp := 3417
	rez := solvePart2("smallTestFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}
