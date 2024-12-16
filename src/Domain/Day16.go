package Domain

import (
	"fmt"
	"sort"
	"time"
)

type D16Point struct {
	X, Y int
}

type D16Path struct {
	Score   int
	Head    D16Guard
	Visited map[D16Point]bool
	Trail   []D16Point
}

func (d D16Path) Clone() D16Path {
	path := D16Path{
		Head:    d.Head,
		Score:   d.Score,
		Visited: map[D16Point]bool{},
		Trail:   []D16Point{},
	}

	path.Trail = append(path.Trail, d.Trail...)

	for k, v := range d.Visited {
		path.Visited[k] = v
	}

	return path
}

type D16Tile rune

const (
	D16_WALL  D16Tile = '#'
	D16_OPEN  D16Tile = '.'
	D16_START D16Tile = 'S'
	D16_END   D16Tile = 'E'
)

type D16Direction rune

const (
	D16_NORTH D16Direction = 'N'
	D16_SOUTH D16Direction = 'S'
	D16_EAST  D16Direction = 'E'
	D16_WEST  D16Direction = 'W'
)

type D16Map struct {
	LivePaths D16Paths
	Map       [][]D16Tile
	Visited   map[D16Guard]int
}

type D16Guard struct {
	Position  D16Point
	Direction D16Direction
}

func (d *D16Guard) GetNeighbors() []D16Neighbor {
	neighs := []D16Neighbor{
		{Position: D16Point{X: d.Position.X, Y: d.Position.Y - 1}, Score: 1, Direction: D16_NORTH},
		{Position: D16Point{X: d.Position.X, Y: d.Position.Y + 1}, Score: 1, Direction: D16_SOUTH},
		{Position: D16Point{X: d.Position.X + 1, Y: d.Position.Y}, Score: 1, Direction: D16_EAST},
		{Position: D16Point{X: d.Position.X - 1, Y: d.Position.Y}, Score: 1, Direction: D16_WEST},
	}

	if d.Direction == D16_NORTH {
		neighs[2].Score += 1000
		neighs[3].Score += 1000
		neighs = append(neighs[:1], neighs[2:]...)
	} else if d.Direction == D16_SOUTH {
		neighs[2].Score += 1000
		neighs[3].Score += 1000
		neighs = neighs[1:]
	} else if d.Direction == D16_EAST {
		neighs[0].Score += 1000
		neighs[1].Score += 1000
		neighs = neighs[:3]
	} else if d.Direction == D16_WEST {
		neighs[0].Score += 1000
		neighs[1].Score += 1000
		neighs = append(neighs[:2], neighs[3:]...)
	}

	return neighs
}

func (d *D16Map) Move() *D16Path {
	path := d.LivePaths[0]
	//d.Print(path)

	d.LivePaths = d.LivePaths[1:]

	var foundPath *D16Path
	for _, neighbor := range path.Head.GetNeighbors() {
		if neighbor.Position.X < 0 || neighbor.Position.Y < 0 || neighbor.Position.Y >= len(d.Map) || neighbor.Position.X >= len(d.Map[neighbor.Position.Y]) {
			continue
		}

		if d.Visited[D16Guard{Position: neighbor.Position, Direction: neighbor.Direction}] > 0 && d.Visited[D16Guard{Position: neighbor.Position, Direction: neighbor.Direction}] < path.Score+neighbor.Score {
			//fmt.Printf("REMOVING PATH BCAUSE ARRIVED TO VISITED NODE WITH HIGHER SCORE\n")
			//fmt.Printf("OLD SCORE %d, NEW SCORE %d\n", d.Visited[D16Guard{Position: neighbor.Position, Direction: neighbor.Direction}], path.Score+neighbor.Score)
			continue
		}

		newPath := path.Clone()
		//newPath := path
		if d.Map[neighbor.Position.Y][neighbor.Position.X] != D16_WALL {
			newPath.Head = D16Guard{Position: neighbor.Position, Direction: neighbor.Direction}
			newPath.Trail = append(newPath.Trail, neighbor.Position)
			newPath.Score += neighbor.Score
			newPath.Visited[neighbor.Position] = true
			d.Visited[D16Guard{Position: neighbor.Position, Direction: neighbor.Direction}] = newPath.Score
			//fmt.Printf("best path has %d visited tiles\n", len(newPath.Visited))

			d.LivePaths = append(d.LivePaths, newPath)
		}

		if d.Map[neighbor.Position.Y][neighbor.Position.X] == D16_END {
			foundPath = &newPath
		}
	}

	sort.Sort(d.LivePaths)

	return foundPath
}

func (d *D16Map) Print(path D16Path) {
	fmt.Printf("Score %d, visited: %d, trail: %d\n", path.Score, len(path.Visited), len(path.Trail))
	for i := 0; i < len(d.Map); i++ {
		for j := 0; j < len(d.Map[i]); j++ {
			if path.Visited[D16Point{X: j, Y: i}] {
				fmt.Printf("O")
			} else {
				fmt.Printf("%c", d.Map[i][j])
			}
		}
		fmt.Printf("\n")
	}

	time.Sleep(50 * time.Millisecond)
}

type D16Neighbor struct {
	Position  D16Point
	Score     int
	Direction D16Direction
}

func NewD16Map() D16Map {
	return D16Map{LivePaths: make(D16Paths, 0), Map: make([][]D16Tile, 0), Visited: make(map[D16Guard]int)}
}

type D16Paths []D16Path

func (v D16Paths) Len() int {
	return len(v)
}

func (v D16Paths) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v D16Paths) Less(i, j int) bool {
	return v[i].Score < v[j].Score
}
