package day5

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func newLine(s string) line {
	re := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
	matches := re.FindAllStringSubmatch(s, -1)[0]

	x1, _ := strconv.Atoi(matches[1])
	x2, _ := strconv.Atoi(matches[3])
	y1, _ := strconv.Atoi(matches[2])
	y2, _ := strconv.Atoi(matches[4])

	res := line{
		x1, y1, x2, y2,
	}
	return res
}

func (s line) isHoriz() bool {
	return s.x1 == s.x2
}

func (s line) isVert() bool {
	return s.y1 == s.y2
}

func toPoint(x int, y int) int {
	return 1000*y + x
}

func min(a int, b int) int {
	if a < b { return a }
	return b
}

func max(a int, b int) int {
	if a > b { return a }
	return b
}

func (s line) toListOfPoints() []int {
	var result []int
	var xval, yval int

	if s.x1 == s.x2 {
		for y:=min(s.y1, s.y2); y <= max(s.y1, s.y2); y++ {
			result = append(result, toPoint(s.x1, y))
		}
	} else if s.y1 == s.y2 {
		for x:=min(s.x1, s.x2); x <= max(s.x1, s.x2); x++ {
			result = append(result, toPoint(x, s.y1))
		}
	} else {
		if s.x1 < s.x2 {
			xval = 1
		} else {
			xval = -1
		}

		if s.y1 < s.y2 {
			yval = 1
		} else {
			yval = -1
		}

		for i:=0; i<=max(s.x1, s.x2) - min(s.x1, s.x2); i++ {
			result = append(result, toPoint(s.x1 + (i*xval), s.y1 + (i*yval)))
		}
	}

	// fmt.Printf("%v -> %v\n", s, result)
	return result
}

func step1(lines []line) {
	nbIntersect := 0
	points := make([]int, 1000*1000, 1000*1000)


	for _, line := range lines {
		if line.isHoriz() || line.isVert() {
			listOfPoints := line.toListOfPoints()
			for _, point := range listOfPoints {
				points[point]++
			}
		}
	}

	for _, point := range points {
		if point > 1 {
			nbIntersect++
		}
	}

	fmt.Println(nbIntersect)
}

func step2(lines []line) {
	nbIntersect := 0
	points := make([]int, 1000*1000, 1000*1000)

	for _, line := range lines {
		listOfPoints := line.toListOfPoints()
		for _, point := range listOfPoints {
			points[point]++
		}
	}

	for _, point := range points {
		if point > 1 {
			nbIntersect++
		}
	}

	fmt.Println(nbIntersect)
}

func Solve() {
	var lines []line

	for _, l := range strings.Split(data, "\n") {
		lines = append(lines, newLine(l))
	}

	step1(lines)
	step2(lines)
}