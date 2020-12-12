package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const fileName = "mainFile.txt"

const floor = '.'
const empty = 'L'
const occupied = '#'

func main() {
	part1 := solvePart1(fileName)
	part2 := solvePart2(fileName)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln:%v\n", part1, part2)
}

func solvePart2(fileName string) int {
	lines := readFile(fileName)
	board := buildArrayFromInput(lines)
	output := cycleContinuouslyPart2(&board)
	return output
}

func cycleContinuouslyPart2(board *[][]rune) int {
	newBoard := board
	change := false
	for true {
		printBoard(newBoard)
		newBoard, change = cycleThroughPart2(newBoard)
		if !change {
			break
		}
	}

	return countOccupiedSeats(newBoard)
}

func cycleThroughPart2(board *[][]rune) (*[][]rune, bool) {
	newBoard := make([][]rune, len(*board))
	madeChange := false
	for i, array := range *board {
		tmp := make([]rune, len(array))
		newBoard[i] = tmp
		for j, seat := range array {
			val := getNewFieldFromBoardPart2(i, j, board)
			newBoard[i][j] = val
			if val != seat {
				madeChange = true
			}
		}
	}
	return &newBoard, madeChange
}

func getNewFieldFromBoardPart2(i int, j int, boardPtr *[][]rune) rune {
	board := *boardPtr
	testSeat := board[i][j]

	if testSeat == floor {
		return floor
	}
	surrounding := getSurroundingPart2(i, j, boardPtr)

	if testSeat == occupied {
		return getNewOccupiedSeat(surrounding, 5)
	}

	if testSeat == empty {
		return getNewEmptySeat(surrounding)
	}
	return ' '
}

func getSurroundingPart2(i int, j int, boardPtr *[][]rune) []rune {
	surrounding := []rune{}
	surrounding = append(surrounding, goDirection(i, j, boardPtr, -1, -1))
	surrounding = append(surrounding, goDirection(i, j, boardPtr, -1, 0))
	surrounding = append(surrounding, goDirection(i, j, boardPtr, -1, 1))
	surrounding = append(surrounding, goDirection(i, j, boardPtr, 0, -1))
	surrounding = append(surrounding, goDirection(i, j, boardPtr, 0, 1))
	surrounding = append(surrounding, goDirection(i, j, boardPtr, 1, 0))
	surrounding = append(surrounding, goDirection(i, j, boardPtr, 1, -1))
	surrounding = append(surrounding, goDirection(i, j, boardPtr, 1, 1))

	return surrounding
}

func goDirection(i int, j int, boardPtr *[][]rune, directI int, directJ int) rune {
	board := *boardPtr
	i = i + directI
	j = j + directJ
	for i >= 0 && j >= 0 && i < len(board) && j < len(board[i]) {
		seat := board[i][j]
		if seat != floor {
			return seat
		}
		i = i + directI
		j = j + directJ
	}
	return ' '
}

func solvePart1(fileName string) int {
	lines := readFile(fileName)
	board := buildArrayFromInput(lines)
	output := cycleContinuously(&board)
	return output
}

func cycleContinuously(board *[][]rune) int {
	newBoard := board
	change := false
	for true {
		newBoard, change = cycleThrough(newBoard)
		if !change {
			break
		}
	}

	return countOccupiedSeats(newBoard)
}

func countOccupiedSeats(board *[][]rune) int {
	occupiedSeats := 0
	for _, array := range *board {
		for _, seat := range array {
			if seat == occupied {
				occupiedSeats++
			}
		}
	}
	return occupiedSeats
}

func cycleThrough(board *[][]rune) (*[][]rune, bool) {
	newBoard := make([][]rune, len(*board))
	madeChange := false
	for i, array := range *board {
		tmp := make([]rune, len(array))
		newBoard[i] = tmp
		for j, seat := range array {
			val := getNewFieldFromBoard(i, j, board)
			newBoard[i][j] = val
			if val != seat {
				madeChange = true
			}
		}
	}
	return &newBoard, madeChange
}

func printBoard(board *[][]rune) {
	for _, array := range *board {
		for _, field := range array {
			fmt.Print(string(field))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func getNewFieldFromBoard(i int, j int, boardPtr *[][]rune) rune {
	board := *boardPtr
	testSeat := board[i][j]

	if testSeat == floor {
		return floor
	}
	surrounding := getSurrounding(i, j, boardPtr)

	if testSeat == occupied {
		return getNewOccupiedSeat(surrounding, 4)
	}

	if testSeat == empty {
		return getNewEmptySeat(surrounding)
	}
	return ' '
}

func getNewOccupiedSeat(surrounding []rune, threshold int) rune {
	occupiedCount := 0
	for _, seat := range surrounding {
		if seat == occupied {
			occupiedCount++
			if occupiedCount >= threshold {
				return empty
			}
		}

	}
	return occupied
}

func getNewEmptySeat(surrounding []rune) rune {
	for _, seat := range surrounding {
		if seat == occupied {
			return empty
		}
	}
	return occupied
}

func getSurrounding(i int, j int, boardPtr *[][]rune) []rune {
	surrounding := []rune{}
	above := getAbove(i, j, boardPtr)
	below := getBelow(i, j, boardPtr)
	row := getRow(i, j, boardPtr)
	surrounding = append(surrounding, above...)
	surrounding = append(surrounding, below...)
	surrounding = append(surrounding, row...)

	return surrounding
}

func getAbove(i int, j int, boardPtr *[][]rune) []rune {
	board := *boardPtr
	maxj := len(board[i]) - 1
	min := j - 1
	max := j + 1
	output := []rune{}
	i--

	if i < 0 {
		return []rune{}
	}
	if min >= 0 {
		output = append(output, board[i][min])
	}

	if max <= maxj {
		output = append(output, board[i][max])
	}

	output = append(output, board[i][j])
	return output
}

func getBelow(i int, j int, boardPtr *[][]rune) []rune {
	board := *boardPtr
	maxj := len(board[i]) - 1

	min := j - 1
	max := j + 1
	output := []rune{}
	i++

	if i >= len(board) {
		return []rune{}
	}

	if min >= 0 {
		output = append(output, board[i][min])
	}

	if max <= maxj {
		output = append(output, board[i][max])
	}

	output = append(output, board[i][j])

	return output
}

func getRow(i int, j int, boardPtr *[][]rune) []rune {
	board := *boardPtr
	output := []rune{}
	maxj := len(board[i]) - 1
	left := j - 1
	right := j + 1

	if left >= 0 {
		output = append(output, board[i][left])
	}
	if right <= maxj {
		output = append(output, board[i][right])
	}

	return output
}

func buildArrayFromInput(lines []string) [][]rune {
	board := make([][]rune, len(lines))
	for i, line := range lines {
		tmp := make([]rune, len(line))
		board[i] = tmp
		for j, r := range line {
			board[i][j] = r
		}
	}
	return board
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
