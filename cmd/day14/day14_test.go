package day14

import (
	"strings"
	"testing"

	"github.com/d1r7y/advent_2023/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParsePlatform(t *testing.T) {
	type testCase struct {
		lines            []string
		expectedPlatform *Platform
	}

	testCases := []testCase{
		{lines: []string{
			"..#.",
			".O#O",
			"...O",
			".#.."},
			expectedPlatform: &Platform{
				Bounds: utilities.NewSize2D(4, 4),
				Columns: []Column{
					{Empty, Empty, Empty, Empty},
					{Empty, Rounded, Empty, Cube},
					{Cube, Cube, Empty, Empty},
					{Empty, Rounded, Rounded, Empty},
				},
			},
		},
		{lines: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#...."},
			expectedPlatform: &Platform{
				Bounds: utilities.NewSize2D(10, 10),
				Columns: []Column{
					{Rounded, Rounded, Empty, Rounded, Empty, Rounded, Empty, Empty, Cube, Cube},
					{Empty, Empty, Empty, Rounded, Rounded, Empty, Empty, Empty, Empty, Rounded},
					{Empty, Rounded, Empty, Empty, Empty, Cube, Rounded, Empty, Empty, Rounded},
					{Empty, Rounded, Empty, Cube, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Cube, Empty, Rounded, Empty, Empty, Empty, Empty, Empty, Empty},
					{Cube, Empty, Cube, Empty, Empty, Rounded, Cube, Empty, Cube, Cube},
					{Empty, Empty, Cube, Empty, Empty, Empty, Rounded, Empty, Cube, Empty},
					{Empty, Empty, Empty, Empty, Rounded, Cube, Empty, Rounded, Cube, Empty},
					{Empty, Empty, Empty, Empty, Cube, Empty, Empty, Empty, Empty, Empty},
					{Empty, Cube, Empty, Rounded, Empty, Cube, Rounded, Empty, Empty, Empty},
				},
			},
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedPlatform, ParsePlatform(test.lines))
	}
}

func TestPlatformDescribe(t *testing.T) {
	type testCase struct {
		lines []string
	}

	testCases := []testCase{
		{lines: []string{
			"..#.",
			".O#O",
			"...O",
			".#.."},
		},
	}

	for _, test := range testCases {
		p := ParsePlatform(test.lines)
		assert.Equal(t, strings.Join(test.lines, "\n"), p.Describe())
	}
}

func TestPlatformTiltNorth(t *testing.T) {
	type testCase struct {
		lines               []string
		expectedDescription string
	}

	testCases := []testCase{
		{lines: []string{
			"..#.",
			".O#O",
			"...O",
			".#.."},
			expectedDescription: `
.O#O
..#O
....
.#..`,
		},
		{lines: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#...."},
			expectedDescription: `
OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`,
		},
	}

	for _, test := range testCases {
		p := ParsePlatform(test.lines)
		p.TiltNorth()
		assert.Equal(t, strings.TrimSpace(test.expectedDescription), p.Describe())
	}
}

func TestPlatformTiltSouth(t *testing.T) {
	type testCase struct {
		lines               []string
		expectedDescription string
	}

	testCases := []testCase{
		{lines: []string{
			"..#.",
			".O#O",
			"...O",
			".#.."},
			expectedDescription: `
..#.
..#.
.O.O
.#.O`,
		},
		{lines: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#...."},
			expectedDescription: `
.....#....
....#....#
...O.##...
...#......
O.O....O#O
O.#..O.#.#
O....#....
OO....OO..
#OO..###..
#OO.O#...O`,
		},
	}

	for _, test := range testCases {
		p := ParsePlatform(test.lines)
		p.TiltSouth()
		assert.Equal(t, strings.TrimSpace(test.expectedDescription), p.Describe())
	}
}

func TestPlatformTiltEast(t *testing.T) {
	type testCase struct {
		lines               []string
		expectedDescription string
	}

	testCases := []testCase{
		{lines: []string{
			"..#.",
			".O#O",
			"...O",
			".#.."},
			expectedDescription: `
..#.
.O#O
...O
.#..`,
		},
		{lines: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#...."},
			expectedDescription: `
....O#....
.OOO#....#
.....##...
.OO#....OO
......OO#.
.O#...O#.#
....O#..OO
.........O
#....###..
#..OO#....`,
		},
	}

	for _, test := range testCases {
		p := ParsePlatform(test.lines)
		p.TiltEast()
		assert.Equal(t, strings.TrimSpace(test.expectedDescription), p.Describe())
	}
}

func TestPlatformTiltWest(t *testing.T) {
	type testCase struct {
		lines               []string
		expectedDescription string
	}

	testCases := []testCase{
		{lines: []string{
			"..#.",
			".O#O",
			"...O",
			".#.."},
			expectedDescription: `
..#.
O.#O
O...
.#..`,
		},
		{lines: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#...."},
			expectedDescription: `
O....#....
OOO.#....#
.....##...
OO.#OO....
OO......#.
O.#O...#.#
O....#OO..
O.........
#....###..
#OO..#....`,
		},
	}

	for _, test := range testCases {
		p := ParsePlatform(test.lines)
		p.TiltWest()
		assert.Equal(t, strings.TrimSpace(test.expectedDescription), p.Describe())
	}
}

func TestPlatformLoadTiltNorth(t *testing.T) {
	type testCase struct {
		lines        []string
		expectedLoad int
	}

	testCases := []testCase{
		{lines: []string{
			"..#.",
			".O#O",
			"...O",
			".#.."},
			expectedLoad: 11},
		{lines: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#...."},
			expectedLoad: 136},
	}

	for _, test := range testCases {
		p := ParsePlatform(test.lines)
		p.TiltNorth()
		assert.Equal(t, test.expectedLoad, p.Load())
	}
}

func TestPlatformTiltCycle1(t *testing.T) {
	type testCase struct {
		lines               []string
		expectedDescription string
	}

	testCases := []testCase{
		{lines: []string{
			"..#.",
			".O#O",
			"...O",
			".#.."},
			expectedDescription: `
..#.
..#.
...O
O#.O`,
		},
		{lines: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#...."},
			expectedDescription: `
.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`,
		},
	}

	for _, test := range testCases {
		p := ParsePlatform(test.lines)
		p.TiltCycle()
		assert.Equal(t, strings.TrimSpace(test.expectedDescription), p.Describe())
	}
}

func TestPlatformTiltCycle2(t *testing.T) {
	type testCase struct {
		lines               []string
		expectedDescription string
	}

	testCases := []testCase{
		{lines: []string{
			"..#.",
			".O#O",
			"...O",
			".#.."},
			expectedDescription: `
..#.
..#.
...O
O#.O`,
		},
		{lines: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#...."},
			expectedDescription: `
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`,
		},
	}

	for _, test := range testCases {
		p := ParsePlatform(test.lines)
		p.TiltCycle()
		p.TiltCycle()
		assert.Equal(t, strings.TrimSpace(test.expectedDescription), p.Describe())
	}
}

func TestPlatformTiltCycle3(t *testing.T) {
	type testCase struct {
		lines               []string
		expectedDescription string
	}

	testCases := []testCase{
		{lines: []string{
			"..#.",
			".O#O",
			"...O",
			".#.."},
			expectedDescription: `
..#.
..#.
...O
O#.O`,
		},
		{lines: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#...."},
			expectedDescription: `
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`,
		},
	}

	for _, test := range testCases {
		p := ParsePlatform(test.lines)
		p.TiltCycle()
		p.TiltCycle()
		p.TiltCycle()
		assert.Equal(t, strings.TrimSpace(test.expectedDescription), p.Describe())
	}
}
