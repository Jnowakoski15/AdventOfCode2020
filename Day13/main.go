package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const fileName = "mainFile.txt"

type bus struct {
	firstDepart int
	ID          int
}

type smartBus struct {
	pos   int
	ID    int
	delta int
}

func newSmartBus(pos int, ID int, delta int) *smartBus {
	return &smartBus{pos: 0, ID: ID, delta: delta}
}

func newBus(startTime int, startID int) bus {
	start := 0
	for i := 0; i <= (startTime + startID); i++ {
		if i%startID == 0 {
			start = i
		}
	}

	return bus{firstDepart: start, ID: startID}
}

func (b *smartBus) checkAndDepart(time int) bool {
	if (time % b.ID) == 0 {
		b.pos = time
		return true
	}
	return false
}

func main() {
	part1 := solvePart1(fileName)
	part2 := solvePart2(fileName)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln:%v\n", part1, part2)
}

func solvePart2(fileName string) int {
	lines := readFile(fileName)
	time, _ := strconv.Atoi(lines[0])
	busIds := strings.Split(lines[1], ",")

	buses := []*smartBus{}
	max := 0
	for i, busID := range busIds {
		if busID != "x" {
			busID, _ := strconv.Atoi(busID)
			if busID > max {
				max = busID
			}
			buses = append(buses, newSmartBus(time, busID, i))
		}
	}

	return findAlignedBuses(time, buses)
}

func findAlignedBuses(start int, buses []*smartBus) int {
	time := start
	for {
		timeIter := 1
		valid := true
		for _, bus := range buses {
			if (time+bus.delta)%bus.ID != 0 {
				valid = false
				break
			}
			timeIter *= bus.ID
		}
		if valid {
			return time
		}
		time += timeIter
	}

}

func solvePart1(fileName string) int {
	lines := readFile(fileName)
	time, _ := strconv.Atoi(lines[0])
	busIds := strings.Split(lines[1], ",")

	buses := []bus{}

	for _, busId := range busIds {
		if busId != "x" {
			bus, _ := strconv.Atoi(busId)
			buses = append(buses, newBus(time, bus))
		}
	}

	return findEarliestBusResult(time, buses)
}

func findEarliestBusResult(depart int, buses []bus) int {
	lowest := math.MaxInt64
	ID := 0
	for _, bus := range buses {
		if bus.firstDepart < lowest {
			lowest = bus.firstDepart
			ID = bus.ID
		}
	}
	return (lowest - depart) * ID
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
