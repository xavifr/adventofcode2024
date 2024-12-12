package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
)

type Day12 struct {
	Map Domain.D12Map
}

func (d *Day12) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	grassRegions := make([]Domain.D12Region, 0)
	nextRegion := []Domain.D12Point{{X: 0, Y: 0}}

	for len(nextRegion) > 0 {
		if d.Map.GetVisited(nextRegion[0]) {
			nextRegion = nextRegion[1:]
			continue
		}

		sameRegion := []Domain.D12Point{nextRegion[0]}
		nextRegion = nextRegion[1:]

		grassRegion := Domain.D12Region{Cells: make([]Domain.D12Point, 0)}

		for len(sameRegion) > 0 {
			point := sameRegion[0]
			sameRegion = sameRegion[1:]

			grassRegion.Cells = append(grassRegion.Cells, point)

			for _, neigh := range d.Map.GetNeighbors(point) {
				if d.Map.GetChar(neigh) == d.Map.GetChar(point) {
					sameRegion = append(sameRegion, neigh)
				} else {
					nextRegion = append(nextRegion, neigh)
				}
			}
		}

		grassRegions = append(grassRegions, grassRegion)
	}

	fmt.Printf("PART1: len regions %d\n", len(grassRegions))
	return nil
}

func (d *Day12) Part2(input *bufio.Scanner) error {
	return nil
}

func (d *Day12) parseInput(input *bufio.Scanner) {
	d.Map = Domain.NewD12Map()

	row := 0
	for input.Scan() {
		line := input.Text()

		for col, c := range line {
			d.Map.Map[Domain.D12Point{X: col, Y: row}] = &Domain.D12Cell{Visited: false, Char: c}
		}

		row++
	}
}
