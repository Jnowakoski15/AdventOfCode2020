package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const fileName = "mainFile.txt"

type ferry struct {
	direction int
	lat       int
	long      int
	waypoint  *ferryWaypoint
}

type ferryWaypoint struct {
	lat  int
	long int
}

type movement struct {
	x int
	y int
}

var turnMap = map[string]int{
	"L": -1,
	"R": 1,
}

var directionMap = map[int]string{
	0: "N",
	1: "E",
	2: "S",
	3: "W",
}

var movementMap = map[int]*movement{
	0: {x: 0, y: 1},
	1: {x: 1, y: 0},
	2: {x: 0, y: -1},
	3: {x: -1, y: 0},
}

func getNewFerry() *ferry {
	return &ferry{direction: 1, lat: 0, long: 0}
}

func getNewFerryWithWaypoint() *ferry {
	return &ferry{direction: 1, lat: 0, long: 0, waypoint: &ferryWaypoint{lat: 10, long: 1}}
}

func getNewFerryWaypoint() *ferry {
	return &ferry{lat: 0, long: 0}
}

func (f *ferry) changeDirection(turn string, val int) {
	turnVal := turnMap[turn]
	percentage := val / 90

	for i := 0; i < percentage; i++ {
		f.direction += turnVal

		if f.direction > 3 {
			f.direction = 0
		}
		if f.direction < 0 {
			f.direction = 3
		}
	}
}

var sinMap = map[int]int{
	90:  1,
	180: 0,
	270: -1,
	360: 0,
}

var cosMap = map[int]int{
	90:  0,
	180: -1,
	270: 0,
	360: 1,
}

func (f *ferry) changeWaypointPosition(turn string, val int) {
	s := sinMap[val]
	c := cosMap[val]
	lat := f.waypoint.lat
	long := f.waypoint.long
	if turn == "R" {
		f.waypoint.lat = lat*c + long*s
		f.waypoint.long = -lat*s + long*c
	}

	if turn == "L" {
		f.waypoint.lat = lat*c - long*s
		f.waypoint.long = lat*s + long*c
	}
}

func (f *ferry) doAction(cmd string, value int) {
	switch cmd {
	case "F":
		movement := movementMap[f.direction]
		f.lat += movement.x * value
		f.long += movement.y * value
	case "L":
		f.changeDirection(cmd, value)
	case "R":
		f.changeDirection(cmd, value)
	case "N":
		f.long += value
	case "E":
		f.lat += value
	case "S":
		f.long -= value
	case "W":
		f.lat -= value
	}
}

func (f *ferry) doWaypointAction(cmd string, value int) {
	switch cmd {
	case "F":
		f.lat += f.waypoint.lat * value
		f.long += f.waypoint.long * value
	case "L":
		f.changeWaypointPosition(cmd, value)
	case "R":
		f.changeWaypointPosition(cmd, value)
	case "N":
		f.waypoint.long += value
	case "E":
		f.waypoint.lat += value
	case "S":
		f.waypoint.long -= value
	case "W":
		f.waypoint.lat -= value
	}
}

func (f *ferry) getMahhattenValue() int {
	lat := f.lat
	long := f.long
	if lat < 0 {
		lat = lat * -1
	}
	if long < 0 {
		long = long * -1
	}
	return lat + long
}

func main() {
	part1 := solvePart1(fileName)
	part2 := solvePart2(fileName)
	fmt.Printf("Part1 Soln: %v\nPart2 Soln: %v\n", part1, part2)
}

func solvePart2(fileName string) int {
	inputs := readFile(fileName)
	ferry := getNewFerryWithWaypoint()

	for _, inputString := range inputs {
		action := inputString[0:1]
		number, _ := strconv.Atoi(inputString[1:])
		ferry.doWaypointAction(action, number)
	}

	return ferry.getMahhattenValue()
}

func solvePart1(fileName string) int {
	inputs := readFile(fileName)
	ferry := getNewFerry()

	for _, inputString := range inputs {
		action := inputString[0:1]
		number, _ := strconv.Atoi(inputString[1:])
		ferry.doAction(action, number)
	}

	return ferry.getMahhattenValue()
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
