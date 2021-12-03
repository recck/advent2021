package main

import (
	"advent2021/util"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := os.Args[1]

	lines := util.ReadFileIntoSlice(fileName)

	width, diagnostics := generateDiagnostics(lines)
	powerConsumption := computePowerConsumption(width, diagnostics)
	oxygenRating := computeRating(lines, diagnostics, "oxygen")
	co2Rating := computeRating(lines, diagnostics, "co2")

	fmt.Printf("Part 1 - Power Consumption: %d\n", powerConsumption)
	fmt.Printf("Part 2 - Life Support Rating: %d\n", oxygenRating * co2Rating)
}

func generateDiagnostics(lines []string) (int, map[int]map[string]int) {
	width := len(lines[0])
	diagnostics := make(map[int]map[string]int, width)

	for _, line := range lines {
		for i := 0; i < width; i++ {
			if diagnostics[i] == nil {
				diagnostics[i] = map[string]int{"0": 0, "1": 0}
			}
			diagnostics[i][string(line[i])]++
		}
	}

	return width, diagnostics
}

func computePowerConsumption(width int, bits map[int]map[string]int) int64 {
	gamma := make([]string, width)
	epsilon := make([]string, width)

	for w, bit := range bits {
		if bit["1"] > bit["0"] {
			gamma[w] = "1"
			epsilon[w] = "0"
		} else {
			gamma[w] = "0"
			epsilon[w] = "1"
		}
	}

	return util.BinDec(strings.Join(gamma, "")) * util.BinDec(strings.Join(epsilon, ""))
}

func getBitMatch(diagnostics map[int]map[string]int, index int, rating string) string {
	bitMatch := ""
	if rating == "oxygen" {
		bitMatch = "0"
		if diagnostics[index]["1"] >= diagnostics[index]["0"] {
			bitMatch = "1"
		}
	} else {
		bitMatch = "1"
		if diagnostics[index]["0"] <= diagnostics[index]["1"] {
			bitMatch = "0"
		}
	}

	return bitMatch
}

func computeRating(bits []string, diagnostics map[int]map[string]int, rating string) int64 {
	var curMatches []string

	for index := 0; len(bits) != 1 && index < len(bits[0]); index++ {
		bitMatch := getBitMatch(diagnostics, index, rating)

		for _, bit := range bits {
			if string(bit[index]) == bitMatch {
				curMatches = append(curMatches, bit)
			}
		}

		bits = curMatches
		_, diagnostics = generateDiagnostics(bits)
	}

	return util.BinDec(bits[0])
}
