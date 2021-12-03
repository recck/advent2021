package util

import (
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileAllAtOnce(fileName string) string {
	contents, err := os.ReadFile(fileName)
	Check(err)

	return strings.TrimSpace(string(contents))
}

func ReadFileIntoSlice(fileName string, delim ...string) []string {
	contents := ReadFileAllAtOnce(fileName)
	var splitter = "\n"

	if len(delim) > 0 {
		splitter = delim[0]
	}

	return strings.Split(contents, splitter)
}

func ReadFileIntoIntSlice(fileName string, delim ...string) []int {
	var splitter = "\n"

	if len(delim) > 0 {
		splitter = delim[0]
	}

	lines := ReadFileIntoSlice(fileName, splitter)
	intArray := make([]int, len(lines))

	for i, v := range lines {
		vInt, err := strconv.Atoi(v)
		Check(err)
		intArray[i] = vInt
	}

	return intArray
}

func SliceSum(input []int) int {
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

func BinDec(bin string) int64 {
	dec, err := strconv.ParseInt(bin, 2, 64)
	Check(err)
	return dec
}
