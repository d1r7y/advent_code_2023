package day13

import (
	"fmt"
	"log"
	"strings"

	"github.com/d1r7y/advent_2023/utilities"
)

type Terrain byte

type TerrainRow []Terrain

const (
	Ash Terrain = iota
	Rock
)

func (t Terrain) Describe() string {
	if t == Ash {
		return "."
	} else if t == Rock {
		return "#"
	}

	log.Panicf("unexpected terrain: %d\n", t)
	return "$"
}

func (t Terrain) Invert() Terrain {
	if t == Ash {
		return Rock
	} else if t == Rock {
		return Ash
	}

	log.Panicf("unexpected terrain: %d\n", t)
	return Ash
}

type ReflectionAxis byte

const (
	None ReflectionAxis = iota
	Vertical
	Horizontal
)

type Reflection struct {
	Axis     ReflectionAxis
	Position int
}

func (r Reflection) Describe() string {
	if r.Axis == Vertical {
		return fmt.Sprintf("Vertical @ %d", r.Position)
	} else if r.Axis == Horizontal {
		return fmt.Sprintf("Horizontal @ %d", r.Position)
	}

	return "None"
}

type Landscape struct {
	Bounds utilities.Size2D
	Ground []TerrainRow
}

func (l *Landscape) Describe() string {
	description := ""
	for _, r := range l.Ground {
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

func (l *Landscape) GetReflectionCore(getExcludedReflection func() Reflection, equalTerrain func(count int, t1 TerrainRow, t2 TerrainRow) bool) Reflection {
	reflection := Reflection{Axis: None}
	excludedReflection := Reflection{Axis: None}

	if getExcludedReflection != nil {
		excludedReflection = getExcludedReflection()
	}

	// Check for horizontal reflection.
NextRow:
	for i := 0; i < l.Bounds.Height-1; i++ {
		upper := i
		lower := i + 1

		for {
			if !equalTerrain(l.Bounds.Width, l.Ground[upper], l.Ground[lower]) {
				continue NextRow
			}

			upper--
			if upper < 0 {
				break
			}

			lower++
			if lower == l.Bounds.Height {
				break
			}
		}

		if excludedReflection.Axis != Horizontal || excludedReflection.Position != i {
			// If we got here, then we have a horizontal reflection.
			reflection.Axis = Horizontal
			reflection.Position = i
			return reflection
		}
	}

	createColumn := func(column int) TerrainRow {
		t := make(TerrainRow, 0)

		for _, r := range l.Ground {
			t = append(t, r[column])
		}

		return t
	}

	// Check for vertical reflection.
NextColumn:
	for i := 0; i < l.Bounds.Width-1; i++ {
		left := i
		right := i + 1

		for {
			t1 := createColumn(left)
			t2 := createColumn(right)

			if !equalTerrain(l.Bounds.Height, t1, t2) {
				continue NextColumn
			}

			left--
			if left < 0 {
				break
			}

			right++
			if right == l.Bounds.Width {
				break
			}
		}

		if excludedReflection.Axis != Vertical || excludedReflection.Position != i {
			// If we got here, then we have a vertical reflection.
			reflection.Axis = Vertical
			reflection.Position = i
			return reflection
		}
	}

	return reflection
}

func (l *Landscape) GetReflectionSmudged(excludedReflection Reflection) Reflection {
	getExcludedReflection := func() Reflection {
		return excludedReflection
	}

	equalTerrain := func(count int, t1 TerrainRow, t2 TerrainRow) bool {
		for i := 0; i < count; i++ {
			if t1[i] != t2[i] {
				return false
			}
		}

		return true
	}

	for y := 0; y < l.Bounds.Height; y++ {
		for x := 0; x < l.Bounds.Width; x++ {
			l.Ground[y][x] = l.Ground[y][x].Invert()
			reflection := l.GetReflectionCore(getExcludedReflection, equalTerrain)
			l.Ground[y][x] = l.Ground[y][x].Invert()
			if reflection.Axis != None {
				return reflection
			}
		}
	}

	log.Println(l.Describe())
	log.Panic("couldn't find alternative reflection\n")
	return Reflection{}
}

func (l *Landscape) GetReflection() Reflection {
	equalTerrain := func(count int, t1 TerrainRow, t2 TerrainRow) bool {
		for i := 0; i < count; i++ {
			if t1[i] != t2[i] {
				return false
			}
		}

		return true
	}

	return l.GetReflectionCore(nil, equalTerrain)
}

func NewLandscape() *Landscape {
	return &Landscape{}
}

func ParseLandscape(lines []string) *Landscape {
	landscape := NewLandscape()

	landscape.Ground = make([]TerrainRow, 0)

	for y, line := range lines {
		row := make(TerrainRow, 0)

		for _, c := range line {
			switch c {
			case '.':
				row = append(row, Ash)
			case '#':
				row = append(row, Rock)
			}

			if y == 0 {
				landscape.Bounds.Width++
			}
		}

		landscape.Ground = append(landscape.Ground, row)

		landscape.Bounds.Height++
	}

	return landscape
}

func day13(fileContents string) error {
	// Part 1: Find the line of reflection in each of the patterns in your notes.
	// What number do you get after summarizing all of your notes?

	lineBundles := make([][]string, 0)
	currentBundle := make([]string, 0)

	for _, line := range strings.Split(strings.TrimSpace(fileContents), "\n") {
		if line != "" {
			currentBundle = append(currentBundle, line)
		} else {
			lineBundles = append(lineBundles, currentBundle)
			currentBundle = make([]string, 0)
		}
	}

	if len(currentBundle) > 0 {
		lineBundles = append(lineBundles, currentBundle)
	}

	noteSummary := 0

	for _, bundles := range lineBundles {
		landscape := ParseLandscape(bundles)

		reflection := landscape.GetReflection()

		if reflection.Axis == Vertical {
			noteSummary += reflection.Position + 1
		} else if reflection.Axis == Horizontal {
			noteSummary += 100 * (reflection.Position + 1)
		} else {
			log.Panicf("no reflection found for '%s'\n", strings.Join(bundles, "\n"))
		}
	}

	log.Printf("Note summary: %d\n", noteSummary)

	// Part 2: In each pattern, fix the smudge and find the different line of reflection.
	// What number do you get after summarizing the new reflection line in each pattern in your notes?
	noteSummarySmudged := 0

	for _, bundles := range lineBundles {
		landscape := ParseLandscape(bundles)

		excludedReflection := landscape.GetReflection()
		reflection := landscape.GetReflectionSmudged(excludedReflection)

		if reflection.Axis == Vertical {
			noteSummarySmudged += reflection.Position + 1
		} else if reflection.Axis == Horizontal {
			noteSummarySmudged += 100 * (reflection.Position + 1)
		} else {
			log.Panicf("no reflection found for '%s'\n", strings.Join(bundles, "\n"))
		}
	}

	log.Printf("Note summary for smudged mirrors: %d\n", noteSummarySmudged)

	return nil
}
