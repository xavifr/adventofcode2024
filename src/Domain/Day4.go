package Domain

type Day4Soup struct {
	Chars [][]rune
}

type Day4Position struct {
	X int
	Y int
}

type Day4Direction int

const (
	D4_N Day4Direction = iota
	D4_S
	D4_E
	D4_W
	D4_SE
	D4_SW
	D4_NE
	D4_NW
)

func NewDay4Soup() Day4Soup {
	return Day4Soup{
		Chars: make([][]rune, 0),
	}
}

func (d *Day4Soup) Append(line string) {
	d.Chars = append(d.Chars, []rune(line))
}

func (d *Day4Soup) FindChar(char rune) []Day4Position {
	positions := make([]Day4Position, 0)

	for i := 0; i < len(d.Chars); i++ {
		for j := 0; j < len(d.Chars[i]); j++ {
			if d.Chars[i][j] == char {
				positions = append(positions, Day4Position{X: i, Y: j})
			}
		}
	}

	return positions
}

func (d *Day4Soup) GetChar(x, y int) rune {
	if x < 0 || x >= len(d.Chars) {
		return ' '
	}
	if y < 0 || y >= len(d.Chars[x]) {
		return ' '
	}

	return d.Chars[x][y]
}

func (d *Day4Soup) GetWord(x, y, length int, direction Day4Direction) string {
	word := string(d.GetChar(x, y))

	for i := 1; i < length; i++ {
		switch direction {
		case D4_N:
			x--
		case D4_S:
			x++
		case D4_E:
			y++
		case D4_W:
			y--
		case D4_SE:
			x++
			y++
		case D4_SW:
			x++
			y--
		case D4_NE:
			x--
			y++
		case D4_NW:
			x--
			y--
		}
		word += string(d.GetChar(x, y))

	}

	return word
}
