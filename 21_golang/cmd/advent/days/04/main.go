package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/katsuragawaa/advent-of-code-2021-go/cmd/advent/days/04/bingo"
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
	game := bingo.Start(parse())
	for i := range game.Numbers {
		for j := range game.Boards {
			board := game.Boards[j]
			board.Update(game.Numbers[i])
			if board.Bingo {
				fmt.Println("BINGO", game.Numbers[i])
				n, err := strconv.Atoi(game.Numbers[i])
				if err != nil {
					fmt.Println("Not a number")
				}
				return n * board.SumUncheck()
			}
		}
	}
	return 0
}

func part2() int {
	game := bingo.Start(parse())
	for i := range game.Numbers {
		for j := range game.Boards {
			if game.Boards[j].Bingo {
				continue
			}

			game.Boards[j].Update(game.Numbers[i])
			if game.Boards[j].Bingo {
				game.Winners++
			}

			if game.Winners >= len(game.Boards) {
				n, err := strconv.Atoi(game.Numbers[i])
				fmt.Printf("Last one, bingo %d, board %d\n", n, j)
				if err != nil {
					fmt.Println("Not a number")
				}
				return n * game.Boards[j].SumUncheck()
			}
		}
	}
	return 0
}

func parse() (values []string) {
	return strings.Split(input, "\n")
}
