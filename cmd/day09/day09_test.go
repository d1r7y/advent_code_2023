package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	type testCase struct {
		line            string
		expectedNumbers []int
	}

	testCases := []testCase{
		{"0 3 6 9 12 15", []int{0, 3, 6, 9, 12, 15}},
		{"0 3 6 9 -12 15", []int{0, 3, 6, 9, -12, 15}},
		{"10 13 16 21 30 45", []int{10, 13, 16, 21, 30, 45}},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedNumbers, ParseLine(test.line))
	}
}

func TestGetDifferences(t *testing.T) {
	type testCase struct {
		numbers             []int
		expectedDifferences []int
	}

	testCases := []testCase{
		{[]int{0, 3, 6, 9, 12, 15}, []int{3, 3, 3, 3, 3}},
		{[]int{1, 3, 6, 10, 15, 21}, []int{2, 3, 4, 5, 6}},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedDifferences, GetDifferences(test.numbers))
	}
}

func TestIsZeroDifferences(t *testing.T) {
	type testCase struct {
		differences  []int
		expectedZero bool
	}

	testCases := []testCase{
		{[]int{0, 3, 6, 9, 12, 15}, false},
		{[]int{1, 3, 6, 10, 15, 21}, false},
		{[]int{0, 0, 0}, true},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedZero, IsZeroDifferences(test.differences))
	}
}

func TestCalculateNextNumberForward(t *testing.T) {
	type testCase struct {
		numbers            []int
		expectedNextNumber int
	}

	testCases := []testCase{
		{[]int{0, 3, 6, 9, 12, 15}, 18},
		{[]int{1, 3, 6, 10, 15, 21}, 28},
		{[]int{10, 13, 16, 21, 30, 45}, 68},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedNextNumber, CalculateNextNumberForward(test.numbers))
	}
}

func TestCalculateNextNumberBackward(t *testing.T) {
	type testCase struct {
		numbers            []int
		expectedNextNumber int
	}

	testCases := []testCase{
		{[]int{10, 13, 16, 21, 30, 45}, 5},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedNextNumber, CalculateNextNumberBackward(test.numbers))
	}
}
