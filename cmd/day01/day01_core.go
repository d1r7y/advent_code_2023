/*
Copyright © 2022 Cameron Esfahani
*/
package day01

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

func DigitFromString(str string) (int, int, error) {
	type strDigitLookup struct {
		str   string
		digit int
	}

	lt := []strDigitLookup{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}

	currentIndex := 0
	s := str[currentIndex:]
	for s != "" {
		if unicode.IsDigit(rune(s[0])) {
			return -1, currentIndex, fmt.Errorf("unexpected character '%s'", s)
		}

		for _, e := range lt {
			if strings.HasPrefix(s, e.str) {
				return e.digit, currentIndex + len(e.str), nil
			}
		}

		currentIndex++
		s = str[currentIndex:]
	}

	return -1, currentIndex, fmt.Errorf("unknown string '%s'", s)
}

func CalculateCalibrationValue(line string) (int, error) {
	firstDigit := -1
	lastDigit := -1

	for i := 0; i < len(line); i++ {
		c := line[i]
		digit := -1

		if unicode.IsDigit(rune(c)) {
			digit = int(c - '0')
		} else {
			possibleDigit, _, err := DigitFromString(line[i:])
			if err == nil {
				digit = possibleDigit
			}
		}

		if digit >= 0 {
			if firstDigit < 0 {
				firstDigit = digit
			} else {
				lastDigit = digit
			}
		}
	}

	if firstDigit < 0 {
		return -1, fmt.Errorf("no digits")
	}
	if lastDigit < 0 {
		lastDigit = firstDigit
	}

	return (firstDigit * 10) + lastDigit, nil
}

func ParseCalibrationValues(text string) (int, error) {
	calibrationSum := 0

	if text == "" {
		return calibrationSum, nil
	}

	for _, line := range strings.Split(text, "\n") {
		v, err := CalculateCalibrationValue(line)
		if err != nil {
			return calibrationSum, err
		}

		calibrationSum += v
	}

	return calibrationSum, nil
}

func day01(fileContent string) error {
	calibrationSum, err := ParseCalibrationValues(string(fileContent))
	if err != nil {
		log.Fatal(err)
	}

	// Part 1: What is the sum of all of the calibration values?
	log.Printf("Sum of all calibration values: %d\n", calibrationSum)

	return nil
}
