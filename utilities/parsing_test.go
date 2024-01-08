package utilities

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIntList(t *testing.T) {
	type parseIntListTestCase struct {
		line            string
		prefix          string
		expectedIntList []int
	}

	testCases := []parseIntListTestCase{
		{"Time:        35     69     68     87", "Time: ", []int{35, 69, 68, 87}},
		{"Distance:   213   1168   1086   1248", "Distance: ", []int{213, 1168, 1086, 1248}},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedIntList, ParseIntList(strings.TrimPrefix(test.line, test.prefix)))
	}
}

func TestParseIntListRemovingAllWhitespace(t *testing.T) {
	type parseIntListTestCase struct {
		line            string
		prefix          string
		expectedIntList []int
	}

	testCases := []parseIntListTestCase{
		{"Time:        35     69     68     87", "Time: ", []int{35696887}},
		{"Distance:   213   1168   1086   1248", "Distance: ", []int{213116810861248}},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedIntList, ParseIntListRemovingAllWhitespace(strings.TrimPrefix(test.line, test.prefix)))
	}
}

func TestParseIntListNegative(t *testing.T) {
	type parseIntListTestCase struct {
		line            string
		prefix          string
		expectedIntList []int
	}

	testCases := []parseIntListTestCase{
		{"Time:        -35     69     68     87", "Time: ", []int{-35, 69, 68, 87}},
		{"Distance:   213   1168   1086   -1248", "Distance: ", []int{213, 1168, 1086, -1248}},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedIntList, ParseIntList(strings.TrimPrefix(test.line, test.prefix)))
	}
}
