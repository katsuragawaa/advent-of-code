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

func parse() (values []string) {
	values = append(values, strings.Split(input, "\n")...)
	return values
}

func main() {
	fmt.Println("part one:", part1())
	fmt.Println("part two:", part2())
}

func part1() int {
	data := parse()

	sum := 0
	for _, line := range data {
		nums := []string{}
		for _, char := range strings.Split(line, "") {
			_, err := strconv.Atoi(char)
			if err != nil {
				continue
			}
			nums = append(nums, char)
		}

		first := nums[0]
		last := nums[len(nums)-1]

		v, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}

		sum += v
	}

	return sum
}

func part2() int {
	valid := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	keys := make([]string, len(valid))
	i := 0
	for k := range valid {
		keys[i] = k
		i++
	}

	sum := 0
	data := parse()
	for _, line := range data {
		first := ""
		last := ""
		a := ""
		for _, char := range strings.Split(line, "") {
			_, err := strconv.Atoi(char)
			if err == nil {
				first = char
				break
			} else {
				a += char
				found := false
				num := ""
				for _, k := range keys {
					if strings.Contains(a, k) {
						found = true
						num = valid[k]
						break
					}
				}
				if found {
					first = num
					break
				}
			}
		}

		a = ""
		for _, char := range strings.Split(reverse(line), "") {
			_, err := strconv.Atoi(char)
			if err == nil {
				last = char
				break
			} else {
				a += char
				found := false
				num := ""
				for _, k := range keys {
					if strings.Contains(a, reverse(k)) {
						found = true
						num = valid[k]
						break
					}
				}
				if found {
					last = num
					break
				}
			}
		}

		v, _ := strconv.Atoi(first + last)
		sum += v
	}

	return sum
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
