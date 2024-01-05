package day06

import (
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
		assert.Equal(t, test.expectedIntList, ParseIntList(test.line, test.prefix))
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
		assert.Equal(t, test.expectedIntList, ParseIntListRemovingAllWhitespace(test.line, test.prefix))
	}
}

func TestParseRaces(t *testing.T) {
	content := `
	Time:        35     69     68     87
	Distance:   213   1168   1086   1248`

	races := ParseRaces(content, true)

	assert.Equal(t, []Race{{35, 213}, {69, 1168}, {68, 1086}, {87, 1248}}, races)
}

func TestSpeedAtTime(t *testing.T) {
	type speedAtTimeTestCase struct {
		pressTime     int
		raceTime      int
		expectedSpeed int
	}

	testCases := []speedAtTimeTestCase{
		{0, 7, 0},
		{1, 7, 1},
		{2, 7, 2},
		{3, 7, 3},
		{4, 7, 4},
		{5, 7, 5},
		{6, 7, 6},
		{7, 7, 0},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedSpeed, SpeedAtTime(test.pressTime, test.raceTime))
	}
}

func TestDistanceAtTime(t *testing.T) {
	type distanceAtTimeTestCase struct {
		speed            int
		time             int
		expectedDistance int
	}

	testCases := []distanceAtTimeTestCase{
		{0, 7, 0},
		{1, 6, 6},
		{2, 5, 10},
		{3, 4, 12},
		{4, 3, 12},
		{5, 2, 10},
		{6, 1, 6},
		{7, 0, 0},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedDistance, DistanceAtTime(test.speed, test.time))
	}
}

func TestDoesBeatRecord(t *testing.T) {
	type distanceAtTimeTestCase struct {
		pressTime      int
		raceTime       int
		recordDistance int
		expectedBeat   bool
	}

	testCases := []distanceAtTimeTestCase{
		{0, 7, 9, false},
		{1, 7, 9, false},
		{2, 7, 9, true},
		{3, 7, 9, true},
		{4, 7, 9, true},
		{5, 7, 9, true},
		{6, 7, 9, false},
		{7, 7, 9, false},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedBeat, DoesBeatRecord(test.pressTime, test.raceTime, test.recordDistance))
	}
}

func TestDetermineWinningWays(t *testing.T) {
	race := Race{Time: 7, Distance: 9}

	beatRecordCount := 0

	for pt := 0; pt < race.Time; pt++ {
		if DoesBeatRecord(pt, race.Time, race.Distance) {
			beatRecordCount++
		}
	}

	assert.Equal(t, 4, beatRecordCount)
}

func TestDetermineTotalWinningWaysPart1(t *testing.T) {
	content := `
	Time:      7  15   30
	Distance:  9  40  200`

	races := ParseRaces(content, true)

	totalWinningWays := 1

	for _, r := range races {
		beatRecordCount := 0
		for pt := 0; pt < r.Time; pt++ {
			if DoesBeatRecord(pt, r.Time, r.Distance) {
				beatRecordCount++
			}
		}

		totalWinningWays *= beatRecordCount
	}

	assert.Equal(t, 288, totalWinningWays)
}

func TestDetermineTotalWinningWaysPart2(t *testing.T) {
	content := `
	Time:      7  15   30
	Distance:  9  40  200`

	races := ParseRaces(content, false)

	totalWinningWays := 1

	for _, r := range races {
		beatRecordCount := 0
		for pt := 0; pt < r.Time; pt++ {
			if DoesBeatRecord(pt, r.Time, r.Distance) {
				beatRecordCount++
			}
		}

		totalWinningWays *= beatRecordCount
	}

	assert.Equal(t, 71503, totalWinningWays)
}
