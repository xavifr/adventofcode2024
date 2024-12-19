package main

import (
	"adventofcode2024/Application"
	"adventofcode2024/Repository"
	"fmt"
	"os"
	"strconv"
)

func main() {
	daysRepo := Repository.NewDaysRepository("./input/")
	_ = daysRepo.Add(1, &Application.Day1{})
	_ = daysRepo.Add(2, &Application.Day2{})
	_ = daysRepo.Add(3, &Application.Day3{})
	_ = daysRepo.Add(4, &Application.Day4{})
	_ = daysRepo.Add(5, &Application.Day5{})
	_ = daysRepo.Add(6, &Application.Day6{})
	_ = daysRepo.Add(7, &Application.Day7{})
	_ = daysRepo.Add(8, &Application.Day8{})
	_ = daysRepo.Add(9, &Application.Day9{})
	_ = daysRepo.Add(10, &Application.Day10{})
	_ = daysRepo.Add(11, &Application.Day11{})
	_ = daysRepo.Add(12, &Application.Day12{})
	_ = daysRepo.Add(13, &Application.Day13{})
	_ = daysRepo.Add(14, &Application.Day14{})
	_ = daysRepo.Add(15, &Application.Day15{})
	_ = daysRepo.Add(16, &Application.Day16{})
	_ = daysRepo.Add(17, &Application.Day17{})
	_ = daysRepo.Add(18, &Application.Day18{})
	_ = daysRepo.Add(19, &Application.Day19{})
	/*_ = daysRepo.Add(2, Application.NewDay2(12, 13, 14))*/

	if len(os.Args) > 1 {
		dayString := os.Args[1]
		day, errConv := strconv.Atoi(dayString)
		if errConv != nil || day <= 0 {
			fmt.Println("Day is not a valid number")
			os.Exit(1)
		}

		errExecute := daysRepo.Execute(day)
		if errExecute != nil {
			fmt.Printf("Error executing day %d: %s", day, errExecute)
			os.Exit(1)
		}
		os.Exit(0)
	}

	var errExecute error
	actDay := 0
	for errExecute == nil {
		actDay++
		fmt.Printf("Executing day %d\n", actDay)
		errExecute = daysRepo.Execute(actDay)
	}

	if errExecute != nil {
		fmt.Printf("Error executing day %d: %s", actDay, errExecute)
		os.Exit(1)
	}

	os.Exit(0)
}
