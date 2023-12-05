package main

import (
	"AdventOfCode/2023/Day1"
	"fmt"
)

func main() {
	ExecuteDay1Problem1()
	ExecuteDay1Problem2()
}

func ExecuteDay1Problem1() {
	calibrateValue, err := Day1.FindCalibrateValue(false)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("First day result: %d\n", calibrateValue)
}

func ExecuteDay1Problem2() {
	calibrateValue, err := Day1.FindCalibrateValue(true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Second day result: %d\n", calibrateValue)
}
