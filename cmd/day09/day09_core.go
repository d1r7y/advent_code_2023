package day09

import (
	"log"
	"strings"

	"github.com/d1r7y/advent_2023/utilities"
)

func ParseLine(line string) []int {
	return utilities.ParseIntList(line)
}

func GetDifferences(numbers []int) []int {
	differences := make([]int, 0)

	for i := 0; i < len(numbers)-1; i++ {
		differences = append(differences, numbers[i+1]-numbers[i])
	}

	return differences
}

func IsZeroDifferences(differences []int) bool {
	for _, d := range differences {
		if d != 0 {
			return false
		}
	}

	return true
}

func CalculateNextNumberForward(numbers []int) int {
	previousDifferences := make([][]int, 0)
	previousDifferences = append(previousDifferences, numbers)
	currentNumbers := numbers

	for {
		differences := GetDifferences(currentNumbers)
		if IsZeroDifferences(differences) {
			nextValue := 0
			for i := len(previousDifferences) - 1; i >= 0; i-- {
				pd := previousDifferences[i]
				nextValue = pd[len(pd)-1] + nextValue
			}

			return nextValue
		} else {
			previousDifferences = append(previousDifferences, differences)
			currentNumbers = differences
		}
	}
}

func CalculateNextNumberBackward(numbers []int) int {
	previousDifferences := make([][]int, 0)
	previousDifferences = append(previousDifferences, numbers)
	currentNumbers := numbers

	for {
		differences := GetDifferences(currentNumbers)
		if IsZeroDifferences(differences) {
			nextValue := 0
			for i := len(previousDifferences) - 1; i >= 0; i-- {
				pd := previousDifferences[i]
				nextValue = pd[0] - nextValue
			}

			return nextValue
		} else {
			previousDifferences = append(previousDifferences, differences)
			currentNumbers = differences
		}
	}
}

func day09(fileContents string) error {
	// Part 1: Analyze your OASIS report and extrapolate the next value for each history. What is the sum of these extrapolated values?
	nextNumbersForwardSum := 0

	for _, line := range strings.Split(string(fileContents), "\n") {
		numbers := ParseLine(line)
		nextNumber := CalculateNextNumberForward(numbers)
		nextNumbersForwardSum += nextNumber
	}

	log.Printf("Sum of forward extrapolated values: %d\n", nextNumbersForwardSum)

	// Part 2: Analyze your OASIS report and extrapolate the next value for each history. What is the sum of these extrapolated values?
	nextNumbersBackwardSum := 0

	for _, line := range strings.Split(string(fileContents), "\n") {
		numbers := ParseLine(line)
		nextNumber := CalculateNextNumberBackward(numbers)
		nextNumbersBackwardSum += nextNumber
	}

	log.Printf("Sum of backward extrapolated values: %d\n", nextNumbersBackwardSum)

	return nil
}
