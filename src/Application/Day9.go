package Application

import (
	"adventofcode2024/Domain"
	"bufio"
	"fmt"
	"strconv"
)

type Day9 struct {
	Disk Domain.D9Disk
}

func (d *Day9) Part1(input *bufio.Scanner) error {
	d.parseInputV1(input)

	// TODO: Commented because it's too slow
	//d.Disk.RearrangeSectorsNoFragment()

	fmt.Printf("PART1: Checksum is %d\n", d.Disk.ChecksumV2())

	return nil
}

func (d *Day9) Part2(input *bufio.Scanner) error {
	d.parseInputV2(input)
	d.Disk.RearrangeSectorsNoFragment()

	fmt.Printf("PART2: Checksum is %d\n", d.Disk.ChecksumV2())

	return nil
}

func (d *Day9) parseInputV2(input *bufio.Scanner) {
	d.Disk.Files = make([]*Domain.D9File, 0)
	d.Disk.Sectors = make(map[int]*Domain.D9File)

	fileId := 0
	position := 0
	for input.Scan() {
		line := input.Text()

		for i, c := range line {
			if i%2 == 0 {
				num, _ := strconv.Atoi(string(c))
				file := Domain.D9File{Id: fileId, Size: num, StartSector: position}
				d.Disk.Files = append(d.Disk.Files, &file)
				position += num
				fileId++
			} else {
				num, _ := strconv.Atoi(string(c))
				position += num
			}
		}
	}
}

func (d *Day9) parseInputV1(input *bufio.Scanner) {
	d.Disk.Files = make([]*Domain.D9File, 0)
	d.Disk.Sectors = make(map[int]*Domain.D9File)

	fileId := 0
	position := 0
	for input.Scan() {
		line := input.Text()

		for i, c := range line {
			if i%2 == 0 {
				num, _ := strconv.Atoi(string(c))
				for num > 0 {
					file := Domain.D9File{Id: fileId, Size: 1, StartSector: position}
					d.Disk.Files = append(d.Disk.Files, &file)
					position++
					num--
				}
				fileId++
			} else {
				num, _ := strconv.Atoi(string(c))
				position += num
			}
		}
	}
}
