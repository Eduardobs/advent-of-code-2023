package main

import (
	"AdventOfCode/2023/Day1"
	"AdventOfCode/2023/Day2"
	"AdventOfCode/2023/Day3"
	"fmt"
)

func main() {
	Execute(Day1.FindCalibratedValueWithNumberString, "Day 1, problem 1 result")
	Execute(Day1.FindCalibratedValueWithoutNumberString, "Day 1, problem 2 result")
	Execute(Day2.CheckPossibleGames, "Day 2, problem 1 result")
	Execute(Day2.SumPowerSets, "Day 2, problem 2 result")
	Execute(Day3.GetSumEngine, "Day 3, problem 1 result")
	Execute(Day3.GetSumEngineRatio, "Day 3, problem 2 result")
}

func Execute(funcExecute func() (int64, error), successMsg string) {
	result, err := funcExecute()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s: %d\n", successMsg, result)
}
