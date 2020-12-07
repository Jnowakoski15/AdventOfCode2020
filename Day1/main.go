package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	l := readInFileToIntArray("intList.txt")
	//l := []int{1721, 979, 366, 299, 675, 1456}
	i := 0
	j := 0
	k := 0
	target := 2020

	for i < len(l) {
		j = 0
		for j < len(l) {
			k = 0
			for k < len(l) {
				if i == j && j == k {
					k++
					continue
				}
				val := l[i] + l[j] + l[k]
				if val == target {
					mult := l[i] * l[j] * l[k]
					fmt.Printf("Answer: %v + %v + %v if you multiply them you get: %v\n", l[i], l[j], l[k], mult)
					return
				}
				k++
			}
			j++
		}
		i++
	}
}

func readInFileToIntArray(name string) []int {
	output := []int{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		output = append(output, val)
	}
	return output
}
