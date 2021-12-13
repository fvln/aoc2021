package day8

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data.txt
var data string


func step2() {
	// Fuck it
}

func step1() {
	result := 0

	for _, entry := range strings.Split(data, "\n") {
		segments := strings.Fields(entry)
		//fmt.Printf("%v\n", segments)
		for i:=11; i<15; i++ {
			//fmt.Printf("lend=%d\n", len(segments[i]))
			switch len(segments[i]) {
			case 2:	// Digit 1
				result++
			case 3:	// Digit 7
				result++
			case 4:	// Digit 4
				result++
			case 7:	// Digit 8
				result++
			}
		}
	}

	fmt.Println(result)
}

func Solve() {
	step1()
	step2()
}