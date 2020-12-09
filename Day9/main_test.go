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
