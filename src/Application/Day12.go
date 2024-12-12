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

	grassRegions := d.getRegions()

	sum := 0
	for _, region := range grassRegions {
		sum += region.GetArea() * region.GetPerimeter()
	}

	fmt.Printf("PART1: sum num %d\n", sum)
	return nil
}

func popElement(slice map[Domain.D12Point]bool) (Domain.D12Point, bool) {
	for key := range slice {
		delete(slice, key)
		return key, true
	}

	return Domain.D12Point{}, false
}
func (d *Day12) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	grassRegions := d.getRegions()

	sum := 0
	for _, region := range grassRegions {
		sum += region.GetArea() * region.GetSides()
	}

	fmt.Printf("PART2: sum num %d\n", sum)
	return nil
}

func (d *Day12) getRegions() []Domain.D12Region {
	grassRegions := make([]Domain.D12Region, 0)

	for {
		nextRegion := d.Map.GetFirstNonVisited()
		if nextRegion == nil {
			break
		}

		//fmt.Printf("New region %c at %d:%d\n", d.Map.GetChar(*nextRegion), (*nextRegion).X, (*nextRegion).Y)

		sameRegion := map[Domain.D12Point]bool{*nextRegion: true}

		grassRegion := Domain.D12Region{Cells: make(map[Domain.D12Point]bool, 0)}

		for {
			point, found := popElement(sameRegion)
			if !found {
				break
			}

			d.Map.SetVisited(point)

			grassRegion.Cells[point] = true

			for _, neigh := range d.Map.GetNeighbors(point) {
				if d.Map.GetChar(neigh) == d.Map.GetChar(point) {
					//fmt.Printf("  add point at %d:%d\n", neigh.X, neigh.Y)
					sameRegion[neigh] = true
				}
			}
		}

		grassRegions = append(grassRegions, grassRegion)
	}

	return grassRegions
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
