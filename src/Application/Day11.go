package Application

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day11 struct {
	Stones    []int
	MapStones map[int]int64
}

func (d *Day11) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	cycles := 25

	for cycles > 0 {
		cycles--

		newStones := make([]int, 0)
		for _, stone := range d.Stones {
			newStones = append(newStones, replaceStone(stone)...)
		}
		d.Stones = newStones
	}

	fmt.Printf("PART1: Number of stones is %d\n", len(d.Stones))
	return nil
}

func replaceStone(num int) []int {
	if num == 0 {
		return []int{1}
	}

	if num > 1 && int(math.Log10(float64(num))+1)%2 == 0 {
		str := fmt.Sprintf("%d", num)

		num1, _ := strconv.Atoi(str[:(len(str) / 2)])
		num2, _ := strconv.Atoi(str[(len(str) / 2):])
		return []int{num1, num2}
	}
	return []int{num * 2024}

}

func (d *Day11) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	cycles := 75

	for cycles > 0 {
		newMapStones := make(map[int]int64)
		for stone, occurrences := range d.MapStones {
			newStones := replaceStone(stone)
			for _, newStone := range newStones {
				newMapStones[newStone] += occurrences
			}
		}

		d.MapStones = newMapStones

		cycles--
	}

	sumStones := int64(0)
	for _, occurrences := range d.MapStones {
		sumStones += occurrences
	}

	fmt.Printf("PART2: Number of stones is %d, distinct %d\n", sumStones, len(d.MapStones))
	return nil
}

func (d *Day11) parseInput(input *bufio.Scanner) {
	d.Stones = make([]int, 0)
	d.MapStones = make(map[int]int64)

	for input.Scan() {
		line := input.Text()

		for _, stone := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(stone)
			d.Stones = append(d.Stones, num)
			if _, ok := d.MapStones[num]; !ok {
				d.MapStones[num] = 1
			} else {
				d.MapStones[num]++
			}
		}
	}
}
