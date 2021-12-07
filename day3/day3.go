package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NB_DIGITS = 12

func loadData() [][]int {
	f, _ := os.Open("day3/data.txt")
	scanner := bufio.NewScanner(f)
	var result [][]int

	for scanner.Scan() {
		line := scanner.Text()
		digits := make([]int, NB_DIGITS)

		for i, digit := range strings.Split(line, "") {
			if digit == "1" {
				digits[i] = 1
			} else {
				digits[i] = -1
			}
		}

		result = append(result, digits)
	}
	return result
}

func step1(data [][]int) {
	digitSums := make([]int, NB_DIGITS)
	gamma := 0
	epsilon := 0

	for _, digits:= range data {
		for i, digit := range digits {
			digitSums[i] += digit
		}
	}

	for i:=0; i < NB_DIGITS; i++  {
		if digitSums[NB_DIGITS - i - 1] > 0 {
			gamma += (1 << i)
		} else {
			epsilon += (1 << i)
		}
	}

	fmt.Printf("Digit sums: %v\n", digitSums)
	println("Day1 - step 1: ", gamma, epsilon, gamma * epsilon)
}

func filterDigits(filter int, digits []int, offset int, reverse bool) bool {
	if reverse {
		if filter >= 0 && digits[offset] == 1 {
			return false
		}
		if filter < 0 && digits[offset] == -1 {
			return false
		}
	} else {
		if filter >= 0 && digits[offset] == -1 {
			return false
		}
		if filter < 0 && digits[offset] == 1 {
			return false
		}
	}
	return true
}

func filterListOnNthDigit(offset int, data [][]int, reverse bool) []int {
	result := make([][]int, 0)
	digitSum := 0

	for _, digits:= range data {
		digitSum += digits[offset]
	}

	for _, digits:= range data {
		if filterDigits(digitSum, digits, offset, reverse) {
			result = append(result, digits)
		}
	}

	fmt.Printf("Filtering offset %d with value %d, got %d results from %d values\n", offset, digitSum, len(result), len(data))
	if len(result) == 1 {
		return result[0]
	}

	if offset >= NB_DIGITS -1 {
		return nil
	}

	return filterListOnNthDigit(offset + 1, result, reverse)
}


func step2(data [][]int) {
	oxy := 0
	filteredData := filterListOnNthDigit(0, data, false)
	if filteredData != nil {
		fmt.Printf("Oxy %v\n", filteredData)
		for j:=0; j < NB_DIGITS; j++ {
			if filteredData[NB_DIGITS-j-1] > 0 {
				oxy += (1 << j)
			}
		}
	}


	co2 := 0
	filteredData = filterListOnNthDigit(0, data, true)
	if filteredData != nil {
		fmt.Printf("co2 %v\n", filteredData)
		for j:=0; j < NB_DIGITS; j++ {
			if filteredData[NB_DIGITS-j-1] > 0 {
				co2 += (1 << j)
			}
		}
	}

	println("Day1 - step 2: ", oxy, co2, oxy * co2)
}

func Solve() {
	data := loadData()
	step1(data)
	step2(data)
}