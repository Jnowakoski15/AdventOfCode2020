package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const fileName = "mainFile.txt"

func main() {
	part1 := solvePart1(fileName)
	part2 := solvePart2(fileName)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln: %v\n", part1, part2)
}

func solvePart2(fileName string) int {
	lines := readFile(fileName)
	adapters := convertToIntList(lines)
	adapters = createPart2List(adapters)
	return countCombos(adapters)
}

func createPart2List(adapters []int) []int {
	adapters = append([]int{0}, adapters...)
	length := len(adapters)
	end := adapters[length-1] + 3
	adapters = append(adapters, end)
	return adapters
}

func countCombos(adapters []int) int {
	countsStartingAt := map[int]int{}

	var f func(starts int) int
	f = func(starts int) int {
		if value, ok := countsStartingAt[starts]; ok {
			return value
		}

		subInts := adapters[starts:]

		if len(subInts) <= 1 {
			return 1
		}

		val := subInts[0]
		withinThree := findIndexForChange(subInts, val)

		count := 0
		for _, index := range withinThree {
			fmt.Printf("Calling for start %v + index %v\n", starts, index)
			count += f(starts + index)
			fmt.Printf("Ending for start %v + index %v + count %v \n", starts, index, count)
		}
		countsStartingAt[starts] = count
		return count
	}

	return f(0)
}

func findIndexForChange(ints []int, val int) []int {
	idxs := []int{}
	for idx, i := range ints {
		if i > val && i <= val+3 {
			idxs = append(idxs, idx)
		}
	}
	return idxs
}

func solvePart1(fileName string) int {
	lines := readFile(fileName)
	adapters := convertToIntList(lines)
	oneJump, threeJump := findAdapterJumps(adapters)
	return oneJump * threeJump
}

func findAdapterJumps(adapters []int) (int, int) {
	oneJump := 0
	threeJump := 1

	for i := 0; i < len(adapters); i++ {
		var prev int
		if i == 0 {
			prev = 0
		} else {
			prev = adapters[i-1]
		}
		cur := adapters[i]
		diff := cur - prev

		switch diff {
		case 1:
			oneJump++
		case 2: //
		case 3:
			threeJump++
		default:
			fmt.Errorf("Jump was more then 3: %v -> %v", prev, cur)
		}
	}
	return oneJump, threeJump
}

func convertToIntList(lines []string) []int {
	adapters := []int{}
	for _, line := range lines {
		adapter, _ := strconv.Atoi(line)
		adapters = append(adapters, adapter)
	}
	sort.Ints(adapters)
	return adapters
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
