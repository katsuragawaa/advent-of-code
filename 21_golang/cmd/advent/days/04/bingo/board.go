package bingo

import (
	"strconv"
	"strings"
)

type Game struct {
	Numbers []string
	Boards  []Board
	Winners int
}

type Board struct {
	Numbers [][]string
	Bingo   bool
}

func Start(input []string) Game {
	nums := input[0]
	boards := input[1:]
	return Game{
		draw(nums),
		getBoards(boards),
		0,
	}
}

func draw(strN string) []string {
	return strings.Split(strN, ",")
}

func getBoards(input []string) []Board {
	var numbers [][]string
	var boards []Board
	i := 0
	for _, line := range input {
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "  ", " ")
		if line == "" {
			i = 0
			continue
		}

		row := strings.Split(line, " ")
		numbers = append(numbers, row)

		if i >= 4 {
			boards = append(boards, Board{numbers, false})
			numbers = nil
		}
		i++
	}
	return boards
}

func (b *Board) Update(n string) {
	for i, row := range b.Numbers {
		for j, num := range row {
			if num == n {
				b.Numbers[i][j] = "x"
				if b.check(j) {
					b.Bingo = true
				}
			}
		}
	}
}

func (b *Board) SumUncheck() int {
	acc := 0
	for _, row := range b.Numbers {
		for _, num := range row {
			n, err := strconv.Atoi(num)
			if err != nil {
				continue
			}
			acc += n
		}
	}
	return acc
}

func (b *Board) check(i int) bool {
	col := make([]string, 0, 5)
	for _, row := range b.Numbers {
		if onlyXs(row) {
			return true
		}
		col = append(col, row[i])
	}

	return onlyXs(col)
}

func onlyXs(a []string) bool {
	b := []string{"x", "x", "x", "x", "x"}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
