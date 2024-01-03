package day20

import (
	"fmt"
	"log"
	"strings"
)

type WrappedList struct {
	List []int
}

func ParseWrappedList(fileContents string) *WrappedList {
	wl := &WrappedList{List: make([]int, 0)}

	for _, line := range strings.Split(fileContents, "\n") {
		var number int
		count, err := fmt.Sscanf(line, "%d", &number)
		if err != nil {
			log.Panic("invalid number")
		}
		if count != 1 {
			log.Panic("invalid line")
		}

		wl.List = append(wl.List, number)
	}

	return wl
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (w *WrappedList) NewIndex(index, delta int) int {
	if delta == 0 {
		return index
	}

	newIndex := mod(index+delta, len(w.List))

	if delta < 0 {
		if newIndex > index {
			// We wrapped.
			newIndex--
		} else if newIndex == 0 {
			// Something something leapfrog versus swap...
			newIndex = len(w.List) - 1
		}
	} else if delta > 0 {
		if newIndex < index {
			// We wrapped.
			newIndex++
		} else if newIndex == len(w.List)-1 {
			// Something something leapfrog versus swap...
			newIndex = 0
		}
	}

	return newIndex
}

func (w *WrappedList) Move(index int, delta int) int {
	newIndex := w.NewIndex(index, delta)

	if newIndex == index {
		return index
	}

	value := w.List[index]

	if newIndex > index {
		for i := index; i < newIndex; i++ {
			w.List[i] = w.List[i+1]
		}
	} else {
		for i := index; i > newIndex; i-- {
			w.List[i] = w.List[i-1]
		}
	}

	w.List[newIndex] = value

	return newIndex
}

func (w *WrappedList) Mix() {
	handled := make([]bool, len(w.List))

	for index := 0; index < len(w.List); {
		if handled[index] {
			index++
			continue
		}

		d := w.List[index]
		newIndex := w.Move(index, d)

		if newIndex > index {
			for i := index; i < newIndex; i++ {
				handled[i] = handled[i+1]
			}
		} else {
			for i := index; i > newIndex; i-- {
				handled[i] = handled[i-1]
			}
		}

		handled[newIndex] = true
	}
}

func (w *WrappedList) GetCoordinates() [3]int {
	for i, n := range w.List {
		if n == 0 {
			c1 := w.List[mod(i+1000, len(w.List))]
			c2 := w.List[mod(i+2000, len(w.List))]
			c3 := w.List[mod(i+3000, len(w.List))]
			return [3]int{c1, c2, c3}
		}
	}

	log.Panic("Couldn't find 0")
	return [3]int{0, 0, 0}
}

func (w *WrappedList) Describe() string {
	str := ""
	for i, n := range w.List {
		if i != 0 {
			str += ", "
		}
		str += fmt.Sprint(n)
	}

	return str
}

func day20(fileContents string) error {
	wl := ParseWrappedList(fileContents)

	// Part 1: Mix the input file to decrypt it.  Get the coordinates.
	wl.Mix()

	sum := 0

	for _, c := range wl.GetCoordinates() {
		sum += c
	}

	fmt.Printf("Sum of grove coordinates: %d\n", sum)

	// Part 2: Ignore the surfaces that are trapped within the droplets.  What is the exterior
	// surface area of the lava droplet?

	return nil
}
