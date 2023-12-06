package main

import (
	"AdventOfCode/2023/Day1"
	"AdventOfCode/2023/Day2"
	"AdventOfCode/2023/Day3"
	"AdventOfCode/2023/Day4"
	"fmt"
)

func main() {
	Execute(Day1.FirstProblem, "Day 1, problem 1 result")
	Execute(Day1.SecondProblem, "Day 1, problem 2 result")
	Execute(Day2.FirstProblem, "Day 2, problem 1 result")
	Execute(Day2.SecondProblem, "Day 2, problem 2 result")
	Execute(Day3.FirstProblem, "Day 3, problem 1 result")
	Execute(Day3.SecondProblem, "Day 3, problem 2 result")
	Execute(Day4.FirstProblem, "Day 4, problem 1 result")
	Execute(Day4.SecondProblem, "Day 4, problem 2 result")
}

func Execute(funcExecute func() (int64, error), successMsg string) {
	result, err := funcExecute()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s: %d\n", successMsg, result)
}
