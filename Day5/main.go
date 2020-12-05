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
	rowStep := 128
	colLower := 0
	colStep := 8

	for _, v := range line {
		if v == front || v == back {
			rowStep = rowStep >> 1
			if v == back {
				rowLower += rowStep
			}
		}
		if v == left || v == right {
			colStep = colStep >> 1
			if v == right {
				colLower += colStep
			}
		}
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
