package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
)

type Day15 struct {
	Warehouse Domain.D15Warehouse
}

func (d *Day15) Part1(input *bufio.Scanner) error {
	d.parseInputPart1(input)

	for len(d.Warehouse.Movements) > 0 {
		d.Warehouse.Move()
	}

	fmt.Printf("PART1: Warehouse value: %d\n", d.Warehouse.Value())

	return nil
}

func (d *Day15) Part2(input *bufio.Scanner) error {
	d.parseInputPart2(input)

	for len(d.Warehouse.Movements) > 0 {
		d.Warehouse.Move2()
		//d.Warehouse.Print()
	}

	fmt.Printf("PART2: Warehouse value: %d\n", d.Warehouse.Value())

	return nil
}

func (d *Day15) parseInputPart1(input *bufio.Scanner) {
	d.Warehouse = Domain.NewD15Warehouse()

	row := 0
	for input.Scan() {
		line := input.Text()

		if line == "" {
			continue
		}

		if line[0] == '#' {
			for col, c := range line {
				if c == '.' {
					continue
				}

				if c == '@' {
					d.Warehouse.Robot = Domain.D15Point{X: col, Y: row}
				} else {
					d.Warehouse.Map[Domain.D15Point{X: col, Y: row}] = Domain.D15Object(c)
				}
			}
			row++
		} else {
			for _, c := range line {
				d.Warehouse.Movements = append(d.Warehouse.Movements, Domain.D15Movement(c))
			}
		}
	}

}

func (d *Day15) parseInputPart2(input *bufio.Scanner) {
	d.Warehouse = Domain.NewD15Warehouse()

	colSz := 0
	row := 0
	for input.Scan() {
		line := input.Text()

		if line == "" {
			continue
		}

		if line[0] == '#' {
			colSz = len(line)
			for col, c := range line {
				if c == '.' {
					continue
				}

				if c == '@' {
					d.Warehouse.Robot = Domain.D15Point{X: col * 2, Y: row}
				} else if c == '#' {
					d.Warehouse.Map[Domain.D15Point{X: col * 2, Y: row}] = Domain.D15Object(c)
					d.Warehouse.Map[Domain.D15Point{X: col*2 + 1, Y: row}] = Domain.D15Object(c)
				} else {
					// Box
					d.Warehouse.Map[Domain.D15Point{X: col * 2, Y: row}] = Domain.D15Object(c)
				}
			}
			row++
		} else {
			for _, c := range line {
				d.Warehouse.Movements = append(d.Warehouse.Movements, Domain.D15Movement(c))
			}
		}
	}

	d.Warehouse.Size = Domain.D15Point{X: colSz * 2, Y: row}

}
