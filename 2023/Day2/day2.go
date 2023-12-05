package Day2

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func CheckPossibleGames() (int, error) {

	gameList, _ := loadInput()

	possiblesGames := 0
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

func checkGame(game string) int {
	gameLine := strings.Split(game, ":")
	gameID, _ := strconv.Atoi(strings.Split(gameLine[0], " ")[1])
	gameSets := strings.Split(gameLine[1], ";")

	for _, gameSet := range gameSets {
		if !checkIsPossibleSet(gameSet) {
			return 0
		}
	}

	return gameID
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

		if count > colors[color] {
			return false
		}
	}

	return true
}

func SumPowerSets() (int, error) {
	gameList, _ := loadInput()

	sumPower := 0
	for i := 0; i < len(gameList); i++ {
		sumPower += CalculatePower(gameList[i])
	}

	return sumPower, nil
}

func CalculatePower(game string) int {
	maxColors := make(map[string]int)
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

			if count > maxColors[color] {
				maxColors[color] = count
			}
		}
	}

	totalPower := 1
	for key := range colors {
		totalPower *= maxColors[key]
	}

	return totalPower
}
