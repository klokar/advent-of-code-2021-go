package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const path = "./c2/moves.txt"

type submarine struct {
	forward int
	aim     int
	depth   int
}

type move struct {
	direction string
	length    int
}

func main() {
	moves, err := readMoves(path)
	if err != nil {
		fmt.Println("could not read input")
		return
	}

	subm := submarine{forward: 0, depth: 0}
	for _, move := range *moves {
		switch move.direction {
		case "forward":
			subm.forward += move.length
			subm.depth += subm.aim * move.length
		case "up":
			subm.aim -= move.length
		case "down":
			subm.aim += move.length
		}
	}

	fmt.Println("End location: ", subm, subm.depth*subm.forward)
}

func readMoves(path string) (*[]move, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	moves := make([]move, 0, 2000)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		length, _ := strconv.Atoi(fields[1])

		moves = append(moves, move{direction: fields[0], length: length})
	}

	return &moves, scanner.Err()
}
