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
