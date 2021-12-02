package main

import (
	"advent2021/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	depth int
	horizontal int
	aim int
}

func (sub Submarine) calculate() int {
	return sub.depth * sub.horizontal
}

func (sub *Submarine) handleForward(value int, handleAim bool) {
	sub.horizontal += value
	if handleAim {
		sub.depth += sub.aim * value
	}
}

func (sub *Submarine) handleVertical(value int, isUp bool, handleAim bool) {
	modifier := 1
	if isUp {
		modifier = -1
	}

	if handleAim {
		sub.aim += value * modifier
	} else {
		sub.depth += value * modifier
	}
}

func main() {
	fileName := os.Args[1]

	lines := util.ReadFileIntoSlice(fileName)

	subP1 := Submarine{}
	subP2 := Submarine{}

	for _, line := range lines {
		// part 1
		updateSubmarine(&subP1, line, false)
		// part 2
		updateSubmarine(&subP2, line, true)
	}

	fmt.Printf("part 1: %d\n", subP1.calculate())
	fmt.Printf("part 2: %d\n", subP2.calculate())
}

func updateSubmarine(sub *Submarine, line string, handleAim bool) {
	chunks := strings.Split(line, " ")
	direction := chunks[0]
	num, err := strconv.Atoi(chunks[1])
	util.Check(err)

	switch direction {
	case "forward":
		sub.handleForward(num, handleAim)
		break
	case "up", "down":
		sub.handleVertical(num, direction == "up", handleAim)
		break
	}
}
