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

func part1() int64 {
	binaries := parse()
	var biGamma string
	var biEpsilon string
	length := len(binaries[0])

	for i := 0; i < length; i++ {
		zeros := 0
		ones := 0
		for _, bit := range binaries {
			if string(bit[i]) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if zeros > ones {
			biGamma += "0"
			biEpsilon += "1"
		} else {
			biGamma += "1"
			biEpsilon += "0"
		}
	}

	gamma, err := strconv.ParseInt(biGamma, 2, 64)
	if err != nil {
		fmt.Printf("Not possible to parse Gamma %s binaries into decimal\n", biGamma)
	}

	epsilon, err := strconv.ParseInt(biEpsilon, 2, 64)
	if err != nil {
		fmt.Printf("Not possible to parse Epsilon %s binaries into decimal\n", biEpsilon)
	}

	return gamma * epsilon
}

func part2() int64 {
	binaries := parse()

	biOxygen := oxygenRatings(binaries, 0)
	oxygen, err := strconv.ParseInt(biOxygen, 2, 64)
	if err != nil {
		fmt.Printf("Not possible to parse Oxygen %s binaries into decimal\n", biOxygen)
	}

	biCo2 := co2Rating(binaries, 0)
	co2, err := strconv.ParseInt(biCo2, 2, 64)
	if err != nil {
		fmt.Printf("Not possible to parse CO2 %s binaries into decimal\n", biOxygen)
	}

	return oxygen * co2
}

func oxygenRatings(ratings []string, i int) string {
	if len(ratings) == 1 {
		return ratings[0]
	}

	var zeros []string
	var ones []string
	for _, rate := range ratings {
		if rate[i:i+1] == "0" {
			zeros = append(zeros, rate)
		} else {
			ones = append(ones, rate)
		}
	}

	if len(zeros) > len(ones) {
		return oxygenRatings(zeros, i+1)
	}

	return oxygenRatings(ones, i+1)
}

func co2Rating(ratings []string, i int) string {
	if len(ratings) == 1 {
		return ratings[0]
	}

	var zeros []string
	var ones []string
	for _, rate := range ratings {
		if rate[i:i+1] == "0" {
			zeros = append(zeros, rate)
		} else {
			ones = append(ones, rate)
		}
	}

	if len(zeros) <= len(ones) {
		return co2Rating(zeros, i+1)
	}

	return co2Rating(ones, i+1)
}

func parse() (values []string) {
	return strings.Split(input, "\n")
}
