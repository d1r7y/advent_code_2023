package day11

import (
	"strings"
	"testing"

	"github.com/d1r7y/advent_2023/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseUniverse(t *testing.T) {
	content := `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	universe := ParseUniverse(strings.Split(strings.TrimSpace(content), "\n"))
	assert.Equal(t, &Universe{
		Bounds: utilities.NewSize2D(10, 10),
		Galaxies: []utilities.Point2D{
			utilities.NewPoint2D(3, 0),
			utilities.NewPoint2D(7, 1),
			utilities.NewPoint2D(0, 2),
			utilities.NewPoint2D(6, 4),
			utilities.NewPoint2D(1, 5),
			utilities.NewPoint2D(9, 6),
			utilities.NewPoint2D(7, 8),
			utilities.NewPoint2D(0, 9),
			utilities.NewPoint2D(4, 9),
		}}, universe)
}

func TestUniverseUnpopulatedRows(t *testing.T) {
	content := `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	universe := ParseUniverse(strings.Split(strings.TrimSpace(content), "\n"))

	assert.Equal(t, []int{3, 7}, universe.UnpopulatedRows())
}

func TestUniverseUnpopulatedColumns(t *testing.T) {
	content := `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	universe := ParseUniverse(strings.Split(strings.TrimSpace(content), "\n"))

	assert.Equal(t, []int{2, 5, 8}, universe.UnpopulatedColumns())
}

func TestUniverseDescribe(t *testing.T) {
	content := `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	universe := ParseUniverse(strings.Split(strings.TrimSpace(content), "\n"))
	assert.Equal(t, strings.TrimSpace(content), universe.Describe())
}

func TestUniverseExpand(t *testing.T) {
	content := `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	universe := ParseUniverse(strings.Split(strings.TrimSpace(content), "\n"))
	universe.Expand(2)

	assert.Equal(t, utilities.NewSize2D(13, 12), universe.Bounds)

	expectedContent := `
....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`

	assert.Equal(t, strings.TrimSpace(expectedContent), universe.Describe())
}

func TestUniverseExpandOlder(t *testing.T) {
	content := `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	universe := ParseUniverse(strings.Split(strings.TrimSpace(content), "\n"))
	universe.Expand(3)

	assert.Equal(t, utilities.NewSize2D(16, 14), universe.Bounds)

	expectedContent := `
.....#..........
...........#....
#...............
................
................
................
..........#.....
.#..............
...............#
................
................
................
...........#....
#.....#.........`

	assert.Equal(t, strings.TrimSpace(expectedContent), universe.Describe())
}

func TestGetPartners(t *testing.T) {
	type testCase struct {
		id               int
		expectedPartners []int
	}

	testCases := []testCase{
		{0, []int{}},
		{1, []int{0}},
		{2, []int{0, 1}},
		{3, []int{0, 1, 2}},
		{4, []int{0, 1, 2, 3}},
		{5, []int{0, 1, 2, 3, 4}},
		{6, []int{0, 1, 2, 3, 4, 5}},
		{7, []int{0, 1, 2, 3, 4, 5, 6}},
		{8, []int{0, 1, 2, 3, 4, 5, 6, 7}},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedPartners, GetPartners(test.id))
	}
}

func TestUniverseGalaxyDistance(t *testing.T) {
	content := `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	type testCase struct {
		id1              int
		id2              int
		expectedDistance int
	}

	testCases := []testCase{
		{4, 8, 9},
		{0, 6, 15},
		{2, 5, 17},
		{7, 8, 5},
	}

	universe := ParseUniverse(strings.Split(strings.TrimSpace(content), "\n"))
	universe.Expand(2)

	for _, test := range testCases {
		// Make sure the distance is the same from id1->id2 as well as id2->id1.
		assert.Equal(t, test.expectedDistance, universe.GalaxyDistance(test.id1, test.id2))
		assert.Equal(t, test.expectedDistance, universe.GalaxyDistance(test.id2, test.id1))
	}
}

func TestSumGalaxyDistances(t *testing.T) {
	content := `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	universe := ParseUniverse(strings.Split(strings.TrimSpace(content), "\n"))
	universe.Expand(2)

	sumGalaxyDistances := 0

	for id := 0; id < universe.GetNumGalaxies(); id++ {
		for _, p := range GetPartners(id) {
			sumGalaxyDistances += universe.GalaxyDistance(id, p)
		}
	}

	assert.Equal(t, 374, sumGalaxyDistances)
}

func TestSumGalaxyDistancesOlder(t *testing.T) {
	content := `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	type testCase struct {
		expansion                  int
		expectedsumGalaxyDistances int
	}

	testCases := []testCase{
		{2, 374},
		{10, 1030},
		{100, 8410},
	}

	for _, test := range testCases {
		universe := ParseUniverse(strings.Split(strings.TrimSpace(content), "\n"))
		universe.Expand(test.expansion)

		sumGalaxyDistances := 0

		for id := 0; id < universe.GetNumGalaxies(); id++ {
			for _, p := range GetPartners(id) {
				sumGalaxyDistances += universe.GalaxyDistance(id, p)
			}
		}

		assert.Equal(t, test.expectedsumGalaxyDistances, sumGalaxyDistances)
	}
}
