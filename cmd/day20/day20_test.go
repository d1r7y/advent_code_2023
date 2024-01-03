package day20

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescribe(t *testing.T) {
	type testCase struct {
		str                 string
		expectedDescription string
	}
	testCases := []testCase{
		{
			str: `1
2
-3
3
-2
0
4`,
			expectedDescription: "1, 2, -3, 3, -2, 0, 4",
		},
	}

	for _, test := range testCases {
		wl := ParseWrappedList(test.str)
		assert.Equal(t, test.expectedDescription, wl.Describe())
	}
}

func TestNewIndex(t *testing.T) {
	type testCase struct {
		index            int
		delta            int
		expectedNewIndex int
	}
	testCases := []testCase{
		{
			index:            0,
			delta:            1,
			expectedNewIndex: 1,
		},
		{
			index:            0,
			delta:            2,
			expectedNewIndex: 2,
		},
		{
			index:            1,
			delta:            -3,
			expectedNewIndex: 4,
		},
		{
			index:            2,
			delta:            3,
			expectedNewIndex: 5,
		},
		{
			index:            2,
			delta:            -2,
			expectedNewIndex: 6,
		},
		{
			index:            3,
			delta:            0,
			expectedNewIndex: 3,
		},
		{
			index:            5,
			delta:            4,
			expectedNewIndex: 3,
		},
	}

	str := `1
2
-3
3
-2
0
4`

	for _, test := range testCases {
		wl := ParseWrappedList(str)
		assert.Equal(t, test.expectedNewIndex, wl.NewIndex(test.index, test.delta), fmt.Sprintf("index=%d delta=%d", test.index, test.delta))
	}
}

func TestMove(t *testing.T) {
	type testCase struct {
		str          string
		index        int
		delta        int
		expectedList string
	}
	testCases := []testCase{
		{
			str: `1
2
-3
3
-2
0
4`,
			index:        0,
			delta:        1,
			expectedList: "2, 1, -3, 3, -2, 0, 4",
		},
		{
			str: `2
1
-3
3
-2
0
4`,
			index:        0,
			delta:        2,
			expectedList: "1, -3, 2, 3, -2, 0, 4",
		},
		{
			str: `1
-3
2
3
-2
0
4`,
			index:        1,
			delta:        -3,
			expectedList: "1, 2, 3, -2, -3, 0, 4",
		},
		{
			str: `1
2
3
-2
-3
0
4`,
			index:        2,
			delta:        3,
			expectedList: "1, 2, -2, -3, 0, 3, 4",
		},
		{
			str: `1
2
-2
-3
0
3
4`,
			index:        2,
			delta:        -2,
			expectedList: "1, 2, -3, 0, 3, 4, -2",
		},
		{
			str: `1
2
-3
0
3
4
-2`,
			index:        3,
			delta:        0,
			expectedList: "1, 2, -3, 0, 3, 4, -2",
		},
		{
			str: `1
2
-3
0
3
4
-2`,
			index:        5,
			delta:        4,
			expectedList: "1, 2, -3, 4, 0, 3, -2",
		},
	}

	for _, test := range testCases {
		wl := ParseWrappedList(test.str)
		wl.Move(test.index, test.delta)
		assert.Equal(t, test.expectedList, wl.Describe(), fmt.Sprintf("index=%d delta=%d", test.index, test.delta))
	}
}

func TestMix(t *testing.T) {
	type testCase struct {
		str          string
		expectedList string
	}
	testCases := []testCase{
		{
			str: `1
2
-3
3
-2
0
4`,
			expectedList: "1, 2, -3, 4, 0, 3, -2",
		},
	}

	for _, test := range testCases {
		wl := ParseWrappedList(test.str)
		wl.Mix()
		assert.Equal(t, test.expectedList, wl.Describe())
	}
}

func TestGetCoordinates(t *testing.T) {
	str := `1
2
-3
4
0
3
-2`

	wl := ParseWrappedList(str)
	assert.Equal(t, [3]int{4, -3, 2}, wl.GetCoordinates())
}
