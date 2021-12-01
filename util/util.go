package util

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileAllAtOnce(fileName string) string {
	contents, err := os.ReadFile(fileName)
	check(err)

	return strings.TrimSpace(string(contents))
}

func ReadFileReturnArray(fileName string, delim ...string) []string {
	contents := ReadFileAllAtOnce(fileName)
	var splitter = "\n"

	if len(delim) > 0 {
		splitter = delim[0]
	}

	return strings.Split(contents, splitter)
}

func ReadFileReturnArrayInts(fileName string, delim ...string) []int {
	var splitter = "\n"

	if len(delim) > 0 {
		splitter = delim[0]
	}

	lines := ReadFileReturnArray(fileName, splitter)
	intArray := make([]int, len(lines))

	for i, v := range lines {
		vInt, err := strconv.Atoi(v)
		check(err)
		intArray[i] = vInt
	}

	return intArray
}

func ArraySum(input []int) int {
	total := 0

	for _, v := range input {
		total += v
	}

	return total
}

func IntMin(x int, y int) int {
	if x > y {
		return y
	}

	return x
}
