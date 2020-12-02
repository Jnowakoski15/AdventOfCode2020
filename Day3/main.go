package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const input = "mainFile.txt"
const open = '.'
const tree = '#'

func main() {
	part1 := findTrees(input)
	part2 := findPart2Trees(input)
	fmt.Printf("Part1 Found %v trees with cheap toboggan\n", part1)
	fmt.Printf("Part2 Result is %v\n", part2)
}

func findPart2Trees(input string) uint64 {
	chanSize := 5
	result := uint64(1)
	c1 := make(chan int, chanSize)
	lines := readFile(input)
	mountain := build2DArray(lines)
	go channelTobogganTreeEncounter(mountain, 1, 1, c1)
	go channelTobogganTreeEncounter(mountain, 1, 3, c1)
	go channelTobogganTreeEncounter(mountain, 1, 5, c1)
	go channelTobogganTreeEncounter(mountain, 1, 7, c1)
	go channelTobogganTreeEncounter(mountain, 2, 1, c1)

	for i := 0; i < chanSize; i++ {
		chanRez := <-c1
		fmt.Printf("Channel got value : %v\n", chanRez)
		result *= uint64(chanRez)
	}
	return result
}

func findTrees(fileName string) int {
	lines := readFile(fileName)
	mountain := build2DArray(lines)
	treesFound := cheapTobogganTreeEncounterCount(mountain)
	return treesFound

}

func channelTobogganTreeEncounter(mnt [][]rune, down int, right int, c1 chan int) {
	c1 <- variableTobogganTreeEncounter(mnt, down, right)
}

func variableTobogganTreeEncounter(mnt [][]rune, down int, right int) int {
	j := right
	trees := 0
	rowMaxLen := len(mnt[0])
	for i := 1; i < len(mnt); i += down {
		row := mnt[i]
		val := row[j]

		if val == tree {
			trees++
		}

		j += right
		if j >= rowMaxLen {
			j = j - rowMaxLen
		}
	}
	fmt.Printf("Slope down: %v, right: %v Trees Found: %v\n", down, right, trees)
	return trees
}

func cheapTobogganTreeEncounterCount(mnt [][]rune) int {
	return variableTobogganTreeEncounter(mnt, 1, 3)
}

func build2DArray(lines []string) [][]rune {
	output := make([][]rune, len(lines))
	for i, line := range lines {
		lineArr := make([]rune, len(line))
		for i, r := range line {
			lineArr[i] = r
			//lineArr = append(lineArr, r)
		}
		output[i] = lineArr
	}
	return output
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
