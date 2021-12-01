package main

import (
	"advent2021/util"
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]

	lines := util.ReadFileReturnArrayInts(fileName)

	increasesP1 := 0
	prevP1 := lines[0]

	increasesP2 := 0
	prevP2 := util.ArraySum(lines[0:3])
	maxLen := len(lines)

	for i, v := range lines {
		// part 1
		if v > prevP1 {
			increasesP1++
		}

		prevP1 = v

		// part 2
		leftBound := i
		rightBound := util.IntMin(maxLen, leftBound + 3)
		window := util.ArraySum(lines[leftBound:rightBound])
		if window > prevP2 {
			increasesP2++
		}

		prevP2 = window
	}

	fmt.Printf("part 1: %d\npart 2: %d\n", increasesP1, increasesP2)
}
