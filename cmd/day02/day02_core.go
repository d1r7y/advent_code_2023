package day02

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type CubePull struct {
	RedCount   int
	GreenCount int
	BlueCount  int
}

type Game struct {
	Id    int
	Pulls []CubePull
}

func ParseGame(line string) (Game, error) {
	game := Game{Pulls: make([]CubePull, 0)}

	gameAndRemainingRE := regexp.MustCompile(`\s*Game\s(\d+):(.*)`)
	gameAndRemainingMatches := gameAndRemainingRE.FindSubmatch([]byte(line))
	if gameAndRemainingMatches == nil {
		return game, fmt.Errorf("invalid game: '%s'", line)
	}
	gameID, err := strconv.Atoi(string(gameAndRemainingMatches[1]))
	if err != nil {
		return game, err
	}
	if gameID <= 0 {
		return game, fmt.Errorf("invalid id")
	}

	game.Id = gameID

	for _, p := range strings.Split(string(gameAndRemainingMatches[2]), ";") {
		var pull CubePull

		for _, m := range strings.Split(p, ",") {
			countAndColorRE := regexp.MustCompile(`(\d+)\s(red|green|blue)`)
			countAndColor := countAndColorRE.FindSubmatch([]byte(m))
			count, err := strconv.Atoi(string(countAndColor[1]))
			if err != nil {
				return game, err
			}

			if count <= 0 {
				return game, fmt.Errorf("invalid count")
			}

			color := string(countAndColor[2])

			switch color {
			case "red":
				pull.RedCount += count
			case "green":
				pull.GreenCount += count
			case "blue":
				pull.BlueCount += count
			}
		}

		game.Pulls = append(game.Pulls, pull)
	}

	return game, nil
}

func ParseGames(text string) ([]Game, error) {
	games := make([]Game, 0)

	if text == "" {
		return games, nil
	}

	for _, line := range strings.Split(text, "\n") {
		if line != "" {
			g, err := ParseGame(line)
			if err != nil {
				return games, err
			}

			games = append(games, g)
		}
	}

	return games, nil
}

func IsGamePossible(g Game, maxRed, maxGreen, maxBlue int) bool {
	for _, p := range g.Pulls {
		if p.RedCount > maxRed {
			return false
		}
		if p.GreenCount > maxGreen {
			return false
		}
		if p.BlueCount > maxBlue {
			return false
		}
	}

	return true
}

func PossibleGameSum(games []Game, maxRed, maxGreen, maxBlue int) int {
	gameIDSum := 0

	for _, g := range games {
		if IsGamePossible(g, maxRed, maxGreen, maxBlue) {
			gameIDSum += g.Id
		}
	}

	return gameIDSum
}

func GameMinimumCubes(g Game) (red int, green int, blue int) {
	for _, p := range g.Pulls {
		if p.RedCount > red {
			red = p.RedCount
		}
		if p.GreenCount > green {
			green = p.GreenCount
		}
		if p.BlueCount > blue {
			blue = p.BlueCount
		}
	}

	return red, green, blue
}

func GamePower(g Game) int {
	red, green, blue := GameMinimumCubes(g)
	return red * green * blue
}

func GamePowerSum(games []Game) int {
	powerSum := 0

	for _, g := range games {
		powerSum += GamePower(g)
	}

	return powerSum
}

func day02(fileContents string) error {
	games, err := ParseGames(string(fileContents))
	if err != nil {
		return err
	}

	// Part 1: Determine which games would have been possible if the bag had been loaded with
	// only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
	gameIDSum := PossibleGameSum(games, 12, 13, 14)

	log.Printf("Sum of possible game IDs: %d\n", gameIDSum)

	// Part 2: For each game, find the minimum set of cubes that must have been present. What is the
	// sum of the power of these sets?

	powerSum := GamePowerSum(games)

	log.Printf("Sum of game powers: %d\n", powerSum)

	return nil
}
