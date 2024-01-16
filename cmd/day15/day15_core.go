package day15

import (
	"log"
	"strconv"
	"strings"
)

type Lens struct {
	Name        string
	FocalLength int
}

type Box struct {
	Number int
	Lenses []Lens
}

func (b Box) FocusingPower() int {
	focusingPower := 0

	for i, l := range b.Lenses {
		focusingPower += (b.Number + 1) * (i + 1) * l.FocalLength
	}

	return focusingPower
}

type BoxLine [256]Box

func NewBoxLine() *BoxLine {
	bl := &BoxLine{}

	for i := 0; i < len(bl); i++ {
		bl[i].Number = i
		bl[i].Lenses = make([]Lens, 0)
	}

	return bl
}

func (bl *BoxLine) RemoveLens(name string) {
	boxIndex := Hash(name)

	for i, b := range bl[boxIndex].Lenses {
		if b.Name == name {
			lenses := make([]Lens, 0)
			lenses = append(lenses, bl[boxIndex].Lenses[:i]...)
			bl[boxIndex].Lenses = append(lenses, bl[boxIndex].Lenses[i+1:]...)

			return
		}
	}
}

func (bl *BoxLine) SetLens(name string, focalLength int) {
	boxIndex := Hash(name)

	for i := range bl[boxIndex].Lenses {
		if bl[boxIndex].Lenses[i].Name == name {
			bl[boxIndex].Lenses[i].FocalLength = focalLength

			return
		}
	}

	// Lens name didn't exist.  Add it to the end of the list.
	bl[boxIndex].Lenses = append(bl[boxIndex].Lenses, Lens{Name: name, FocalLength: focalLength})
}

func (bl *BoxLine) TotalFocusingPower() int {
	totalFocusingPower := 0

	for _, b := range bl {
		totalFocusingPower += b.FocusingPower()
	}

	return totalFocusingPower
}

func Hash(str string) byte {
	return RunningHash(0, str)
}

func RunningHash(currentHash byte, str string) byte {
	for _, c := range str {
		currentHash += byte(c)
		interim := int(currentHash) * 17
		currentHash = byte(interim % 256)
	}

	return currentHash
}

func SumInitializationSequence(str string) int {
	sum := 0
	for _, s := range strings.Split(str, ",") {
		sum += int(Hash(s))
	}

	return sum
}

func SumFocusingPowerFromInitializationSequence(str string) int {
	bl := NewBoxLine()

	for _, is := range strings.Split(str, ",") {
		// "is" consists of a label and either:
		// * an = followed by a number
		// * a -
		equal := strings.Split(is, "=")
		if len(equal) == 2 {
			fl, err := strconv.Atoi(equal[1])
			if err != nil {
				log.Fatal(err)
			}

			bl.SetLens(equal[0], fl)
		} else {
			label := strings.ReplaceAll(is, "-", "")
			bl.RemoveLens(label)
		}
	}

	return bl.TotalFocusingPower()
}

func day15(fileContents string) error {
	// Part 1: Run the HASH algorithm on each step in the initialization sequence.
	// What is the sum of the results? (The initialization sequence is one long line;
	// be careful when copy-pasting it.)
	sum := SumInitializationSequence(fileContents)

	log.Printf("Sum of initialization sequence hash: %d\n", sum)

	// Part 2: With the help of an over-enthusiastic reindeer in a hard hat,
	// follow the initialization sequence. What is the focusing power of the
	// resulting lens configuration?
	focusingPower := SumFocusingPowerFromInitializationSequence(fileContents)

	log.Printf("Focusing power of resulting lens configuration: %d\n", focusingPower)

	return nil
}
