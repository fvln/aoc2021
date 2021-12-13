package day6

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string


func step1(data []int) {
	fishes := make([]int, len(data))
	copy(fishes, data)

	for i:=0; i<80; i++ {
		newFishes := make([]int, 0)
		for idx, value := range fishes {
			if value == 0 {
				fishes[idx] = 6
				newFishes = append(newFishes, 8)
			} else {
				fishes[idx] = value - 1
			}
		}

		fishes = append(fishes, newFishes...)
		// fmt.Printf("%v\n", fishes)
	}

	fmt.Println("Number of fishes after 80 rounds:", len(fishes))
}

func step2(data []int) {
	fishes := make([]int, 9)

	for _, value := range data {
		fishes[value]++
	}

	for i:=0; i<256; i++ {
		newFishes := make([]int, 9)
		for i:=1; i <9; i++ {
			newFishes[i-1] = fishes[i]
		}
		newFishes[8] = fishes[0]
		newFishes[6] += fishes[0]

		fishes = newFishes
	}

	sum := 0
	for i:=0; i <9; i++ {
		sum += fishes[i]
	}

	fmt.Println("Number of fishes after 256 rounds:", sum)

}

func Solve() {
	var fishes []int

	for _, l := range strings.Split(data, ",") {
		value, _ := strconv.Atoi(l)
		fishes = append(fishes, value)
	}

	step1(fishes)
	step2(fishes)
}