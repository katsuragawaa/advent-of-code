package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed "input.txt"
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	fmt.Println("part one:", part1())
	fmt.Println("part two:", part2())
}

func part1() int {
	depths := parse()
	count := 0
	last := depths[0]
	for _, d := range depths[1:] {
		if last < d {
			count++
		}
		last = d
	}

	return count
}

func part2() int {
	depths := parse()
	count := 0
	lastSum := depths[0] + depths[1] + depths[2]
	for i := range depths[1:] {
		if i >= len(depths)-2 {
			break
		}
		sum := depths[i] + depths[i+1] + depths[i+2]
		if lastSum < sum {
			count++
		}
		lastSum = sum
	}

	return count
}

func parse() (values []int) {
	for _, l := range strings.Split(input, "\n") {
		val, err := strconv.Atoi(l)
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
		values = append(values, val)
	}

	return values
}
