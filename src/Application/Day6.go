package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
)

type Day6 struct {
	Map Domain.D6Map
}

func (d *Day6) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	//_ = d.Map.MoveGuard()

	fmt.Printf("PART1: Guard moved %d\n", d.Map.VisitedPonts)

	return nil
}

func (d *Day6) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	countErrs := 0

	pivotMap := Domain.D6Map{Points: make([][]Domain.D6Object, len(d.Map.Points)), Guard: Domain.D6Guard{X: d.Map.Guard.X, Y: d.Map.Guard.Y, Direction: d.Map.Guard.Direction}}
	copy(pivotMap.Points, d.Map.Points)
	for k := 0; k < len(d.Map.Points); k++ {
		pivotMap.Points[k] = make([]Domain.D6Object, len(d.Map.Points[k]))
		copy(pivotMap.Points[k], d.Map.Points[k])
	}

	pivotMap.MoveGuard()

	for i := 0; i < len(d.Map.Points); i++ {
		for j := 0; j < len(d.Map.Points[i]); j++ {
			if pivotMap.Points[i][j] == Domain.D6_TRASH || pivotMap.Points[i][j] == Domain.D6_EMPTY || i == d.Map.Guard.Y && j == d.Map.Guard.X {
				continue
			}

			clonedMap := Domain.D6Map{Points: make([][]Domain.D6Object, len(d.Map.Points)), Guard: Domain.D6Guard{X: d.Map.Guard.X, Y: d.Map.Guard.Y, Direction: d.Map.Guard.Direction}}
			copy(clonedMap.Points, d.Map.Points)
			for k := 0; k < len(d.Map.Points); k++ {
				clonedMap.Points[k] = make([]Domain.D6Object, len(d.Map.Points[k]))
				copy(clonedMap.Points[k], d.Map.Points[k])
			}

			clonedMap.Points[i][j] = Domain.D6_OBSTACLE

			//clonedMap.Print()

			err := clonedMap.MoveGuard()
			if err != nil {
				//fmt.Printf("Increased error count with obstacle at %d,%d\n", i, j)
				countErrs++
			}
		}
	}

	fmt.Printf("PART2: Error count %d\n", countErrs)

	return nil
}

func (d *Day6) parseInput(input *bufio.Scanner) {
	d.Map = Domain.D6Map{Points: make([][]Domain.D6Object, 0), Guard: Domain.D6Guard{}}

	for input.Scan() {

		line := input.Text()

		row := make([]Domain.D6Object, len(line))
		for col, c := range line {
			switch c {
			case rune(Domain.D6_EMPTY):
				row[col] = Domain.D6_EMPTY
			case rune(Domain.D6_TRASH):
				row[col] = Domain.D6_TRASH
			default:
				d.Map.Guard.X = col
				d.Map.Guard.Y = len(d.Map.Points)
				d.Map.Guard.Direction = Domain.D6Direction(c)
				row[col] = Domain.D6_EMPTY
			}
		}

		d.Map.Points = append(d.Map.Points, row)
	}
}
