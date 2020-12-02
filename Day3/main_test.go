package main

import "testing"

func testPart1(t *testing.T) {
	const input = "testFile.txt"
	rez := findTrees(input)
	exp := 7
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func TestPart2(t *testing.T) {
	const input = "testFile.txt"
	rez := findPart2Trees(input)
	exp := uint64(336)
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}
