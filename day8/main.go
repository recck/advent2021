package main

import (
	"advent2021/util"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Signal struct {
	patterns []string
	output []string
}

type Number struct {
	segments string
	totalSegments int
}

func main() {
	fileName := os.Args[1]

	input := util.ReadFileIntoSlice(fileName)

	numbers := getNumbers()
	signals := parseInput(input)

	fmt.Printf("Part 1: %d\n", part1(signals, numbers))
	fmt.Printf("Part 2: %d\n", part2(signals))
}

func part1(signals []Signal, numbers []Number) int {
	total := 0

	matchingNumbers := []int{
		numbers[1].totalSegments,
		numbers[4].totalSegments,
		numbers[7].totalSegments,
		numbers[8].totalSegments,
	}

	for _, signal := range signals {
		for _, output := range signal.output {
			if util.IntSliceContains(matchingNumbers, len(output)) {
				total++
			}
		}
	}

	return total
}

func part2(signals []Signal) int {
	total := 0

	for _, signal := range signals {
		numbers := getNumbers()
		for _, pattern := range signal.patterns {
			switch len(pattern) {
			case 2: numbers[1].segments = pattern
			case 3: numbers[7].segments = pattern
			case 4: numbers[4].segments = pattern
			case 7 :numbers[8].segments = pattern
			}
		}

		for _, pattern := range signal.patterns {
			switch len(pattern) {
			case 5:
				if util.StringContainsLetters(numbers[1].segments, pattern) {
					numbers[3].segments = pattern
				} else if util.MatchingCharacters(pattern, numbers[4].segments) == 3 {
					numbers[5].segments = pattern
				} else {
					numbers[2].segments = pattern
				}
			case 6:
				if !util.StringContainsLetters(numbers[1].segments, pattern) {
					numbers[6].segments = pattern
				} else if !util.StringContainsLetters(numbers[4].segments, pattern) {
					numbers[0].segments = pattern
				} else {
					numbers[9].segments = pattern
				}
			}
		}

		outputValue := ""
		for _, output := range signal.output {
			for i := range numbers {
				if util.SortString(output) == util.SortString(numbers[i].segments) {
					outputValue += strconv.Itoa(i)
				}
			}
		}

		outputAsInt, err := strconv.Atoi(outputValue)
		util.Check(err)
		total += outputAsInt
	}

	return total
}

func parseInput(input []string) []Signal {
	signals := make([]Signal, len(input))

	re := regexp.MustCompile(`([a-g]+)`)

	for i, line := range input {
		matches := re.FindAllString(line, 14)
		signals[i] = Signal{
			patterns: matches[:10],
			output: matches[10:],
		}
	}

	return signals
}

func getNumbers() []Number {
	numbers := make([]Number, 10)
	numbers[0] = Number{totalSegments: 6}
	numbers[1] = Number{totalSegments: 2}
	numbers[2] = Number{totalSegments: 5}
	numbers[3] = Number{totalSegments: 5}
	numbers[4] = Number{totalSegments: 4}
	numbers[5] = Number{totalSegments: 5}
	numbers[6] = Number{totalSegments: 6}
	numbers[7] = Number{totalSegments: 3}
	numbers[8] = Number{totalSegments: 7}
	numbers[9] = Number{totalSegments: 6}

	return numbers
}
