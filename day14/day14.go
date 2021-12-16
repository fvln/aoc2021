package day14

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data.txt
var data string


func step1(transfos map[string][]string, nbTemplates map[string]int, nbLetters map[string]int, nbIterations int) {
	for i:=0; i<nbIterations; i++ {

		diff := make(map[string]int, 0)

		for template, count := range nbTemplates {
			if count <= 0 {
				continue
			}

			newValues, ok := transfos[template]
			if ok {
				diff[newValues[0]] += count
				diff[newValues[1]] += count
				diff[template] -= count

				nbLetters[string(newValues[0][1])] += count
			}
			// fmt.Printf("Removed %d*%s, got %s + %s\n", count, template, newValues[0], newValues[1])
		}

		//fmt.Printf("Diff: %v\n", diff)
		for k, v := range diff {
			nbTemplates[k] += v
		}
		//fmt.Printf("Round %d: %v\n", i, nbTemplates)
		//fmt.Printf("Round %d: %v\n", i, nbLetters)
	}

	leastCommonCount := 999999999999999
	mostCommonCount := 0

	for _, count := range nbLetters {
		if count > mostCommonCount {
			mostCommonCount = count
		}
		if count > 0 && count < leastCommonCount {
			leastCommonCount = count
		}
	}

	fmt.Println(mostCommonCount, leastCommonCount, mostCommonCount - leastCommonCount)
}


func Solve() {
	transfos := make(map[string][]string)
	nbTemplates := make(map[string]int)
	nbLetters := make (map[string]int)

	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, " -> ")
		transfos[parts[0]] = []string{
			string(parts[0][0]) + parts[1],
			parts[1] + string(parts[0][1]),
		}
	}

	//input := "NNCB"
	input := "SNPVPFCPPKSBNSPSPSOF"
	for i:=0; i<len(input)-1; i++ {
		nbTemplates[input[i:i+2]]++
		nbLetters[string(input[i])]++
	}
	nbLetters[string(input[len(input) - 1])]++

	//fmt.Printf("%v\n", transfos)
	//fmt.Printf("%v\n", nbTemplates)

	//step1(transfos, nbTemplates, nbLetters, 10)
	step1(transfos, nbTemplates, nbLetters, 40)
}