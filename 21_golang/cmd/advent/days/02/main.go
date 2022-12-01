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
	dirs := parse()
	x := 0
	y := 0
	for _, dir := range dirs {
		val, err := strconv.Atoi(dir[1])
		if err != nil {
			fmt.Println("Second value of the slice isn't a integer")
		}
		switch dir[0] {
		case "forward":
			x += val
		case "down":
			y += val
		case "up":
			y -= val
		default:
			fmt.Println("Not a valid direction")
		}
	}

	return x * y
}

func part2() int {
	dirs := parse()
	x := 0
	y := 0
	aim := 0
	for _, dir := range dirs {
		val, err := strconv.Atoi(dir[1])
		if err != nil {
			fmt.Println("Second value of the slice isn't a integer")
		}
		switch dir[0] {
		case "forward":
			x += val
			y += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		default:
			fmt.Println("Not a valid direction")
		}
	}

	return x * y
}

func parse() (values [][]string) {
	for _, l := range strings.Split(input, "\n") {
		val := strings.Split(l, " ")
		values = append(values, val)
	}

	return values
}
