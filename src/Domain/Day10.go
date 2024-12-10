package Domain

type D10Point struct {
	X, Y int
}
type D10TrailMap struct {
	Trailheads []D10Point
	Map        [][]int
}

func (d *D10TrailMap) GetScore(point D10Point, withEnds bool) int {
	trailends := make(map[D10Point]bool, 0)

	return d.recurse(point, 0, withEnds, trailends)
}

func (d *D10TrailMap) getPoint(point D10Point) int {
	return d.Map[point.Y][point.X]
}

func (d *D10TrailMap) recurse(point D10Point, carry int, withEnds bool, trailends map[D10Point]bool) int {
	if d.getPoint(point) == 9 {
		if _, ok := trailends[point]; !withEnds || !ok {
			trailends[point] = true
			carry += 1
		}
		return carry
	}

	for _, neighbour := range d.getNeighbours(point) {
		if d.getPoint(neighbour) == d.getPoint(point)+1 {
			carry = d.recurse(neighbour, carry, withEnds, trailends)
		}
	}

	return carry
}

func (d *D10TrailMap) getNeighbours(point D10Point) []D10Point {
	points := make([]D10Point, 0)

	if point.X-1 >= 0 {
		points = append(points, D10Point{X: point.X - 1, Y: point.Y})
	}

	if point.X+1 < len(d.Map[point.Y]) {
		points = append(points, D10Point{X: point.X + 1, Y: point.Y})
	}

	if point.Y-1 >= 0 {
		points = append(points, D10Point{X: point.X, Y: point.Y - 1})
	}

	if point.Y+1 < len(d.Map) {
		points = append(points, D10Point{X: point.X, Y: point.Y + 1})
	}

	return points
}
