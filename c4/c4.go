package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	path = "./c4/input.txt"
)

type board struct {
	nums   map[int]string
	marked map[string]bool
	winSeq int
	winNum int
}

func (b *board) mark(called int) {
	if loc, ok := b.nums[called]; ok {
		b.marked[loc] = true
	}
}

func (b board) hasWon() bool {
	for i := 0; i < 5; i++ {
		xMatches := 0
		yMatches := 0
		for j := 0; j < 5; j++ {
			if b.marked[key(i, j)] {
				xMatches++
			}
			if b.marked[key(j, i)] {
				yMatches++
			}
		}
		if xMatches == 5 || yMatches == 5 {
			return true
		}
	}

	return false
}

func (b board) score() int {
	sum := 0
	for num, loc := range b.nums {
		if b.marked[loc] == false {
			sum += num
		}
	}

	return sum * b.winNum
}

func main() {
	numbers, boards, err := read(path)
	if err != nil {
		panic("Could not parse input")
	}

	winSeq := 1
	for _, num := range *numbers {
		for i, board := range *boards {
			if !board.hasWon() {
				board.mark(num)
				if board.hasWon() {
					board.winSeq = winSeq
					board.winNum = num
					(*boards)[i] = board
					winSeq++
				}
			}
		}
	}

	maxNr := len(*boards)
	for _, board := range *boards {
		if board.winSeq == maxNr {
			fmt.Println("Last winning board: ", board)
			fmt.Println("Last winning board score: ", board.score())
			return
		}
	}
}

func read(path string) (*[]int, *[]board, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	parsingNumbers := true
	numbers := make([]int, 0, 100)
	boards := make([]board, 0, 30)

	boardRow := 0
	brd := emptyBoard()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			parsingNumbers = false
			continue
		}

		if parsingNumbers {
			for _, s := range strings.Split(scanner.Text(), ",") {
				num, err := strconv.Atoi(s)
				if err != nil {
					panic(fmt.Sprintf("Could not parse %s", s))
				}

				numbers = append(numbers, num)
			}
		} else {
			for i, s := range strings.Fields(scanner.Text()) {
				num, err := strconv.Atoi(s)
				if err != nil {
					panic(fmt.Sprintf("Could not parse %s", s))
				}

				brd.nums[num] = key(boardRow, i)
				brd.marked[key(boardRow, i)] = false
			}

			boardRow++

			// Reset board if 5th row was processed
			if boardRow == 5 {
				boards = append(boards, brd)
				brd = emptyBoard()
				boardRow = 0
			}
		}
	}

	return &numbers, &boards, scanner.Err()
}

func emptyBoard() board {
	return board{nums: make(map[int]string), marked: make(map[string]bool)}
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
