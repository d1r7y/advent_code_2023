package day07

import (
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Bid int

type Card int

const (
	Joker Card = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Cards [5]Card

func (c Cards) Describe() string {
	description := ""
	for _, v := range c {
		switch v {
		case Joker:
			description += "J"
		case Jack:
			description += "J"
		case Two:
			description += "2"
		case Three:
			description += "3"
		case Four:
			description += "4"
		case Five:
			description += "5"
		case Six:
			description += "6"
		case Seven:
			description += "7"
		case Eight:
			description += "8"
		case Nine:
			description += "9"
		case Ten:
			description += "T"
		case Queen:
			description += "Q"
		case King:
			description += "K"
		case Ace:
			description += "A"
		}
	}

	return description
}

type Strength int

const (
	HighCard Strength = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Strength Strength
	Cards    Cards
}

func (h Hand) Describe() string {
	description := h.Cards.Describe()

	switch h.Strength {
	case HighCard:
		description += " high card"
	case OnePair:
		description += " one pair"
	case TwoPair:
		description += " two pair"
	case ThreeOfAKind:
		description += " three of a kind"
	case FullHouse:
		description += " full house"
	case FourOfAKind:
		description += " four of a kind"
	case FiveOfAKind:
		description += " five of a kind"
	}

	return description
}

type HandAndBid struct {
	Hand Hand
	Bid  Bid
}

func getCardMap() map[string]Card {
	return map[string]Card{"2": Two, "3": Three, "4": Four, "5": Five, "6": Six, "7": Seven, "8": Eight, "9": Nine, "T": Ten, "J": Jack, "Q": Queen, "K": King, "A": Ace}
}

func getCardMapJokers() map[string]Card {
	return map[string]Card{"2": Two, "3": Three, "4": Four, "5": Five, "6": Six, "7": Seven, "8": Eight, "9": Nine, "T": Ten, "J": Joker, "Q": Queen, "K": King, "A": Ace}
}

func CalculateCardsStrength(cards Cards) Strength {
	cardCount := [int(Ace) + 1]int{}

	for _, card := range cards {
		cardCount[card]++
	}

	threeOfAKindSeen := false

	distinctCardsSeen := 0
	for _, count := range cardCount {
		if count != 0 {
			distinctCardsSeen++

			if count == 3 {
				threeOfAKindSeen = true
			}
		}
	}

	strength := HighCard

	switch distinctCardsSeen {
	case 5:
		strength = HighCard
	case 4:
		strength = OnePair
	case 3:
		if threeOfAKindSeen {
			strength = ThreeOfAKind
		} else {
			strength = TwoPair
		}
	case 2:
		if threeOfAKindSeen {
			strength = FullHouse
		} else {
			strength = FourOfAKind
		}
	case 1:
		strength = FiveOfAKind
	}

	return strength
}

func CalculateCardsStrengthJokers(cards Cards) Strength {
	cardCount := [int(Ace) + 1]int{}

	jokerCount := 0

	for _, card := range cards {
		if card == Joker {
			jokerCount++
			continue
		}

		cardCount[card]++
	}

	// If we didn't find any jokers, then we can use the non joker calculations.
	if jokerCount == 0 {
		return CalculateCardsStrength(cards)
	}

	// We have at least one joker.

	threeOfAKindSeen := false

	distinctCardsSeen := 0
	for _, count := range cardCount {
		if count != 0 {
			distinctCardsSeen++

			if count == 3 {
				threeOfAKindSeen = true
			}
		}
	}

	strength := HighCard

	switch distinctCardsSeen {
	case 4:
		// 1 joker and 4 non matching cards.
		strength = OnePair
	case 3:
		// 2 joker and 3 non matching cards, or 1 joker and a pair and another card.
		strength = ThreeOfAKind
	case 2:
		if jokerCount == 1 {
			// 1 joker and 2 pairs, or 1 joker and three of a kind and another card.
			if threeOfAKindSeen {
				strength = FourOfAKind
			} else {
				strength = FullHouse
			}
		} else {
			// 3 jokers and a non pair, or 2 jokers and a pair and another card
			strength = FourOfAKind
		}
	case 1:
		// 4 jokers and one card, or 3 jokers and a pair, or 2 jokers and three of a kind, or 1 joker and four of a kind.
		strength = FiveOfAKind
	case 0:
		// 5 jokers.
		strength = FiveOfAKind
	}

	return strength
}

func CompareHands(h1 Hand, h2 Hand) int {
	if h1.Strength != h2.Strength {
		if h1.Strength < h2.Strength {
			return -1
		} else {
			return 1
		}
	}

	// They have the same strength...Start comparing cards, left to right, and highest card wins.

	for i := 0; i < len(h1.Cards); i++ {
		if h1.Cards[i] != h2.Cards[i] {
			if h1.Cards[i] < h2.Cards[i] {
				return -1
			} else {
				return 1
			}
		}
	}

	// Identical hands.
	return 0
}

func ParseCards(str string, jokers bool) Cards {
	cardsRE := regexp.MustCompile(`[2-9TJQKA]`)
	cardsMatches := cardsRE.FindAllString(str, -1)

	cards := Cards{}

	if len(cardsMatches) != 5 {
		log.Panicf("unexpected hand: '%s'\n", str)
	}

	for i, c := range cardsMatches {
		if jokers {
			cards[i] = getCardMapJokers()[c]
		} else {
			cards[i] = getCardMap()[c]
		}
	}

	return cards
}

func ParseHand(str string, jokers bool) Hand {
	hand := Hand{}

	hand.Cards = ParseCards(str, jokers)
	if jokers {
		hand.Strength = CalculateCardsStrengthJokers(hand.Cards)
	} else {
		hand.Strength = CalculateCardsStrength(hand.Cards)
	}

	return hand
}

func ParseBid(str string) Bid {
	bid, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return Bid(bid)
}

func ParseHandAndBid(line string, jokers bool) (Hand, Bid) {
	cardsAndBid := strings.Fields(strings.TrimSpace(line))

	hand := ParseHand(cardsAndBid[0], jokers)
	bid := ParseBid(cardsAndBid[1])

	return hand, bid
}

func day07(fileContents string) error {
	// Part 1: Find the rank of every hand in your set. What are the total winnings?
	handAndBidList1 := make([]HandAndBid, 0)

	for _, line := range strings.Split(string(fileContents), "\n") {
		hb := HandAndBid{}

		hb.Hand, hb.Bid = ParseHandAndBid(line, false)
		handAndBidList1 = append(handAndBidList1, hb)
	}

	sort.Slice(handAndBidList1, func(i, j int) bool {
		return CompareHands(handAndBidList1[i].Hand, handAndBidList1[j].Hand) < 0
	})

	totalWinnings1 := 0

	for rank, hb := range handAndBidList1 {
		totalWinnings1 += int(hb.Bid) * (rank + 1)
	}

	log.Printf("Total winnings: %d\n", totalWinnings1)

	// Part 2: Using the new joker rule, find the rank of every hand in your set. What are the new total winnings?
	handAndBidList2 := make([]HandAndBid, 0)

	for _, line := range strings.Split(string(fileContents), "\n") {
		hb := HandAndBid{}

		hb.Hand, hb.Bid = ParseHandAndBid(line, true)
		handAndBidList2 = append(handAndBidList2, hb)
	}

	sort.Slice(handAndBidList2, func(i, j int) bool {
		return CompareHands(handAndBidList2[i].Hand, handAndBidList2[j].Hand) < 0
	})

	totalWinnings2 := 0

	for rank, hb := range handAndBidList2 {
		totalWinnings2 += int(hb.Bid) * (rank + 1)
	}

	log.Printf("Total winnings with jokers: %d\n", totalWinnings2)

	return nil
}
