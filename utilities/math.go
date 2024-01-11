package utilities

func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func AbsoluteDifference(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

type Size2D struct {
	Width  int
	Height int
}

func NewSize2D(Width int, Height int) Size2D {
	return Size2D{Width, Height}
}

type Point2D struct {
	X int
	Y int
}

func NewPoint2D(X int, Y int) Point2D {
	return Point2D{X, Y}
}

func ManhattanDistance(p1, p2 Point2D) int {
	return AbsoluteDifference(p1.X, p2.X) + AbsoluteDifference(p1.Y, p2.Y)
}

func IsLeft(start Point2D, end Point2D, point Point2D) int {
	val := (end.X-start.X)*(point.Y-start.Y) - (point.X-start.X)*(end.Y-start.Y)
	return val
}

func PointInPolyCrossing(point Point2D, vertices []Point2D) bool {
	crossing := 0

	for i := 0; i < len(vertices)-1; i++ {
		if (vertices[i].Y <= point.Y && vertices[i+1].Y > point.Y) ||
			(vertices[i].Y > point.Y && vertices[i+1].Y <= point.Y) {
			vt := (point.Y - vertices[i].Y) / (vertices[i+1].Y - vertices[i].Y)
			if point.X < vertices[i].X+vt*(vertices[i+1].X-vertices[i].X) {
				crossing++
			}
		}

	}

	return (crossing & 1) != 0
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
