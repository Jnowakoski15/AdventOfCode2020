package main

import "testing"

func Test1(t *testing.T) {
	exp := 5
	rez := solvePart1("testFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func Test2(t *testing.T) {
	exp := 8
	rez := solvePart2("testFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestReal1(t *testing.T) {
	exp := 1331
	rez := solvePart1("mainFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestReal2(t *testing.T) {
	exp := 1121
	rez := solvePart2("mainFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}
