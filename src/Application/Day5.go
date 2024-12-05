package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Day5 struct {
	Printer Domain.D5Printer
}

func (d *Day5) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	middleSum := 0

	for _, book := range d.Printer.Books {
		correctOrder := true

		for i, page := range book {

			for j := i + 1; j < len(book); j++ {
				if rules, ok := d.Printer.Rules[book[j]]; ok {
					if slices.Index(rules, page) != -1 {
						correctOrder = false
						break
					}
				}
			}
			if !correctOrder {
				break
			}
		}

		if correctOrder {
			middleSum += book[(len(book)-1)/2]
		}

	}

	fmt.Printf("PART1: Middle sum: %d\n", middleSum)

	return nil
}

func (d *Day5) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	middleSum := 0

	for _, book := range d.Printer.Books {
		origCorrectOrder := true

		for {
			correctOrder := true
			for i := range book {

				if rules, ok := d.Printer.Rules[book[i]]; ok {
					moveTo := -1
					for j := i - 1; j >= 0; j-- {
						if slices.Index(rules, book[j]) != -1 {
							origCorrectOrder = false
							correctOrder = false
							moveTo = j
						}
					}

					if moveTo >= 0 {
						book[moveTo], book[i] = book[i], book[moveTo]
					}
				}
			}

			if correctOrder {
				break
			}
		}

		if !origCorrectOrder {
			middleSum += book[(len(book)-1)/2]
		}

	}

	fmt.Printf("PART2: Middle sum: %d\n", middleSum)

	return nil
}

func (d *Day5) parseInput(input *bufio.Scanner) {
	scanning := "rules"
	ruleExpr, _ := regexp.Compile(`^(\d+)\|(\d+)$`)

	d.Printer.Books = make([][]int, 0)
	d.Printer.Rules = make(map[int][]int, 0)

	for input.Scan() {
		line := input.Text()

		if line == "" && scanning == "rules" {
			scanning = "books"
			continue
		}

		if scanning == "rules" {
			matches := ruleExpr.FindStringSubmatch(line)
			if len(matches) > 0 {
				num1, _ := strconv.Atoi(matches[1])
				num2, _ := strconv.Atoi(matches[2])
				d.Printer.Rules[num1] = append(d.Printer.Rules[num1], num2)
			}
		} else if scanning == "books" {
			pages := strings.Split(line, ",")
			intPages := make([]int, 0)
			for _, page := range pages {
				num1, _ := strconv.Atoi(page)
				intPages = append(intPages, num1)

			}

			d.Printer.Books = append(d.Printer.Books, intPages)
		}
	}
}
