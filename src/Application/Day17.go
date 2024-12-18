package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Day17 struct {
	Machine Domain.D17Machine
}

func (d *Day17) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	d.Machine.Run()

	fmt.Printf("PART1: State: %s, IP: %d\n", d.Machine.State, d.Machine.IP)
	fmt.Printf("OUTPUT: ")
	for i, out := range d.Machine.GetOutput() {
		if i > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("%d", out)
	}
	fmt.Printf("\n")

	return nil
}

func (d *Day17) Part2(input *bufio.Scanner) error {
	// based on https://github.com/michel-kraemer/adventofcode-rust/blob/main/2024/day17/src/main.rs
	
	d.parseInput(input)

	if len(d.Machine.Program) == 6 { // skip demo
		return nil
	}

	iniB := d.Machine.RegB
	iniC := d.Machine.RegC

	factors := make([]int, len(d.Machine.Program))

	for {
		testA := int64(0)
		for i, j := range factors {
			shifter := int(math.Pow(8, float64(i)))
			shifted := shifter * j
			testA += int64(shifted)
		}

		d.Machine.RegA = testA
		d.Machine.RegB = iniB
		d.Machine.RegC = iniC
		d.Machine.State = Domain.D17_RUNNING
		d.Machine.IP = 0
		d.Machine.Output = make([]int, 0)
		d.Machine.Run()

		if reflect.DeepEqual(d.Machine.Program, d.Machine.Output) {
			break
		}

		for i := len(d.Machine.Program) - 1; i >= 0; i-- {
			if len(d.Machine.Output) < i+1 {
				factors[i]++
				break
			}

			if d.Machine.Program[i] != d.Machine.Output[i] {
				factors[i]++
				break
			}
		}
	}

	willUseA := int64(0)
	for i, j := range factors {
		shifter := int(math.Pow(8, float64(i)))
		shifted := shifter * j
		willUseA += int64(shifted)
	}

	fmt.Printf("PART2: Will use A: %d\n", willUseA)
	return nil
}

func (d *Day17) parseInput(input *bufio.Scanner) {
	d.Machine = Domain.NewD17Machine()

	exprRegA, _ := regexp.Compile(`^Register A: (\d+)$`)
	exprRegB, _ := regexp.Compile(`^Register B: (\d+)$`)
	exprRegC, _ := regexp.Compile(`^Register C: (\d+)$`)
	exprProg, _ := regexp.Compile(`Program: ([\d,]+)$`)

	for input.Scan() {
		line := input.Text()

		matches := exprRegA.FindStringSubmatch(line)
		if len(matches) > 0 {
			num, _ := strconv.Atoi(matches[1])
			d.Machine.RegA = int64(num)
		}

		matches = exprRegB.FindStringSubmatch(line)
		if len(matches) > 0 {
			num, _ := strconv.Atoi(matches[1])
			d.Machine.RegB = int64(num)
		}

		matches = exprRegC.FindStringSubmatch(line)
		if len(matches) > 0 {
			num, _ := strconv.Atoi(matches[1])
			d.Machine.RegC = int64(num)
		}

		matches = exprProg.FindStringSubmatch(line)
		if len(matches) > 0 {
			for _, numStr := range strings.Split(matches[1], ",") {
				num, _ := strconv.Atoi(numStr)
				d.Machine.Program = append(d.Machine.Program, num)
			}
		}
	}
}
