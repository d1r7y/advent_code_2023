package day04

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	Number         int
	Count          int
	WinningNumbers []int
	MyNumbers      []int
}

func ParseCard(line string) (*Card, error) {
	re := regexp.MustCompile(`Card\s+(\d+):(.*)`)
	cardNumberAndRemainingMatches := re.FindStringSubmatch(line)
	if cardNumberAndRemainingMatches == nil {
		return nil, fmt.Errorf("invalid card: '%s'", line)
	}

	cardNumber, err := strconv.Atoi(string(cardNumberAndRemainingMatches[1]))
	if err != nil {
		return nil, err
	}

	// Split the remaining matches around |.  Left side is list of winning numbers, right side is list of numbers you have.
	numbers := strings.Split(string(cardNumberAndRemainingMatches[2]), "|")

	winningNumbers := make([]int, 0)
	myNumbers := make([]int, 0)

	for _, n := range strings.Fields(numbers[0]) {
		number, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		winningNumbers = append(winningNumbers, number)
	}

	for _, n := range strings.Fields(numbers[1]) {
		number, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		myNumbers = append(myNumbers, number)
	}

	card := &Card{
		Number:         cardNumber,
		Count:          1,
		WinningNumbers: winningNumbers,
		MyNumbers:      myNumbers,
	}

	return card, nil
}

func (c *Card) WinningMatches() int {
	winningMatches := 0

	contains := func(list []int, value int) bool {
		for _, lv := range list {
			if lv == value {
				return true
			}
		}
		return false
	}

	for _, wn := range c.WinningNumbers {
		if contains(c.MyNumbers, wn) {
			winningMatches++
		}
	}

	return winningMatches
}

func (c *Card) Worth() int {
	winningMatches := c.WinningMatches()

	return int(math.Pow(2, float64(winningMatches-1)))
}

func day04(fileContents string) error {
	cards := make([]*Card, 0)
	for _, line := range strings.Split(fileContents, "\n") {
		if line != "" {
			card, err := ParseCard(line)
			if err != nil {
				log.Fatal(err)
			}

			cards = append(cards, card)
		}
	}

	// Part 1: How many points are they worth in total?
	totalPoints := 0

	for _, card := range cards {
		totalPoints += card.Worth()
	}

	log.Printf("Cards worth a total of %d points\n", totalPoints)

	// Part 2: Including the original set of scratchcards, how many total scratchcards do you end up with?
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

	log.Printf("Total cards won %d\n", totalCardsWon)

	return nil
}
