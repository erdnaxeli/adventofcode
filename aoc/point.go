package aoc

type Point struct {
	X int
	Y int
	Z int
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
// It supposes a grid where X goes from west to east, and Y from north to south.
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
