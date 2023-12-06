package Day4

import (
	"errors"
	"os"
	"slices"
	"strings"
)

func FirstProblem() (int64, error) {
	cardsLine, _ := loadInput()

	total := 0

	for _, line := range cardsLine {
		parts := strings.Split(line, ":")
		parts = strings.Split(parts[1], "|")

		winning := strings.Fields(parts[0])
		ours := strings.Fields(parts[1])

		var haveWinning int
		for _, n := range ours {
			if slices.Contains(winning, n) {
				haveWinning++
			}
		}

		if haveWinning >= 1 {
			total += 1 << (haveWinning - 1)
		}
	}

	return int64(total), nil
}

func SecondProblem() (int64, error) {
	total := 0

	cardsLine, _ := loadInput()

	cards := make([]string, 0)
	for _, card := range cardsLine {
		cards = append(cards, card)
	}

	countsOfCards := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		countsOfCards[i] = 1
	}

	for i, card := range cards {
		parts := strings.Split(card, ":")
		parts = strings.Split(parts[1], "|")

		winning := strings.Fields(parts[0])
		ours := strings.Fields(parts[1])

		var haveWinning int
		for _, n := range ours {
			if slices.Contains(winning, n) {
				haveWinning++
			}
		}

		for j := 1; j <= haveWinning; j++ {
			countsOfCards[i+j] += countsOfCards[i]
		}

		total += countsOfCards[i]
	}

	return int64(total), nil
}

func loadInput() ([]string, error) {
	content, err := os.ReadFile("./2023/Day4/input.txt")
	if err != nil {
		return nil, errors.New("file not found")
	}

	return strings.Split(string(content), "\n"), nil
}
