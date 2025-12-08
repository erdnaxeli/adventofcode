package aoc

import (
	"math"
)

type Point struct {
	X int
	Y int
	Z int
}

func NewPointXY(x int, y int) Point {
	return NewPointXYZ(x, y, 0)
}

func NewPointXYZ(x int, y int, z int) Point {
	return Point{X: x, Y: y, Z: z}
}

func (p Point) Add(o Point) Point {
	return Point{
		p.X + o.X,
		p.Y + o.Y,
		p.Z + o.Z,
	}
}

func (p Point) MultI(i int) Point {
	return Point{
		p.X * i,
		p.Y * i,
		p.Z * i,
	}
}

func (p Point) RotateRightZ() Point {
	return Point{
		p.Y,
		-p.X,
		p.Z,
	}
}

func (p Point) RotateLeftZ() Point {
	return Point{
		-p.Y,
		p.X,
		p.Z,
	}
}

// MoveNorth return a new point one step to the north.
//
// It supposes a grid where X goes from north to south, and Y from east to west.
func (p Point) MoveNorth() Point {
	return Point{p.X - 1, p.Y, p.Z}
}

func (p Point) MoveWest() Point {
	return Point{p.X, p.Y - 1, p.Z}
}

func (p Point) MoveSouth() Point {
	return Point{p.X + 1, p.Y, p.Z}
}

func (p Point) MoveEast() Point {
	return Point{p.X, p.Y + 1, p.Z}
}

func (p1 Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) +
		math.Pow(float64(p1.Y-p2.Y), 2) +
		math.Pow(float64(p1.Z-p2.Z), 2))
}
