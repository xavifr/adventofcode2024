package Domain

import (
	"fmt"
	"strings"
	"time"
)

type D15Point struct {
	X int
	Y int
}

func (p D15Point) Move(movement D15Movement) D15Point {
	if movement == UP {
		return D15Point{X: p.X, Y: p.Y - 1}
	} else if movement == DOWN {
		return D15Point{X: p.X, Y: p.Y + 1}
	} else if movement == LEFT {
		return D15Point{X: p.X - 1, Y: p.Y}
	} else if movement == RIGHT {
		return D15Point{X: p.X + 1, Y: p.Y}
	}

	return D15Point{}

}

type D15Object rune

const (
	ROBOT D15Object = '@'
	WALL  D15Object = '#'
	BOX   D15Object = 'O'
	SPACE D15Object = '.'
)

type D15Movement rune

const (
	UP    D15Movement = '^'
	DOWN  D15Movement = 'v'
	LEFT  D15Movement = '<'
	RIGHT D15Movement = '>'
)

type D15Warehouse struct {
	Size      D15Point
	Map       map[D15Point]D15Object
	Movements []D15Movement
	Robot     D15Point
}

func (d *D15Warehouse) Value() int {
	sum := 0
	for p, v := range d.Map {
		if v == BOX {
			sum += p.X + p.Y*100
		}
	}
	return sum
}

func (d *D15Warehouse) Move() bool {
	nextMove := d.Movements[0]
	d.Movements = d.Movements[1:]

	prediction, possible := d.GetPrediction(nextMove)

	if !possible {
		return false
	}

	for i := len(prediction) - 1; i >= 0; i-- {
		d.Map[prediction[i].Move(nextMove)] = d.Map[prediction[i]]
		delete(d.Map, prediction[i])
	}

	d.Robot = d.Robot.Move(nextMove)
	return true
}

func (d *D15Warehouse) GetPrediction(movement D15Movement) ([]D15Point, bool) {
	robot := d.Robot
	output := []D15Point{}

	//foundBox := false
	for {
		robot = robot.Move(movement)

		switch d.Map[robot] {
		case WALL:
			return []D15Point{}, false
		case BOX:
			output = append(output, robot)
			//foundBox = true
		default: // SPACE
			return output, true
		}
	}
}

func (d *D15Warehouse) Move2() bool {
	nextMove := d.Movements[0]
	d.Movements = d.Movements[1:]

	prediction, possible := d.GetPrediction2(nextMove)

	if !possible {
		return false
	}

	newPoints := map[D15Point]D15Object{}
	for i := len(prediction) - 1; i >= 0; i-- {
		newPoints[prediction[i].Move(nextMove)] = d.Map[prediction[i]]
		delete(d.Map, prediction[i])
	}

	for point, object := range newPoints {
		d.Map[point] = object
	}

	d.Robot = d.Robot.Move(nextMove)
	return true
}

func (d *D15Warehouse) GetPrediction2(movement D15Movement) ([]D15Point, bool) {
	boxesMoved := map[D15Point]bool{}
	pointsBeingPushed := map[D15Point]bool{d.Robot: true}

	//foundBox := false
	for len(pointsBeingPushed) > 0 {
		for point := range pointsBeingPushed {
			delete(pointsBeingPushed, point)
			point = point.Move(movement)

			if d.Map[point] == WALL {
				return []D15Point{}, false
			} else if d.Map[point] == BOX {
				boxesMoved[point] = true
				if movement == UP || movement == DOWN {
					pointsBeingPushed[point] = true
					pointsBeingPushed[D15Point{X: point.X + 1, Y: point.Y}] = true
				} else if movement == RIGHT {
					pointsBeingPushed[D15Point{X: point.X + 1, Y: point.Y}] = true
				} else if movement == LEFT {
					pointsBeingPushed[point] = true
				}
			} else if (movement == LEFT || movement == UP || movement == DOWN) && d.Map[D15Point{X: point.X - 1, Y: point.Y}] == BOX {
				boxesMoved[D15Point{X: point.X - 1, Y: point.Y}] = true
				if movement == UP || movement == DOWN {
					pointsBeingPushed[D15Point{X: point.X - 1, Y: point.Y}] = true
					pointsBeingPushed[point] = true
				} else if movement == LEFT {
					pointsBeingPushed[D15Point{X: point.X - 1, Y: point.Y}] = true
				}
			}
		}
	}

	arrayOfBoxes := []D15Point{}
	for point := range boxesMoved {
		arrayOfBoxes = append(arrayOfBoxes, point)
	}

	return arrayOfBoxes, true
}

func (d *D15Warehouse) Print() {
	printMap := make([]string, d.Size.Y)
	for i := range printMap {
		printMap[i] = strings.Repeat(".", d.Size.X)
	}

	for point, object := range d.Map {
		line := printMap[point.Y]
		if object == BOX {
			line = line[:point.X] + "[]" + line[point.X+2:]
		} else {
			line = line[:point.X] + string(object) + line[point.X+1:]

		}

		printMap[point.Y] = line
	}

	robotChar := "@"
	if len(d.Movements) > 0 {
		robotChar = string(d.Movements[0])
	}
	lineRobot := printMap[d.Robot.Y]
	lineRobot = lineRobot[:d.Robot.X] + robotChar + lineRobot[d.Robot.X+1:]

	printMap[d.Robot.Y] = lineRobot

	fmt.Printf("%s\n", strings.Join(printMap, "\n"))
	if len(d.Movements) > 0 {
		fmt.Printf("Next move is %c\n", d.Movements[0])
	}

	time.Sleep(1 * time.Millisecond)
}

func NewD15Warehouse() D15Warehouse {
	return D15Warehouse{
		Map:       make(map[D15Point]D15Object),
		Movements: make([]D15Movement, 0),
	}
}
