package day10

import (
	"log"
	"strings"
	"testing"

	"github.com/d1r7y/advent_2023/utilities"
	"github.com/stretchr/testify/assert"
)

func TestNewGrid(t *testing.T) {
	type testCase struct {
		content      []string
		expectedGrid *Grid
	}

	testCases := []testCase{
		{[]string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		},
			&Grid{StartPosition: utilities.NewPoint2D(1, 1),
				Bounds: utilities.NewSize2D(5, 5),
				Rows: []Row{
					{Ground, Ground, Ground, Ground, Ground},
					{Ground, BendSouthEastPipe, HorizontalPipe, BendSouthWestPipe, Ground},
					{Ground, VerticalPipe, Ground, VerticalPipe, Ground},
					{Ground, BendNorthEastPipe, HorizontalPipe, BendNorthWestPipe, Ground},
					{Ground, Ground, Ground, Ground, Ground},
				}},
		},
		{[]string{
			"-L|F7",
			"7S-7|",
			"L|7||",
			"-L-J|",
			"L|-JF",
		},
			&Grid{StartPosition: utilities.NewPoint2D(1, 1),
				Bounds: utilities.NewSize2D(5, 5),
				Rows: []Row{
					{HorizontalPipe, BendNorthEastPipe, VerticalPipe, BendSouthEastPipe, BendSouthWestPipe},
					{BendSouthWestPipe, BendSouthEastPipe, HorizontalPipe, BendSouthWestPipe, VerticalPipe},
					{BendNorthEastPipe, VerticalPipe, BendSouthWestPipe, VerticalPipe, VerticalPipe},
					{HorizontalPipe, BendNorthEastPipe, HorizontalPipe, BendNorthWestPipe, VerticalPipe},
					{BendNorthEastPipe, VerticalPipe, HorizontalPipe, BendNorthWestPipe, BendSouthEastPipe},
				}},
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedGrid, ParseGrid(test.content))
	}
}

func TestGridDescribe(t *testing.T) {
	type testCase struct {
		content []string
	}

	testCases := []testCase{
		{[]string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		}},
		{[]string{
			"-L|F7",
			"7S-7|",
			"L|7||",
			"-L-J|",
			"L|-JF",
		}},
	}

	for _, test := range testCases {
		grid := ParseGrid(test.content)
		assert.Equal(t, test.content, strings.Split(grid.Describe(), "\n"))
	}
}

func TestGridGetNeighborTile(t *testing.T) {
	type testCase struct {
		content      []string
		position     utilities.Point2D
		direction    Direction
		expectedTile Tile
	}

	testCases := []testCase{
		{[]string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		}, utilities.NewPoint2D(0, 0), East, Ground},
		{[]string{
			"-L|F7",
			"7S-7|",
			"L|7||",
			"-L-J|",
			"L|-JF",
		}, utilities.NewPoint2D(0, 1), East, BendSouthEastPipe},
		{[]string{
			"-L|F7",
			"7S-7|",
			"L|7||",
			"-L-J|",
			"L|-JF",
		}, utilities.NewPoint2D(3, 0), East, BendSouthWestPipe},
		{[]string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		}, utilities.NewPoint2D(2, 2), East, VerticalPipe},
		{[]string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		}, utilities.NewPoint2D(2, 2), North, HorizontalPipe},
		{[]string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		}, utilities.NewPoint2D(2, 2), South, HorizontalPipe},
		{[]string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		}, utilities.NewPoint2D(4, 3), West, BendNorthWestPipe},
	}

	for _, test := range testCases {
		grid := ParseGrid(test.content)
		assert.Equal(t, test.expectedTile, grid.GetNeighborTile(test.position, test.direction))
	}
}

func TestCanTileExit(t *testing.T) {
	type testCase struct {
		tile         Tile
		direction    Direction
		expectedExit bool
	}

	testCases := []testCase{
		{VerticalPipe, North, true},
		{VerticalPipe, South, true},
		{VerticalPipe, East, false},
		{VerticalPipe, West, false},

		{HorizontalPipe, North, false},
		{HorizontalPipe, South, false},
		{HorizontalPipe, East, true},
		{HorizontalPipe, West, true},

		{BendNorthEastPipe, North, true},
		{BendNorthEastPipe, South, false},
		{BendNorthEastPipe, East, true},
		{BendNorthEastPipe, West, false},

		{BendNorthWestPipe, North, true},
		{BendNorthWestPipe, South, false},
		{BendNorthWestPipe, East, false},
		{BendNorthWestPipe, West, true},

		{BendSouthEastPipe, North, false},
		{BendSouthEastPipe, South, true},
		{BendSouthEastPipe, East, true},
		{BendSouthEastPipe, West, false},

		{BendSouthWestPipe, North, false},
		{BendSouthWestPipe, South, true},
		{BendSouthWestPipe, East, false},
		{BendSouthWestPipe, West, true},

		{Ground, North, false},
		{Ground, South, false},
		{Ground, East, false},
		{Ground, West, false},

		{Start, North, false},
		{Start, South, false},
		{Start, East, false},
		{Start, West, false},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedExit, CanTileExit(test.tile, test.direction))
	}
}

func TestIsTilePip(t *testing.T) {
	type testCase struct {
		tile         Tile
		expectedPipe bool
	}

	testCases := []testCase{
		{VerticalPipe, true},
		{HorizontalPipe, true},
		{BendNorthEastPipe, true},
		{BendNorthWestPipe, true},
		{BendSouthEastPipe, true},
		{BendSouthWestPipe, true},
		{Ground, false},
		{Start, false},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedPipe, IsTilePipe(test.tile))
	}
}

func TestCanConnect(t *testing.T) {
	type testCase struct {
		tile1           Tile
		direction       Direction
		tile2           Tile
		expectedConnect bool
	}

	testCases := []testCase{
		{VerticalPipe, North, VerticalPipe, true},
		{VerticalPipe, North, HorizontalPipe, false},
		{VerticalPipe, North, BendNorthEastPipe, false},
		{VerticalPipe, North, BendNorthWestPipe, false},
		{VerticalPipe, North, BendSouthEastPipe, true},
		{VerticalPipe, North, BendSouthWestPipe, true},

		{VerticalPipe, South, VerticalPipe, true},
		{VerticalPipe, South, HorizontalPipe, false},
		{VerticalPipe, South, BendNorthEastPipe, true},
		{VerticalPipe, South, BendNorthWestPipe, true},
		{VerticalPipe, South, BendSouthEastPipe, false},
		{VerticalPipe, South, BendSouthWestPipe, false},

		{VerticalPipe, East, VerticalPipe, false},
		{VerticalPipe, East, HorizontalPipe, false},
		{VerticalPipe, East, BendNorthEastPipe, false},
		{VerticalPipe, East, BendNorthWestPipe, false},
		{VerticalPipe, East, BendSouthEastPipe, false},
		{VerticalPipe, East, BendSouthWestPipe, false},

		{VerticalPipe, West, VerticalPipe, false},
		{VerticalPipe, West, HorizontalPipe, false},
		{VerticalPipe, West, BendNorthEastPipe, false},
		{VerticalPipe, West, BendNorthWestPipe, false},
		{VerticalPipe, West, BendSouthEastPipe, false},
		{VerticalPipe, West, BendSouthWestPipe, false},

		{HorizontalPipe, North, VerticalPipe, false},
		{HorizontalPipe, North, HorizontalPipe, false},
		{HorizontalPipe, North, BendNorthEastPipe, false},
		{HorizontalPipe, North, BendNorthWestPipe, false},
		{HorizontalPipe, North, BendSouthEastPipe, false},
		{HorizontalPipe, North, BendSouthWestPipe, false},

		{HorizontalPipe, South, VerticalPipe, false},
		{HorizontalPipe, South, HorizontalPipe, false},
		{HorizontalPipe, South, BendNorthEastPipe, false},
		{HorizontalPipe, South, BendNorthWestPipe, false},
		{HorizontalPipe, South, BendSouthEastPipe, false},
		{HorizontalPipe, South, BendSouthWestPipe, false},

		{HorizontalPipe, East, VerticalPipe, false},
		{HorizontalPipe, East, HorizontalPipe, true},
		{HorizontalPipe, East, BendNorthEastPipe, false},
		{HorizontalPipe, East, BendNorthWestPipe, true},
		{HorizontalPipe, East, BendSouthEastPipe, false},
		{HorizontalPipe, East, BendSouthWestPipe, true},

		{HorizontalPipe, West, VerticalPipe, false},
		{HorizontalPipe, West, HorizontalPipe, true},
		{HorizontalPipe, West, BendNorthEastPipe, true},
		{HorizontalPipe, West, BendNorthWestPipe, false},
		{HorizontalPipe, West, BendSouthEastPipe, true},
		{HorizontalPipe, West, BendSouthWestPipe, false},

		{BendNorthEastPipe, North, VerticalPipe, true},
		{BendNorthEastPipe, North, HorizontalPipe, false},
		{BendNorthEastPipe, North, BendNorthEastPipe, false},
		{BendNorthEastPipe, North, BendNorthWestPipe, false},
		{BendNorthEastPipe, North, BendSouthEastPipe, true},
		{BendNorthEastPipe, North, BendSouthWestPipe, true},

		{BendNorthEastPipe, South, VerticalPipe, false},
		{BendNorthEastPipe, South, HorizontalPipe, false},
		{BendNorthEastPipe, South, BendNorthEastPipe, false},
		{BendNorthEastPipe, South, BendNorthWestPipe, false},
		{BendNorthEastPipe, South, BendSouthEastPipe, false},
		{BendNorthEastPipe, South, BendSouthWestPipe, false},

		{BendNorthEastPipe, East, VerticalPipe, false},
		{BendNorthEastPipe, East, HorizontalPipe, true},
		{BendNorthEastPipe, East, BendNorthEastPipe, false},
		{BendNorthEastPipe, East, BendNorthWestPipe, true},
		{BendNorthEastPipe, East, BendSouthEastPipe, false},
		{BendNorthEastPipe, East, BendSouthWestPipe, true},

		{BendNorthEastPipe, West, VerticalPipe, false},
		{BendNorthEastPipe, West, HorizontalPipe, false},
		{BendNorthEastPipe, West, BendNorthEastPipe, false},
		{BendNorthEastPipe, West, BendNorthWestPipe, false},
		{BendNorthEastPipe, West, BendSouthEastPipe, false},
		{BendNorthEastPipe, West, BendSouthWestPipe, false},

		{BendNorthWestPipe, North, VerticalPipe, true},
		{BendNorthWestPipe, North, HorizontalPipe, false},
		{BendNorthWestPipe, North, BendNorthEastPipe, false},
		{BendNorthWestPipe, North, BendNorthWestPipe, false},
		{BendNorthWestPipe, North, BendSouthEastPipe, true},
		{BendNorthWestPipe, North, BendSouthWestPipe, true},

		{BendNorthWestPipe, South, VerticalPipe, false},
		{BendNorthWestPipe, South, HorizontalPipe, false},
		{BendNorthWestPipe, South, BendNorthEastPipe, false},
		{BendNorthWestPipe, South, BendNorthWestPipe, false},
		{BendNorthWestPipe, South, BendSouthEastPipe, false},
		{BendNorthWestPipe, South, BendSouthWestPipe, false},

		{BendNorthWestPipe, East, VerticalPipe, false},
		{BendNorthWestPipe, East, HorizontalPipe, false},
		{BendNorthWestPipe, East, BendNorthEastPipe, false},
		{BendNorthWestPipe, East, BendNorthWestPipe, false},
		{BendNorthWestPipe, East, BendSouthEastPipe, false},
		{BendNorthWestPipe, East, BendSouthWestPipe, false},

		{BendNorthWestPipe, West, VerticalPipe, false},
		{BendNorthWestPipe, West, HorizontalPipe, true},
		{BendNorthWestPipe, West, BendNorthEastPipe, true},
		{BendNorthWestPipe, West, BendNorthWestPipe, false},
		{BendNorthWestPipe, West, BendSouthEastPipe, true},
		{BendNorthWestPipe, West, BendSouthWestPipe, false},

		{BendSouthEastPipe, North, VerticalPipe, false},
		{BendSouthEastPipe, North, HorizontalPipe, false},
		{BendSouthEastPipe, North, BendNorthEastPipe, false},
		{BendSouthEastPipe, North, BendNorthWestPipe, false},
		{BendSouthEastPipe, North, BendSouthEastPipe, false},
		{BendSouthEastPipe, North, BendSouthWestPipe, false},

		{BendSouthEastPipe, South, VerticalPipe, true},
		{BendSouthEastPipe, South, HorizontalPipe, false},
		{BendSouthEastPipe, South, BendNorthEastPipe, true},
		{BendSouthEastPipe, South, BendNorthWestPipe, true},
		{BendSouthEastPipe, South, BendSouthEastPipe, false},
		{BendSouthEastPipe, South, BendSouthWestPipe, false},

		{BendSouthEastPipe, East, VerticalPipe, false},
		{BendSouthEastPipe, East, HorizontalPipe, true},
		{BendSouthEastPipe, East, BendNorthEastPipe, false},
		{BendSouthEastPipe, East, BendNorthWestPipe, true},
		{BendSouthEastPipe, East, BendSouthEastPipe, false},
		{BendSouthEastPipe, East, BendSouthWestPipe, true},

		{BendSouthEastPipe, West, VerticalPipe, false},
		{BendSouthEastPipe, West, HorizontalPipe, false},
		{BendSouthEastPipe, West, BendNorthEastPipe, false},
		{BendSouthEastPipe, West, BendNorthWestPipe, false},
		{BendSouthEastPipe, West, BendSouthEastPipe, false},
		{BendSouthEastPipe, West, BendSouthWestPipe, false},

		{BendSouthWestPipe, North, VerticalPipe, false},
		{BendSouthWestPipe, North, HorizontalPipe, false},
		{BendSouthWestPipe, North, BendNorthEastPipe, false},
		{BendSouthWestPipe, North, BendNorthWestPipe, false},
		{BendSouthWestPipe, North, BendSouthEastPipe, false},
		{BendSouthWestPipe, North, BendSouthWestPipe, false},

		{BendSouthWestPipe, South, VerticalPipe, true},
		{BendSouthWestPipe, South, HorizontalPipe, false},
		{BendSouthWestPipe, South, BendNorthEastPipe, true},
		{BendSouthWestPipe, South, BendNorthWestPipe, true},
		{BendSouthWestPipe, South, BendSouthEastPipe, false},
		{BendSouthWestPipe, South, BendSouthWestPipe, false},

		{BendSouthWestPipe, East, VerticalPipe, false},
		{BendSouthWestPipe, East, HorizontalPipe, false},
		{BendSouthWestPipe, East, BendNorthEastPipe, false},
		{BendSouthWestPipe, East, BendNorthWestPipe, false},
		{BendSouthWestPipe, East, BendSouthEastPipe, false},
		{BendSouthWestPipe, East, BendSouthWestPipe, false},

		{BendSouthWestPipe, West, VerticalPipe, false},
		{BendSouthWestPipe, West, HorizontalPipe, true},
		{BendSouthWestPipe, West, BendNorthEastPipe, true},
		{BendSouthWestPipe, West, BendNorthWestPipe, false},
		{BendSouthWestPipe, West, BendSouthEastPipe, true},
		{BendSouthWestPipe, West, BendSouthWestPipe, false},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedConnect, CanConnect(test.tile1, test.direction, test.tile2))
	}
}

func TestGridDistance(t *testing.T) {
	type testCase struct {
		content          []string
		expectedDistance int
	}

	testCases := []testCase{
		{[]string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		}, 4},
		{[]string{
			"..F7.",
			".FJ|.",
			"SJ.L7",
			"|F--J",
			"LJ...",
		}, 8},
	}

	for _, test := range testCases {
		grid := ParseGrid(test.content)

		distance := 0

		grid.TraverseLoop(func(p utilities.Point2D, d Direction, t Tile) bool {
			distance++
			return true
		})

		assert.Equal(t, test.expectedDistance, distance/2)
	}
}

func TestGridArea(t *testing.T) {
	type testCase struct {
		content      []string
		expectedArea int
	}

	testCases := []testCase{
		{[]string{
			"...........",
			".S-------7.",
			".|F-----7|.",
			".||.....||.",
			".||.....||.",
			".|L-7.F-J|.",
			".|..|.|..|.",
			".L--J.L--J.",
			"...........",
		}, 4},
		{[]string{
			"..........",
			".S------7.",
			".|F----7|.",
			".||....||.",
			".||....||.",
			".|L-7F-J|.",
			".|..||..|.",
			".L--JL--J.",
			"..........",
		}, 4},
		{[]string{
			".F----7F7F7F7F-7....",
			".|F--7||||||||FJ....",
			".||.FJ||||||||L7....",
			"FJL7L7LJLJ||LJ.L-7..",
			"L--J.L7...LJS7F-7L7.",
			"....F-J..F7FJ|L7L7L7",
			"....L7.F7||L7|.L7L7|",
			".....|FJLJ|FJ|F7|.LJ",
			"....FJL-7.||.||||...",
			"....L---J.LJ.LJLJ...",
		}, 8},
	}

	for _, test := range testCases {
		grid := ParseGrid(test.content)

		visited := NewDistances(grid.Bounds)

		areas := NewDistances(grid.Bounds)

		vertices := make([]utilities.Point2D, 0)

		grid.TraverseLoop(func(p utilities.Point2D, d Direction, t Tile) bool {
			vertices = append(vertices, p)
			visited.SetDistance(p, 10)
			return true
		})

		log.Println("\n", visited.Describe())

		vertices = append(vertices, grid.StartPosition)
		area := 0

		for y := 0; y < visited.Bounds.Height; y++ {
			for x := 0; x < visited.Bounds.Width; x++ {
				d := visited.GetDistance(utilities.NewPoint2D(x, y))
				if d < 0 {
					if utilities.PointInPolyCrossing(utilities.NewPoint2D(x, y), vertices) {
						areas.SetDistance(utilities.NewPoint2D(x, y), 10)
						area++
					} else {
						areas.SetDistance(utilities.NewPoint2D(x, y), 0)
					}
				}
			}
		}

		log.Println("\n", areas.Describe())
		assert.Equal(t, test.expectedArea, area)
	}
}
