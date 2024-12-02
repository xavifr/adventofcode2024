package Domain

import "bufio"

type Day interface {
	Part1(*bufio.Scanner) error
	Part2(*bufio.Scanner) error
}
