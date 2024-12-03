package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Day3 struct {
	operations []Domain.Day3Operation
}

func (d *Day3) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	accumulated := 0

	for _, op := range d.operations {
		if op.Operation == Domain.D3OpMul {
			accumulated += op.Num1 * op.Num2
		}
	}

	fmt.Printf("PART1: Accumulated value: %d\n", accumulated)
	return nil
}

func (d *Day3) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	accumulated := 0
	enabled := true
	for _, op := range d.operations {
		if op.Operation == Domain.D3OpMul && enabled {
			accumulated += op.Num1 * op.Num2
		} else if op.Operation == Domain.D3OpEn {
			enabled = true
		} else if op.Operation == Domain.D3OpDis {
			enabled = false
		}
	}

	fmt.Printf("PART2: Accumulated value: %d\n", accumulated)
	return nil
}

func (d *Day3) parseInput(input *bufio.Scanner) {
	d.operations = make([]Domain.Day3Operation, 0)

	expr, _ := regexp.Compile(`(do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\))`)
	mulExp, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doExp, _ := regexp.Compile(`do\(\)`)
	dontExp, _ := regexp.Compile(`don't\(\)`)
	for input.Scan() {
		line := input.Text()

		matches := expr.FindAllString(line, -1)
		for _, match := range matches {
			matchNums := mulExp.FindStringSubmatch(match)
			if len(matchNums) > 0 {
				num1, _ := strconv.Atoi(matchNums[1])
				num2, _ := strconv.Atoi(matchNums[2])

				d.operations = append(d.operations, Domain.Day3Operation{
					Operation: Domain.D3OpMul,
					Num1:      num1,
					Num2:      num2,
				})
				continue
			}

			matchDo := doExp.FindStringSubmatch(match)
			if len(matchDo) > 0 {
				d.operations = append(d.operations, Domain.Day3Operation{
					Operation: Domain.D3OpEn,
				})
				continue
			}

			matchDont := dontExp.FindStringSubmatch(match)
			if len(matchDont) > 0 {
				d.operations = append(d.operations, Domain.Day3Operation{
					Operation: Domain.D3OpDis,
				})
				continue
			}
		}
	}
}
