package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const fileName = "mainFile.txt"

func main() {
	part1 := solvePart1(fileName)
	part2 := solvePart2(fileName)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln: %v \n", part1, part2)
}

func solvePart2(fileName string) int {
	ranges, myTicketString, nearbyTicketsString := readFile(fileName)
	nameRangeMap := getMapOfNameToPossibleInts(ranges)
	myTicket := getTickets(myTicketString)[0]
	nearTicketsList := getTickets(nearbyTicketsString)
	validNumMap := getValidNumMap(ranges)
	nearTicketsList = removeInvalidTickets(nearTicketsList, validNumMap)
	possiblitiesPerCol := getPossiblitiesPerCol(nameRangeMap, myTicket, nearTicketsList)
	fieldsPerCol := getFieldsPerColmn(possiblitiesPerCol)

	return calculateDeparture(fieldsPerCol, myTicket)

}

func removeInvalidTickets(nearTicketsList [][]int, validMap map[int]struct{}) [][]int {
	outputList := [][]int{}
	for _, ticket := range nearTicketsList {
		isValid := true
		for _, ticketNum := range ticket {
			_, ok := validMap[ticketNum]
			if !ok {
				isValid = false
			}
		}
		if isValid {
			outputList = append(outputList, ticket)
		}
	}
	return outputList
}

func calculateDeparture(fieldsPerCol map[int]string, myTicket []int) int {
	output := 1
	departure := "departure"

	for k, v := range fieldsPerCol {
		if strings.Contains(v, departure) {
			output *= myTicket[k]
		}
	}

	return output
}

func findSingleItemList(possiblitiesPerCol map[int]map[string]struct{}, skippedKeys map[int]struct{}) (int, string, bool) {

	for K, V := range possiblitiesPerCol {
		_, ok := skippedKeys[K]
		if !ok {
			if len(V) == 1 {
				for value := range V { //only 1 element to return
					return K, value, false
				}

			}
		}
	}
	return -1, "", true
}

func getFieldsPerColmn(possiblitiesPerCol map[int]map[string]struct{}) map[int]string {
	finalMap := make(map[int]string)
	skipMap := make(map[int]struct{})
	for {
		singelePos, singleVal, isComplete := findSingleItemList(possiblitiesPerCol, skipMap)
		if isComplete {
			break
		}
		skipMap[singelePos] = struct{}{}
		possiblitiesPerCol = reduceMap(possiblitiesPerCol, singleVal)
		finalMap[singelePos] = singleVal
	}
	fmt.Println(possiblitiesPerCol)
	return finalMap
}

func reduceMap(possiblitiesPerCol map[int]map[string]struct{}, singleVal string) map[int]map[string]struct{} {
	output := make(map[int]map[string]struct{})

	for k, v := range possiblitiesPerCol {
		if len(v) != 1 {
			tmpMap := make(map[string]struct{})
			for setVal := range v {
				if singleVal != setVal {
					tmpMap[setVal] = struct{}{}
				}
			}
			output[k] = tmpMap
			continue
		}
		output[k] = v
	}
	return output
}

func getPossiblitiesPerCol(nameRangeMap map[string]map[int]struct{}, myTicket []int, nearTicketList [][]int) map[int]map[string]struct{} {

	fieldSet := getPossibleFields(nameRangeMap)
	//nearTicketList = append(nearTicketList, myTicket)
	outputMap := make(map[int]map[string]struct{})
	for i := 0; i < len(nearTicketList[0]); i++ {
		tmpSet := getPossibleFields(nameRangeMap)
		for j := 0; j < len(nearTicketList); j++ {
			val := nearTicketList[j][i]
			for fieldName := range fieldSet {
				rangeSet := nameRangeMap[fieldName]
				_, ok := rangeSet[val]
				if !ok {
					_, tmpSetOk := tmpSet[fieldName]
					if tmpSetOk {
						delete(tmpSet, fieldName)
					}

				}
			}
		}
		outputMap[i] = tmpSet
	}

	return outputMap
}

func getPossibleFields(nameRangeMap map[string]map[int]struct{}) map[string]struct{} {
	keys := make(map[string]struct{})
	for k := range nameRangeMap {
		keys[k] = struct{}{}
	}
	return keys
}

func getMapOfNameToPossibleInts(input string) map[string]map[int]struct{} {
	outputMap := make(map[string]map[int]struct{})
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		tmpMap := make(map[int]struct{})
		match := getParams(`(?P<Name>.*)\: (?P<start1>\d+)-(?P<end1>\d+) or (?P<start2>\d+)-(?P<end2>\d+)`, line)
		name, _ := match["Name"]
		start1, _ := strconv.Atoi(match["start1"])
		end1, _ := strconv.Atoi(match["end1"])

		start2, _ := strconv.Atoi(match["start2"])
		end2, _ := strconv.Atoi(match["end2"])

		for i := start1; i <= end1; i++ {
			tmpMap[i] = struct{}{}
		}

		for i := start2; i <= end2; i++ {
			tmpMap[i] = struct{}{}
		}
		outputMap[name] = tmpMap

	}
	return outputMap
}

func solvePart1(fileName string) int {
	ranges, _, nearbyTicketsString := readFile(fileName)
	validNumMap := getValidNumMap(ranges)
	nearTicketsList := getTickets(nearbyTicketsString)
	invalidNums := getInvalidNums(nearTicketsList, validNumMap)
	return addList(invalidNums)
}

func addList(numList []int) int {
	output := 0
	for _, val := range numList {
		output += val
	}
	return output
}

func getInvalidNums(ticketList [][]int, validMap map[int]struct{}) []int {
	outputList := []int{}
	for _, ticket := range ticketList {
		for _, ticketNum := range ticket {
			_, ok := validMap[ticketNum]
			if !ok {
				outputList = append(outputList, ticketNum)
			}
		}
	}
	return outputList
}

func getTickets(input string) [][]int {
	lines := strings.Split(input, "\n")
	length := len(lines)
	output := make([][]int, length-1)

	//skip first line its the header of nearby tickets: or your tickets:
	for i, line := range lines {
		if i != 0 {
			output[i-1] = getTicket(line)
		}
	}
	return output
}

func getTicket(ticket string) []int {
	tmp := []int{}
	stringNumbers := strings.Split(ticket, ",")
	for _, stringNum := range stringNumbers {
		num, _ := strconv.Atoi(stringNum)
		tmp = append(tmp, num)
	}
	return tmp
}

func getValidNumMap(input string) map[int]struct{} {
	outputMap := make(map[int]struct{})
	lines := strings.Split(input, "\n")

	for _, line := range lines {

		match := getParams(`(?P<Name>\w+)\: (?P<start1>\d+)-(?P<end1>\d+) or (?P<start2>\d+)-(?P<end2>\d+)`, line)
		start1, _ := strconv.Atoi(match["start1"])
		end1, _ := strconv.Atoi(match["end1"])

		start2, _ := strconv.Atoi(match["start2"])
		end2, _ := strconv.Atoi(match["end2"])

		for i := start1; i <= end1; i++ {
			outputMap[i] = struct{}{}
		}

		for i := start2; i <= end2; i++ {
			outputMap[i] = struct{}{}
		}
	}
	return outputMap
}

func getParams(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}

func readFile(filename string) (string, string, string) {
	bytesRead, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n\n")

	return lines[0], lines[1], lines[2]
}
