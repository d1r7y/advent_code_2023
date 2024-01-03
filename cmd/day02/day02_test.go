package day02

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGame(t *testing.T) {
	type parseGameTest struct {
		str           string
		expectedError bool
		gameID        int
	}

	tests := []parseGameTest{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", false, 1},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", false, 2},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false, 3},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false, 4},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", false, 5},
		{"Game zzz: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true, -1},
		{"	Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", false, 1},
	}

	for _, test := range tests {
		game, err := ParseGame(test.str)
		if test.expectedError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.gameID, game.Id)
		}
	}
}

func TestIsGamePossible(t *testing.T) {
	type isGamePossibleTest struct {
		game             Game
		maxRed           int
		maxGreen         int
		maxBlue          int
		expectedPossible bool
	}

	tests := []isGamePossibleTest{
		{Game{Id: 1, Pulls: []CubePull{
			{RedCount: 1, GreenCount: 2, BlueCount: 3},
		}}, math.MaxInt, math.MaxInt, math.MaxInt, true},
		{Game{Id: 1, Pulls: []CubePull{
			{RedCount: 2, GreenCount: 2, BlueCount: 3},
		}}, 1, math.MaxInt, math.MaxInt, false},
		{Game{Id: 1, Pulls: []CubePull{
			{RedCount: 1, GreenCount: 2, BlueCount: 3},
			{RedCount: 2, GreenCount: 2, BlueCount: 3},
		}}, 1, math.MaxInt, math.MaxInt, false},
		{Game{Id: 1, Pulls: []CubePull{
			{RedCount: 0, GreenCount: 0, BlueCount: 0},
		}}, 0, 0, 0, true},
		{Game{Id: 1, Pulls: []CubePull{
			{RedCount: 3, GreenCount: 4, BlueCount: 5},
		}}, 3, 4, 5, true},

		// From sample data
		{Game{Id: 1, Pulls: []CubePull{
			{RedCount: 4, GreenCount: 0, BlueCount: 3},
			{RedCount: 1, GreenCount: 2, BlueCount: 6},
			{RedCount: 0, GreenCount: 2, BlueCount: 0},
		}}, 12, 13, 14, true},
		{Game{Id: 2, Pulls: []CubePull{
			{RedCount: 0, GreenCount: 2, BlueCount: 1},
			{RedCount: 1, GreenCount: 3, BlueCount: 4},
			{RedCount: 0, GreenCount: 1, BlueCount: 1},
		}}, 12, 13, 14, true},
		{Game{Id: 3, Pulls: []CubePull{
			{RedCount: 20, GreenCount: 8, BlueCount: 6},
			{RedCount: 4, GreenCount: 13, BlueCount: 5},
			{RedCount: 1, GreenCount: 5, BlueCount: 0},
		}}, 12, 13, 14, false},
		{Game{Id: 4, Pulls: []CubePull{
			{RedCount: 3, GreenCount: 1, BlueCount: 6},
			{RedCount: 6, GreenCount: 3, BlueCount: 0},
			{RedCount: 14, GreenCount: 3, BlueCount: 15},
		}}, 12, 13, 14, false},
		{Game{Id: 4, Pulls: []CubePull{
			{RedCount: 6, GreenCount: 3, BlueCount: 1},
			{RedCount: 1, GreenCount: 2, BlueCount: 2},
		}}, 12, 13, 14, true},
	}

	for _, test := range tests {
		possible := IsGamePossible(test.game, test.maxRed, test.maxGreen, test.maxBlue)
		assert.Equal(t, test.expectedPossible, possible)
	}
}

func TestPossibleGameSum(t *testing.T) {
	content := `
	Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	games, err := ParseGames(content)
	assert.NoError(t, err)

	sum := PossibleGameSum(games, 12, 13, 14)
	assert.Equal(t, 8, sum)
}

func TestGameMinimumCubes(t *testing.T) {
	type parseGameTest struct {
		str           string
		expectedRed   int
		expectedGreen int
		expectedBlue  int
	}

	tests := []parseGameTest{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 4, 2, 6},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 1, 3, 4},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 20, 13, 6},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 14, 3, 15},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 6, 3, 2},
	}

	for _, test := range tests {
		g, err := ParseGame(test.str)
		assert.NoError(t, err)

		red, green, blue := GameMinimumCubes(g)
		assert.Equal(t, test.expectedRed, red)
		assert.Equal(t, test.expectedGreen, green)
		assert.Equal(t, test.expectedBlue, blue)
	}
}

func TestGamePower(t *testing.T) {
	type parseGameTest struct {
		str           string
		expectedPower int
	}

	tests := []parseGameTest{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
	}

	for _, test := range tests {
		g, err := ParseGame(test.str)
		assert.NoError(t, err)

		power := GamePower(g)
		assert.Equal(t, test.expectedPower, power)
	}
}

func TestGamePowerSum(t *testing.T) {
	content := `
	Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	games, err := ParseGames(content)
	assert.NoError(t, err)

	powerSum := GamePowerSum(games)
	assert.Equal(t, 2286, powerSum)
}
