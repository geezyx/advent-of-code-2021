package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var inputs []string
	for scanner.Scan() {
		inputs = append(inputs, strings.Trim(scanner.Text(), "\n"))
	}
	fmt.Println("Part1:")
	fmt.Println(part1(inputs))
	fmt.Println("Part2:")
	fmt.Println(part2(inputs))
}

type bit struct {
	zeroes int
	ones   int
}

func (b bit) common() string {
	if b.zeroes == b.ones {
		return "1"
	}
	if b.zeroes > b.ones {
		return "0"
	}
	return "1"
}

func (b bit) uncommon() string {
	if b.zeroes == b.ones {
		return "0"
	}
	if !(b.zeroes > b.ones) {
		return "0"
	}
	return "1"
}

func filterInputs(inputs []string, index int, value string) []string {
	filtered := []string{}
	for _, in := range inputs {
		inVals := strings.Split(in, "")
		if inVals[index] == value {
			filtered = append(filtered, in)
		}
	}
	return filtered
}

func parseBits(inputs []string) []bit {
	var bits = make([]bit, len(inputs[0]))
	for _, input := range inputs {
		for i, b := range strings.Split(input, "") {
			switch b {
			case "0":
				bits[i].zeroes++
			case "1":
				bits[i].ones++
			}
		}
	}
	return bits
}

func part1(inputs []string) int64 {
	bits := parseBits(inputs)

	var gamma, epsilon string
	for _, b := range bits {
		gamma += b.common()
		epsilon += b.uncommon()
	}

	gammaVal, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilonVal, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	return gammaVal * epsilonVal
}

func parseO2(inputs []string, index int) int64 {
	if len(inputs) == 1 {
		res, err := strconv.ParseInt(inputs[0], 2, 64)
		if err != nil {
			panic(err)
		}
		return res
	}
	bits := parseBits(inputs)
	newInputs := filterInputs(inputs, index, bits[index].common())
	return parseO2(newInputs, index+1)
}

func parseCO2(inputs []string, index int) int64 {
	if len(inputs) == 1 {
		res, err := strconv.ParseInt(inputs[0], 2, 64)
		if err != nil {
			panic(err)
		}
		return res
	}
	bits := parseBits(inputs)
	newInputs := filterInputs(inputs, index, bits[index].uncommon())
	return parseCO2(newInputs, index+1)
}

func part2(inputs []string) int64 {
	o2 := parseO2(inputs, 0)
	co2 := parseCO2(inputs, 0)

	return o2 * co2
}
