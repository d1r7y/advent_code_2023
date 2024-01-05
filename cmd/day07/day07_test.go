package day07

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareHands(t *testing.T) {
	type testCase struct {
		hand1              Hand
		hand2              Hand
		expectedComparison int
	}

	tests := []testCase{
		{ParseHand("AAAAA", false), ParseHand("AAAAA", false), 0},
		{ParseHand("AAAAA", false), ParseHand("KKKKK", false), 1},
		{ParseHand("22222", false), ParseHand("KKKKK", false), -1},

		{ParseHand("AAAA2", false), ParseHand("AAAA3", false), -1},
		{ParseHand("AAAA3", false), ParseHand("AAAA3", false), 0},
		{ParseHand("AAAA3", false), ParseHand("AAAA2", false), 1},

		{ParseHand("AAKKK", false), ParseHand("KKAAA", false), 1},
		{ParseHand("KKAAA", false), ParseHand("AAAKK", false), -1},
		{ParseHand("AAAKK", false), ParseHand("AAAKK", false), 0},

		{ParseHand("33332", false), ParseHand("2AAAA", false), 1},

		{ParseHand("77888", false), ParseHand("77788", false), 1},

		{ParseHand("3579J", false), ParseHand("J9753", false), -1},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedComparison, CompareHands(test.hand1, test.hand2))
	}
}

func TestCompareHandsJoker(t *testing.T) {
	type testCase struct {
		hand1              Hand
		hand2              Hand
		expectedComparison int
	}

	tests := []testCase{
		{ParseHand("AAAAA", true), ParseHand("AAAAA", true), 0},
		{ParseHand("AAAAA", true), ParseHand("KKKKK", true), 1},
		{ParseHand("22222", true), ParseHand("KKKKK", true), -1},

		{ParseHand("AAAA2", true), ParseHand("AAAA3", true), -1},
		{ParseHand("AAAA3", true), ParseHand("AAAA3", true), 0},
		{ParseHand("AAAA3", true), ParseHand("AAAA2", true), 1},

		{ParseHand("AAKKK", true), ParseHand("KKAAA", true), 1},
		{ParseHand("KKAAA", true), ParseHand("AAAKK", true), -1},
		{ParseHand("AAAKK", true), ParseHand("AAAKK", true), 0},

		{ParseHand("33332", true), ParseHand("2AAAA", true), 1},

		{ParseHand("77888", true), ParseHand("77788", true), 1},

		{ParseHand("3579J", true), ParseHand("J9753", true), 1},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedComparison, CompareHands(test.hand1, test.hand2))
	}
}

func TestParseHandAndBid(t *testing.T) {
	type testCase struct {
		line         string
		jokers       bool
		expectedHand Hand
		expectedBid  Bid
	}

	tests := []testCase{
		{"32T3K 765", false, Hand{Cards: Cards{Three, Two, Ten, Three, King}, Strength: OnePair}, 765},
		{"T55J5 684", false, Hand{Cards: Cards{Ten, Five, Five, Jack, Five}, Strength: ThreeOfAKind}, 684},
		{"KK677 28", false, Hand{Cards: Cards{King, King, Six, Seven, Seven}, Strength: TwoPair}, 28},
		{"QQQJA 483", false, Hand{Cards: Cards{Queen, Queen, Queen, Jack, Ace}, Strength: ThreeOfAKind}, 483},

		{"32T3K 765", true, Hand{Cards: Cards{Three, Two, Ten, Three, King}, Strength: OnePair}, 765},
		{"T55J5 684", true, Hand{Cards: Cards{Ten, Five, Five, Joker, Five}, Strength: FourOfAKind}, 684},
		{"KK677 28", true, Hand{Cards: Cards{King, King, Six, Seven, Seven}, Strength: TwoPair}, 28},
		{"QQQJA 483", true, Hand{Cards: Cards{Queen, Queen, Queen, Joker, Ace}, Strength: FourOfAKind}, 483},
	}

	for _, test := range tests {
		hand, bid := ParseHandAndBid(test.line, test.jokers)
		assert.Equal(t, test.expectedHand, hand)
		assert.Equal(t, test.expectedBid, bid)
	}
}

func TestCalculateCardsStrength(t *testing.T) {
	type testCase struct {
		line             string
		jokers           bool
		expectedStrength Strength
	}

	tests := []testCase{
		{"32T3K", false, OnePair},
		{"T55J5", false, ThreeOfAKind},
		{"KK677", false, TwoPair},
		{"QQQJA", false, ThreeOfAKind},

		{"32T3K", true, OnePair},
		{"T55J5", true, FourOfAKind},
		{"KK677", true, TwoPair},
		{"QQQJA", true, FourOfAKind},

		{"J2345", true, OnePair},
		{"JJ234", true, ThreeOfAKind},
		{"J2234", true, ThreeOfAKind},
		{"JJJ24", true, FourOfAKind},
		{"JJ224", true, FourOfAKind},
		{"JJ222", true, FiveOfAKind},
		{"JJJJ2", true, FiveOfAKind},
		{"JJJ22", true, FiveOfAKind},
		{"JJ222", true, FiveOfAKind},
		{"J2222", true, FiveOfAKind},
		{"J2277", true, FullHouse},
	}

	for _, test := range tests {
		cards := ParseCards(test.line, test.jokers)
		if test.jokers {
			assert.Equal(t, test.expectedStrength, CalculateCardsStrengthJokers(cards), test.line)
		} else {
			assert.Equal(t, test.expectedStrength, CalculateCardsStrength(cards), test.line)
		}
	}
}

func TestHandDescribe(t *testing.T) {
	hand := ParseHand("J6JKJ", true)
	assert.Equal(t, "J6JKJ four of a kind", hand.Describe())
}

func TestTotalWinningsJoker(t *testing.T) {
	content := `
	32T3K 765
	T55J5 684
	KK677 28
	KTJJT 220
	QQQJA 483`
	handAndBidList := make([]HandAndBid, 0)

	for _, line := range strings.Split(content, "\n") {
		hb := HandAndBid{}

		if line != "" {
			hb.Hand, hb.Bid = ParseHandAndBid(line, true)
			handAndBidList = append(handAndBidList, hb)
		}
	}

	sort.Slice(handAndBidList, func(i, j int) bool {
		return CompareHands(handAndBidList[i].Hand, handAndBidList[j].Hand) < 0
	})

	totalWinnings := 0

	for rank, hb := range handAndBidList {
		totalWinnings += int(hb.Bid) * (rank + 1)
	}

	assert.Equal(t, 5905, totalWinnings)
}
