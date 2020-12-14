package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	mask := ""
	memoryMap := make(map[string]string)
	for _, line := range lines {
		newMask, memAddress, val := parseInput(line)
		// No Mask found so we must push memory into address
		if newMask == "" {
			val := getPaddedBitString(val)
			memAddress := getPaddedBitString(memAddress)
			memAddress = applyV2Mask(mask, memAddress)
			memLocations := getMemoryLocationsV2(memAddress)

			for _, location := range memLocations {
				memoryMap[location] = val
			}

		} else {
			mask = newMask
		}
	}
	return memoryToInt(memoryMap)
}

func getPaddedBitString(input string) string {
	val := decimalToBitString(input)
	val = padWithZero(val, 36)
	return val
}

func memoryToIntV2(memMap map[string]string) int {
	output := 0
	for _, v := range memMap {
		xVals := countXFound(v)
		memAddressVal := caluclateAllFloatingX(v, xVals)
		output += memAddressVal
	}
	return output
}

func getMemoryLocationsV2(address string) []string {
	xVals := countXFound(address)
	bitsArray := generateBitComboArray(xVals)
	var newAddress string
	addressList := []string{}
	for _, bitArray := range bitsArray {
		newAddress = address
		for _, bit := range bitArray {
			newAddress = strings.Replace(newAddress, "X", string(bit), 1)
		}
		addressList = append(addressList, newAddress)
	}
	return addressList
}

func caluclateAllFloatingX(value string, xCount int) int {
	bitsArray := generateBitComboArray(xCount)
	var newVal string
	outputVal := 0
	for _, bitArray := range bitsArray {
		newVal = value
		for _, bit := range bitArray {
			newVal = strings.Replace(newVal, "X", string(bit), 1)
		}
		outputVal += bitStringToInt(newVal)
	}
	return outputVal
}

func generateBitComboArray(xCount int) [][]rune {
	combosIntSize := int(math.Exp2(float64(xCount))) - 1
	combos := make([][]rune, combosIntSize+1)
	for i := 0; i <= combosIntSize; i++ {
		bits := strconv.FormatInt(int64(i), 2)
		bits = padWithZero(bits, xCount)
		combos[i] = []rune(bits)
	}
	return combos
}

func countXFound(val string) int {
	count := 0
	for _, ch := range val {
		if ch == 'X' {
			count++
		}
	}
	return count
}

func applyV2Mask(mask, val string) string {
	valArray := []rune(val)
	maskArray := []rune(mask)

	for i := len(val) - 1; i >= 0; i-- {
		maskVal := maskArray[i]
		if maskVal == '1' || maskVal == 'X' {
			valArray[i] = maskVal
		}
	}
	return string(valArray)
}

func solvePart1(fileName string) int {
	lines := readFile(fileName)
	mask := ""
	memoryMap := make(map[string]string)
	for _, line := range lines {
		newMask, memAdress, val := parseInput(line)
		// No Mask found so we must push memory into address
		if newMask == "" {
			val := decimalToBitString(val)
			val = padWithZero(val, 36)
			val = applyMask(mask, val)
			memoryMap[memAdress] = val
		} else {
			mask = newMask
		}
	}
	return memoryToInt(memoryMap)
}

func memoryToInt(memMap map[string]string) int {
	output := 0
	for _, v := range memMap {
		bitInt, _ := strconv.ParseInt(v, 2, 64)
		output += int(bitInt)
	}
	return output
}

func decimalToBitString(v string) string {
	intVal, _ := strconv.Atoi(v)
	return strconv.FormatInt(int64(intVal), 2)
}

func bitStringToInt(v string) int {
	count := 0.0
	output := 0
	for i := len(v) - 1; i >= 0; i-- {
		if v[i] == '1' {
			output += int(math.Exp2(count))
		}
		count++
	}
	return output
}

func padWithZero(v string, zeros int) string {
	bitSize := zeros
	delta := bitSize - len(v)
	zeroPad := []rune{}
	zeroRune := '0'
	for i := 0; i < delta; i++ {
		zeroPad = append(zeroPad, zeroRune)
	}

	vArray := []rune(v)
	outputAr := append(zeroPad, vArray...)
	return string(outputAr)
}

func applyMask(mask, val string) string {
	valArray := []rune(val)
	maskArray := []rune(mask)

	for i := len(val) - 1; i >= 0; i-- {
		maskVal := maskArray[i]
		if maskVal == '0' || maskVal == '1' {
			valArray[i] = maskVal
		}
	}
	return string(valArray)
}

func parseInput(input string) (string, string, string) {
	splitInput := strings.Split(input, " = ")

	if splitInput[0] == "mask" {
		return splitInput[1], "", ""
	}
	address := getMemAddress(splitInput[0])
	val := splitInput[1]
	return "", address, val
}

func getMemAddress(mem string) string {
	outputAr := []rune{}
	capture := false
	for _, ch := range mem {
		if capture && ch != ']' {
			outputAr = append(outputAr, ch)
		}
		if ch == '[' {
			capture = true
		}
		if ch == ']' {
			return string(outputAr)
		}
	}
	return string(outputAr)
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
