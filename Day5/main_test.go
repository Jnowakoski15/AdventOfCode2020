package main

import "testing"

func Test1(t *testing.T) {
	exp := 820
	rez := solvePart1("testFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestVerifyPart1(t *testing.T) {
	exp := 998
	rez := solvePart1("mainFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func TestVerifyPart2(t *testing.T) {
	exp := 676
	rez := solvePart2("mainFile.txt")

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}
