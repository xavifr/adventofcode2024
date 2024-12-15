package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Day14 struct {
	Bathroom Domain.D14Bathroom
}

func (d *Day14) Part1(input *bufio.Scanner) error {
	d.Bathroom = Domain.NewD14Bathroom()

	d.parseInput(input)

	cycles := 100
	for cycles > 0 {
		d.Bathroom.Move()
		cycles--
	}

	fmt.Printf("PART1: Bathroom value: %d\n", d.Bathroom.Calc())
	return nil
}

func (d *Day14) Part2(input *bufio.Scanner) error {
	d.Bathroom = Domain.NewD14Bathroom()

	d.parseInput(input)

	d.Bathroom.Print()

	cycles := 0
	for cycles < 1000 {
		fmt.Printf("Cycles : %d\n", cycles)
		d.Bathroom.Move()
		d.Bathroom.Print()
		cycles++
		//time.Sleep(25 * time.Millisecond)
	}

	fmt.Printf("PART2: Bathroom value: %d\n", d.Bathroom.Calc())
	return nil
}

func (d *Day14) parseInput(input *bufio.Scanner) {
	robotExpr, e := regexp.Compile(`^p=(\d+),(\d+) v=(\-?\d+),(\-?\d+)$`)
	if e != nil {
		fmt.Printf("Err : %v\n", e)
		return
	}

	numRobots := 0

	for input.Scan() {
		line := input.Text()
		matches := robotExpr.FindStringSubmatch(line)

		if len(matches) > 0 {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			vx, _ := strconv.Atoi(matches[3])
			vy, _ := strconv.Atoi(matches[4])

			point := Domain.D14Point{X: x, Y: y}
			d.Bathroom.Robots[point] = append(d.Bathroom.Robots[point], Domain.D14Robot{
				Vector: Domain.D14Point{X: vx, Y: vy},
			})
			numRobots++
		} else {
			fmt.Printf("ERR PARSE ROBOT %s\n", line)
		}
	}

	if numRobots < 20 {
		d.Bathroom.Size = Domain.D14Point{X: 11, Y: 7}
	} else {
		d.Bathroom.Size = Domain.D14Point{X: 101, Y: 103}
	}
}
