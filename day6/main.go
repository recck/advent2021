package main

import (
	"advent2021/util"
	"fmt"
	"os"
)

const SIMULATION_DAYS_TOTAL = 256
const SIMULATION_DAYS_P1 = 80
const DEFAULT_TIMER = 8

func main() {
	fileName := os.Args[1]

	input := util.SplitIntoIntSlice(util.ReadFileAllAtOnce(fileName), ",")
	fishTracker := make([]int, DEFAULT_TIMER + 1)

	for _, fish := range input {
		fishTracker[fish]++
	}

	for i := 0; i < SIMULATION_DAYS_TOTAL; i++ {
		if i == SIMULATION_DAYS_P1 {
			fmt.Printf("Part 1: %d\n", util.SliceSum(fishTracker))
		}

		resettingFishCount := fishTracker[0] // fish that will reset
		fishTracker = append(fishTracker[1:], 0)
		fishTracker[6] += resettingFishCount // fish that will spawn a new fish reset to 6
		fishTracker[8] += resettingFishCount // fish that reset == fish that spawned
	}

	fmt.Printf("Part 2: %d\n", util.SliceSum(fishTracker))
}
