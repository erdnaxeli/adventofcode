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
