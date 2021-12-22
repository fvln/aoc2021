package main

import (
	"aoc2021/day18"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day18.Solve()

	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s\n", elapsed)
}
