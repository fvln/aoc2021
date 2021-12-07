package day2

import (
"bufio"
"os"
	"strconv"
	"strings"
)

type move struct {
	horizontalPos int
	depth int
}

func (m *move) add(n move) {
	m.horizontalPos += n.horizontalPos
	m.depth += n.depth
}

func loadData() []move {
	f, _ := os.Open("day2/data.txt")
	scanner := bufio.NewScanner(f)
	result := []move{}

	for scanner.Scan() {
		command := scanner.Text()
		parts := strings.Split(command, " ")
		if len(parts) != 2 {
			continue
		}

		value, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		switch parts[0] {
		case "forward":
			result = append(result, move{horizontalPos: value, depth: 0})
			break

		case "up":
			result = append(result, move{horizontalPos: 0, depth: -value})
			break

		case "down":
			result = append(result, move{horizontalPos: 0, depth: value})
			break
		}
	}
	return result
}

func step1(data []move) {
	totalMoves := move{0,0}

	for _, move:= range data {
		totalMoves.add(move)
	}

	println("Day1 - step 1: ", totalMoves.horizontalPos * totalMoves.depth)
}

func step2(data []move) {
	totalMoves := move{0,0}
	aim := 0

	for _, move:= range data {
		if move.horizontalPos != 0 {
			totalMoves.horizontalPos += move.horizontalPos
			totalMoves.depth += move.horizontalPos * aim
		} else {
			aim += move.depth
		}
	}

	println("Day1 - step 2: ", totalMoves.horizontalPos * totalMoves.depth)
}

func Solve() {
	data := loadData()
	step1(data)
	step2(data)
}