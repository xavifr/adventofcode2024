package Domain

type D8Point struct {
	X, Y int
}

func (d *D8Point) GetAntinode(second D8Point) []D8Point {
	vectorX, vectorY := second.X-d.X, second.Y-d.Y

	return []D8Point{
		{X: d.X - vectorX, Y: d.Y - vectorY},
		{X: second.X + vectorX, Y: second.Y + vectorY},
	}
}
func (d *D8Point) GetVector(second D8Point) D8Point {
	return D8Point{
		X: second.X - d.X,
		Y: second.Y - d.Y,
	}
}

type D8Map struct {
	Antinodes      [][]bool
	AntinodesCount int
	Frequencies    map[rune][]D8Point
}

func (d *D8Map) AddAntinode(antinode D8Point) bool {
	if antinode.X < 0 || antinode.Y < 0 || antinode.Y >= len(d.Antinodes) || antinode.X >= len(d.Antinodes[antinode.Y]) {
		return false
	}

	if !d.Antinodes[antinode.Y][antinode.X] {
		d.Antinodes[antinode.Y][antinode.X] = true
		d.AntinodesCount++
		return true
	}

	return true
}

func (d *D8Map) AddAntinodeVector(center, vector D8Point) bool {
	posX, posY := center.X, center.Y

	for posX >= 0 && posY >= 0 && posY < len(d.Antinodes) && posX < len(d.Antinodes[posY]) {
		if !d.Antinodes[posY][posX] {
			d.Antinodes[posY][posX] = true
			d.AntinodesCount++
		}

		posX -= vector.X
		posY -= vector.Y
	}

	posX, posY = center.X, center.Y
	for posX >= 0 && posY >= 0 && posY < len(d.Antinodes) && posX < len(d.Antinodes[posY]) {
		if !d.Antinodes[posY][posX] {
			d.Antinodes[posY][posX] = true
			d.AntinodesCount++
		}

		posX += vector.X
		posY += vector.Y
	}

	return true
}
