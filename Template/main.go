package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const fileName = "mainFile.txt"

func main() {
	part1 := solvePart1(fileName)
	fmt.Printf("Part1 Soln: %v", part1)
}

func solvePart1(fileName string) int {
	return 0
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
