package day10

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/d1r7y/advent_2023/utilities"
)

type Tile byte

func (t Tile) Describe() string {
	switch t {
	case VerticalPipe:
		return "|"
	case HorizontalPipe:
		return "-"
	case BendNorthEastPipe:
		return "L"
	case BendNorthWestPipe:
		return "J"
	case BendSouthWestPipe:
		return "7"
	case BendSouthEastPipe:
		return "F"
	case Ground:
		return "."
	case Start:
		return "S"
	}

	log.Panicf("unknown tile %d\n", t)
	return " "
}

const (
	VerticalPipe Tile = iota
	HorizontalPipe
	BendNorthEastPipe
	BendNorthWestPipe
	BendSouthWestPipe
	BendSouthEastPipe
	Ground
	Start
)

type Direction byte

func (d Direction) Describe() string {
	switch d {
	case North:
		return "N"
	case South:
		return "S"
	case East:
		return "E"
	case West:
		return "W"
	}

	log.Panicf("unknown directions %d\n", d)
	return " "
}

const (
	North Direction = iota
	East
	South
	West
)

type Row []Tile

type Distances struct {
	Bounds utilities.Size2D
	Rows   [][]int
}

func NewDistances(bounds utilities.Size2D) *Distances {
	distances := &Distances{
		Bounds: bounds,
		Rows:   make([][]int, bounds.Height),
	}

	for y := 0; y < distances.Bounds.Height; y++ {
		row := make([]int, bounds.Width)
		for x := 0; x < distances.Bounds.Width; x++ {
			row[x] = -1
		}

		distances.Rows[y] = row
	}

	return distances
}

func (d *Distances) Describe() string {
	str := ""

	for y, row := range d.Rows {
		for _, distance := range row {

			if distance < 0 {
				str += "."
			} else {
				if distance == 0 {
					str += "0"
				} else {
					str += fmt.Sprintf("%d", int(math.Log(float64(distance))))
				}
			}
		}

		if y < d.Bounds.Height-1 {
			str += "\n"
		}
	}

	return str
}

func (d *Distances) validatePoint(p utilities.Point2D) {
	if p.Y < 0 || p.Y >= d.Bounds.Height {
		log.Panicf("invalid position %d,%d\n", p.X, p.Y)
	}
	if p.X < 0 || p.X >= d.Bounds.Width {
		log.Panicf("invalid position %d,%d\n", p.X, p.Y)
	}
}

func (d *Distances) SetDistance(p utilities.Point2D, distance int) {
	d.validatePoint(p)

	d.Rows[p.Y][p.X] = distance
}

func (d *Distances) GetDistance(p utilities.Point2D) int {
	d.validatePoint(p)

	return d.Rows[p.Y][p.X]
}

type Grid struct {
	StartPosition utilities.Point2D
	Bounds        utilities.Size2D
	Rows          []Row
}

func ParseGrid(lines []string) *Grid {

	grid := &Grid{}
	grid.Rows = make([]Row, 0)

	for y, line := range lines {
		row := make([]Tile, 0)

		for x, t := range line {
			var tile Tile

			switch t {
			case '|':
				tile = VerticalPipe
			case '-':
				tile = HorizontalPipe
			case 'L':
				tile = BendNorthEastPipe
			case 'J':
				tile = BendNorthWestPipe
			case '7':
				tile = BendSouthWestPipe
			case 'F':
				tile = BendSouthEastPipe
			case '.':
				tile = Ground
			case 'S':
				tile = Start

				grid.StartPosition.X = x
				grid.StartPosition.Y = y
			}

			if y == 0 {
				grid.Bounds.Width++
			}
			row = append(row, tile)
		}

		grid.Bounds.Height++
		grid.Rows = append(grid.Rows, row)
	}

	// Now determine what pipe is at the starting location.  Find the two directions leading out of the start node.
	exitDirections := make([]Direction, 0)

	grid.ForEachNeighbor(grid.StartPosition, func(p utilities.Point2D, d Direction, t Tile) bool {
		if CanTileExit(t, InvertDirection(d)) {
			exitDirections = append(exitDirections, d)
		}

		return true
	})

	if len(exitDirections) != 2 {
		log.Panicf("unexpected number of exit directions: %v\n", exitDirections)
	}

	startTile := TileFromDirections(exitDirections[0], exitDirections[1])

	grid.SetTile(grid.StartPosition, startTile)

	return grid
}

func (g *Grid) validatePoint(p utilities.Point2D) {
	if p.Y < 0 || p.Y >= g.Bounds.Height {
		log.Panicf("invalid point %d,%d\n", p.X, p.Y)
	}
	if p.X < 0 || p.X >= g.Bounds.Width {
		log.Panicf("invalid point %d,%d\n", p.X, p.Y)
	}
}

func (g *Grid) ForEachNeighbor(p utilities.Point2D, callback func(p utilities.Point2D, d Direction, t Tile) bool) {
	g.validatePoint(p)

	if p.Y > 0 {
		// North
		if !callback(utilities.NewPoint2D(p.X, p.Y-1), North, g.GetTile(utilities.NewPoint2D(p.X, p.Y-1))) {
			return
		}
	}

	if p.Y < g.Bounds.Height-1 {
		// South
		if !callback(utilities.NewPoint2D(p.X, p.Y+1), South, g.GetTile(utilities.NewPoint2D(p.X, p.Y+1))) {
			return
		}
	}

	if p.X < g.Bounds.Width-1 {
		// East
		if !callback(utilities.NewPoint2D(p.X+1, p.Y), East, g.GetTile(utilities.NewPoint2D(p.X+1, p.Y))) {
			return
		}
	}

	if p.X > 0 {
		// West
		if !callback(utilities.NewPoint2D(p.X-1, p.Y), West, g.GetTile(utilities.NewPoint2D(p.X-1, p.Y))) {
			return
		}
	}
}

func (g *Grid) TraverseLoop(callback func(p utilities.Point2D, d Direction, t Tile) bool) {
	currentPosition := g.StartPosition
	currentTile := g.GetTile(currentPosition)
	var currentDirection Direction

	// Pick an initial starting direction.
	switch currentTile {
	case VerticalPipe:
		currentDirection = North
	case HorizontalPipe:
		currentDirection = East
	case BendNorthEastPipe:
		currentDirection = South
	case BendNorthWestPipe:
		currentDirection = South
	case BendSouthEastPipe:
		currentDirection = North
	case BendSouthWestPipe:
		currentDirection = North
	default:
		log.Panicf("unexpected tile: %s\n", currentTile.Describe())
	}

	for {
		if !callback(currentPosition, currentDirection, currentTile) {
			break
		}

		currentDirection = UnusedExitDirection(currentTile, currentDirection)
		currentPosition = UpdatePosition(currentPosition, currentDirection)
		currentTile = g.GetTile(currentPosition)

		// We're back to the beginning.
		if currentPosition == g.StartPosition {
			break
		}
	}
}

func (g *Grid) GetTile(p utilities.Point2D) Tile {
	g.validatePoint(p)

	return g.Rows[p.Y][p.X]
}

func (g *Grid) SetTile(p utilities.Point2D, t Tile) {
	g.validatePoint(p)

	g.Rows[p.Y][p.X] = t
}

func (g *Grid) GetNeighborTile(p utilities.Point2D, d Direction) Tile {
	// Make sure we aren't going out of bounds.
	switch d {
	case North:
		if p.Y == 0 {
			log.Panicf("invalid direction %d from position %d,%d\n", d, p.X, p.Y)
		}
		return g.Rows[p.Y-1][p.X]
	case East:
		if p.X == g.Bounds.Width-1 {
			log.Panicf("invalid direction %d from position %d,%d\n", d, p.X, p.Y)
		}
		return g.Rows[p.Y][p.X+1]
	case South:
		if p.Y == g.Bounds.Height-1 {
			log.Panicf("invalid direction %d from position %d,%d\n", d, p.X, p.Y)
		}
		return g.Rows[p.Y+1][p.X]
	case West:
		if p.X == 0 {
			log.Panicf("invalid direction %d from position %d,%d\n", d, p.X, p.Y)
		}
		return g.Rows[p.Y][p.X-1]
	}

	log.Panicf("unexpected direction %d\n", d)
	return Start
}

func (g *Grid) Describe() string {
	str := ""

	for y, row := range g.Rows {
		for x, t := range row {
			ts := t.Describe()

			if g.StartPosition.X == x && g.StartPosition.Y == y {
				ts = "S"
			}

			str += ts
		}

		if y < g.Bounds.Height-1 {
			str += "\n"
		}
	}

	return str
}

func IsTilePipe(t Tile) bool {
	switch t {
	case VerticalPipe:
	case HorizontalPipe:
	case BendNorthEastPipe:
	case BendNorthWestPipe:
	case BendSouthEastPipe:
	case BendSouthWestPipe:
	default:
		return false
	}

	return true
}

func TileFromDirections(d1 Direction, d2 Direction) Tile {
	if d1 == d2 {
		log.Panicf("both directions are the same %d\n", d1)
	}

	if (d1 == North && d2 == South) || (d1 == South && d2 == North) {
		return VerticalPipe
	}
	if (d1 == West && d2 == East) || (d1 == East && d2 == West) {
		return HorizontalPipe
	}
	if (d1 == North && d2 == East) || (d1 == East && d2 == North) {
		return BendNorthEastPipe
	}
	if (d1 == North && d2 == West) || (d1 == West && d2 == North) {
		return BendNorthWestPipe
	}
	if (d1 == South && d2 == East) || (d1 == East && d2 == South) {
		return BendSouthEastPipe
	}
	if (d1 == South && d2 == West) || (d1 == West && d2 == South) {
		return BendSouthWestPipe
	}

	log.Panicf("unexpected directions %d and %d\n", d1, d2)

	return Start
}

func InvertDirection(d Direction) Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	}

	log.Panicf("unexpected direction: %d\n", d)
	return North
}

func UnusedExitDirection(t Tile, d Direction) Direction {
	switch t {
	case VerticalPipe:
		if d == North {
			return North
		} else if d == South {
			return South
		}
		log.Panicf("unexpected tile/direction: %s %s\n", t.Describe(), d.Describe())
	case HorizontalPipe:
		if d == East {
			return East
		} else if d == West {
			return West
		}
		log.Panicf("unexpected tile/direction: %s %s\n", t.Describe(), d.Describe())
	case BendNorthEastPipe:
		if d == South {
			return East
		} else if d == West {
			return North
		}
		log.Panicf("unexpected tile/direction: %s %s\n", t.Describe(), d.Describe())
	case BendNorthWestPipe:
		if d == South {
			return West
		} else if d == East {
			return North
		}
		log.Panicf("unexpected tile/direction: %s %s\n", t.Describe(), d.Describe())
	case BendSouthEastPipe:
		if d == North {
			return East
		} else if d == West {
			return South
		}
		log.Panicf("unexpected tile/direction: %s %s\n", t.Describe(), d.Describe())
	case BendSouthWestPipe:
		if d == North {
			return West
		} else if d == East {
			return South
		}
		log.Panicf("unexpected tile/direction: %s %s\n", t.Describe(), d.Describe())
	}

	log.Panicf("unexpected tile: %s\n", t.Describe())
	return North
}

func CanTileExit(t Tile, d Direction) bool {
	switch t {
	case VerticalPipe:
		if d == North || d == South {
			return true
		}
	case HorizontalPipe:
		if d == East || d == West {
			return true
		}
	case BendNorthEastPipe:
		if d == North || d == East {
			return true
		}
	case BendNorthWestPipe:
		if d == North || d == West {
			return true
		}
	case BendSouthEastPipe:
		if d == South || d == East {
			return true
		}
	case BendSouthWestPipe:
		if d == South || d == West {
			return true
		}
	}

	return false
}

func UpdatePosition(p utilities.Point2D, d Direction) utilities.Point2D {
	newPosition := p

	switch d {
	case North:
		newPosition.Y--
	case South:
		newPosition.Y++
	case East:
		newPosition.X++
	case West:
		newPosition.X--
	}

	return newPosition
}

// CanConnect Can t1 connect to t2 if t2 is a d neighbor of t1?
func CanConnect(t1 Tile, d Direction, t2 Tile) bool {
	if !IsTilePipe(t1) || !IsTilePipe(t2) {
		log.Panicf("unexpected tile type: %d\n", t1)
	}

	switch t1 {
	case VerticalPipe:
		if d == North {
			if t2 == BendSouthEastPipe || t2 == BendSouthWestPipe || t2 == VerticalPipe {
				return true
			}
		} else if d == South {
			if t2 == BendNorthEastPipe || t2 == BendNorthWestPipe || t2 == VerticalPipe {
				return true
			}
		}
	case HorizontalPipe:
		if d == East {
			if t2 == BendNorthWestPipe || t2 == BendSouthWestPipe || t2 == HorizontalPipe {
				return true
			}
		} else if d == West {
			if t2 == BendNorthEastPipe || t2 == BendSouthEastPipe || t2 == HorizontalPipe {
				return true
			}
		}
	case BendNorthEastPipe:
		if d == North && (t2 == BendSouthEastPipe || t2 == BendSouthWestPipe || t2 == VerticalPipe) {
			return true
		}
		if d == East && (t2 == BendSouthWestPipe || t2 == BendNorthWestPipe || t2 == HorizontalPipe) {
			return true
		}
	case BendNorthWestPipe:
		if d == North && (t2 == BendSouthEastPipe || t2 == BendSouthWestPipe || t2 == VerticalPipe) {
			return true
		}
		if d == West && (t2 == BendNorthEastPipe || t2 == BendSouthEastPipe || t2 == HorizontalPipe) {
			return true
		}
	case BendSouthEastPipe:
		if d == South && (t2 == BendNorthEastPipe || t2 == BendNorthWestPipe || t2 == VerticalPipe) {
			return true
		}
		if d == East && (t2 == BendNorthWestPipe || t2 == BendSouthWestPipe || t2 == HorizontalPipe) {
			return true
		}
	case BendSouthWestPipe:
		if d == South && (t2 == BendNorthEastPipe || t2 == BendNorthWestPipe || t2 == VerticalPipe) {
			return true
		}
		if d == West && (t2 == BendNorthEastPipe || t2 == BendSouthEastPipe || t2 == HorizontalPipe) {
			return true
		}
	}

	return false
}

func day10(fileContents string) error {
	// Part 1: Find the single giant loop starting at S. How many steps along the loop does it take
	// to get from the starting position to the point farthest from the starting position?
	grid := ParseGrid(strings.Split(strings.TrimSpace(fileContents), "\n"))

	distances := NewDistances(grid.Bounds)
	distance := 0

	grid.TraverseLoop(func(p utilities.Point2D, d Direction, t Tile) bool {
		distances.SetDistance(p, distance)
		distance++
		return true
	})

	log.Printf("Steps to furthest point from starting position: %d\n", distance/2)

	// Part 2: Figure out whether you have time to search for the nest by calculating the area
	// within the loop. How many tiles are enclosed by the loop?
	vertices := make([]utilities.Point2D, 0)

	visited := NewDistances(grid.Bounds)

	grid.TraverseLoop(func(p utilities.Point2D, d Direction, t Tile) bool {
		visited.SetDistance(p, 10)
		vertices = append(vertices, p)
		return true
	})

	vertices = append(vertices, grid.StartPosition)

	area := 0

	for y := 0; y < visited.Bounds.Height; y++ {
		for x := 0; x < visited.Bounds.Width; x++ {
			d := visited.GetDistance(utilities.NewPoint2D(x, y))
			if d < 0 {
				if utilities.PointInPolyCrossing(utilities.NewPoint2D(x, y), vertices) {
					area++
				}
			}
		}
	}

	log.Printf("Area enclosed by loop: %d\n", area)

	return nil
}
