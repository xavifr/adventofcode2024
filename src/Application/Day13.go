package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Day13 struct {
	prizes []Domain.D13Prize
}

func (d *Day13) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	sumCoins := 0
	for _, prize := range d.prizes {
		if minA, minB, possible := prize.Resolve(false); possible && minA <= 100 && minB <= 100 {
			sumCoins += 3*minA + 1*minB
		}
	}

	fmt.Printf("PART1: Sum of coins: %d\n", sumCoins)

	return nil
}

func (d *Day13) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	sumCoins := 0
	for _, prize := range d.prizes {
		if minA, minB, possible := prize.Resolve(true); possible {
			sumCoins += 3*minA + 1*minB
		}
	}

	fmt.Printf("PART2: Sum of coins: %d\n", sumCoins)

	return nil
}

func (d *Day13) parseInput(input *bufio.Scanner) {
	d.prizes = make([]Domain.D13Prize, 0)

	exprButtonA, _ := regexp.Compile(`^Button A: X\+(\d+). Y\+(\d+)$`)
	exprButtonB, _ := regexp.Compile(`^Button B: X\+(\d+). Y\+(\d+)$`)
	exprPrize, _ := regexp.Compile(`^Prize: X=(\d+). Y=(\d+)$`)

	var lastPriceX, lastPriceY, lastAX, lastAY, lastBX, lastBY int
	for input.Scan() {
		line := input.Text()

		if matches := exprButtonA.FindStringSubmatch(line); matches != nil {
			lastAX, _ = strconv.Atoi(matches[1])
			lastAY, _ = strconv.Atoi(matches[2])
		} else if matches := exprButtonB.FindStringSubmatch(line); matches != nil {
			lastBX, _ = strconv.Atoi(matches[1])
			lastBY, _ = strconv.Atoi(matches[2])
		} else if matches := exprPrize.FindStringSubmatch(line); matches != nil {
			lastPriceX, _ = strconv.Atoi(matches[1])
			lastPriceY, _ = strconv.Atoi(matches[2])
			d.prizes = append(d.prizes, Domain.NewD13Prize(lastPriceX, lastPriceY, lastAX, lastAY, lastBX, lastBY))
		}
	}
}
