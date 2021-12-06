package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	path          = "./c6/input.txt"
	newBirthTimer = 9
	oldBirthTimer = 7
	iterations    = 256
)

func main() {
	fish, err := read(path)
	if err != nil {
		panic("Could not parse input")
	}

	for i := 0; i < iterations; i++ {
		var fishC [newBirthTimer]int
		for j := newBirthTimer - 1; j >= 0; j-- {
			if j > 0 {
				fishC[j-1] = fish[j]
			} else {
				fishC[newBirthTimer-1] = fish[j]
				fishC[oldBirthTimer-1] += fish[j]
			}
		}

		fish = fishC
	}

	sum := 0
	for _, count := range fish {
		sum += count
	}

	fmt.Println("Total number of fish: ", sum)
}

func read(path string) ([newBirthTimer]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return [newBirthTimer]int{}, err
	}
	defer file.Close()

	var fish [newBirthTimer]int

	scn := bufio.NewScanner(file)
	scn.Scan()
	timers := strings.Split(scn.Text(), ",")
	for _, timer := range timers {
		num, _ := strconv.Atoi(timer)
		fish[num]++
	}

	return fish, scn.Err()
}
