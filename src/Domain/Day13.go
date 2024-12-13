package Domain

import "math"

type D13Vector struct {
	X, Y int
}

type D13Prize struct {
	X, Y    int
	ButtonA D13Vector
	ButtonB D13Vector
}


func (d D13Prize) Resolve(withExtraCost bool) (int, int, bool) {
	if withExtraCost {
		d.Y += 10000000000000
		d.X += 10000000000000
	}

	minB := ((float64(d.Y) * float64(d.ButtonA.X)) - (float64(d.ButtonA.Y) * float64(d.X))) / ((-1 * float64(d.ButtonA.Y) * float64(d.ButtonB.X)) + (float64(d.ButtonA.X) * float64(d.ButtonB.Y)))

	if minB != math.Floor(minB) {
		return 0, 0, false
	}

	minA := (float64(d.X) - float64(d.ButtonB.X)*minB) / float64(d.ButtonA.X)

	if minA != math.Floor(minA) {
		return 0, 0, false
	}

	return int(minA), int(minB), true
}

func NewD13Prize(x, y int, buttonAX, buttonAY, buttonBX, buttonBY int) D13Prize {
	return D13Prize{
		X: x,
		Y: y,
		ButtonA: D13Vector{
			X: buttonAX,
			Y: buttonAY,
		},
		ButtonB: D13Vector{
			X: buttonBX,
			Y: buttonBY,
		},
	}
}
