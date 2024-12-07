package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day7 struct {
	Operations []Domain.D7Operation
}

func (d *Day7) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	sumPossibleOperations := 0

	for _, operation := range d.Operations {
		if operation.IsPossibleTwo() {
			sumPossibleOperations += operation.Result
		}
	}

	fmt.Printf("PART1: possible operations %v\n", sumPossibleOperations)

	return nil
}

func (d *Day7) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	sumPossibleOperations := 0

	for _, operation := range d.Operations {
		if operation.IsPossibleThree() {
			sumPossibleOperations += operation.Result
		}
	}

	fmt.Printf("PART2: possible operations %v\n", sumPossibleOperations)

	return nil
}

func (d *Day7) parseInput(input *bufio.Scanner) {
	d.Operations = make([]Domain.D7Operation, 0)

	operationExpression, _ := regexp.Compile(`^(\d+): (.*)$`)

	for input.Scan() {
		line := input.Text()

		matches := operationExpression.FindStringSubmatch(line)
		if len(matches) > 0 {
			num1, _ := strconv.Atoi(matches[1])

			operation := Domain.D7Operation{
				Result:   num1,
				Operands: make([]int, 0),
			}

			rest := strings.Split(matches[2], " ")
			for _, operand := range rest {
				num2, _ := strconv.Atoi(operand)
				operation.Operands = append(operation.Operands, num2)
			}

			d.Operations = append(d.Operations, operation)
		}
	}
}
