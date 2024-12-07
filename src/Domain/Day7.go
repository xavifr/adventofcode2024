package Domain

import (
	"fmt"
	"math"
	"strconv"
)

type D7Operation struct {
	Result   int
	Operands []int
}

func (d *D7Operation) IsPossibleTwo() bool {
	for i := 0.0; i < math.Pow(2, float64(len(d.Operands)-1)); i++ {
		if d.testPossibilityTwo(int(i)) {
			return true
		}
	}

	return false
}

func (d *D7Operation) IsPossibleThree() bool {
	for i := 0.0; i < math.Pow(3, float64(len(d.Operands)-1)); i++ {
		if d.testPossibilityThree(int(i)) {
			return true
		}
	}

	return false
}

func (d *D7Operation) testPossibilityTwo(test int) bool {
	accumulator := d.Operands[0]
	for i, operand := range d.Operands[1:] {
		if test>>i&1 == 0 {
			accumulator += operand
		} else {
			accumulator *= operand
		}

		if accumulator == d.Result {
			return true
		} else if accumulator > d.Result {
			break
		}
	}
	return false
}

func (d *D7Operation) testPossibilityThree(test int) bool {
	accumulator := d.Operands[0]
	for i, operand := range d.Operands[1:] {
		tval := ternary(test, i)
		switch tval {
		case 0:
			accumulator += operand
		case 1:
			accumulator *= operand
		case 2:
			t := fmt.Sprintf("%d%d", accumulator, operand)
			accumulator, _ = strconv.Atoi(t)
		}

		if accumulator > d.Result {
			break
		}
	}

	return accumulator == d.Result
}

func ternary(number, position int) int {
	for position > 0 {
		number /= 3
		position--
	}

	return number % 3
}
