package main

import (
	"advent2021/util"
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type Line struct {
	left Coordinate
	right Coordinate
}

type Diagram struct {
	lines []Line
	width int
	height int
}

func (l Line) isHorizontalOrVertical() bool {
	return l.left.x == l.right.x || l.left.y == l.right.y
}

func main() {
	fileName := os.Args[1]

	lines := util.ReadFileIntoSlice(fileName)
	diagram := makeDiagram(lines)

	fmt.Printf("Part 1: %d\n", computeOverlappingPoints(diagram, false))
	fmt.Printf("Part 2: %d\n", computeOverlappingPoints(diagram, true))
}

func makeDiagram(lines []string) Diagram {
	var diagram Diagram
	width := 0
	height := 0

	// x1,y1 -> x2,y2
	for _, line := range lines {
		coordinates := strings.Split(strings.TrimSpace(line), " -> ")
		left := util.SplitIntoIntSlice(coordinates[0], ",")
		right := util.SplitIntoIntSlice(coordinates[1], ",")
		diagram.lines = append(diagram.lines, Line{
			Coordinate{left[0], left[1]},
			Coordinate{right[0], right[1]},
		})
		width = util.IntMax(width, util.IntMax(left[0], right[0]))
		height = util.IntMax(height, util.IntMax(left[1], right[1]))
	}

	diagram.width = width + 1
	diagram.height = height + 1

	return diagram
}

func computeOverlappingPoints(diagram Diagram, includeDiagonals bool) int {
	board := make([][]int, diagram.height)
	for y := range board {
		board[y] = make([]int, diagram.width)
	}

	for _, line := range diagram.lines {
		if !line.isHorizontalOrVertical() && !includeDiagonals {
			continue
		}

		x0, y0, x1, y1 := line.left.x, line.left.y, line.right.x, line.right.y
		dx := util.IntAbs(x1 - x0)
		stepX := -1
		if line.left.x < line.right.x {
			stepX = 1
		}
		dy := util.IntAbs(y1 - y0) * -1
		stepY := -1
		if line.left.y < line.right.y {
			stepY = 1
		}
		e := dx + dy

		for {
			board[y0][x0]++

			if x0 == x1 && y0 == y1 {
				break
			}

			e2 := e * 2

			if e2 >= dy {
				e += dy
				x0 += stepX
			}

			if e2 <= dx {
				e += dx
				y0 += stepY
			}
		}
	}

	greaterThan2 := 0

	for _, xs := range board {
		for _, x := range xs {
			if x >= 2 {
				greaterThan2++
			}
		}
	}

	return greaterThan2
}
