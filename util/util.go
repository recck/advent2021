package util

import (
	"os"
	"sort"
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

func RemoveBlankLines(lines []string) []string {
	var output []string

	for _, line := range lines {
		if line == "" {
			continue
		}

		output = append(output, strings.TrimSpace(line))
	}

	return output
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

func SplitIntoIntSlice(line string, delim ...string) []int {
	var splitter = " "

	if len(delim) > 0 {
		splitter = delim[0]
	}

	values := strings.Split(line, splitter)

	intArray := make([]int, len(values))

	for i, v := range values {
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

func SliceProduct(input []int) int {
	total := 1

	for _, v := range input {
		total *= v
	}

	return total
}

func IntMin(x int, y int) int {
	if x > y {
		return y
	}

	return x
}

func IntMax(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func IntAbs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func BinDec(bin string) int64 {
	dec, err := strconv.ParseInt(bin, 2, 64)
	Check(err)
	return dec
}

func IntSliceContains(slice []int, check int) bool {
	for _, v := range slice {
		if v == check {
			return true
		}
	}

	return false
}

func StringSliceContains(slice []string, check string) bool {
	for _, v := range slice {
		if v == check {
			return true
		}
	}

	return false
}

func StringSubSlice(sliceA []string, sliceB []string) bool {
	for _, v := range sliceA {
		if !StringSliceContains(sliceB, v) {
			return false
		}
	}

	return true
}

func StringContainsLetters(stringA string, stringB string) bool {
	return StringSubSlice(strings.Split(stringA, ""), strings.Split(stringB, ""))
}

func MatchingCharacters(stringA string, stringB string) int {
	total := 0

	for _, c := range strings.Split(stringA, "") {
		if strings.Contains(stringB, c) {
			total++
		}
	}

	return total
}

func SortString(input string) string {
	s := strings.Split(input, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
