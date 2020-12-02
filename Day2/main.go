package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const regEx = `(?P<min>\d+)\-(?P<max>\d+) (?P<letter>\w)\: (?P<pass>\w+)`
const file = "RealFile.txt"

func main() {
	count1, count2 := NumPassingPasswords(file)
	fmt.Println("Number of compliant passwords for part 1: " + strconv.Itoa(count1))
	fmt.Println("Number of compliant passwords for part 2: " + strconv.Itoa(count2))
}

//NumPassingPasswords determines if a password passes its policy requirement
func NumPassingPasswords(fileName string) (int, int) {
	part1Count := 0
	part2Count := 0
	lines := fetchFile(fileName)
	for _, line := range lines {
		if line == "\n" || line == "" {
			break
		}
		min, max, letter, password := captureData(line)
		if isCompliantPart1(min, max, letter, password) {
			part1Count++
		}

		if isCompliantPart2(min, max, letter, password) {
			part2Count++
		}
	}

	return part1Count, part2Count
}

func captureData(input string) (int, int, rune, string) {
	var regex = regexp.MustCompile(regEx)
	match := regex.FindStringSubmatch(input)
	result := make(map[string]string)

	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	min, _ := strconv.Atoi(result["min"])
	max, _ := strconv.Atoi(result["max"])
	letter := result["letter"]
	pass := result["pass"]
	return min, max, []rune(letter)[0], pass
}

func fetchFile(filename string) []string {
	bytesRead, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	return lines
}

func isCompliantPart2(first int, second int, testRune rune, pass string) bool {
	r1 := rune(pass[first-1])
	r2 := rune(pass[second-1])

	if r1 == r2 {
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
