package main

import (
	"advent2021/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]

	lines := util.ReadFileIntoSlice(fileName)

	width, diagnostics := generateDiagnostics(lines)
	powerConsumption := computePowerConsumption(width, diagnostics)
	oxygenRating := computeOxygenGeneratorRating(lines, width, diagnostics)
	co2Rating := computeCO2Rating(lines, width, diagnostics)

	fmt.Printf("Part 1 - Power Consumption: %d\n", powerConsumption)
	fmt.Printf("Part 2 - Oxygen Rating: %d\n", oxygenRating)
	fmt.Printf("Part 2 - CO2 Rating: %d\n", co2Rating)
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

	gammaDec, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	epsilonDec, _ := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)

	return gammaDec * epsilonDec
}

func computeOxygenGeneratorRating(bits []string, width int, diagnostics map[int]map[string]int) int64 {
	index := 0

	for len(bits) != 1 {
		var curMatches []string

		bitMatch := "0"
		if diagnostics[index]["1"] >= diagnostics[index]["0"] {
			bitMatch = "1"
		}

		for _, bit := range bits {
			if string(bit[index]) == bitMatch {
				curMatches = append(curMatches, bit)
			}
		}

		bits = curMatches
		_, diagnostics = generateDiagnostics(bits)
		index++
		index = util.IntMin(index, width)
	}

	oxygenDec, _ := strconv.ParseInt(bits[0], 2, 64)

	return oxygenDec
}

func computeCO2Rating(bits []string, width int, diagnostics map[int]map[string]int) int64 {
	index := 0

	for len(bits) != 1 {
		var curMatches []string

		bitMatch := "1"
		if diagnostics[index]["0"] <= diagnostics[index]["1"] {
			bitMatch = "0"
		}

		for _, bit := range bits {
			if string(bit[index]) == bitMatch {
				curMatches = append(curMatches, bit)
			}
		}

		bits = curMatches
		_, diagnostics = generateDiagnostics(bits)
		index++
		index = util.IntMin(index, width)
	}

	oxygenDec, _ := strconv.ParseInt(bits[0], 2, 64)

	return oxygenDec
}
