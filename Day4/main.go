package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

//URL : https://adventofcode.com/2020/day/4
const fileName = "mainFile.txt"

var mandatoryField = map[string]bool{
	"byr": true,
	"iyr": true,
	"eyr": true,
	"hgt": true,
	"hcl": true,
	"ecl": true,
	"pid": true,
}

func isInBounds(test string, min int, max int) bool {
	val, err := strconv.Atoi(test)
	if err != nil {
		return false
	}
	if val >= min && val <= max {
		return true
	}

	return false
}

//four digits; at least 1920 and at most 2002.
var birthYearTest = func(test string) bool {
	return isInBounds(test, 1920, 2002)
}

//four digits; at least 2010 and at most 2020.
var issueYearTest = func(test string) bool {
	return isInBounds(test, 2010, 2020)
}

//four digits; at least 2020 and at most 2030.
var experationYearTest = func(test string) bool {
	return isInBounds(test, 2020, 2030)
}

//followed by either cm or in:
//If cm, the number must be at least 150 and at most 193.
//If in, the number must be at least 59 and at most 76.
var heightTest = func(test string) bool {
	cleanString := test
	if strings.Contains(test, "cm") {
		cleanString = test[0 : len(test)-2]
		return isInBounds(cleanString, 150, 193)
	}
	if strings.Contains(test, "in") {
		cleanString = test[0 : len(test)-2]
		return isInBounds(cleanString, 59, 76)
	}
	return false
}

//a # followed by exactly six characters 0-9 or a-f.
//Example: hcl:#888785
var hairColorTest = func(test string) bool {
	var rex = `^#([a-fA-F0-9]{6})$`
	var regex = regexp.MustCompile(rex)
	return regex.Match([]byte(test))
}

//(Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
var eyeColorTest = func(test string) bool {
	eyeColorMap := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	_, ok := eyeColorMap[test]
	return ok
}

//(Passport ID) - a nine-digit number, including leading zeroes.
var passportIDTest = func(test string) bool {
	var rex = `^\d{9}$`
	var regex = regexp.MustCompile(rex)
	return regex.Match([]byte(test))
}

var strictMandatoryFieldMap = map[string]func(string) bool{
	"byr": birthYearTest,
	"iyr": issueYearTest,
	"eyr": experationYearTest,
	"hgt": heightTest,
	"hcl": hairColorTest,
	"ecl": eyeColorTest,
	"pid": passportIDTest,
	"cid": func(in string) bool { return true },
}

func main() {
	part1 := part1Evaluation(fileName)
	part2 := part2Evaluation(fileName)
	fmt.Printf("Part1: %v\nPart2: %v\n", part1, part2)
}

func part2Evaluation(fileName string) int {
	lines := readFile(fileName)
	approvedCount := 0
	for _, line := range lines {
		passed := passesAllPassportChecks(line)
		if passed {
			approvedCount++
		}
	}

	return approvedCount
}

func part1Evaluation(fileName string) int {
	lines := readFile(fileName)
	approvedCount := 0
	for _, line := range lines {
		passed := containsAllMandatoryFields(line)
		if passed {
			approvedCount++
		}
	}

	return approvedCount
}

func passesAllPassportChecks(input string) bool {
	arrOfKeyValue := strings.Fields(input)

	if !containsAllMandatoryFields(input) {
		fmt.Println("Didn't Pass Mandatory fields: " + input)
		return false
	}

	for _, v := range arrOfKeyValue {
		vArray := strings.Split(v, ":")
		key := vArray[0]
		value := vArray[1]
		testFunc, ok := strictMandatoryFieldMap[key]
		if !ok {
			return false
		}
		if !testFunc(value) {
			return false
		}

	}
	return true
}

func containsAllMandatoryFields(input string) bool {
	for k := range mandatoryField {
		if !strings.Contains(input, k) {
			return false
		}
	}
	return true
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
