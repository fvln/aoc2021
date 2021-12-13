package day9

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string



func isLowPoint(matrix [][]int, height, width, i, j int) bool {
	value := matrix[i][j]

	if i > 0 {
		if value >= matrix[i-1][j] {
			return false
		}
	}
	if j > 0 {
		if value >= matrix[i][j-1] {
			return false
		}
	}
	if i + 1 < height {
		if value >= matrix[i+1][j] {
			return false
		}
	}
	if j + 1 < width {
		if value >= matrix[i][j+1] {
			return false
		}
	}

	fmt.Println("lowPoints at line", i, "column", j)
	return true
}

func step1(matrix [][]int, height int, width int) {
	lowPoints := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if isLowPoint(matrix, height, width, i, j) {
				lowPoints += (matrix[i][j] + 1)
			}
		}
	}
	fmt.Println(lowPoints)
}

func step2(matrix [][]int, height int, width int) {
	//for t:=1; t < 100000; t++ {
	//	runOneStep(matrix, height, width)
	//	if isSynced(matrix, height, width) {
	//		fmt.Println(t)
	//		return
	//	}
	//}
}

func Solve() {
	var matrix [][]int
	var height, width int

	for _, line := range strings.Split(data, "\n") {
		row := make([]int, len(line))
		width = len(line)
		for i, digit := range strings.Split(line, "") {
			row[i], _ = strconv.Atoi(digit)
		}
		matrix = append(matrix, row)
		height++
	}

	fmt.Println(height, width)
	step1(matrix, height, width)
	//step2(matrix, height, width)
}