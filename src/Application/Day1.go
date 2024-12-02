package Application

import (
	"bufio"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

type Day1 struct {
	list1 []int
	list2 []int

	map1 map[int]int
	map2 map[int]int
}

func (d *Day1) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	slices.Sort(d.list1)
	slices.Sort(d.list2)

	var dist []int
	for i := 0; i < len(d.list1); i++ {
		diff := d.list1[i] - d.list2[i]
		if diff < 0 {
			diff *= -1
		}

		dist = append(dist, diff)
	}

	sumDist := 0
	for i := 0; i < len(dist); i++ {
		sumDist += dist[i]
	}

	fmt.Printf("PART1: Sum of distances: %d\n", sumDist)

	return nil
}

func (d *Day1) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	similarities := 0
	for num, times := range d.map1 {
		if _, ok := d.map2[num]; ok {
			similarities += times * num * d.map2[num]
		}
	}

	fmt.Printf("PART1: Sum of similarities: %d\n", similarities)

	return nil
}

func (d *Day1) parseInput(input *bufio.Scanner) {
	expr, _ := regexp.Compile(`(\d+)\s+(\d+)`)

	d.list1 = make([]int, 0)
	d.list2 = make([]int, 0)
	d.map1 = make(map[int]int)
	d.map2 = make(map[int]int)

	for input.Scan() {
		line := input.Text()
		matches := expr.FindStringSubmatch(line)
		num1, _ := strconv.Atoi(matches[1])
		num2, _ := strconv.Atoi(matches[2])
		d.list1 = append(d.list1, num1)
		d.list2 = append(d.list2, num2)

		if _, ok := d.map1[num1]; !ok {
			d.map1[num1] = 0
		}

		if _, ok := d.map2[num2]; !ok {
			d.map2[num2] = 0
		}

		d.map1[num1]++
		d.map2[num2]++
	}

}
