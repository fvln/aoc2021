package day10

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed data.txt
var data string

func checkLine(line string) int {
	stack := make([]string, 0)

	for _, letter := range strings.Split(line, "") {
		//fmt.Printf("%v\n", stack)

		switch letter {
		case "(":
			stack = append(stack, ")")
			continue
		case "[":
			stack = append(stack, "]")
			continue
		case "<":
			stack = append(stack, ">")
			continue
		case "{":
			stack = append(stack, "}")
			continue
		}

		if stack[len(stack)-1] == letter {
			stack = stack[:len(stack)-1]
			continue
		}

		switch letter {
		case ")": return 3
		case "]": return 57
		case "}": return 1197
		case ">": return 25137
		}
	}
	return 0
}

func completeLine(line string) int {
	stack := make([]string, 0)
	score := 0

	for _, letter := range strings.Split(line, "") {
		//fmt.Printf("%v\n", stack)

		switch letter {
		case "(":
			stack = append(stack, ")")
			continue
		case "[":
			stack = append(stack, "]")
			continue
		case "<":
			stack = append(stack, ">")
			continue
		case "{":
			stack = append(stack, "}")
			continue
		}

		if stack[len(stack)-1] == letter {
			stack = stack[:len(stack)-1]
			continue
		} else {
			// mismatch
			return 0
		}
	}

	fmt.Printf("Missing: %v\n", stack)

	for i:=len(stack)-1; i>= 0; i-- {
		score = 5 * score
		switch stack[i] {
		case ")": score += 1
		case "]": score += 2
		case "}": score += 3
		case ">": score += 4
		}
	}

	fmt.Println(score)
	return score
}


func step1(lines []string) {
	score := 0

	for _, line := range lines {
		newScore := checkLine(line)
		// fmt.Println(newScore)
		score += newScore
	}

	fmt.Println(score)
}

func step2(lines []string) {
	var scores []int

	for _, line := range lines {
		newScore := completeLine(line)
		// fmt.Println(newScore)
		if newScore > 0 {
			scores = append(scores, newScore)
		}
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	fmt.Println(scores[(len(scores)-1) / 2])
}

func Solve() {
	lines := strings.Split(data, "\n")

	step1(lines)
	step2(lines)
}