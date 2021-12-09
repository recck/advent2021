package main

import (
	"advent2021/util"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Cell struct {
	value int
	isLowest bool
	x int
	y int
}

func (c Cell) String() string {
	if c.isLowest {
		return "x"
	} else {
		return "."
	}
}

type Heightmap struct {
	cells [][]Cell
	sizeX int
	sizeY int
}

func (h Heightmap) String() string {
	output := ""
	for y := range h.cells {
		output += fmt.Sprintln(h.cells[y])
	}
	return output
}

func (h Heightmap) getCellAt(x int, y int) *Cell {
	valid := x >= 0 && x < h.sizeX && y >= 0 && y < h.sizeY

	if valid {
		return &h.cells[y][x]
	}

	return nil
}

func (h Heightmap) determineLowest(cell *Cell) {
	pX, nX, pY, nY := cell.x - 1, cell.x + 1, cell.y - 1, cell.y + 1
	isLowest := true
	topCell, bottomCell := h.getCellAt(cell.x, pY), h.getCellAt(cell.x, nY)
	leftCell, rightCell := h.getCellAt(pX, cell.y), h.getCellAt(nX, cell.y)
	cellsToCheck := []*Cell{topCell, bottomCell, leftCell, rightCell}

	for _, cellToCheck := range cellsToCheck {
		if cellToCheck != nil && cellToCheck.value <= cell.value {
			isLowest = false
		}
	}

	cell.isLowest = isLowest
}

func main() {
	fileName := os.Args[1]

	input := util.ReadFileIntoSlice(fileName)
	heightmap := makeHeightmap(input)

	fmt.Printf("Part 1: %d\n", part1(heightmap))
	fmt.Printf("Part 2: %d\n", part2(heightmap))
}

func part1(heightmap Heightmap) int {
	total := 0

	for y := 0; y < heightmap.sizeY; y++ {
		for x := 0; x < heightmap.sizeX; x++ {
			cell := &heightmap.cells[y][x]
			heightmap.determineLowest(cell)
			if cell.isLowest {
				total += cell.value + 1
			}
		}
	}

	return total
}

func part2(heightmap Heightmap) int {
	var basins []int
	basinTracker := make([][]bool, heightmap.sizeY)

	for y := 0; y < heightmap.sizeY; y++ {
		basinTracker[y] = make([]bool, heightmap.sizeX)
	}

	for y := 0; y < heightmap.sizeY; y++ {
		for x := 0; x < heightmap.sizeX; x++ {
			basinTotal := getBasin(heightmap, &basinTracker, x, y)
			if basinTotal > 0 {
				basins = append(basins, basinTotal)
			}
		}
	}

	sort.Slice(basins, func(i int, j int) bool {
		return basins[j] < basins[i]
	})

	return util.SliceProduct(basins[:3])
}

func getBasin(heightmap Heightmap, basinTracker *[][]bool, x int, y int) int {
	if (*basinTracker)[y][x] {
		return 0
	}

	cell := heightmap.getCellAt(x, y)
	if cell == nil || cell.value >= 9 {
		(*basinTracker)[y][x] = true
		return 0
	}

	cells := []Cell{*cell}
	basinTotal := 0
	for len(cells) > 0 {
		curCell := cells[0]
		cells = cells[1:]

		if curCell.value >= 9 || (*basinTracker)[curCell.y][curCell.x] {
			continue
		}

		(*basinTracker)[curCell.y][curCell.x] = true
		basinTotal += 1

		pX, nX, pY, nY := curCell.x - 1, curCell.x + 1, curCell.y - 1, curCell.y + 1
		topCell, bottomCell := heightmap.getCellAt(curCell.x, pY), heightmap.getCellAt(curCell.x, nY)
		leftCell, rightCell := heightmap.getCellAt(pX, curCell.y), heightmap.getCellAt(nX, curCell.y)
		cellsToAdd := []*Cell{topCell, bottomCell, leftCell, rightCell}
		for _, cellToAdd := range cellsToAdd {
			if cellToAdd != nil {
				cells = append(cells, *cellToAdd)
			}
		}
	}

	return basinTotal
}

func makeHeightmap(input []string) Heightmap {
	heightmap := Heightmap{
		cells: make([][]Cell, len(input)),
		sizeY: len(input),
		sizeX: len(input[0]),
	}

	for i := range heightmap.cells {
		heightmap.cells[i] = make([]Cell, len(input[0]))
	}

	for y, line := range input {
		for x, value := range strings.Split(line, "") {
			valueAsInt, err := strconv.Atoi(value)
			util.Check(err)
			heightmap.cells[y][x] = Cell{
				value: valueAsInt,
				x: x,
				y: y,
				isLowest: false,
			}
		}
	}

	return heightmap
}
