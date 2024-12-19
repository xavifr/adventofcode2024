package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"strings"
)

type Day19 struct {
	Store Domain.D19Store
}

func (d *Day19) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	towelsPossible := 0

	for i := 0; i < len(d.Store.Patterns); i++ {
		if d.Store.FindPattern(d.Store.Patterns[i]) {
			towelsPossible++
		}
	}

	fmt.Printf("PART1: List of possible towels is %d\n", towelsPossible)
	return nil
}

func (d *Day19) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	towelsPossible := 0

	for i := 0; i < len(d.Store.Patterns); i++ {
		towelsPossible += d.Store.FindPatternRecurse(d.Store.Patterns[i])
	}

	fmt.Printf("PART2: List of possibilities is %d\n", towelsPossible)
	return nil
}

func (d *Day19) parseInput(input *bufio.Scanner) {
	d.Store = Domain.D19Store{
		Towels: make(map[string]bool),

		Patterns: make([]string, 0),
	}

	for input.Scan() {
		line := input.Text()

		if strings.Contains(line, ",") {
			availableTowels := strings.Split(line, ", ")
			for _, availableTowel := range availableTowels {
				d.Store.Towels[availableTowel] = true
			}
		} else if len(line) > 0 {
			d.Store.Patterns = append(d.Store.Patterns, line)
		}
	}
}
