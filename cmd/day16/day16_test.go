package day16

import (
	"strings"
	"testing"

	"github.com/d1r7y/advent_2023/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseGrid(t *testing.T) {
	content := `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	grid := ParseGrid(content)

	assert.Equal(t, utilities.NewSize2D(10, 10), grid.Bounds)
	assert.Equal(t, []TileRow{
		{Empty, VerticalSplitter, Empty, Empty, Empty, LeftMirror, Empty, Empty, Empty, Empty},
		{VerticalSplitter, Empty, HorizontalSplitter, Empty, LeftMirror, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, VerticalSplitter, HorizontalSplitter, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, VerticalSplitter, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, LeftMirror},
		{Empty, Empty, Empty, Empty, RightMirror, Empty, LeftMirror, LeftMirror, Empty, Empty},
		{Empty, HorizontalSplitter, Empty, HorizontalSplitter, RightMirror, Empty, Empty, VerticalSplitter, Empty, Empty},
		{Empty, VerticalSplitter, Empty, Empty, Empty, Empty, HorizontalSplitter, VerticalSplitter, Empty, LeftMirror},
		{Empty, Empty, RightMirror, RightMirror, Empty, VerticalSplitter, Empty, Empty, Empty, Empty},
	}, grid.Rows)
}

func TestGridDescribe(t *testing.T) {
	content := `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	grid := ParseGrid(content)

	assert.Equal(t, strings.TrimSpace(content), grid.Describe())
}

func TestUpdatePhotonEmpty(t *testing.T) {
	content := `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	grid := ParseGrid(content)

	type testCase struct {
		photon           Photon
		expectedContinue bool
		expectedPhoton   Photon
	}

	testCases := []testCase{
		{NewPhoton(0, 0, East), true, NewPhoton(1, 0, East)},
		{NewPhoton(0, 0, West), false, NewPhoton(0, 0, West)},
		{NewPhoton(0, 0, North), false, NewPhoton(0, 0, North)},
		{NewPhoton(0, 0, South), true, NewPhoton(0, 1, South)},

		{NewPhoton(grid.Bounds.Width-1, 0, East), false, NewPhoton(grid.Bounds.Width-1, 0, East)},
		{NewPhoton(grid.Bounds.Width-1, 0, West), true, NewPhoton(grid.Bounds.Width-2, 0, West)},
		{NewPhoton(grid.Bounds.Width-1, 0, North), false, NewPhoton(grid.Bounds.Width-1, 0, North)},
		{NewPhoton(grid.Bounds.Width-1, 0, South), true, NewPhoton(grid.Bounds.Width-1, 1, South)},

		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, East), false, NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, East)},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, West), true, NewPhoton(grid.Bounds.Width-2, grid.Bounds.Height-1, West)},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, North), true, NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-2, North)},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, South), false, NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, South)},

		{NewPhoton(0, grid.Bounds.Height-1, East), true, NewPhoton(1, grid.Bounds.Height-1, East)},
		{NewPhoton(0, grid.Bounds.Height-1, West), false, NewPhoton(0, grid.Bounds.Height-1, West)},
		{NewPhoton(0, grid.Bounds.Height-1, North), true, NewPhoton(0, grid.Bounds.Height-2, North)},
		{NewPhoton(0, grid.Bounds.Height-1, South), false, NewPhoton(0, grid.Bounds.Height-1, South)},
	}

	for _, test := range testCases {
		cont, newPhoton := grid.UpdatePhotonEmpty(test.photon)

		assert.Equal(t, test.expectedContinue, cont)
		assert.Equal(t, test.expectedPhoton, newPhoton)
	}
}

func TestUpdatePhotonMirror(t *testing.T) {
	content := `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	grid := ParseGrid(content)

	type testCase struct {
		photon           Photon
		tile             Tile
		expectedContinue bool
		expectedPhoton   Photon
	}

	testCases := []testCase{
		{NewPhoton(0, 0, East), LeftMirror, true, NewPhoton(0, 1, South)},
		{NewPhoton(0, 0, East), RightMirror, false, NewPhoton(0, 0, East)},
		{NewPhoton(1, 0, West), LeftMirror, false, NewPhoton(1, 0, West)},
		{NewPhoton(1, 0, West), RightMirror, true, NewPhoton(1, 1, South)},
		{NewPhoton(1, 0, North), LeftMirror, true, NewPhoton(0, 0, West)},
		{NewPhoton(1, 0, North), RightMirror, true, NewPhoton(2, 0, East)},
		{NewPhoton(1, 0, South), LeftMirror, true, NewPhoton(2, 0, East)},
		{NewPhoton(1, 0, South), RightMirror, true, NewPhoton(0, 0, West)},

		{NewPhoton(grid.Bounds.Width-1, 0, East), LeftMirror, true, NewPhoton(grid.Bounds.Width-1, 1, South)},
		{NewPhoton(grid.Bounds.Width-1, 0, East), RightMirror, false, NewPhoton(grid.Bounds.Width-1, 0, East)},
		{NewPhoton(grid.Bounds.Width-1, 0, West), LeftMirror, false, NewPhoton(grid.Bounds.Width-1, 0, West)},
		{NewPhoton(grid.Bounds.Width-1, 0, West), RightMirror, true, NewPhoton(grid.Bounds.Width-1, 1, South)},
		{NewPhoton(grid.Bounds.Width-1, 0, North), LeftMirror, true, NewPhoton(grid.Bounds.Width-2, 0, West)},
		{NewPhoton(grid.Bounds.Width-1, 0, North), RightMirror, false, NewPhoton(grid.Bounds.Width-1, 0, North)},
		{NewPhoton(grid.Bounds.Width-1, 0, South), LeftMirror, false, NewPhoton(grid.Bounds.Width-1, 0, South)},
		{NewPhoton(grid.Bounds.Width-1, 0, South), RightMirror, true, NewPhoton(grid.Bounds.Width-2, 0, West)},

		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, East), LeftMirror, false, NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, East)},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, East), RightMirror, true, NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-2, North)},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, West), LeftMirror, true, NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-2, North)},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, West), RightMirror, false, NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, West)},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, North), LeftMirror, true, NewPhoton(grid.Bounds.Width-2, grid.Bounds.Height-1, West)},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, North), RightMirror, false, NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, North)},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, South), LeftMirror, false, NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, South)},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, South), RightMirror, true, NewPhoton(grid.Bounds.Width-2, grid.Bounds.Height-1, West)},

		{NewPhoton(0, grid.Bounds.Height-1, East), LeftMirror, false, NewPhoton(0, grid.Bounds.Height-1, East)},
		{NewPhoton(0, grid.Bounds.Height-1, East), RightMirror, true, NewPhoton(0, grid.Bounds.Height-2, North)},
		{NewPhoton(0, grid.Bounds.Height-1, West), LeftMirror, true, NewPhoton(0, grid.Bounds.Height-2, North)},
		{NewPhoton(0, grid.Bounds.Height-1, West), RightMirror, false, NewPhoton(0, grid.Bounds.Height-1, West)},
		{NewPhoton(0, grid.Bounds.Height-1, North), LeftMirror, false, NewPhoton(0, grid.Bounds.Height-1, North)},
		{NewPhoton(0, grid.Bounds.Height-1, North), RightMirror, true, NewPhoton(1, grid.Bounds.Height-1, East)},
		{NewPhoton(0, grid.Bounds.Height-1, South), LeftMirror, true, NewPhoton(1, grid.Bounds.Height-1, East)},
		{NewPhoton(0, grid.Bounds.Height-1, South), RightMirror, false, NewPhoton(0, grid.Bounds.Height-1, South)},
	}

	for _, test := range testCases {
		cont, newPhoton := grid.UpdatePhotonMirror(test.photon, test.tile)

		assert.Equal(t, test.expectedContinue, cont)
		assert.Equal(t, test.expectedPhoton, newPhoton)
	}
}

func TestUpdatePhotonSplitter(t *testing.T) {
	content := `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	grid := ParseGrid(content)

	type testCase struct {
		photon          Photon
		tile            Tile
		expectedPhotons []Photon
	}

	testCases := []testCase{
		{NewPhoton(0, 0, East), VerticalSplitter, []Photon{NewPhoton(0, 1, South)}},
		{NewPhoton(0, 0, East), HorizontalSplitter, []Photon{NewPhoton(1, 0, East)}},
		{NewPhoton(1, 0, West), VerticalSplitter, []Photon{NewPhoton(1, 1, South)}},
		{NewPhoton(1, 0, West), HorizontalSplitter, []Photon{NewPhoton(0, 0, West)}},
		{NewPhoton(1, 0, North), VerticalSplitter, []Photon{}},
		{NewPhoton(1, 0, North), HorizontalSplitter, []Photon{NewPhoton(0, 0, West), NewPhoton(2, 0, East)}},
		{NewPhoton(1, 0, South), VerticalSplitter, []Photon{NewPhoton(1, 1, South)}},
		{NewPhoton(1, 0, South), HorizontalSplitter, []Photon{NewPhoton(0, 0, West), NewPhoton(2, 0, East)}},

		{NewPhoton(grid.Bounds.Width-1, 0, East), VerticalSplitter, []Photon{NewPhoton(grid.Bounds.Width-1, 1, South)}},
		{NewPhoton(grid.Bounds.Width-1, 0, East), HorizontalSplitter, []Photon{}},
		{NewPhoton(grid.Bounds.Width-1, 0, West), VerticalSplitter, []Photon{NewPhoton(grid.Bounds.Width-1, 1, South)}},
		{NewPhoton(grid.Bounds.Width-1, 0, West), HorizontalSplitter, []Photon{NewPhoton(grid.Bounds.Width-2, 0, West)}},
		{NewPhoton(grid.Bounds.Width-1, 0, North), VerticalSplitter, []Photon{}},
		{NewPhoton(grid.Bounds.Width-1, 0, North), HorizontalSplitter, []Photon{NewPhoton(grid.Bounds.Width-2, 0, West)}},
		{NewPhoton(grid.Bounds.Width-1, 0, South), VerticalSplitter, []Photon{NewPhoton(grid.Bounds.Width-1, 1, South)}},
		{NewPhoton(grid.Bounds.Width-1, 0, South), HorizontalSplitter, []Photon{NewPhoton(grid.Bounds.Width-2, 0, West)}},

		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, East), VerticalSplitter, []Photon{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-2, North)}},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, East), HorizontalSplitter, []Photon{}},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, West), VerticalSplitter, []Photon{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-2, North)}},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, West), HorizontalSplitter, []Photon{NewPhoton(grid.Bounds.Width-2, grid.Bounds.Height-1, West)}},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, North), VerticalSplitter, []Photon{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-2, North)}},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, North), HorizontalSplitter, []Photon{NewPhoton(grid.Bounds.Width-2, grid.Bounds.Height-1, West)}},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, South), VerticalSplitter, []Photon{}},
		{NewPhoton(grid.Bounds.Width-1, grid.Bounds.Height-1, South), HorizontalSplitter, []Photon{NewPhoton(grid.Bounds.Width-2, grid.Bounds.Height-1, West)}},

		{NewPhoton(0, grid.Bounds.Height-1, East), VerticalSplitter, []Photon{NewPhoton(0, grid.Bounds.Height-2, North)}},
		{NewPhoton(0, grid.Bounds.Height-1, East), HorizontalSplitter, []Photon{NewPhoton(1, grid.Bounds.Height-1, East)}},
		{NewPhoton(0, grid.Bounds.Height-1, West), VerticalSplitter, []Photon{NewPhoton(0, grid.Bounds.Height-2, North)}},
		{NewPhoton(0, grid.Bounds.Height-1, West), HorizontalSplitter, []Photon{}},
		{NewPhoton(0, grid.Bounds.Height-1, North), VerticalSplitter, []Photon{NewPhoton(0, grid.Bounds.Height-2, North)}},
		{NewPhoton(0, grid.Bounds.Height-1, North), HorizontalSplitter, []Photon{NewPhoton(1, grid.Bounds.Height-1, East)}},
		{NewPhoton(0, grid.Bounds.Height-1, South), VerticalSplitter, []Photon{}},
		{NewPhoton(0, grid.Bounds.Height-1, South), HorizontalSplitter, []Photon{NewPhoton(1, grid.Bounds.Height-1, East)}},
	}

	for _, test := range testCases {
		photons := grid.UpdatePhotonSplitter(test.photon, test.tile)

		assert.Equal(t, test.expectedPhotons, photons)
	}
}

func TestGetEnergizedTiles(t *testing.T) {
	content := `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	expectedDescription := `
######....
.#...#....
.#...#####
.#...##...
.#...##...
.#...##...
.#..####..
########..
.#######..
.#...#.#..`

	grid := ParseGrid(content)

	initialPhoton := Photon{
		Position:  utilities.NewPoint2D(0, 0),
		Direction: East,
	}

	assert.Equal(t, strings.TrimSpace(expectedDescription), GetEnergizedTiles(grid, initialPhoton))
}

func TestGetEnergizedTilesCount(t *testing.T) {
	content := `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	grid := ParseGrid(content)

	initialPhoton := Photon{
		Position:  utilities.NewPoint2D(0, 0),
		Direction: East,
	}

	assert.Equal(t, 46, GetEnergizedTilesCount(grid, initialPhoton))
}

func TestGetMaxEnergizedTilesCount(t *testing.T) {
	content := `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	grid := ParseGrid(content)

	assert.Equal(t, 51, GetMaxEnergizedTilesCount(grid))
}
