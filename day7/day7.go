package day7

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

func min(a int, b int) int {
	if a < b { return a }
	return b
}

func max(a int, b int) int {
	if a > b { return a }
	return b
}

func moveToPos(data []int, target int) int {
	score := 0
	for _, pos := range data {
		score += max(target - pos, pos - target)
	}
	return score
}

func moveToPosStep2(data []int, target int, movingCosts []int) int {
	score := 0
	for _, pos := range data {
		cost := movingCosts[max(target - pos, pos - target)]
		//fmt.Printf("Move from pos %d to %d, cost=%d\n", pos, target, cost)
		score += cost
	}
	return score
}

func computeMovingCosts(max int) []int {
	result := make([]int, max+1)
	result[0] = 0

	for i:=1; i<=max; i++ {
		result[i] = result[i-1] + i
	}

	return result
}

func step1(data []int) {
	bestPos := 0
	bestScore := 999999

	for pos:=data[0]; pos<=data[len(data)-1]; pos++ {
		score := moveToPos(data, pos)
		if score < bestScore {
			bestScore = score
			bestPos = pos
		}
	}

	fmt.Println("Best position:", bestPos, "with score", bestScore)
}

func step2(data []int) {
	bestPos := 0
	bestScore := 999999999

	maxMovement := data[len(data)-1] - data[0] //slice is ordered
	movingCosts := computeMovingCosts(maxMovement)
	fmt.Printf("%v\n", movingCosts)

	for pos:=data[0]; pos<=data[len(data)-1]; pos++ {
		score := moveToPosStep2(data, pos, movingCosts)
		if score < bestScore {
			bestScore = score
			bestPos = pos
		}
	}

	fmt.Println("Best position:", bestPos, "with score", bestScore)
}

func Solve() {
	var horiz []int

	for _, l := range strings.Split(data, ",") {
		value, _ := strconv.Atoi(l)
		horiz = append(horiz, value)
	}

	sort.Slice(horiz, func(i, j int) bool {
		return horiz[i] < horiz[j]
	})

	//fmt.Printf("%v\n", horiz)
	step1(horiz)
	step2(horiz)
}