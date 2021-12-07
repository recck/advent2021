package main

import (
	"advent2021/util"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	fileName := os.Args[1]

	input := util.SplitIntoIntSlice(util.ReadFileAllAtOnce(fileName), ",")

	sort.Ints(input)
	middleIndex := len(input) / 2
	median := 0

	if len(input) % 2 == 1 {
		median = input[middleIndex]
	} else {
		median = (input[middleIndex - 1] + input[middleIndex]) / 2
	}

	total := 0

	for i := range input {
		total += util.IntAbs(input[i] - median)
	}

	fmt.Printf("Part 1: %d\n", total)

	nonMedianTotal := math.MaxInt64

	for j := 0; j < input[len(input) - 1]; j++ {
		curTotal := 0
		for i := range input {
			diff := util.IntAbs(j - input[i])
			curTotal += (diff * (diff + 1)) / 2
		}

		nonMedianTotal = util.IntMin(curTotal, nonMedianTotal)
	}

	fmt.Printf("Part 2: %d\n", nonMedianTotal)
}
