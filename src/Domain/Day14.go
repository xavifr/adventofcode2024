package Domain

type D14Point struct {
	X, Y int
}

type D14Robot struct {
	Vector D14Point
}

func (d *D14Robot) Move(position D14Point) D14Point {
	return D14Point{
		X: position.X + d.Vector.X,
		Y: position.Y + d.Vector.Y,
	}
}

type D14Bathroom struct {
	Robots map[D14Point][]D14Robot
	Size   D14Point
}

func (d *D14Bathroom) Calc() int {
	var q1, q2, q3, q4 int

	for point, robots := range d.Robots {
		if point.X >= 0 && point.X < d.Size.X/2 {
			if point.Y >= 0 && point.Y < d.Size.Y/2 {
				//fmt.Printf("append point %d:%d to Q1\n", point.X, point.Y)
				q1 += len(robots)
			} else if point.Y > d.Size.Y/2 {
				//fmt.Printf("append point %d:%d to Q2\n", point.X, point.Y)
				q2 += len(robots)
			}
		} else if point.X > d.Size.X/2 {
			if point.Y >= 0 && point.Y < d.Size.Y/2 {
				//fmt.Printf("append point %d:%d to Q3\n", point.X, point.Y)
				q3 += len(robots)
			} else if point.Y > d.Size.Y/2 {
				//fmt.Printf("append point %d:%d to Q4\n", point.X, point.Y)
				q4 += len(robots)
			}
		}

	}
	return q1 * q2 * q3 * q4
}

func (d *D14Bathroom) Print() {
	for y := 0; y < d.Size.Y; y++ {
		for x := 0; x < d.Size.X; x++ {
			if _, ok := d.Robots[D14Point{X: x, Y: y}]; ok {
				print("#")
			} else {
				print(".")
			}
		}
		print("\n")
	}
}

func (d *D14Bathroom) Move() {
	newRobots := map[D14Point][]D14Robot{}

	for point, robots := range d.Robots {
		for _, robot := range robots {
			newPosition := robot.Move(point)
			if newPosition.X < 0 {
				newPosition.X = d.Size.X + newPosition.X
			}

			for newPosition.X >= d.Size.X {
				newPosition.X = newPosition.X - d.Size.X
			}

			if newPosition.Y < 0 {
				newPosition.Y = d.Size.Y + newPosition.Y
			}

			for newPosition.Y >= d.Size.Y {
				newPosition.Y = newPosition.Y - d.Size.Y
			}

			newRobots[newPosition] = append(newRobots[newPosition], robot)
		}
	}

	d.Robots = newRobots
}

func NewD14Bathroom() D14Bathroom {
	return D14Bathroom{
		Size:   D14Point{},
		Robots: make(map[D14Point][]D14Robot, 0),
	}
}
