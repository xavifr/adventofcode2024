package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
)

type Day4 struct {
	Soup Domain.Day4Soup
}

func (d *Day4) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	words := 0

	directions := []Domain.Day4Direction{Domain.D4_N, Domain.D4_S, Domain.D4_E, Domain.D4_W, Domain.D4_NE, Domain.D4_NW, Domain.D4_SE, Domain.D4_SW}
	positions := d.Soup.FindChar('X')
	for _, pos := range positions {
		for _, direction := range directions {
			word := d.Soup.GetWord(pos.X, pos.Y, 4, direction)
			if word == "XMAS" {
				words++
			}
		}
	}

	fmt.Printf("PART1: Number of XMAS: %d\n", words)

	return nil
}

func (d *Day4) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	words := 0

	positions := d.Soup.FindChar('A')
	for _, pos := range positions {
		found := 0
		word := d.Soup.GetWord(pos.X-1, pos.Y-1, 3, Domain.D4_SE)
		if word == "MAS" || word == "SAM" {
			found++
		}

		word = d.Soup.GetWord(pos.X-1, pos.Y+1, 3, Domain.D4_SW)
		if word == "MAS" || word == "SAM" {
			found++
		}

		if found == 2 {
			words++
		}
	}

	fmt.Printf("PART2: Number of X-MAS: %d\n", words)
	return nil
}

func (d *Day4) parseInput(input *bufio.Scanner) {
	d.Soup = Domain.NewDay4Soup()

	for input.Scan() {
		line := input.Text()
		d.Soup.Append(line)
	}
}
