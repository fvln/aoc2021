package main

import (
	"aoc2021/day17"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day17.Solve()

	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s\n", elapsed)
}
