package Day1

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func FindCalibrateValue(checkNumberString bool) (int, error) {
	content, err := os.ReadFile("./2023/Day1/input.txt")
	if err != nil {
		return -1, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")

	total := 0
	for i := 0; i < len(lines); i++ {
		total += getNumbers(lines[i], checkNumberString)
	}

	return total, nil
}

func getNumbers(line string, checkNumberString bool) int {
	firstNumber := -1
	lastNumber := -1

	for index, char := range line {
		if checkNumberString {
			if number, err := checkStringNumber(line[index:]); err == nil {
				if firstNumber == -1 {
					firstNumber = number
				}
				lastNumber = number

				continue
			}
		}

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

func checkStringNumber(line string) (int, error) {
	numbersMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for key := range numbersMap {
		if strings.HasPrefix(line, key) {
			return numbersMap[key], nil
		}
	}
	return 0, errors.New("number written in full not found")
}
