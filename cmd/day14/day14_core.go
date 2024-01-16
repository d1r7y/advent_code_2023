package day14

import (
	"log"
	"strings"
	"sync"

	"github.com/d1r7y/advent_2023/utilities"
)

const ColumnsPerGoRoutine = 5

type Rock byte

const (
	Empty Rock = iota
	Rounded
	Cube
)

func (r Rock) Describe() string {
	switch r {
	case Empty:
		return "."
	case Rounded:
		return "O"
	case Cube:
		return "#"
	}

	log.Panicf("unknown rock type: %d\n", r)
	return "X"
}

type Column []Rock

type Direction byte

type Platform struct {
	Bounds  utilities.Size2D
	Columns []Column
}

func ParsePlatform(lines []string) *Platform {
	platform := &Platform{}

	platform.Bounds.Width = len(lines[0])
	platform.Bounds.Height = len(lines)

	platform.Columns = make([]Column, platform.Bounds.Width)

	for i := 0; i < platform.Bounds.Width; i++ {
		platform.Columns[i] = make(Column, platform.Bounds.Height)
	}

	for c := 0; c < platform.Bounds.Height; c++ {
		for r := 0; r < platform.Bounds.Width; r++ {
			var rock Rock
			switch lines[c][r] {
			case '.':
				rock = Empty
			case 'O':
				rock = Rounded
			case '#':
				rock = Cube
			default:
				log.Fatalf("unknown rock type: %d\n", lines[c][r])
			}

			platform.Columns[r][c] = rock
		}
	}

	return platform
}

func (p *Platform) Describe() string {
	str := ""

	for r := 0; r < p.Bounds.Width; r++ {
		if str != "" {
			str += "\n"
		}
		for c := 0; c < p.Bounds.Height; c++ {
			str += p.Columns[c][r].Describe()
		}
	}

	return str
}

func (p *Platform) TiltNorth() {
	processColumns := func(wg *sync.WaitGroup, startColumnIndex int, endColumnIndex int) {
		if wg != nil {
			defer wg.Done()
		}

		for ci := startColumnIndex; ci <= endColumnIndex; ci++ {
			column := p.Columns[ci]
			for i := 1; i < p.Bounds.Height; i++ {
				rock := column[i]

				if rock != Rounded {
					continue
				}

				// Rounded rock.
				// See how far we can slide it before:
				// - We either hit the top (i=0).
				// - We hit a cube rock.
				// - We hit another rounded rock.
				si := i - 1
				for ; si >= 0; si-- {
					if column[si] != Empty {
						break
					}
				}

				if si >= 0 {
					column[i] = Empty
					column[si+1] = rock
				} else {
					column[0] = rock
					column[i] = Empty
				}
			}

		}
	}

	if len(p.Columns) < ColumnsPerGoRoutine {
		processColumns(nil, 0, len(p.Columns)-1)
		return
	}

	var wg sync.WaitGroup

	groups := len(p.Columns) / ColumnsPerGoRoutine
	for i := 0; i < groups; i++ {
		wg.Add(1)
		go processColumns(&wg, i*ColumnsPerGoRoutine, i*ColumnsPerGoRoutine+ColumnsPerGoRoutine-1)
	}

	remainder := len(p.Columns) % ColumnsPerGoRoutine
	if remainder > 0 {
		wg.Add(remainder)
		go processColumns(&wg, groups*ColumnsPerGoRoutine, groups*ColumnsPerGoRoutine+remainder)
	}

	wg.Wait()
}

func (p *Platform) TiltSouth() {
	processColumns := func(wg *sync.WaitGroup, startColumnIndex int, endColumnIndex int) {
		if wg != nil {
			defer wg.Done()
		}

		for ci := startColumnIndex; ci <= endColumnIndex; ci++ {
			column := p.Columns[ci]
			for i := p.Bounds.Height - 2; i >= 0; i-- {
				rock := column[i]

				if rock != Rounded {
					continue
				}

				// Rounded rock.
				// See how far we can slide it before:
				// - We either hit the bottom (i=p.Bounds.Height-1).
				// - We hit a cube rock.
				// - We hit another rounded rock.
				si := i + 1
				for ; si <= p.Bounds.Height-1; si++ {
					if column[si] != Empty {
						break
					}
				}

				if si <= p.Bounds.Height-1 {
					column[i] = Empty
					column[si-1] = rock
				} else {
					column[p.Bounds.Height-1] = rock
					column[i] = Empty
				}
			}
		}
	}

	if len(p.Columns) < ColumnsPerGoRoutine {
		processColumns(nil, 0, len(p.Columns)-1)
		return
	}

	var wg sync.WaitGroup

	groups := len(p.Columns) / ColumnsPerGoRoutine
	for i := 0; i < groups; i++ {
		wg.Add(1)
		go processColumns(&wg, i*ColumnsPerGoRoutine, i*ColumnsPerGoRoutine+ColumnsPerGoRoutine-1)
	}

	remainder := len(p.Columns) % ColumnsPerGoRoutine
	if remainder > 0 {
		wg.Add(remainder)
		go processColumns(&wg, groups*ColumnsPerGoRoutine, groups*ColumnsPerGoRoutine+remainder)
	}

	wg.Wait()
}

func (p *Platform) TiltEast() {
	processRows := func(wg *sync.WaitGroup, startRowIndex int, endRowIndex int) {
		if wg != nil {
			defer wg.Done()
		}

		for ri := startRowIndex; ri <= endRowIndex; ri++ {
			for j := p.Bounds.Width - 2; j >= 0; j-- {
				rock := p.Columns[j][ri]

				if rock != Rounded {
					continue
				}

				// Rounded rock.
				// See how far we can slide it before:
				// - We either hit the right (j=p.Bounds.Width-1).
				// - We hit a cube rock.
				// - We hit another rounded rock.
				si := j + 1
				for ; si <= p.Bounds.Width-1; si++ {
					if p.Columns[si][ri] != Empty {
						break
					}
				}

				if si <= p.Bounds.Width-1 {
					p.Columns[j][ri] = Empty
					p.Columns[si-1][ri] = rock
				} else {
					p.Columns[p.Bounds.Width-1][ri] = rock
					p.Columns[j][ri] = Empty
				}
			}
		}
	}

	if len(p.Columns[0]) < ColumnsPerGoRoutine {
		processRows(nil, 0, len(p.Columns[0])-1)
		return
	}

	var wg sync.WaitGroup

	groups := len(p.Columns[0]) / ColumnsPerGoRoutine
	for i := 0; i < groups; i++ {
		wg.Add(1)
		go processRows(&wg, i*ColumnsPerGoRoutine, i*ColumnsPerGoRoutine+ColumnsPerGoRoutine-1)
	}

	remainder := len(p.Columns[0]) % ColumnsPerGoRoutine
	if remainder > 0 {
		wg.Add(remainder)
		go processRows(&wg, groups*ColumnsPerGoRoutine, groups*ColumnsPerGoRoutine+remainder)
	}

	wg.Wait()
}

func (p *Platform) TiltWest() {
	processRows := func(wg *sync.WaitGroup, startRowIndex int, endRowIndex int) {
		if wg != nil {
			defer wg.Done()
		}

		for ri := startRowIndex; ri <= endRowIndex; ri++ {
			for j := 1; j <= p.Bounds.Width-1; j++ {
				rock := p.Columns[j][ri]

				if rock != Rounded {
					continue
				}

				// Rounded rock.
				// See how far we can slide it before:
				// - We either hit the left (j=0).
				// - We hit a cube rock.
				// - We hit another rounded rock.
				si := j - 1
				for ; si >= 0; si-- {
					if p.Columns[si][ri] != Empty {
						break
					}
				}

				if si >= 0 {
					p.Columns[j][ri] = Empty
					p.Columns[si+1][ri] = rock
				} else {
					p.Columns[0][ri] = rock
					p.Columns[j][ri] = Empty
				}
			}
		}
	}

	if len(p.Columns[0]) < ColumnsPerGoRoutine {
		processRows(nil, 0, len(p.Columns[0])-1)
		return
	}

	var wg sync.WaitGroup

	groups := len(p.Columns[0]) / ColumnsPerGoRoutine
	for i := 0; i < groups; i++ {
		wg.Add(1)
		go processRows(&wg, i*ColumnsPerGoRoutine, i*ColumnsPerGoRoutine+ColumnsPerGoRoutine-1)
	}

	remainder := len(p.Columns[0]) % ColumnsPerGoRoutine
	if remainder > 0 {
		wg.Add(remainder)
		go processRows(&wg, groups*ColumnsPerGoRoutine, groups*ColumnsPerGoRoutine+remainder)
	}

	wg.Wait()
}

func (p *Platform) TiltCycle() {
	p.TiltNorth()
	p.TiltWest()
	p.TiltSouth()
	p.TiltEast()
}

func (p *Platform) Load() int {
	totalLoad := 0

	for _, column := range p.Columns {
		for i := 0; i < p.Bounds.Height; i++ {
			if column[i] == Rounded {
				totalLoad += p.Bounds.Height - i
			}
		}
	}

	return totalLoad
}

func day14(fileContents string) error {
	platform := ParsePlatform(strings.Split(fileContents, "\n"))

	// Part 1: Tilt the platform so that the rounded rocks all roll north.
	// Afterward, what is the total load on the north support beams?
	platform.TiltNorth()

	log.Printf("Total load on north support beams: %d\n", platform.Load())

	// Part 2: Run the spin cycle for 1000000000 cycles. Afterward, what is the
	// total load on the north support beams?
	for i := 0; i < 1000000000; i++ {
		if i%100000 == 0 {
			log.Printf("Cycle %d\n", i)
		}

		platform.TiltCycle()
	}

	log.Printf("Total load on north support beams after spin cycles: %d\n", platform.Load())

	return nil
}
