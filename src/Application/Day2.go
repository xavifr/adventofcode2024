package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Day2 struct {
	reports []Domain.Day2Report
}

func (d *Day2) Part1(input *bufio.Scanner) error {
	return nil
	d.parseInput(input)

	numSafe := 0
	for _, report := range d.reports {
		if report.Safe(false) {
			//fmt.Printf("PART1: Safe report: %v\n", report)
			numSafe++
		}
	}

	fmt.Printf("PART1: Number of safe reports: %d\n", numSafe)
	return nil
}

func (d *Day2) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	numSafe := 0
	for _, report := range d.reports {
		if report.Safe(true) {
			//fmt.Printf("PART2: Safe report: %v\n", report)
			numSafe++
		}
	}

	fmt.Printf("PART2: Number of safe reports: %d\n", numSafe)
	return nil
}

func (d *Day2) parseInput(input *bufio.Scanner) {
	d.reports = make([]Domain.Day2Report, 0)

	for input.Scan() {
		line := input.Text()

		levelsStr := strings.Split(line, " ")
		levels := make([]int, 0)
		for _, levelStr := range levelsStr {
			lvl, _ := strconv.Atoi(levelStr)
			levels = append(levels, lvl)
		}

		d.reports = append(d.reports, Domain.Day2Report{Levels: levels})
	}
}
