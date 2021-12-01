package day1

import (
	"bufio"
	"os"
	"strconv"
)

func loadData() []int {
	f, _ := os.Open("day1/data.txt")
	scanner := bufio.NewScanner(f)
	result := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		if err == nil {
			result = append(result, value)
		}

	}
	return result
}

func step1(data []int) {
	increases := 0
	previousValue := 999999999

	for _, value:= range data {
		if value > previousValue {
			increases++
		}
		previousValue = value
	}

	println("Day1 - step 1: ", increases)
}

type window struct {
	values []int
	size int
	nbValues int
}

func newWindow(size int) *window {
	return &window{
		values: make([]int, 0),
		size:  size,
		nbValues: 0,
	}
}

func (w *window) push(value int) {
	if w.nbValues >= w.size {
		// Pop first value + append
		w.values = w.values[1:]
		w.values = append(w.values, value)
	} else {
		w.values = append(w.values, value)
		w.nbValues++
	}
}

func (w *window) sum() (int, bool) {
	if w.nbValues < w.size {
		// Window is not full
		return 0, false
	}

	sum := 0
	for _, value := range w.values {
		sum += value
	}
	return sum, true
}

func step2(data []int) {
	increases := 0
	previousSum := 999999999
	window := newWindow(3)

	for _, value:= range data {
		window.push(value)
		newSum, ok := window.sum()
		if !ok {
			continue
		}

		if newSum > previousSum {
			increases += 1
		}
		previousSum = newSum
	}

	println("Day1 - step 2: ", increases)
}

func Solve() {
	data := loadData()
	step1(data)
	step2(data)
}