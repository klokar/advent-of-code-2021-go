package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	path = "./c3/input.txt"
	bits = 12
)

func main() {
	binaries, err := read(path)
	if err != nil {
		fmt.Println("Could not parse input!")
		return
	}

	incidence := calculateIncidence(binaries)

	gamma, epsilon := gammaEpsilon(incidence)
	og, scr := generatorScrubber(binaries, incidence)

	fmt.Println("Gamma: ", gamma)
	fmt.Println("Epsilon: ", epsilon)
	fmt.Println("Oxygen Generator: ", og)
	fmt.Println("Co2 Scrubber: ", scr)
	fmt.Println(gamma*epsilon, og*scr)
}

func generatorScrubber(binaries *[]string, incidence *[bits][2]int) (int64, int64) {
	oxyKeep := make([]string, len(*binaries))
	scrKeep := make([]string, len(*binaries))
	copy(oxyKeep, *binaries)
	copy(scrKeep, *binaries)

	oxyIncidence := *incidence
	scrIncidence := *incidence

	for i := 0; i < bits; i++ {

		// Set most common Oxygen Generator bit incidence
		mostCommon := mostCommonBit(oxyIncidence[i])

		// Iterate over remaining binaries and remove non-matching
		for j := 0; j < len(oxyKeep) && len(oxyKeep) > 1; j++ {
			if string(oxyKeep[j][i]) != mostCommon {
				oxyKeep = append(oxyKeep[:j], oxyKeep[j+1:]...)
				oxyIncidence = *calculateIncidence(&oxyKeep)
				j--
			}
		}

		// Set most common CO2 scrubber rating bit incidence
		mostCommon = mostCommonBit(scrIncidence[i])

		// Iterate over remaining binaries and remove non-matching
		for j := 0; j < len(scrKeep) && len(scrKeep) > 1; j++ {
			if string(scrKeep[j][i]) == mostCommon {
				scrKeep = append(scrKeep[:j], scrKeep[j+1:]...)
				scrIncidence = *calculateIncidence(&scrKeep)
				j--
			}
		}
	}

	og, _ := strconv.ParseInt(oxyKeep[0], 2, 64)
	scr, _ := strconv.ParseInt(scrKeep[0], 2, 64)

	return og, scr
}

func mostCommonBit(bitIncidence [2]int) string {
	if bitIncidence[0] > bitIncidence[1] {
		return "0"
	}

	return "1"
}

func gammaEpsilon(incidence *[bits][2]int) (int64, int64) {
	gammaS := ""
	epsilonS := ""
	for _, bitIncidence := range incidence {
		if bitIncidence[0] > bitIncidence[1] {
			gammaS += "0"
			epsilonS += "1"
		} else {
			gammaS += "1"
			epsilonS += "0"
		}
	}

	gammaD, _ := strconv.ParseInt(gammaS, 2, 64)
	epsilonD, _ := strconv.ParseInt(epsilonS, 2, 64)

	return gammaD, epsilonD
}

func calculateIncidence(binaries *[]string) *[bits][2]int {
	var incidence [bits][2]int

	for i := 0; i < bits; i++ {
		for _, binary := range *binaries {
			if string(binary[i]) == "1" {
				incidence[i][1] += 1
			} else {
				incidence[i][0] += 1
			}
		}
	}

	return &incidence
}

func read(path string) (*[]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	binaries := make([]string, 0, 1000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		binaries = append(binaries, scanner.Text())
	}

	return &binaries, scanner.Err()
}
