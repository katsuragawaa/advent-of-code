package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	a := func() (result int) {
		for _, line := range strings.Fields(string(input)) {
			result += 10 * int(line[strings.IndexAny(line, "123456789")]-'0')
			result += int(line[strings.LastIndexAny(line, "123456789")] - '0')
		}
		return
	}()
	fmt.Println(a)

	b := func() (result int) {
		r := strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
			"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e")
		for _, line := range strings.Fields(string(input)) {
			line = r.Replace(r.Replace(line))
			result += 10 * int(line[strings.IndexAny(line, "123456789")]-'0')
			result += int(line[strings.LastIndexAny(line, "123456789")] - '0')
		}
		return
	}()
	fmt.Println(b)
}
