package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const fileName = "mainFile.txt"

const add = "acc"
const nop = "nop"
const jmp = "jmp"

//URL : https://adventofcode.com/2020/day/8
func main() {
	part1 := solvePart1(fileName)
	part2 := solvePart2(fileName)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln: %v\n", part1, part2)
}

func solvePart1(fileName string) int {
	lines := readFile(fileName)
	val, _ := runCode(lines)
	return val
}

func solvePart2(fileName string) int {
	lines := readFile(fileName)
	complete := false
	last := -1
	var val int
	for complete != true {
		val, last, complete = runModifiedCode(lines, last)
	}
	return val
}

func runModifiedCode(lines []string, lastLineModfied int) (int, int, bool) {
	accum := 0
	visitied := make(map[int]string)
	count := 0
	swapped := false
	modded := lastLineModfied
	for i := 0; i < len(lines); i++ {
		count++
		code := lines[i]
		_, ok := visitied[i]
		if ok {
			return accum, modded, false
		}
		visitied[i] = code
		op, increase, val := getOpVal(code)
		if (swapped == false && count > lastLineModfied) && (op == jmp || op == nop) {
			swapped = true
			modded = count
			switch op {
			case nop:
				op = jmp
			case jmp:
				op = nop
			}
		}

		switch op {
		case add:
			if increase {
				accum += val
			} else {
				accum -= val
			}
			continue
		case jmp:
			if increase {
				i += val - 1
			} else {
				i -= val + 1
			}
		}
	}
	return accum, modded, true
}

func runCode(lines []string) (int, bool) {
	accum := 0
	visitied := make(map[int]string)
	count := 0
	for i := 0; i < len(lines); i++ {
		count++
		code := lines[i]
		_, ok := visitied[i]
		if ok {
			return accum, false
		}
		visitied[i] = code
		op, increase, val := getOpVal(code)

		switch op {
		case add:
			if increase {
				accum += val
			} else {
				accum -= val
			}
			continue
		case jmp:
			if increase {
				i += val - 1
			} else {
				i -= val + 1
			}
		}
	}
	return accum, true
}

func getOpVal(code string) (string, bool, int) {
	split := strings.Split(code, " ")
	increase := strings.Contains(split[1], "+")
	valStr := split[1]
	valStr = valStr[1:]
	val, _ := strconv.Atoi(valStr)
	return split[0], increase, val
}

func readFile(filename string) []string {
	bytesRead, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	return lines
}
