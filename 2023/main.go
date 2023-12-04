package main

import (
	"AdventOfCode/2023/Day1"
	"AdventOfCode/2023/Day2"
	"fmt"
)

func main() {
	ExecuteDay1()
	ExecuteDay2()
}

func ExecuteDay1() {
	calibrateValue, err := Day1.FindCalibrateValue()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("First day result: %d\n", calibrateValue)
}

func ExecuteDay2() {
	calibrateValue, err := Day2.FindCalibrateValue()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Second day result: %d\n", calibrateValue)
}
