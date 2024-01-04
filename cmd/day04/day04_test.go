package day04

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCard(t *testing.T) {
	type parseCardTest struct {
		line         string
		expectedCard *Card
	}

	tests := []parseCardTest{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", &Card{Number: 1, Count: 1, WinningNumbers: []int{41, 48, 83, 86, 17}, MyNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53}}},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", &Card{Number: 2, Count: 1, WinningNumbers: []int{13, 32, 20, 16, 61}, MyNumbers: []int{61, 30, 68, 82, 17, 32, 24, 19}}},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", &Card{Number: 3, Count: 1, WinningNumbers: []int{1, 21, 53, 59, 44}, MyNumbers: []int{69, 82, 63, 72, 16, 21, 14, 1}}},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", &Card{Number: 4, Count: 1, WinningNumbers: []int{41, 92, 73, 84, 69}, MyNumbers: []int{59, 84, 76, 51, 58, 5, 54, 83}}},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", &Card{Number: 5, Count: 1, WinningNumbers: []int{87, 83, 26, 28, 32}, MyNumbers: []int{88, 30, 70, 12, 93, 22, 82, 36}}},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", &Card{Number: 6, Count: 1, WinningNumbers: []int{31, 18, 13, 56, 72}, MyNumbers: []int{74, 77, 10, 23, 35, 67, 36, 11}}},
	}

	for _, test := range tests {
		card, err := ParseCard(test.line)

		assert.NoError(t, err)
		assert.Equal(t, test.expectedCard, card)
	}
}

func TestWinningMatches(t *testing.T) {
	type parseCardTest struct {
		line                   string
		expectedWinningMatches int
	}

	tests := []parseCardTest{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 4},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
	}

	for _, test := range tests {
		card, err := ParseCard(test.line)

		assert.NoError(t, err)
		assert.Equal(t, test.expectedWinningMatches, card.WinningMatches())
	}
}

func TestWorth(t *testing.T) {
	type parseCardTest struct {
		line          string
		expectedWorth int
	}

	tests := []parseCardTest{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
	}

	for _, test := range tests {
		card, err := ParseCard(test.line)

		assert.NoError(t, err)
		assert.Equal(t, test.expectedWorth, card.Worth())
	}
}

func TestTotalCardsWon(t *testing.T) {
	content := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	cards := make([]*Card, 0)
	for _, line := range strings.Split(content, "\n") {
		if line != "" {
			card, err := ParseCard(strings.TrimSpace(line))
			if err != nil {
				log.Fatal(err)
			}

			cards = append(cards, card)
		}
	}

	for i, _ := range cards {
		winningMatches := cards[i].WinningMatches()

		for duplicates := 0; duplicates < cards[i].Count; duplicates++ {
			for j := i + 1; j <= i+winningMatches; j++ {
				cards[j].Count++
			}
		}
	}

	totalCardsWon := 0
	for _, c := range cards {
		totalCardsWon += c.Count
	}

	assert.Equal(t, 30, totalCardsWon)
}
