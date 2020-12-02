package main

import "testing"

func TestPart1(t *testing.T) {
	const input = "testFile.txt"
	rez := findTrees(input)
	exp := 7
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}
}

func TestRealPart1(t *testing.T) {
	const input = "mainFile.txt"
	rez := findTrees(input)
	exp := 171
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

func TestRealPart2(t *testing.T) {
	const input = "mainFile.txt"
	rez := findPart2Trees(input)
	exp := uint64(1206576000)
	if rez != exp {
		t.Errorf("Expected: %v but result was %v\n", exp, rez)
	}

}
