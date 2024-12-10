package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"strconv"
)

type Day10 struct {
	TrailMap Domain.D10TrailMap
}

func (d *Day10) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	score := 0
	for _, trailhead := range d.TrailMap.Trailheads {
		score += d.TrailMap.GetScore(trailhead, true)
	}

	fmt.Printf("PART1: Score is %d\n", score)

	return nil
}

func (d *Day10) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	score := 0
	for _, trailhead := range d.TrailMap.Trailheads {
		score += d.TrailMap.GetScore(trailhead, false)
	}

	fmt.Printf("PART2: Score is %d\n", score)

	return nil
}

func (d *Day10) parseInput(input *bufio.Scanner) {
	d.TrailMap = Domain.D10TrailMap{
		Trailheads: make([]Domain.D10Point, 0),
		Map:        make([][]int, 0),
	}

	for input.Scan() {
		line := input.Text()

		row := make([]int, 0)
		for x, c := range line {
			num, _ := strconv.Atoi(string(c))

			row = append(row, num)
			if num == 0 {
				d.TrailMap.Trailheads = append(d.TrailMap.Trailheads, Domain.D10Point{X: x, Y: len(d.TrailMap.Map)})
			}
		}

		d.TrailMap.Map = append(d.TrailMap.Map, row)
	}
}
