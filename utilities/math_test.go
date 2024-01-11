package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeFactors(t *testing.T) {
	type testCase struct {
		number          int
		expectedFactors []int
	}

	testCases := []testCase{
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{7, []int{7}},
		{800, []int{2, 2, 2, 2, 2, 5, 5}},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedFactors, PrimeFactors(test.number))
	}
}

func TestIntPow(t *testing.T) {
	type testCase struct {
		base          int
		exponent      int
		expectedPower int
	}

	testCases := []testCase{
		{2, 0, 1},
		{3, 3, 27},
		{2, 16, 65536},
		{8, 4, 64 * 64},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedPower, IntPow(test.base, test.exponent))
	}
}
