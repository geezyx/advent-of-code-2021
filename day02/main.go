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
		inputs = append(inputs, scanner.Text())
	}
	fmt.Println("Part1:")
	fmt.Println(part1(inputs))
	fmt.Println("Part2:")
	fmt.Println(part2(inputs))
}

func part1(inputs []string) int64 {
	var depth, horizontal int64
	for _, in := range inputs {
		res := strings.Split(in, " ")
		direction := res[0]
		distance, err := strconv.ParseInt(res[1], 10, 16)
		if err != nil {
			panic(err)
		}
		switch direction {
		case "forward":
			horizontal += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}
	return depth * horizontal
}

func part2(inputs []string) int64 {
	var depth, horizontal int64
	var aim int64
	for _, in := range inputs {
		res := strings.Split(in, " ")
		direction := res[0]
		distance, err := strconv.ParseInt(res[1], 10, 16)
		if err != nil {
			panic(err)
		}
		switch direction {
		case "forward":
			horizontal += distance
			depth += aim * distance
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}
	return depth * horizontal
}
