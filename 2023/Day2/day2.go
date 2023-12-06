package Day2

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colors = map[string]int64{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func FirstProblem() (int64, error) {
	gameList, _ := loadInput()

	var possiblesGames int64 = 0
	for i := 0; i < len(gameList); i++ {
		possiblesGames += checkGame(gameList[i])
	}

	return possiblesGames, nil
}

func loadInput() ([]string, error) {
	content, err := os.ReadFile("./2023/Day2/input.txt")
	if err != nil {
		return nil, errors.New("file not found")
	}

	return strings.Split(string(content), "\n"), nil
}

func checkGame(game string) int64 {
	gameLine := strings.Split(game, ":")
	gameID, _ := strconv.Atoi(strings.Split(gameLine[0], " ")[1])
	gameSets := strings.Split(gameLine[1], ";")

	for _, gameSet := range gameSets {
		if !checkIsPossibleSet(gameSet) {
			return 0
		}
	}

	return int64(gameID)
}

func checkIsPossibleSet(gameSet string) bool {
	balls := strings.Split(gameSet, ",")
	for _, ball := range balls {
		ballSplit := strings.Split(strings.TrimSpace(ball), " ")
		count, err := strconv.Atoi(ballSplit[0])
		color := ballSplit[1]

		if err != nil {
			fmt.Printf("error to get count of %s color", color)
			return false
		}

		if int64(count) > colors[color] {
			return false
		}
	}

	return true
}

func SecondProblem() (int64, error) {
	gameList, _ := loadInput()

	var sumPower int64 = 0
	for i := 0; i < len(gameList); i++ {
		sumPower += CalculatePower(gameList[i])
	}

	return sumPower, nil
}

func CalculatePower(game string) int64 {
	maxColors := make(map[string]int64)
	for key := range colors {
		maxColors[key] = 0
	}

	gameSplit := strings.Split(game, ":")

	gameSets := strings.Split(gameSplit[1], ";")
	for _, gameSet := range gameSets {
		balls := strings.Split(gameSet, ",")
		for _, ball := range balls {
			ballSplit := strings.Split(strings.TrimSpace(ball), " ")
			count, err := strconv.Atoi(ballSplit[0])
			color := ballSplit[1]

			if err != nil {
				fmt.Printf("error to get count of %s color", color)
				return 0
			}

			if int64(count) > maxColors[color] {
				maxColors[color] = int64(count)
			}
		}
	}

	var totalPower int64 = 1
	for key := range colors {
		totalPower *= maxColors[key]
	}

	return totalPower
}
