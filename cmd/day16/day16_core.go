package day16

import (
	"log"
	"math"
	"strings"

	"github.com/d1r7y/advent_2023/utilities"
)

type Direction byte

const (
	North Direction = 1 << 0
	East            = 1 << 1
	South           = 1 << 2
	West            = 1 << 3
)

type Photon struct {
	Direction Direction
	Position  utilities.Point2D
}

func NewPhoton(x, y int, direction Direction) Photon {
	return Photon{Position: utilities.NewPoint2D(x, y), Direction: direction}
}

type Tile byte

func (t Tile) Describe() string {
	switch t {
	case Empty:
		return "."
	case LeftMirror:
		return "\\"
	case RightMirror:
		return "/"
	case VerticalSplitter:
		return "|"
	case HorizontalSplitter:
		return "-"
	}

	log.Panicf("unknown tile: %d\n", t)
	return "*"
}

const (
	Empty Tile = iota
	LeftMirror
	RightMirror
	VerticalSplitter
	HorizontalSplitter
)

type TileRow []Tile

type VisitedTile Direction

func (vt VisitedTile) Describe() string {
	description := ""

	if vt&VisitedTile(North) != 0 {
		description += "north "
	}
	if vt&VisitedTile(East) != 0 {
		description += "east "
	}
	if vt&VisitedTile(South) != 0 {
		description += "south "
	}
	if vt&VisitedTile(West) != 0 {
		description += "north "
	}

	return description
}

type VisitedTileRow []VisitedTile

type Grid struct {
	Photons     *utilities.FIFO[Photon]
	Bounds      utilities.Size2D
	Rows        []TileRow
	VisitedRows []VisitedTileRow
}

func (g *Grid) Describe() string {
	description := ""
	for _, r := range g.Rows {
		str := ""

		for _, t := range r {
			str += t.Describe()
		}
		if description != "" {
			description += "\n"
		}

		description += str
	}

	return description
}

func (g *Grid) validatePosition(position utilities.Point2D) bool {
	if position.X >= g.Bounds.Width {
		return false
	}
	if position.Y >= g.Bounds.Height {
		return false
	}

	return true
}

func (g *Grid) GetTile(position utilities.Point2D) Tile {
	if !g.validatePosition(position) {
		log.Panicf("invalid position %d,%d\n", position.X, position.Y)
	}

	return g.Rows[position.Y][position.X]
}

func (g *Grid) UpdateVisitedTile(position utilities.Point2D, direction Direction) {
	if !g.validatePosition(position) {
		log.Panicf("invalid position %d,%d\n", position.X, position.Y)
	}

	g.VisitedRows[position.Y][position.X] |= VisitedTile(direction)
}

func (g *Grid) HaveVisitedTile(position utilities.Point2D, direction Direction) bool {
	if !g.validatePosition(position) {
		log.Panicf("invalid position %d,%d\n", position.X, position.Y)
	}

	return g.VisitedRows[position.Y][position.X]&VisitedTile(direction) != 0
}

func (g *Grid) UpdatePhotonEmpty(p Photon) (bool, Photon) {
	if p.Direction == North {
		if p.Position.Y > 0 {
			p.Position.Y--
		} else {
			return false, p
		}
	} else if p.Direction == East {
		if p.Position.X < g.Bounds.Width-1 {
			p.Position.X++
		} else {
			return false, p
		}
	} else if p.Direction == South {
		if p.Position.Y < g.Bounds.Height-1 {
			p.Position.Y++
		} else {
			return false, p
		}

	} else if p.Direction == West {
		if p.Position.X > 0 {
			p.Position.X--
		} else {
			return false, p
		}
	}

	return true, p
}

func (g *Grid) UpdatePhotonMirror(p Photon, mirrorTile Tile) (bool, Photon) {
	if (mirrorTile == RightMirror && p.Direction == North) || (mirrorTile == LeftMirror && p.Direction == South) {
		// Direction will be East
		if p.Position.X < g.Bounds.Width-1 {
			p.Direction = East
			p.Position.X++
		} else {
			return false, p
		}
	} else if (mirrorTile == RightMirror && p.Direction == East) || (mirrorTile == LeftMirror && p.Direction == West) {
		// Direction will be North
		if p.Position.Y > 0 {
			p.Direction = North
			p.Position.Y--
		} else {
			return false, p
		}
	} else if (mirrorTile == RightMirror && p.Direction == South) || (mirrorTile == LeftMirror && p.Direction == North) {
		// Direction will be West
		if p.Position.X > 0 {
			p.Direction = West
			p.Position.X--
		} else {
			return false, p
		}
	} else if (mirrorTile == RightMirror && p.Direction == West) || (mirrorTile == LeftMirror && p.Direction == East) {
		// Direction will be South
		if p.Position.Y < g.Bounds.Height-1 {
			p.Direction = South
			p.Position.Y++
		} else {
			return false, p
		}
	}

	return true, p
}

func (g *Grid) UpdatePhotonSplitter(p Photon, splitterTile Tile) []Photon {
	photons := make([]Photon, 0)

	// Handle no split cases
	if splitterTile == VerticalSplitter && (p.Direction == North || p.Direction == South) {
		if c, np := g.UpdatePhotonEmpty(p); c {
			photons = append(photons, np)
		}
		return photons
	}

	if splitterTile == HorizontalSplitter && (p.Direction == East || p.Direction == West) {
		if c, np := g.UpdatePhotonEmpty(p); c {
			photons = append(photons, np)
		}
		return photons
	}

	// Handle split cases
	if splitterTile == VerticalSplitter {
		p1 := p
		p2 := p

		photons := make([]Photon, 0)

		if p1.Position.Y > 0 {
			p1.Direction = North
			p1.Position.Y--
			photons = append(photons, p1)
		}

		if p2.Position.Y < g.Bounds.Height-1 {
			p2.Direction = South
			p2.Position.Y++
			photons = append(photons, p2)
		}

		return photons
	}

	if splitterTile == HorizontalSplitter {
		p1 := p
		p2 := p

		photons := make([]Photon, 0)

		if p1.Position.X > 0 {
			p1.Direction = West
			p1.Position.X--
			photons = append(photons, p1)
		}

		if p2.Position.X < g.Bounds.Width-1 {
			p2.Direction = East
			p2.Position.X++
			photons = append(photons, p2)
		}

		return photons
	}

	return photons
}

func (g *Grid) UpdatePhoton(p Photon, t Tile) []Photon {
	if t == Empty {
		if cont, p := g.UpdatePhotonEmpty(p); cont {
			return []Photon{p}
		}
	} else if t == LeftMirror || t == RightMirror {
		if cont, p := g.UpdatePhotonMirror(p, t); cont {
			return []Photon{p}
		}

	} else if t == VerticalSplitter || t == HorizontalSplitter {
		return g.UpdatePhotonSplitter(p, t)
	}

	return []Photon{}
}

func (g *Grid) StartPhoton(initialPhoton Photon, visitedTile func(position utilities.Point2D)) {
	if !g.validatePosition(initialPhoton.Position) {
		log.Panicf("invalid position %d,%d\n", initialPhoton.Position.X, initialPhoton.Position.Y)
	}

	g.Photons.Push(initialPhoton)

	for {
		if g.Photons.IsEmpty() {
			break
		}

		p := g.Photons.Pop()
		t := g.GetTile(p.Position)

		visitedTile(p.Position)

		g.UpdateVisitedTile(p.Position, p.Direction)

		for _, np := range g.UpdatePhoton(p, t) {
			if !g.HaveVisitedTile(np.Position, np.Direction) {
				g.Photons.Push(np)
			}
		}
	}
}

func (g *Grid) Reset() {
	g.Photons = &utilities.FIFO[Photon]{}

	g.VisitedRows = make([]VisitedTileRow, 0)

	for i := 0; i < g.Bounds.Height; i++ {
		g.VisitedRows = append(g.VisitedRows, make(VisitedTileRow, g.Bounds.Width))
	}
}

func ParseGrid(content string) *Grid {
	grid := &Grid{}

	grid.Photons = &utilities.FIFO[Photon]{}
	grid.Rows = make([]TileRow, 0)
	grid.VisitedRows = make([]VisitedTileRow, 0)

	for y, line := range strings.Split(strings.TrimSpace(content), "\n") {
		row := make(TileRow, 0)

		for _, c := range line {
			switch c {
			case '.':
				row = append(row, Empty)
			case '\\':
				row = append(row, LeftMirror)
			case '/':
				row = append(row, RightMirror)
			case '|':
				row = append(row, VerticalSplitter)
			case '-':
				row = append(row, HorizontalSplitter)
			}

			if y == 0 {
				grid.Bounds.Width++
			}
		}

		grid.Rows = append(grid.Rows, row)
		grid.Bounds.Height++
	}

	for i := 0; i < grid.Bounds.Height; i++ {
		grid.VisitedRows = append(grid.VisitedRows, make(VisitedTileRow, grid.Bounds.Width))
	}

	return grid
}

func GetEnergizedTiles(grid *Grid, initialPhoton Photon) string {
	energizedTiles := make([][]bool, 0)

	for r := 0; r < grid.Bounds.Height; r++ {
		energizedTiles = append(energizedTiles, make([]bool, grid.Bounds.Width))
	}

	getDescription := func() string {
		description := ""
		for y, row := range energizedTiles {
			if y != 0 {
				description += "\n"
			}

			for _, a := range row {
				if a {
					description += "#"
				} else {
					description += "."
				}
			}
		}

		return description
	}

	grid.StartPhoton(initialPhoton, func(position utilities.Point2D) {
		energizedTiles[position.Y][position.X] = true
	})

	return getDescription()
}

func GetEnergizedTilesCount(grid *Grid, initialPhoton Photon) int {
	energizedTiles := make(map[utilities.Point2D]bool)

	grid.StartPhoton(initialPhoton, func(position utilities.Point2D) {
		energizedTiles[position] = true
	})

	energizedTilesCount := 0

	for range energizedTiles {
		energizedTilesCount++
	}

	return energizedTilesCount
}

func GetMaxEnergizedTilesCount(grid *Grid) int {
	maxEnergizedTileCount := math.MinInt

	getMaxTileCount := func(photon Photon) {
		grid.Reset()
		energizedTileCount := GetEnergizedTilesCount(grid, photon)

		if energizedTileCount > maxEnergizedTileCount {
			maxEnergizedTileCount = energizedTileCount
		}
	}

	// Top row
	for x := 0; x < grid.Bounds.Width; x++ {
		getMaxTileCount(Photon{Position: utilities.NewPoint2D(x, 0), Direction: South})
	}

	// Bottom row
	for x := 0; x < grid.Bounds.Width; x++ {
		getMaxTileCount(Photon{Position: utilities.NewPoint2D(x, grid.Bounds.Height-1), Direction: North})
	}

	// Left column
	for y := 0; y < grid.Bounds.Height; y++ {
		getMaxTileCount(Photon{Position: utilities.NewPoint2D(0, y), Direction: East})
	}

	// Right column
	for y := 0; y < grid.Bounds.Height; y++ {
		getMaxTileCount(Photon{Position: utilities.NewPoint2D(grid.Bounds.Width-1, y), Direction: West})
	}

	return maxEnergizedTileCount
}

func day16(fileContents string) error {
	grid := ParseGrid(fileContents)

	// Part 1: The light isn't energizing enough tiles to produce lava; to debug the contraption,
	// you need to start by analyzing the current situation. With the beam starting in the
	// top-left heading right, how many tiles end up being energized?
	initialPhoton := Photon{
		Position:  utilities.NewPoint2D(0, 0),
		Direction: East,
	}

	energizedTilesCount := GetEnergizedTilesCount(grid, initialPhoton)

	log.Printf("Energized tile count: %d\n", energizedTilesCount)

	// Part 2: Find the initial beam configuration that energizes the largest number of tiles;
	// how many tiles are energized in that configuration?
	maximumEnergedTilesCount := GetMaxEnergizedTilesCount(grid)

	log.Printf("Maximum energized tile count: %d\n", maximumEnergedTilesCount)

	return nil
}
