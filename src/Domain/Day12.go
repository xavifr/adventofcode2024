package Domain

type D12Point struct {
	X, Y int
}

func (d *D12Point) GetNeighbors() []D12Point {
	return []D12Point{
		D12Point{X: d.X + 1, Y: d.Y},
		D12Point{X: d.X - 1, Y: d.Y},
		D12Point{X: d.X, Y: d.Y + 1},
		D12Point{X: d.X, Y: d.Y - 1},
	}
}

type D12Region struct {
	Cells []D12Point
}

type D12Cell struct {
	Visited bool
	Char    rune
}

type D12Map struct {
	Map map[D12Point]*D12Cell
}

func (d *D12Map) GetChar(point D12Point) rune {
	return d.Map[point].Char
}

func (d *D12Map) GetVisited(point D12Point) bool {
	return d.Map[point].Visited
}

func (d *D12Map) GetNeighbors(point D12Point) []D12Point {
	points := []D12Point{}

	for _, neigh := range point.GetNeighbors() {
		if nextPoint, ok := d.Map[neigh]; ok {
			if !nextPoint.Visited {
				points = append(points, neigh)
			}

		}
	}

	// TODO: set actual point as visited

	return points
}

func NewD12Map() D12Map {
	return D12Map{Map: make(map[D12Point]*D12Cell, 0)}
}
