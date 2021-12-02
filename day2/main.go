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
		sub.horizontal += num
		if handleAim {
			sub.depth += sub.aim * num
		}
		break
	case "up":
		if handleAim {
			sub.aim -= num
		} else {
			sub.depth -= num
		}
		break
	case "down":
		if handleAim {
			sub.aim += num
		} else {
			sub.depth += num
		}
		break
	}
}
