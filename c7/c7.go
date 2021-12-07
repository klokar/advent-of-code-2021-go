package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const path = "./c7/input.txt"

func main() {
	crabs, err := read(path)
	if err != nil {
		panic("Could not parse input")
	}

	sum := 0
	for _, crab := range crabs {
		sum += crab
	}

	avg := sum / len(crabs)
	fmt.Println("Avg: ", avg)

	fuelUsage := make(map[int]int)

	for i := 0; i <= 10; i++ {
		num1 := avg + i
		num2 := avg - i

		for _, crab := range crabs {
			o1 := offset(crab, num1)
			o2 := offset(crab, num2)
			fuelUsage[num1] += o1 * (o1 + 1) / 2

			// Don't calculate twice for first number
			if num2 != avg {
				fuelUsage[num2] += o2 * (o2 + 1) / 2
			}
		}
	}

	minOffset := 0
	minUsage := math.MaxInt
	for num, usage := range fuelUsage {
		if usage < minUsage {
			minUsage = usage
			minOffset = num
		}
	}

	fmt.Println("Min offset and usage: ", minOffset, minUsage)
	fmt.Println("Usages: ", fuelUsage)
}

func offset(position, targeting int) int {
	if targeting > position {
		return targeting - position
	}

	return position - targeting
}

func read(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return []int{}, err
	}
	defer file.Close()

	crabs := make([]int, 0, 1000)

	scn := bufio.NewScanner(file)
	scn.Scan()
	positions := strings.Split(scn.Text(), ",")
	for _, position := range positions {
		num, _ := strconv.Atoi(position)
		crabs = append(crabs, num)
	}

	return crabs, scn.Err()
}
