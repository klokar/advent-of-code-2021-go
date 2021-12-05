package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	path                 = "./c5/input.txt"
	dangerousOccurrences = 2
)

func main() {
	_, dangers, err := read(path)
	if err != nil {
		panic("Could not parse input")
	}

	fmt.Println("Dangers: ", len(*dangers))
}

func read(path string) (*map[string]int, *[]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	vents := make(map[string]int)
	dangers := make([]string, 0, 100)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		from := strings.Split(parts[0], ",")
		to := strings.Split(parts[1], ",")

		xf, _ := strconv.Atoi(from[0])
		xt, _ := strconv.Atoi(to[0])
		xMin, xMax := minMax(xf, xt)
		yf, _ := strconv.Atoi(from[1])
		yt, _ := strconv.Atoi(to[1])
		yMin, yMax := minMax(yf, yt)

		// Vertical or horizontal
		if xf == xt || yf == yt {
			for x := xMin; x <= xMax; x++ {
				for y := yMin; y <= yMax; y++ {
					vents[key(x, y)] += 1
					if vents[key(x, y)] == dangerousOccurrences {
						dangers = append(dangers, key(x, y))
					}
				}
			}
		}

		// Diagonal
		if xMax-xMin == yMax-yMin {
			xRange := makeRange(xf, xt, xMin, xMax)
			yRange := makeRange(yf, yt, yMin, yMax)

			for i := 0; i < len(xRange); i++ {
				x := xRange[i]
				y := yRange[i]
				vents[key(x, y)] += 1
				if vents[key(x, y)] == dangerousOccurrences {
					dangers = append(dangers, key(x, y))
				}
			}
		}
	}

	return &vents, &dangers, scanner.Err()
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}

	return b, a
}

func makeRange(n1, n2, min, max int) []int {
	a := make([]int, max-min+1)

	for i := range a {
		if n1 < n2 {
			a[i] = n1 + i
		} else {
			a[i] = n1 - i
		}

	}

	return a
}
