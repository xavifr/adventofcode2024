package Domain

type D12Point struct {
	X, Y int
}

func (d *D12Point) GetNeighbors() []D12Point {
	return []D12Point{
		{X: d.X + 1, Y: d.Y},
		{X: d.X - 1, Y: d.Y},
		{X: d.X, Y: d.Y + 1},
		{X: d.X, Y: d.Y - 1},
	}
}

type D12Region struct {
	Cells map[D12Point]bool
}

func (d *D12Region) GetArea() int {
	return len(d.Cells)
}

func (d *D12Region) GetPerimeter() int {
	perimeter := 0
	for point := range d.Cells {
		for _, neigh := range point.GetNeighbors() {
			if _, ok := d.Cells[neigh]; !ok {
				perimeter++
			}
		}
	}

	return perimeter
}

func (d *D12Region) GetSides() int {
	sides := 0

	cells := make(map[D12Point]bool, 0)
	for k, v := range d.Cells {
		cells[k] = v
	}

	for element := range d.Cells {
		if !cells[D12Point{element.X, element.Y + 1}] && !cells[D12Point{element.X + 1, element.Y}] {
			sides++
		}

		if !cells[D12Point{element.X, element.Y + 1}] && !cells[D12Point{element.X - 1, element.Y}] {
			sides++
		}

		if !cells[D12Point{element.X, element.Y - 1}] && !cells[D12Point{element.X - 1, element.Y}] {
			sides++
		}

		if !cells[D12Point{element.X, element.Y - 1}] && !cells[D12Point{element.X + 1, element.Y}] {
			sides++
		}

		if cells[D12Point{element.X, element.Y + 1}] && cells[D12Point{element.X + 1, element.Y}] && !cells[D12Point{element.X + 1, element.Y + 1}] {
			sides++
		}
		if cells[D12Point{element.X, element.Y + 1}] && cells[D12Point{element.X - 1, element.Y}] && !cells[D12Point{element.X - 1, element.Y + 1}] {
			sides++
		}

		if cells[D12Point{element.X, element.Y - 1}] && cells[D12Point{element.X + 1, element.Y}] && !cells[D12Point{element.X + 1, element.Y - 1}] {
			sides++
		}

		if cells[D12Point{element.X, element.Y - 1}] && cells[D12Point{element.X - 1, element.Y}] && !cells[D12Point{element.X - 1, element.Y - 1}] {
			sides++
		}

	}

	return sides
}

type D12Cell struct {
	Visited bool
	Char    rune
}

type D12Map struct {
	Map map[D12Point]*D12Cell
}

func (d *D12Map) GetFirstNonVisited() *D12Point {
	for point, cell := range d.Map {
		if !cell.Visited {
			return &point
		}
	}

	return nil
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

	return points
}

func (d *D12Map) SetVisited(point D12Point) {
	d.Map[point].Visited = true
}

func NewD12Map() D12Map {
	return D12Map{Map: make(map[D12Point]*D12Cell, 0)}
}
