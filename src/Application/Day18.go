package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Day18 struct {
	RAM Domain.D18Ram
}

func (d *Day18) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	var bestPath int
	if len(d.RAM.CorruptionOrder) > 40 {
		// final
		d.RAM.CorruptBytes(1024)
		bestPath = d.RAM.BestPath(Domain.D18Point{0, 0}, Domain.D18Point{70, 70})
	} else {
		// demo
		d.RAM.CorruptBytes(12)
		bestPath = d.RAM.BestPath(Domain.D18Point{0, 0}, Domain.D18Point{6, 6})
	}

	fmt.Printf("PART1: Best path has %d steps\n", bestPath)

	return nil
}

func (d *Day18) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	var corruptedBytes int
	if len(d.RAM.CorruptionOrder) < 40 { // skip demo
		corruptedBytes = 12
	} else {
		corruptedBytes = 1024
	}

	for {
		corruptedBytes++
		d.RAM.CorruptBytes(corruptedBytes)

		var bestPath int
		if len(d.RAM.CorruptionOrder) > 40 {
			// final
			bestPath = d.RAM.BestPath(Domain.D18Point{0, 0}, Domain.D18Point{70, 70})
		} else {
			// demo
			bestPath = d.RAM.BestPath(Domain.D18Point{0, 0}, Domain.D18Point{6, 6})
		}

		if bestPath == -1 { // no solution
			break
		}
	}

	fmt.Printf("PART2: There are no exit path after corrupting %d,%d steps\n", d.RAM.CorruptionOrder[corruptedBytes-1].X, d.RAM.CorruptionOrder[corruptedBytes-1].Y)

	return nil
}

func (d *Day18) parseInput(input *bufio.Scanner) {
	d.RAM = Domain.NewD18Ram()

	for input.Scan() {
		line := input.Text()

		point := strings.Split(line, ",")
		if len(point) != 2 {
			continue
		}

		num1, _ := strconv.Atoi(point[0])
		num2, _ := strconv.Atoi(point[1])
		d.RAM.CorruptionOrder = append(d.RAM.CorruptionOrder, Domain.D18Point{X: num1, Y: num2})
	}

	if len(d.RAM.CorruptionOrder) < 40 {
		d.RAM.Size = Domain.D18Point{7, 7}
	} else {
		d.RAM.Size = Domain.D18Point{71, 71}
	}

}
