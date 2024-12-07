package Domain

import (
	"errors"
	"fmt"
)

type D6Object rune

const (
	D6_EMPTY            D6Object = '.'
	D6_TRASH            D6Object = '#'
	D6_OBSTACLE         D6Object = 'O'
	D6_VISITED_NORTH    D6Object = '^'
	D6_VISITED_SOUTH    D6Object = 'v'
	D6_VISITED_EAST     D6Object = '>'
	D6_VISITED_WEST     D6Object = '<'
	D6_VISITED_MULTIPLE D6Object = '+'
)

type D6Direction rune

const (
	D6_NORTH = '^'
	D6_SOUTH = 'v'
	D6_EAST  = '>'
	D6_WEST  = '<'
)

type D6Guard struct {
	X, Y      int
	Direction D6Direction
}

type D6Map struct {
	Points       [][]D6Object
	Guard        D6Guard
	VisitedPonts int
}

func (d *D6Map) PositionInMap(X, Y int) bool {
	return X >= 0 && Y >= 0 && Y < len(d.Points) && X < len(d.Points[Y])
}

func (d *D6Map) PositionEmpty(X, Y int) bool {
	return d.Points[Y][X] != D6_TRASH && d.Points[Y][X] != D6_OBSTACLE
}

func (d *D6Map) GetPosition(X, Y int) D6Object {
	if !d.PositionInMap(X, Y) {
		return D6_EMPTY
	}
	return d.Points[Y][X]
}

func (d *D6Guard) GuardNextPosition() (int, int) {
	X, Y := d.X, d.Y

	if d.Direction == D6_NORTH {
		Y--
	} else if d.Direction == D6_EAST {
		X++
	} else if d.Direction == D6_SOUTH {
		Y++
	} else if d.Direction == D6_WEST {
		X--
	}

	return X, Y
}

func (d *D6Guard) Move() {
	d.X, d.Y = d.GuardNextPosition()
}

func (d *D6Guard) Rotate() {
	if d.Direction == D6_NORTH {
		d.Direction = D6_EAST
	} else if d.Direction == D6_EAST {
		d.Direction = D6_SOUTH
	} else if d.Direction == D6_SOUTH {
		d.Direction = D6_WEST
	} else if d.Direction == D6_WEST {
		d.Direction = D6_NORTH
	}
}

func (d *D6Map) GuardCanMove() bool {
	if !d.PositionInMap(d.Guard.X, d.Guard.Y) {
		return false
	}

	if !d.PositionInMap(d.Guard.GuardNextPosition()) {
		return false
	}

	return d.PositionEmpty(d.Guard.GuardNextPosition())
}

func (d *D6Map) Print() {
	startX := d.Guard.X - 6
	if startX < 0 {
		startX = 0
	}

	startY := d.Guard.Y - 6
	if startY < 0 {
		startY = 0
	}

	endX := startX + 12
	if endX >= len(d.Points[0]) {
		endX = len(d.Points[0]) - 1
	}
	endY := startY + 12
	if endY >= len(d.Points) {
		endY = len(d.Points) - 1
	}

	fmt.Printf("%c[1J%c[1;1H", rune(033), rune(033))

	for i := startY; i <= endY; i++ {
		for j := startX; j <= endX; j++ {
			if j == d.Guard.X && i == d.Guard.Y {
				fmt.Printf("%c", d.Guard.Direction)
			} else {
				fmt.Printf("%c", d.Points[i][j])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Visited points are :%d\n", d.VisitedPonts)

	fmt.Printf("")
}

func (d *D6Map) MoveGuard() error {
	d.VisitedPonts = 1
	d.Points[d.Guard.Y][d.Guard.X] = D6Object(d.Guard.Direction)

	for d.PositionInMap(d.Guard.X, d.Guard.Y) {
		if !d.PositionInMap(d.Guard.GuardNextPosition()) {
			break
		} else if d.GuardCanMove() {
			d.Guard.Move()
			if d.Points[d.Guard.Y][d.Guard.X] == D6_EMPTY {

				d.Points[d.Guard.Y][d.Guard.X] = D6Object(d.Guard.Direction)
				d.VisitedPonts++
			}

		} else {
			d.Guard.Rotate()
		}

		nextPoint := d.GetPosition(d.Guard.GuardNextPosition())
		if nextPoint == D6Object(d.Guard.Direction) {
			return errors.New("loop detected")
		}

		//d.Print()
	}

	return nil
}
