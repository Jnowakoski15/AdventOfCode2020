package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	count1, count2 := NumPassingPasswords("RealFile.txt")
	fmt.Println("Number of compliant passwords for part 1: " + strconv.Itoa(count1))
	fmt.Println("Number of compliant passwords for part 2: " + strconv.Itoa(count2))
}

//NumPassingPasswords determines if a password passes its policy requirement
func NumPassingPasswords(fileName string) (int, int) {
	part1Count := 0
	part2Count := 0
	bytesRead, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		if line == "\n" || line == "" {
			break
		}
		minMax, letter, password := getSperatedValues(line)
		min, max := getMinAndMax(minMax)
		passwordLetter := rune(letter[0])

		if isCompliantPart1(min, max, passwordLetter, password) {
			part1Count++
		}

		if isCompliantPart2(min, max, passwordLetter, password) {
			part2Count++
		}
	}

	return part1Count, part2Count
}

func isCompliantPart2(first int, second int, testRune rune, pass string) bool {

	length := len(pass)
	r1 := rune(pass[first-1])
	r2 := rune(pass[second-1])

	if r1 == r2 {
		return false
	}

	if first > length || second > length {
		return false
	}

	if r1 == testRune && r2 != testRune {
		return true
	}

	if r1 != testRune && r2 == testRune {
		return true
	}

	return false
}

func isCompliantPart1(min int, max int, testRune rune, pass string) bool {
	count := 0
	for _, r := range pass {
		if r == testRune {
			count++
		}
	}

	if count >= min && count <= max {
		return true
	}

	return false
}

func getSperatedValues(l string) (string, string, string) {
	split := strings.Split(l, " ")
	minMax := split[0]
	letter := split[1]
	password := split[2]

	return minMax, letter, password
}

func getMinAndMax(mm string) (int, int) {
	minMaxSplit := strings.Split(mm, "-")
	min, _ := strconv.Atoi(minMaxSplit[0])
	max, _ := strconv.Atoi(minMaxSplit[1])
	return min, max
}
