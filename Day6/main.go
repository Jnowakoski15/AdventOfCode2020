package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const fileName = "mainFile.txt"

// URL: https://adventofcode.com/2020/day/6
func main() {
	part1 := solvePart1(fileName)
	part2 := solvePart2(fileName)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln: %v\n", part1, part2)
}

func solvePart2(fileName string) int {
	lines := readFile(fileName)
	output := 0
	for _, line := range lines {
		persons := 1
		fmt.Println(line)
		visited := map[rune]int{}
		for _, char := range line {
			if char != '\n' {
				val, ok := visited[char]
				if !ok {
					visited[char] = 1
					continue
				}
				visited[char] = val + 1

			}
			if char == '\n' {
				persons++
			}
		}
		for _, v := range visited {
			if v == persons {
				output++
			}
		}

	}
	return output
}

func solvePart1(fileName string) int {
	lines := readFile(fileName)
	output := 0

	for _, line := range lines {
		fmt.Println(line)
		visited := map[rune]bool{}
		for _, char := range line {
			if char != '\n' {
				_, ok := visited[char]
				if !ok {
					output++
					visited[char] = true
				}
			}
		}
	}
	return output
}

func readFile(filename string) []string {
	bytesRead, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n\n")

	return lines
}
