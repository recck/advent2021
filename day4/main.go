package main

import (
	"advent2021/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const GridSize = 5

type BingoNumber struct {
	value int
	drawn bool
}

type BingoRow struct {
	numbers []BingoNumber
}

type BingoBoard struct {
	rows []BingoRow
	hasWon bool
	winningNumber int
}

func (board *BingoBoard) markNumber(number int) {
	for row := 0; row < GridSize; row++ {
		for cell := 0; cell < GridSize; cell++ {
			if board.rows[row].numbers[cell].value == number {
				board.rows[row].numbers[cell].drawn = true
			}
		}
	}
}

func (board BingoBoard) calculateScore() int {
	sum := 0
	for _, row := range board.rows {
		for _, num := range row.numbers {
			if !num.drawn {
				sum += num.value
			}
		}
	}

	return sum * board.winningNumber
}

func (board BingoBoard) isWinner() bool {
	// check horizontals
	for _, row := range board.rows {
		if numbersDrawn(row.numbers) == GridSize {
			return true
		}
	}

	// check verticals
	for i := 0; i < GridSize; i++ {
		var rowToCheck []BingoNumber
		for _, row := range board.rows {
			rowToCheck = append(rowToCheck, row.numbers[i])
		}

		if numbersDrawn(rowToCheck) == GridSize {
			return true
		}
	}

	// check diagonals
	//var tLBrRow []BingoNumber
	//var tRBlRow []BingoNumber
	//for i := 0; i < GridSize; i++ {
	//	tLBrRow = append(tLBrRow, board.rows[i].numbers[i])
	//	tRBlRow = append(tRBlRow, board.rows[i].numbers[GridSize - i - 1])
	//}
	//
	//if numbersDrawn(tLBrRow) == GridSize || numbersDrawn(tRBlRow) == GridSize {
	//	return true
	//}

	return false
}

func main() {
	fileName := os.Args[1]

	lines := util.ReadFileIntoSlice(fileName)

	numbersToDraw := util.SplitIntoIntSlice(lines[0], ",")
	lines = util.RemoveBlankLines(lines[2:])
	boards := makeBingoBoards(lines)

	var winners []BingoBoard

	for _, numberToDraw := range numbersToDraw {
		for i := 0; i < len(boards); i++ {
			if boards[i].hasWon {
				continue
			}

			boards[i].markNumber(numberToDraw)

			if boards[i].isWinner() {
				boards[i].hasWon = true
				boards[i].winningNumber = numberToDraw
				winners = append(winners, boards[i])
			}
		}
	}

	fmt.Printf("Part 1: %d\n", winners[0].calculateScore())
	fmt.Printf("Part 2: %d\n", winners[len(winners) - 1].calculateScore())
}

func makeBingoBoards(lines []string) []BingoBoard {
	var boards []BingoBoard

	numLines := len(lines)

	for i := 0; i < numLines; i += GridSize {
		board := BingoBoard{hasWon: false}
		for j := i; j < i +GridSize; j++ {
			board.rows = append(board.rows, makeBingoRow(lines[j]))
		}
		boards = append(boards, board)
	}

	return boards
}

func makeBingoRow(line string) BingoRow {
	row := BingoRow{}
	numbers := strings.Fields(line)

	for _, num := range numbers {
		numInt, _ := strconv.Atoi(num)
		bingoNumber := BingoNumber{
			drawn: false,
			value: numInt,
		}
		row.numbers = append(row.numbers, bingoNumber)
	}

	return row
}

func numbersDrawn(numbers []BingoNumber) int {
	drawn := 0
	for _, number := range numbers {
		if number.drawn {
			drawn++
		}
	}

	return drawn
}
