package Domain

type D18Point struct {
	X, Y int
}

func (d *D18Point) GetNeighbors() []D18Point {
	return []D18Point{
		D18Point{X: d.X + 1, Y: d.Y},
		D18Point{X: d.X - 1, Y: d.Y},
		D18Point{X: d.X, Y: d.Y + 1},
		D18Point{X: d.X, Y: d.Y - 1},
	}
}

type D18Ram struct {
	Size            D18Point
	MemoryCorrupted map[D18Point]bool
	CorruptionOrder []D18Point
}

func (d *D18Ram) CorruptBytes(bytes int) bool {
	d.MemoryCorrupted = make(map[D18Point]bool, 0)

	for i := 0; i < bytes; i++ {
		if i >= len(d.CorruptionOrder) {
			return false
		}
		d.MemoryCorrupted[d.CorruptionOrder[i]] = true
	}

	return true
}

func (d *D18Ram) BestPath(start, end D18Point) int {
	visited := map[D18Point]int{start: 0}
	nextPoints := []D18Point{start}

	for {
		if len(nextPoints) == 0 {
			break
		}

		point := nextPoints[0]
		nextPoints = nextPoints[1:]
		for _, neigh := range point.GetNeighbors() {
			if neigh.X < 0 || neigh.Y < 0 || neigh.X >= d.Size.X || neigh.Y >= d.Size.Y || d.MemoryCorrupted[neigh] {
				// out of scope or corrupted (WALL)
				continue
			}

			if visited[neigh] == 0 {
				visited[neigh] = visited[point] + 1
				nextPoints = append(nextPoints, neigh)
			}

			if neigh == end {
				return visited[end]
			}
		}
	}

	return -1
}

func NewD18Ram() D18Ram {
	return D18Ram{Size: D18Point{}, MemoryCorrupted: make(map[D18Point]bool, 0), CorruptionOrder: make([]D18Point, 0)}
}
