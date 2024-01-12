package day13

import (
	"strings"
	"testing"

	"github.com/d1r7y/advent_2023/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseLandscape(t *testing.T) {
	type testCase struct {
		content           string
		expectedLandscape *Landscape
	}

	testCases := []testCase{
		{
			content: `
				#.##..##.
				..#.##.#.
				##......#
				##......#
				..#.##.#.
				..##..##.
				#.#.##.#.`,
			expectedLandscape: &Landscape{
				Bounds: utilities.Size2D{Width: 9, Height: 7},
				Ground: []TerrainRow{
					{Rock, Ash, Rock, Rock, Ash, Ash, Rock, Rock, Ash},
					{Ash, Ash, Rock, Ash, Rock, Rock, Ash, Rock, Ash},
					{Rock, Rock, Ash, Ash, Ash, Ash, Ash, Ash, Rock},
					{Rock, Rock, Ash, Ash, Ash, Ash, Ash, Ash, Rock},
					{Ash, Ash, Rock, Ash, Rock, Rock, Ash, Rock, Ash},
					{Ash, Ash, Rock, Rock, Ash, Ash, Rock, Rock, Ash},
					{Rock, Ash, Rock, Ash, Rock, Rock, Ash, Rock, Ash},
				}},
		},
		{
			content: `
				#...##..#
				#....#..#
				..##..###
				#####.##.
				#####.##.
				..##..###
				#....#..#`,
			expectedLandscape: &Landscape{
				Bounds: utilities.Size2D{Width: 9, Height: 7},
				Ground: []TerrainRow{
					{Rock, Ash, Ash, Ash, Rock, Rock, Ash, Ash, Rock},
					{Rock, Ash, Ash, Ash, Ash, Rock, Ash, Ash, Rock},
					{Ash, Ash, Rock, Rock, Ash, Ash, Rock, Rock, Rock},
					{Rock, Rock, Rock, Rock, Rock, Ash, Rock, Rock, Ash},
					{Rock, Rock, Rock, Rock, Rock, Ash, Rock, Rock, Ash},
					{Ash, Ash, Rock, Rock, Ash, Ash, Rock, Rock, Rock},
					{Rock, Ash, Ash, Ash, Ash, Rock, Ash, Ash, Rock},
				}},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expectedLandscape, ParseLandscape(strings.Split(strings.TrimSpace(testCase.content), "\n")))
	}
}

func TestLandscapeGetReflection(t *testing.T) {
	type testCase struct {
		content            string
		expectedReflection Reflection
	}

	testCases := []testCase{
		{
			content: `
				#.##..##.
				..#.##.#.
				##......#
				##......#
				..#.##.#.
				..##..##.
				#.#.##.#.`,
			expectedReflection: Reflection{Axis: Vertical, Position: 4},
		},
		{
			content: `
				#...##..#
				#....#..#
				..##..###
				#####.##.
				#####.##.
				..##..###
				#....#..#`,
			expectedReflection: Reflection{Axis: Horizontal, Position: 3},
		},
	}

	for _, testCase := range testCases {
		l := ParseLandscape(strings.Split(strings.TrimSpace(testCase.content), "\n"))
		assert.Equal(t, testCase.expectedReflection, l.GetReflection())
	}
}

func TestLandscapeGetReflectionSmudged(t *testing.T) {
	type testCase struct {
		content            string
		expectedReflection Reflection
	}

	testCases := []testCase{
		{
			content: `
				#.##..##.
				..#.##.#.
				##......#
				##......#
				..#.##.#.
				..##..##.
				#.#.##.#.`,
			expectedReflection: Reflection{Axis: Horizontal, Position: 2},
		},
		{
			content: `
				#...##..#
				#....#..#
				..##..###
				#####.##.
				#####.##.
				..##..###
				#....#..#`,
			expectedReflection: Reflection{Axis: Horizontal, Position: 0},
		},
		{
			content: `
				.##..#..#.###
				.##.#...#..##
				..#.######.#.
				.......#.....
				.......#.....
				..#.######.#.
				.##.#...#..##
				.##..#..#.###
				####.#####.#.
				####.#####.#.
				.##.....#.###`,
			expectedReflection: Reflection{Axis: Horizontal, Position: 8},
		},
		{
			content: `
				..##.##.##...
				##.#......###
				..########...
				#####..######
				..#......#...
				.####..####..
				#....##....##
				....#..#.....
				#...####...##
				##.#....#.###
				#...#..#...##
				#####..######
				##..#..#..###
				.#...##...#..
				..#.#..#.#...`,
			expectedReflection: Reflection{Axis: Vertical, Position: 5},
		},
		{
			content: `
				.###.#..#.###..
				#.#.#.##.#.#.##
				..#.######.#...
				..#...##...#...
				#.##.####..#.##
				##.########.###
				..#..####..#...
				#.##......##.##
				##.#..##..#.###`,
			expectedReflection: Reflection{Axis: Vertical, Position: 6},
		},
		{
			content: `
				.##.#....####..
				###....###...##
				.....###..#####
				.....###..#####
				###....###...##
				.##.#....####..
				#.#..#.#...#.##
				##..#.##.#.##..
				.#.#.#.#..#.###
				.#.###.#..#..##
				.#..#.....#.###
				..###.###..##..
				...#.#..#..####
				#.####.##..#.##
				..##.##.###.###
				#.#.##.#..#.###
				#..###...###.#.`,
			expectedReflection: Reflection{Axis: Vertical, Position: 13},
		},
	}

	for _, testCase := range testCases {
		l := ParseLandscape(strings.Split(strings.TrimSpace(testCase.content), "\n"))
		excludedReflection := l.GetReflection()
		assert.Equal(t, testCase.expectedReflection, l.GetReflectionSmudged(excludedReflection), testCase.content)
	}
}

func TestSummarizeNotes(t *testing.T) {
	terrains := []string{
		`
		#.##..##.
		..#.##.#.
		##......#
		##......#
		..#.##.#.
		..##..##.
		#.#.##.#.`,
		`
		#...##..#
		#....#..#
		..##..###
		#####.##.
		#####.##.
		..##..###
		#....#..#`,
	}

	noteSummary := 0

	for _, terrain := range terrains {
		l := ParseLandscape(strings.Split(strings.TrimSpace(terrain), "\n"))
		reflection := l.GetReflection()
		if reflection.Axis == Vertical {
			noteSummary += reflection.Position + 1
		} else if reflection.Axis == Horizontal {
			noteSummary += 100 * (reflection.Position + 1)
		}
	}

	assert.Equal(t, 405, noteSummary)
}

func TestSummarizeNotesSmudged(t *testing.T) {
	terrains := []string{
		`
		#.##..##.
		..#.##.#.
		##......#
		##......#
		..#.##.#.
		..##..##.
		#.#.##.#.`,
		`
		#...##..#
		#....#..#
		..##..###
		#####.##.
		#####.##.
		..##..###
		#....#..#`,
	}

	noteSummary := 0

	for _, terrain := range terrains {
		l := ParseLandscape(strings.Split(strings.TrimSpace(terrain), "\n"))
		excludedReflection := l.GetReflection()
		reflection := l.GetReflectionSmudged(excludedReflection)
		if reflection.Axis == Vertical {
			noteSummary += reflection.Position + 1
		} else if reflection.Axis == Horizontal {
			noteSummary += 100 * (reflection.Position + 1)
		}
	}

	assert.Equal(t, 400, noteSummary)
}
