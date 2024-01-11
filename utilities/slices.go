package utilities

import "math"

func MaxIntInList(l []int) int {
	max := math.MinInt

	for _, e := range l {
		if e > max {
			max = e
		}
	}

	return max
}

func MinIntInList(l []int) int {
	min := math.MaxInt

	for _, e := range l {
		if e < min {
			min = e
		}
	}

	return min
}
