package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var inputs []int64
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 16)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, i)
	}
	fmt.Println("Part1:")
	fmt.Println(part1(inputs))
	fmt.Println("Part2:")
	fmt.Println(part2(inputs))
}

func part1(inputs []int64) int {
	var prev int64
	var increments int
	for _, in := range inputs {
		if in > prev {
			increments++
		}
		prev = in
	}
	return increments - 1
}

func part2(inputs []int64) int {
	var increments int
	for i := 0; i < len(inputs)-3; i++ {
		sum1 := inputs[i] + inputs[i+1] + inputs[i+2]
		sum2 := inputs[i+1] + inputs[i+2] + inputs[i+3]
		if sum2 > sum1 {
			increments++
		}
	}
	return increments
}
