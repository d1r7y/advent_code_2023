package day03

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAdjacent(t *testing.T) {
	type isAdjacentTest struct {
		point            Location
		lineStart        Location
		lineEnd          Location
		expectedAdjacent bool
	}

	tests := []isAdjacentTest{
		{Location{0, 0}, Location{0, 1}, Location{2, 1}, true},
		{Location{0, 1}, Location{0, 1}, Location{2, 1}, true},
		{Location{0, 2}, Location{0, 1}, Location{2, 1}, true},
		{Location{0, 3}, Location{0, 1}, Location{2, 1}, false},
		{Location{0, -1}, Location{0, 1}, Location{2, 1}, false},

		{Location{-1, 0}, Location{0, 1}, Location{2, 1}, true},
		{Location{-1, 1}, Location{0, 1}, Location{2, 1}, true},
		{Location{-1, 2}, Location{0, 1}, Location{2, 1}, true},
		{Location{-1, 3}, Location{0, 1}, Location{2, 1}, false},
		{Location{-1, -1}, Location{0, 1}, Location{2, 1}, false},

		{Location{-2, 0}, Location{0, 1}, Location{2, 1}, false},
		{Location{-2, 1}, Location{0, 1}, Location{2, 1}, false},
		{Location{-2, 2}, Location{0, 1}, Location{2, 1}, false},
		{Location{-2, 3}, Location{0, 1}, Location{2, 1}, false},
		{Location{-2, -1}, Location{0, 1}, Location{2, 1}, false},

		{Location{1, 0}, Location{0, 1}, Location{2, 1}, true},
		{Location{1, 2}, Location{0, 1}, Location{2, 1}, true},

		{Location{2, 0}, Location{0, 1}, Location{2, 1}, true},
		{Location{2, 1}, Location{0, 1}, Location{2, 1}, true},
		{Location{2, 2}, Location{0, 1}, Location{2, 1}, true},
		{Location{2, 3}, Location{0, 1}, Location{2, 1}, false},
		{Location{2, -1}, Location{0, 1}, Location{2, 1}, false},

		{Location{3, 0}, Location{0, 1}, Location{2, 1}, true},
		{Location{3, 1}, Location{0, 1}, Location{2, 1}, true},
		{Location{3, 2}, Location{0, 1}, Location{2, 1}, true},
		{Location{3, 3}, Location{0, 1}, Location{2, 1}, false},
		{Location{3, -1}, Location{0, 1}, Location{2, 1}, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedAdjacent, IsAdjacent(test.point, test.lineStart, test.lineEnd), fmt.Sprintf("P: %d,%d, L: %d,%d - %d,%d", test.point.X, test.point.Y, test.lineStart.X, test.lineStart.Y, test.lineEnd.X, test.lineEnd.Y))
	}
}

func TestParseSchematicLine(t *testing.T) {
	type parseSchematicLineTest struct {
		line            string
		expectedErr     bool
		expectedParts   []Part
		expectedNumbers []Number
	}

	tests := []parseSchematicLineTest{
		{"467..114..", false, []Part{}, []Number{{Number: 467, Start: Location{0, 0}, End: Location{2, 0}}, {Number: 114, Start: Location{5, 0}, End: Location{7, 0}}}},
		{"...*......", false, []Part{{Name: '*', AdjacentNumberIndexes: NewAdjacentNumberIndexes(), Location: Location{3, 0}}}, []Number{}},
		{"..35..633.", false, []Part{}, []Number{{Number: 35, Start: Location{2, 0}, End: Location{3, 0}}, {Number: 633, Start: Location{6, 0}, End: Location{8, 0}}}},
		{"......#...", false, []Part{{Name: '#', AdjacentNumberIndexes: NewAdjacentNumberIndexes(), Location: Location{6, 0}}}, []Number{}},
		{"617*......", false, []Part{{Name: '*', AdjacentNumberIndexes: NewAdjacentNumberIndexes(), Location: Location{3, 0}}}, []Number{{Number: 617, Start: Location{0, 0}, End: Location{2, 0}}}},
		{".....+.58.", false, []Part{{Name: '+', AdjacentNumberIndexes: NewAdjacentNumberIndexes(), Location: Location{5, 0}}}, []Number{{Number: 58, Start: Location{7, 0}, End: Location{8, 0}}}},
		{"..592.....", false, []Part{}, []Number{{Number: 592, Start: Location{2, 0}, End: Location{4, 0}}}},
		{"......755.", false, []Part{}, []Number{{Number: 755, Start: Location{6, 0}, End: Location{8, 0}}}},
		{"...$.*....", false, []Part{{Name: '$', AdjacentNumberIndexes: NewAdjacentNumberIndexes(), Location: Location{3, 0}}, {Name: '*', AdjacentNumberIndexes: NewAdjacentNumberIndexes(), Location: Location{5, 0}}}, []Number{}},
		{".664.598..", false, []Part{}, []Number{{Number: 664, Start: Location{1, 0}, End: Location{3, 0}}, {Number: 598, Start: Location{5, 0}, End: Location{7, 0}}}},
	}

	for _, test := range tests {
		parts, numbers, err := ParseSchematicLine(0, test.line)
		if test.expectedErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)

			assert.Equal(t, parts, test.expectedParts)
			assert.Equal(t, numbers, test.expectedNumbers)
		}
	}
}

func TestParseSchematicPartNumberSum(t *testing.T) {
	content := `
	467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..`

	allParts, allNumbers, err := ParseSchematic(content)
	assert.NoError(t, err)

	// Find the number adjacent to each part.
	for i := range allParts {

		for j, n := range allNumbers {
			if IsAdjacent(allParts[i].Location, n.Start, n.End) {
				allParts[i].AdjacentNumberIndexes.AddAdjacentNumberIndex(j)
			}
		}
	}

	// Part 1: What is the sum of all of the part numbers in the engine schematic?
	partNumberSum := 0
	for _, p := range allParts {
		for _, index := range p.AdjacentNumberIndexes.GetAdjacentNumberIndexes() {
			partNumberSum += allNumbers[index].Number
		}
	}

	assert.Equal(t, 4361, partNumberSum)
}

func TestParseSchematicGearRatioSum(t *testing.T) {
	content := `
	467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..`

	allParts, allNumbers, err := ParseSchematic(content)
	assert.NoError(t, err)

	// Find the number adjacent to each part.
	for i := range allParts {

		for j, n := range allNumbers {
			if IsAdjacent(allParts[i].Location, n.Start, n.End) {
				allParts[i].AdjacentNumberIndexes.AddAdjacentNumberIndex(j)
			}
		}
	}

	// Part 2: What is the sum of all of the gear ratios in your engine schematic?
	gearRatioSum := 0
	for _, p := range allParts {
		if p.Name == Gear {
			// Check to see it has exactly two adjacent part numbers.
			adjacentNumberIndex := p.AdjacentNumberIndexes.GetAdjacentNumberIndexes()
			if len(adjacentNumberIndex) != 2 {
				continue
			}

			number1 := allNumbers[adjacentNumberIndex[0]].Number
			number2 := allNumbers[adjacentNumberIndex[1]].Number
			gearRatioSum += number1 * number2
		}
	}

	assert.Equal(t, 467835, gearRatioSum)
}
