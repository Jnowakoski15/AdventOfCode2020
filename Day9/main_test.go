package main

import "testing"

func Test1(t *testing.T) {
	exp := uint64(127)
	rez := solvePart1("testFile.txt", 5)

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func Test2(t *testing.T) {
	exp := uint64(62)
	rez := solvePart2("testFile.txt", 5)

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func Test1Real(t *testing.T) {
	exp := uint64(25918798)
	rez := solvePart1("mainFile.txt", 25)

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}

func Test2Real(t *testing.T) {
	exp := uint64(3340942)
	rez := solvePart2("mainFile.txt", 25)

	if rez != exp {
		t.Errorf("Rez : %v does not match expected: %v\n", rez, exp)
	}
}
