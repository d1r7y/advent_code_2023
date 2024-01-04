package day03

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const Gear = '*'

type AdjacentNumberIndexes struct {
	Map map[int]bool
}

func NewAdjacentNumberIndexes() *AdjacentNumberIndexes {
	return &AdjacentNumberIndexes{Map: make(map[int]bool)}
}

func (a *AdjacentNumberIndexes) AddAdjacentNumberIndex(index int) {
	a.Map[index] = true
}

func (a *AdjacentNumberIndexes) GetAdjacentNumberIndexes() []int {
	adjacentIndexes := make([]int, 0)
	for k, _ := range a.Map {
		adjacentIndexes = append(adjacentIndexes, k)
	}

	return adjacentIndexes
}

type Location struct {
	X int
	Y int
}

type Part struct {
	Name                  byte
	AdjacentNumberIndexes *AdjacentNumberIndexes
	Location              Location
}

type Number struct {
	Number int
	Start  Location
	End    Location
}

func IsAdjacent(point Location, lineStart Location, lineEnd Location) bool {
	if point.Y < (lineStart.Y-1) || point.Y > (lineStart.Y+1) {
		return false
	}
	if point.X < (lineStart.X-1) || point.X > (lineEnd.X+1) {
		return false
	}

	return true
}

func ParseSchematicLine(y int, line string) ([]Part, []Number, error) {
	re := regexp.MustCompile(`(?:[0-9]+)|(?:[^.])`)
	matches := re.FindAllStringIndex(line, -1)
	if matches == nil {
		return nil, nil, fmt.Errorf("unexpected line '%s'", line)
	}

	parts := make([]Part, 0)
	numbers := make([]Number, 0)

	for _, mi := range matches {
		if unicode.IsDigit(rune(line[mi[0]])) {
			// This match must be a part number.
			partNumber, err := strconv.Atoi(line[mi[0]:mi[1]])
			if err != nil {
				return nil, nil, err
			}

			number := Number{
				Number: partNumber,
				Start:  Location{mi[0], y},
				End:    Location{mi[1] - 1, y},
			}

			numbers = append(numbers, number)
		} else {
			// This match must be a symbol.
			part := Part{
				Name:                  line[mi[0]],
				AdjacentNumberIndexes: NewAdjacentNumberIndexes(),
				Location:              Location{mi[0], y},
			}

			parts = append(parts, part)
		}
	}

	return parts, numbers, nil
}

func ParseSchematic(fileContents string) ([]Part, []Number, error) {
	y := 0

	allParts := make([]Part, 0)
	allNumbers := make([]Number, 0)

	for _, line := range strings.Split(fileContents, "\n") {
		if line != "" {
			parts, numbers, err := ParseSchematicLine(y, strings.TrimSpace(line))
			if err != nil {
				return nil, nil, err
			}

			allParts = append(allParts, parts...)
			allNumbers = append(allNumbers, numbers...)
			y++
		}
	}

	return allParts, allNumbers, nil
}

func day03(fileContents string) error {
	allParts, allNumbers, err := ParseSchematic(fileContents)
	if err != nil {
		log.Fatal(err)
	}

	// Find the number adjacent to each part.
	for i, _ := range allParts {

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

	log.Printf("Sum of part numbers: %d\n", partNumberSum)

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

	log.Printf("Sum of all gear ratios: %d\n", gearRatioSum)

	return nil
}
