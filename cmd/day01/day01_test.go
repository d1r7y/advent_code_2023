package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCalibrationValue(t *testing.T) {
	type calculateCalibrationValueTest struct {
		value int
		line  string
	}

	tests := []calculateCalibrationValueTest{
		{12, "1abc2"},
		{38, "pqr3stu8vwx"},
		{15, "a1b2c3d4e5f"},
		{77, "treb7uchet"},
		{29, "two1nine"},
		{83, "eightwothree"},
		{13, "abcone2threexyz"},
		{24, "xtwone3four"},
		{42, "4nineeightseven2"},
		{14, "zoneight234"},
		{76, "7pqrstsixteen"},
		{91, "vmtkqpjftc9twonej"},
	}

	for _, test := range tests {
		v, err := CalculateCalibrationValue(test.line)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, test.value, v, "unexpected calibration value")
	}
}

func TestDigitFromString(t *testing.T) {
	type testDigitFromStringTest struct {
		str           string
		digit         int
		consumed      int
		expectedError bool
	}

	tests := []testDigitFromStringTest{
		{"zero", -1, 0, true},
		{"one", 1, len("one"), false},
		{"two", 2, len("two"), false},
		{"three", 3, len("three"), false},
		{"four", 4, len("four"), false},
		{"five", 5, len("five"), false},
		{"six", 6, len("six"), false},
		{"seven", 7, len("seven"), false},
		{"eight", 8, len("eight"), false},
		{"nine", 9, len("nine"), false},
		{"ten", -1, 0, true},
		{"sixteen", 6, len("six"), false},
		{"eightwothree", 8, len("eight"), false},
	}

	for _, test := range tests {
		d, c, err := DigitFromString(test.str)
		if test.expectedError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.digit, d)
			assert.Equal(t, test.consumed, c)
		}
	}
}
