package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const fileName = "mainFile.txt"

type bag struct {
	name          string
	containsShiny bool
}

type rule struct {
	bagType string
	count   int
}

func main() {
	part1 := solvePart1(fileName)
	part2 := solvePart2(fileName)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln: %v\n", part1, part2)
}

func solvePart1(fileName string) int {
	lines := readFile(fileName)
	bagMap := buildBagMap(lines)
	output := howManyBagContainShinyGold(bagMap)
	return output
}

func solvePart2(fileName string) int {
	lines := readFile(fileName)
	bagMap := buildBagMap(lines)
	output := howManyBagInsideShinyGold(bagMap)
	return output
}

func howManyBagInsideShinyGold(bagMap map[string][]rule) int {
	count := traverseCountBags(bagMap, "shiny gold bag") - 1
	return count
}

func traverseCountBags(bagMap map[string][]rule, key string) int {
	rules := bagMap[key]
	count := 1
	for _, rule := range rules {
		if rule.count > 0 {
			count += (rule.count * traverseCountBags(bagMap, rule.bagType))
		}
	}
	return count
}

func howManyBagContainShinyGold(bagmap map[string][]rule) int {
	output := 0
	for k := range bagmap {
		containsShinyGold := traverse(bagmap, k, "shiny gold bag")
		if containsShinyGold {
			output++
			fmt.Printf("Bag: %v can hold shiny gold bag\n", k)
		}

	}
	return output
}

func traverse(bagmap map[string][]rule, key string, base string) bool {
	rules := bagmap[key]
	for _, rule := range rules {
		if rule.bagType == base {
			return true
		} else {
			foundBag := traverse(bagmap, rule.bagType, base)
			if foundBag {
				return foundBag
			}
		}
	}
	return false
}

func buildBagMap(lines []string) map[string][]rule {
	bagMap := map[string][]rule{}
	for _, line := range lines {
		bag, rule := getBagAndRules(line)
		bagMap[bag] = rule
	}
	return bagMap
}

func getBagAndRules(line string) (string, []rule) {
	rules := []rule{}
	line = strings.Trim(line, ".")
	split := strings.Split(line, "contain")
	ruleSplit := strings.Split(split[1], ",")

	for _, ruleLine := range ruleSplit {
		ruleLine = strings.Trim(ruleLine, " ")
		var tmpRule rule
		if ruleLine != "no other bags" {
			indexOfFirstSpace := strings.Index(ruleLine, " ")
			count, _ := strconv.Atoi(ruleLine[0:indexOfFirstSpace])
			bagType := ruleLine[indexOfFirstSpace+1:]
			if count != 1 {
				bagType = bagType[0 : len(bagType)-1]
			}
			tmpRule = rule{count: count, bagType: bagType}
		} else {
			tmpRule = rule{count: 0, bagType: ruleLine}
		}
		rules = append(rules, tmpRule)
	}
	key := split[0]
	key = strings.Trim(key, " ")
	key = key[0 : len(key)-1]

	return key, rules
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
