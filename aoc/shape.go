package aoc

type Shape struct {
	points Set[Point]
	xMin   int
	xMax   int
	yMin   int
	yMax   int
}

// NewShape creates a shape with the given points.
//
// The points are expected to be side by side, either horizontally or vertically,
// and the shape is expected to be closed.
func NewShape(points Set[Point]) Shape {
	p := points.First()
	xMin, xMax, yMin, yMax := p.X, p.X, p.Y, p.Y

	for p := range points.Values() {
		if p.X < xMin {
			xMin = p.X
		}

		if p.X > xMax {
			xMax = p.X
		}

		if p.Y < yMin {
			yMin = p.Y
		}

		if p.Y > yMax {
			yMax = p.Y
		}
	}

	return Shape{
		points,
		xMin,
		xMax,
		yMin,
		yMax,
	}
}

// Contains returns true if the Shape s contains the Point p.
//
// If p is part of the points defining the shape, it is considered as contained.
func (s Shape) Contains(p Point) bool {
	if s.points.Contains(p) {
		return true
	}

	/*
		dLeft := p.Y - s.yMin
		dTop := p.X - s.xMin
		dRight := s.yMax - p.Y
		dDown := s.xMax - p.X

		switch min(dLeft, dTop, dRight, dDown) {
		case dLeft:

		case dTop:
		case dRight:
		case dDown:
		}
	*/

	onBorder := false
	borderCrossingCountX := 0
	for x := s.xMin; x <= p.X; x++ {
		if !onBorder {
			if s.points.Contains(NewPointXY(x, p.Y)) {
				onBorder = true
			}
		} else if !s.points.Contains(NewPointXY(x, p.Y)) {
			onBorder = false
			borderCrossingCountX++
		}
	}

	borderCrossingCountY := 0
	for y := s.yMin; y <= p.Y; y++ {
		if !onBorder {
			if s.points.Contains(NewPointXY(p.X, y)) {
				onBorder = true
			}
		} else if !s.points.Contains(NewPointXY(p.X, y)) {
			onBorder = false
			borderCrossingCountY++
		}
	}

	return borderCrossingCountX%2 == 1 && borderCrossingCountY%2 == 1
}
