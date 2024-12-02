package Repository

import (
	"adventofcode2024/Domain"
	"bufio"
	"errors"
	"fmt"
	"os"
)

type InputType string
type ProblemPart string

const (
	DemoInput  InputType = "demo"
	FinalInput           = "input"
)

const (
	Part1 ProblemPart = "part1"
	Part2             = "part2"
)

type DaysRepository struct {
	days      map[int]Domain.Day
	inputPath string
}

func (dr *DaysRepository) Execute(day int) error {
	handler, found := dr.days[day]
	if !found {
		return fmt.Errorf("no handler found for day %d\n", day)
	}

	var errExecute error

	demoFileP1, errDemo1 := dr.getInputPart(day, DemoInput, Part1)
	if errDemo1 == nil {
		defer demoFileP1.Close()

		fmt.Println("----------------------------------------")
		fmt.Printf("Executing DEMO day%d:%s\n", day, Part1)
		fmt.Println("----------------------------------------")
		errExecute = handler.Part1(bufio.NewScanner(demoFileP1))
		if errExecute != nil {
			return fmt.Errorf("error executing day %d:%s:%s => %s", day, DemoInput, Part1, errExecute)
		}
		fmt.Println()
	}

	finalFile, errFinal := dr.getInput(day, FinalInput)
	if errFinal == nil {
		defer finalFile.Close()

		fmt.Println("----------------------------------------")
		fmt.Printf("Executing FINAL day%d:%s\n", day, Part1)
		fmt.Println("----------------------------------------")
		errExecute = handler.Part1(bufio.NewScanner(finalFile))
		if errExecute != nil {
			return fmt.Errorf("error executing day %d:%s:%s => %s", day, FinalInput, Part1, errExecute)
		}
		fmt.Println()
	}

	demoFileP2, errDemo2 := dr.getInputPart(day, DemoInput, Part2)
	if errDemo2 == nil {
		defer demoFileP2.Close()

		fmt.Println("----------------------------------------")
		fmt.Printf("Executing DEMO day%d:%s\n", day, Part2)
		fmt.Println("----------------------------------------")
		errExecute = handler.Part2(bufio.NewScanner(demoFileP2))
		if errExecute != nil {
			return fmt.Errorf("error executing day %d:%s:%s => %s", day, DemoInput, Part2, errExecute)
		}
		fmt.Println()
	}

	if errFinal == nil {
		finalFile.Seek(0, 0)

		fmt.Println("----------------------------------------")
		fmt.Printf("Executing FINAL day%d:%s\n", day, Part2)
		fmt.Println("----------------------------------------")
		errExecute = handler.Part2(bufio.NewScanner(finalFile))
		if errExecute != nil {
			return fmt.Errorf("error executing day %d:%s:%s => %s", day, FinalInput, Part2, errExecute)
		}
		fmt.Println()
	}

	return nil
}

func (dr *DaysRepository) Add(day int, handler Domain.Day) error {
	_, found := dr.days[day]
	if found {
		return errors.New("day already found")
	}

	dr.days[day] = handler
	return nil
}

func (dr *DaysRepository) getInput(day int, inputType InputType) (*os.File, error) {
	fileName := fmt.Sprintf("%s/day%d.%s", dr.inputPath, day, inputType)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (dr *DaysRepository) getInputPart(day int, inputType InputType, part ProblemPart) (*os.File, error) {
	fileName := fmt.Sprintf("%s/day%d.%s_%s", dr.inputPath, day, inputType, part)
	file, err := os.Open(fileName)
	if err != nil {
		return dr.getInput(day, inputType)
	}

	return file, nil
}

func NewDaysRepository(inputPath string) *DaysRepository {
	daysRepository := &DaysRepository{days: map[int]Domain.Day{}, inputPath: inputPath}

	return daysRepository
}
