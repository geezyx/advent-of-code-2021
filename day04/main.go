package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board struct {
	index map[int64]int
	spots []spot
}

func (b *board) pick(num int64) bool {
	if i, found := b.index[num]; found {
		// fmt.Printf("found pick %s at index %d\n", num, i)
		b.spots[i].picked = true
		return true
	}
	return false
}

func (b *board) results(pickedNum int64) int64 {
	var total int64
	for _, spot := range b.spots {
		if !spot.picked {
			total += spot.number
		}
	}
	return total * pickedNum
}

type spot struct {
	number int64
	picked bool
}

func (b board) check() bool {
	return b.checkHorizontals() || b.checkVerticals()
}

func (b board) checkHorizontals() bool {
	for i := 0; i < 5; i++ {
		if b.spots[5*i+0].picked &&
			b.spots[5*i+1].picked &&
			b.spots[5*i+2].picked &&
			b.spots[5*i+3].picked &&
			b.spots[5*i+4].picked {
			return true
		}
	}
	return false
}

func (b board) checkVerticals() bool {
	for i := 0; i < 5; i++ {
		if b.spots[i].picked &&
			b.spots[i+5].picked &&
			b.spots[i+10].picked &&
			b.spots[i+15].picked &&
			b.spots[i+20].picked {
			return true
		}
	}
	return false
}

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

func part1(inputs []string) int64 {
	picks, boards := loadInput(inputs)

	for _, pick := range picks {
		for _, b := range boards {
			if found := b.pick(pick); found {
				if b.check() {
					return b.results(pick)
				}
			}
		}
	}
	return int64(0)
}

func part2(inputs []string) int64 {
	picks, boards := loadInput(inputs)

	for _, pick := range picks {
		nonWinners := []board{}
		for _, b := range boards {
			if found := b.pick(pick); found {
				if b.check() {
					if len(boards) == 1 {
						return boards[0].results(pick)
					}
					continue
				}
			}
			nonWinners = append(nonWinners, b)
		}
		boards = nonWinners
	}

	return int64(0)
}

func loadInput(input []string) ([]int64, []board) {
	picks := strings.Split(input[0], ",")
	var pickList []int64
	for _, p := range picks {
		val, err := strconv.ParseInt(p, 10, 8)
		if err != nil {
			panic(err)
		}
		pickList = append(pickList, val)
	}
	boards := []board{}
	i := 2
	for {
		boards = append(boards, loadBoard(input[i:i+5]))
		i += 6
		if i > len(input) {
			break
		}
	}
	return pickList, boards
}

func loadBoard(lines []string) board {
	b := board{
		index: make(map[int64]int),
		spots: []spot{},
	}
	for i, line := range lines {
		chars := strings.Split(line, "")
		for j := 0; j < 5; j++ {
			num := strings.Join(chars[(3*j):(3*j)+2], "")
			num = strings.TrimSpace(num)
			val, err := strconv.ParseInt(num, 10, 8)
			if err != nil {
				panic(err)
			}
			b.index[val] = (5 * i) + j
			b.spots = append(b.spots, spot{
				number: val,
				picked: false,
			})
		}
	}
	return b
}
