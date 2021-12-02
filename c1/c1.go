package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const path = "./c1/depths.txt"

func main() {
	depths, errors := loadDepths(path)
	if errors != nil {
		fmt.Println("Wrong input!")
		return
	}

	n := 0
	for i := 0; i < len(*depths)-3; i++ {
		curr := aggregateWindow(i, depths)
		next := aggregateWindow(i+1, depths)

		if next > curr {
			n++
		}
	}

	fmt.Println("Changed: ", n)
}

func aggregateWindow(i int, depths *[]int) int {
	slice := (*depths)[i : i+3]
	sum := 0

	for _, depth := range slice {
		sum += depth
	}

	return sum
}

func loadDepths(path string) (*[]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	depths := make([]int, 0, 2000)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, errors.New("non int input")
		}

		depths = append(depths, num)
	}

	return &depths, scanner.Err()
}
