package day13

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

const WIDTH = 2000
const HEIGHT = 2000



func foldMatrixY(matrix [][]bool, xmax *int, ymax *int, fold int) {
	fmt.Println("FoldY", fold, *xmax, *ymax)

	for y := 1; y <= *ymax - fold && fold -y >= 0; y++ {
		for x := 0; x <= *xmax; x++ {
			matrix[fold - y][x] = matrix[fold - y][x] || matrix[fold + y][x]
		}
	}
	*ymax = fold
}

func foldMatrixX(matrix [][]bool, xmax *int, ymax *int, fold int) {
	fmt.Println("FoldX", fold, *xmax, *ymax)

	for x := 1; x <= *xmax - fold && fold - x >= 0; x++ {
		for y := 0; y <= *ymax; y++ {
			matrix[y][fold - x] = matrix[y][fold - x] || matrix[y][fold + x]
		}
	}
	*xmax = fold
}

func step1(matrix [][]bool, xmax int, ymax int) {

	foldMatrixX(matrix, &xmax, &ymax, 655)

	dots := 0
	for y := 0; y <= ymax; y++ {
		for x := 0; x <= xmax; x++ {
			if matrix[y][x] {
				//fmt.Printf("X ")
				dots++
			} else {
				//fmt.Printf(". ")
			}
		}
		//fmt.Println()
	}

	//fmt.Println(dots)
}

func step2(matrix [][]bool, xmax int, ymax int) {

	foldMatrixX(matrix, &xmax, &ymax, 655)
	foldMatrixY(matrix, &xmax, &ymax, 447)
	foldMatrixX(matrix, &xmax, &ymax, 327)
	foldMatrixY(matrix, &xmax, &ymax, 223)
	foldMatrixX(matrix, &xmax, &ymax, 163)
	foldMatrixY(matrix, &xmax, &ymax, 111)
	foldMatrixX(matrix, &xmax, &ymax, 81)
	foldMatrixY(matrix, &xmax, &ymax, 55)
	foldMatrixX(matrix, &xmax, &ymax, 40)
	foldMatrixY(matrix, &xmax, &ymax, 27)
	foldMatrixY(matrix, &xmax, &ymax, 13)
	foldMatrixY(matrix, &xmax, &ymax, 6)

	dots := 0
	for y := 0; y <= ymax; y++ {
		for x := 0; x <= xmax; x++ {
			if matrix[y][x] {
				fmt.Printf("X ")
				dots++
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
}

func max(a int, b int) int {
	if a > b { return a }
	return b
}

func Solve() {
	matrix := make([][]bool, WIDTH)
	xmax := 0
	ymax := 0

	for i, _ := range matrix {
		matrix[i] = make([]bool, HEIGHT)
	}

	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		matrix[y][x] = true

		xmax = max(xmax, x)
		ymax = max(ymax, y)
	}

	step1(matrix, xmax, ymax)
	step2(matrix, xmax, ymax)
}