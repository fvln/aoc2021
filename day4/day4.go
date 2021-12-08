package day4

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

type bingoBoard struct {
	rows [][]string
	cols [][]string
	remaining int
	winner bool
}

func newBingoBoard(lines []string) bingoBoard {
	board := bingoBoard{
		rows: make([][]string, 5),
		cols: make([][]string, 5),
		remaining: 0,
		winner: false,
	}

	for i:=0; i<5; i++ {
		numbers := strings.Fields(lines[i])
		board.rows[i] = numbers

		for j, number := range numbers {
			board.cols[j] = append(board.cols[j], number)
			value, _ := strconv.Atoi(number)
			board.remaining += value
		}
	}

	return board
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (b *bingoBoard) Print() {
	for _, row := range b.rows {
		fmt.Printf("%v\n", row)
	}
	fmt.Printf("Remaining: %d\n", b.remaining)
}

func (b *bingoBoard) Draw(drawn string) bool {
	if b.winner {
		return false
	}

	value, _ := strconv.Atoi(drawn)
	found := false

	for i, row := range b.rows {
		for j, number := range row {
			if number == drawn {
				b.rows[i] = remove(row, j)
				b.remaining -= value

				if len(b.rows[i]) == 0 {
					fmt.Println("Board has won with ROW number", drawn, b.remaining, b.remaining * value)
					b.winner = true
					return true
				}
				found = true
				break
			}
		}
	}

	if found {
		for i, col := range b.cols {
			for j, number := range col {
				if number == drawn {
					b.cols[i] = remove(col, j)
					if len(b.cols[i]) == 0 {
						fmt.Println("Board has won with COL number", drawn, b.remaining, b.remaining*value)
						b.winner = true
						return true
					}
					break
				}
			}
		}
	}

	return false
}

func step1(lines []string) {
	var boards []bingoBoard

	for i:=2; i<len(lines); i+=6 {
		newBoard := newBingoBoard(lines[i:i+5])
		boards = append(boards, newBoard)
	}

	for _, number := range strings.Split(lines[0], ",") {
		fmt.Println("Drawing number", number)
		for i, _ := range boards {
			if boards[i].Draw(number)  {
				boards[i].Print()
				fmt.Println("Winner is board #", i)
				return
			}

		}
	}
}

func step2(lines []string) {
	var boards []bingoBoard

	for i:=2; i<len(lines); i+=6 {
		newBoard := newBingoBoard(lines[i:i+5])
		boards = append(boards, newBoard)
	}

	for _, number := range strings.Split(lines[0], ",") {
		fmt.Println("")
		fmt.Println("Drawing number", number)
		for i, _ := range boards {
			if boards[i].Draw(number)  {
				boards[i].Print()

				fmt.Println("Winner board", i)
			}
		}
	}
}

func Solve() {
	lines := strings.Split(data, "\n")

	step1(lines)
	step2(lines)
}