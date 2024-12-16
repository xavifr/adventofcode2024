package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
)

type Day16 struct {
	Map Domain.D16Map
}

func (d *Day16) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	var path *Domain.D16Path
	for {
		path = d.Map.Move()
		if path != nil {
			break
		}

	}

	fmt.Printf("PART1: Good path has score: %d\n", path.Score)

	return nil
}

func (d *Day16) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	goodPaths := []Domain.D16Path{}
	for len(d.Map.LivePaths) > 0 {
		path := d.Map.Move()
		if path != nil {
			goodPaths = append(goodPaths, *path)

			if len(goodPaths) > 0 && d.Map.LivePaths[0].Score >= goodPaths[0].Score {
				break
			}

		}
	}

	tilesUsed := map[Domain.D16Point]bool{}
	for _, path := range goodPaths {
		/*if d.Map.Map[path.Head.Position.Y][path.Head.Position.X] == Domain.D16_END {
			for _, point := range path.Trail {
				tilesUsed[point] = true
			}
		}*/
		for k := range path.Visited {
			tilesUsed[k] = true
		}
	}

	fmt.Printf("PART2: Good paths had tiles: %d\n", len(tilesUsed))

	return nil
}

func (d *Day16) parseInput(input *bufio.Scanner) {
	d.Map = Domain.NewD16Map()

	for input.Scan() {
		line := input.Text()

		row := []Domain.D16Tile{}
		for col, c := range line {
			row = append(row, Domain.D16Tile(c))
			if Domain.D16Tile(c) == Domain.D16_START {
				d.Map.LivePaths = append(d.Map.LivePaths, Domain.D16Path{
					Head: Domain.D16Guard{
						Position:  Domain.D16Point{X: col, Y: len(d.Map.Map)},
						Direction: Domain.D16_EAST},
					Visited: map[Domain.D16Point]bool{
						{X: col, Y: len(d.Map.Map)}: true,
					},
					Trail: []Domain.D16Point{
						{X: col, Y: len(d.Map.Map)},
					},
				})
			}
		}

		d.Map.Map = append(d.Map.Map, row)
	}
}
