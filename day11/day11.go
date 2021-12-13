package day11

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

const DETONATED_VALUE = -999999

func min(a int, b int) int {
	if a < b { return a }
	return b
}

func max(a int, b int) int {
	if a > b { return a }
	return b
}

func increaseNearbyValues(matrix [][]int, height, width, i, j int) {
	if i > 0 {
		if j > 0 { matrix[i-1][j-1]++ }
		matrix[i-1][j]++
		if j + 1 < width { matrix[i-1][j+1]++ }
	}

	if j > 0 { matrix[i][j-1]++ }
	if j + 1 < width { matrix[i][j+1]++ }

	if i + 1 < height {
		if j > 0 { matrix[i+1][j-1]++ }
		matrix[i+1][j]++
		if j + 1 < width { matrix[i+1][j+1]++ }
	}
}

func runOneStep(matrix [][]int, height, width int) int {
	nbDetonations := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			matrix[i][j]++
		}
	}

	for flashPending := true; flashPending == true; {
		flashPending = false
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if matrix[i][j] >= 10 {
					matrix[i][j] = DETONATED_VALUE
					increaseNearbyValues(matrix, height, width, i, j)
					nbDetonations++
				}
			}
		}

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if matrix[i][j] >= 10 {
					flashPending = true
				}
			}
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if matrix[i][j] < 0 {
				matrix[i][j] = 0
			}
		}
	}

	for i := 0; i < height; i++ {
		fmt.Printf("%v\n", matrix[i])
	}
	fmt.Printf("\n")

	return nbDetonations
}

func step1(matrix [][]int, height int, width int) {
	nbDetonations := 0
	for t:=0; t < 100; t++ {
		nbDetonations += runOneStep(matrix, height, width)
	}

	fmt.Println(nbDetonations)
}

func isSynced(matrix [][]int, height int, width int) bool {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if matrix[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func step2(matrix [][]int, height int, width int) {
	for t:=1; t < 100000; t++ {
		runOneStep(matrix, height, width)
		if isSynced(matrix, height, width) {
			fmt.Println(t)
			return
		}
	}

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

	//step1(matrix, height, width)
	step2(matrix, height, width)
	//step2(lines)
}