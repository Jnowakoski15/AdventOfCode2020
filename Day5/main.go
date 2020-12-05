package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const fileName = "mainFile.txt"

//FFFBBBFRRR

const front = 'F'
const back = 'B'
const right = 'R'
const left = 'L'

func main() {
	part1 := solvePart1(fileName)
	part2 := solvePart2(fileName)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln: %v\n", part1, part2)
}

func solvePart1(fileName string) int {
	output := 0
	lines := readFile(fileName)
	for _, line := range lines {
		ID := getSeatID(line)
		if ID > output {
			output = ID
		}
	}
	return output
}

func solvePart2(fileName string) int {
	lines := readFile(fileName)
	seatList := buildSeatingList(lines)
	ID := searchSeatsForMissing(seatList)
	return ID
}

func searchSeatsForMissing(seatList []int) int {
	sort.Ints(seatList)
	length := len(seatList)
	lastID := seatList[0]
	for i := 1; i < length; i++ {
		ID := seatList[i]
		if lastID+1 != ID {
			return ID - 1
		}
		lastID = ID
	}
	return 0
}

func buildSeatingList(lines []string) []int {
	seatList := []int{}
	for _, line := range lines {
		ID := getSeatID(line)
		seatList = append(seatList, ID)
	}
	return seatList
}

func getSeatID(line string) int {
	ID := 0
	rowLower := 0
	rowUpper := 127
	colUpper := 7
	colLower := 0

	for _, v := range line {
		if v == front || v == back {
			step := ((rowUpper - rowLower) + 1) / 2
			if v == front {
				rowUpper -= step
			}
			if v == back {
				rowLower += step
			}
		}

		if v == left || v == right {
			step := ((colUpper - colLower) + 1) / 2
			if v == left {
				colUpper -= step
			}
			if v == right {
				colLower += step
			}
		}
	}

	if rowLower != rowUpper {
		fmt.Printf("Match fail on row: %v | upper: %v lower %v\n", line, rowUpper, rowLower)
	}

	if colLower != colUpper {
		fmt.Printf("Match fail on col: %v| upper: %v lower %v\n", line, colLower, colUpper)
	}

	ID = rowLower*8 + colLower
	return ID
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
