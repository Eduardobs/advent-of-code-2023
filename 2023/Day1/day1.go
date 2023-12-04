package Day1

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func FindCalibrateValue() (int, error) {
	content, err := os.ReadFile("./2023/Day1/input.txt")
	if err != nil {
		return -1, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")

	total := 0
	for i := 0; i < len(lines); i++ {
		total += getNumbers(lines[i])
	}

	return total, nil
}

func getNumbers(line string) int {
	firstNumber := -1
	lastNumber := -1

	for _, char := range line {

		if number, err := strconv.Atoi(string(char)); err == nil {
			if firstNumber == -1 {
				firstNumber = number
			}
			lastNumber = number
		}
	}

	resultNumber := firstNumber*10 + lastNumber

	return resultNumber
}
