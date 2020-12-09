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
	part1 := solvePart1(fileName, 25)
	part2 := solvePart2(fileName, 25)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln: %v\n", part1, part2)
}

func solvePart2(fileName string, preamble int) uint64 {
	lines := readFile(fileName)
	seenList := []uint64{}
	for i, line := range lines {
		val, _ := strconv.ParseUint(line, 10, 64)
		if i > preamble-1 {
			if !foundInPrevious(val, seenList, preamble) {
				continuousAdd := getAddValuesList(val, seenList)
				return getHighAndLowMult(continuousAdd)
			}
		}
		seenList = append(seenList, val)
	}
	return 0
}

func getHighAndLowMult(list []uint64) uint64 {
	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })
	val1 := list[0]
	val2 := list[len(list)-1]
	return val1 + val2
}

func getAddValuesList(target uint64, seenList []uint64) []uint64 {
	start := 0
	current := uint64(0)

	for i := 0; i < len(seenList); i++ {
		current += seenList[i]
		if current > target {
			current = 0
			start++
			i = start
		}
		if current == target {
			return seenList[start:i]
		}
	}
	return nil
}

func solvePart1(fileName string, preamble int) uint64 {
	lines := readFile(fileName)
	seenList := []uint64{}
	for i, line := range lines {
		val, _ := strconv.ParseUint(line, 10, 64)
		if i > preamble-1 {
			if !foundInPrevious(val, seenList, preamble) {
				return val
			}
		}
		seenList = append(seenList, val)
	}
	return 0
}

func foundInPrevious(val uint64, seenList []uint64, preamble int) bool {
	length := len(seenList)
	seenMap := make(map[uint64]struct{})
	for i := length - 1; i >= length-preamble; i-- {
		tmp := seenList[i]
		x := val - tmp
		_, ok := seenMap[x]
		if ok {
			return true
		}
		seenMap[tmp] = struct{}{}
	}

	x := val - seenList[0]
	_, ok := seenMap[x]
	if ok {
		return true
	}
	return false
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
