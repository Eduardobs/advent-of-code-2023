package main

import (
	"AdventOfCode/2023/Day1"
	"AdventOfCode/2023/Day2"
	"fmt"
)

func main() {
	ExecuteDay1Problem1()
	ExecuteDay1Problem2()
	ExecuteDay2Problem1()
	ExecuteDay2Problem2()
}

func ExecuteDay1Problem1() {
	calibrateValue, err := Day1.FindCalibrateValue(false)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("First day, problem 1 result: %d\n", calibrateValue)
}

func ExecuteDay1Problem2() {
	calibrateValue, err := Day1.FindCalibrateValue(true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("First day, problem 2 result: %d\n", calibrateValue)
}

func ExecuteDay2Problem1() {
	possibleGames, err := Day2.CheckPossibleGames()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Second day, problem 1 result: %d\n", possibleGames)
}

func ExecuteDay2Problem2() {
	possibleGames, err := Day2.SumPowerSets()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Second day, problem 1 result: %d\n", possibleGames)
}
