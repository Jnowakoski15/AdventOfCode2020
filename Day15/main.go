package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const fileName = "mainFile.txt"

func main() {
	part1 := solvePart1(fileName)
	part2 := solvePart2(fileName)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln: %v", part1, part2)
}

type last2 struct {
	earlier int
	later   int
}

func (l *last2) getLastSeen() int {
	return l.later
}

func (l *last2) getValueSeen() int {
	return l.later - l.earlier
}

func (l *last2) pushRecentlySeen(input int) {
	if l.later == 0 {
		l.later = input
	} else {
		tmp := l.later
		l.later = input
		l.earlier = tmp
	}
}

func newLast2(input int) *last2 {
	return &last2{earlier: 0, later: input}
}

func solvePart1(fileName string) int {
	return solve(fileName, 2020)
}

func solvePart2(fileName string) int {
	return solve(fileName, 30000000)
}

func solve(fileName string, spokenEnd int) int {
	numMap := make(map[int]*last2)
	lines := readFile(fileName)
	numStrings := strings.Split(lines[0], ",")

	start := 0
	last := 0
	//Preload numbers
	for i, numString := range numStrings {
		num, _ := strconv.Atoi(numString)
		_, ok := numMap[num]
		if !ok {
			numMap[num] = newLast2(i + 1)
		}
		start = i
		last = num
		//		fmt.Printf("Spoken: %v Value: %v\n", i+1, num)
	}
	val := last
	for i := start + 1; i != spokenEnd; i++ {
		lastTuple := numMap[val]
		if lastTuple.earlier != 0 {
			val = lastTuple.getValueSeen()
			curTuple, ok := numMap[val]
			if !ok {
				numMap[val] = newLast2(i + 1)
			} else {
				curTuple.pushRecentlySeen(i + 1)
			}
		} else {
			zeroTuple, ok := numMap[0]
			if ok {
				zeroTuple.pushRecentlySeen(i + 1)
			} else {
				numMap[0] = newLast2(i + 1)
			}

			val = 0
		}
		//		fmt.Printf("Spoken: %v Value: %v\n", i+1, val)
	}
	return val
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
