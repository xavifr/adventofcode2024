package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
)

type Day8 struct {
	Map Domain.D8Map
}

func (d *Day8) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	for _, frequency := range d.Map.Frequencies {
		for i := 0; i < len(frequency); i++ {
			for j := i + 1; j < len(frequency); j++ {
				for _, antinode := range frequency[i].GetAntinode(frequency[j]) {
					d.Map.AddAntinode(antinode)
				}
			}
		}
	}

	fmt.Printf("PART1: Number of antinodes: %d\n", d.Map.AntinodesCount)
	return nil
}

func (d *Day8) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	for _, frequency := range d.Map.Frequencies {
		for i := 0; i < len(frequency); i++ {
			for j := i + 1; j < len(frequency); j++ {
				d.Map.AddAntinodeVector(frequency[i], frequency[i].GetVector(frequency[j]))
			}
		}
	}

	fmt.Printf("PART2: Number of antinodes: %d\n", d.Map.AntinodesCount)
	return nil
}

func (d *Day8) parseInput(input *bufio.Scanner) {
	d.Map = Domain.D8Map{
		Antinodes:   make([][]bool, 0),
		Frequencies: make(map[rune][]Domain.D8Point, 0),
	}

	row := 0
	for input.Scan() {
		line := input.Text()

		d.Map.Antinodes = append(d.Map.Antinodes, make([]bool, len(line)))

		for col, c := range line {
			if c == '.' {
				continue
			}
			if _, ok := d.Map.Frequencies[c]; !ok {
				d.Map.Frequencies[c] = make([]Domain.D8Point, 0)
			}

			d.Map.Frequencies[c] = append(d.Map.Frequencies[c], Domain.D8Point{X: col, Y: row})
		}

		row++
	}
}
