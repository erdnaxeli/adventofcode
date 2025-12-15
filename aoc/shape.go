package aoc

import (
	"iter"
)

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
//
// The shape must be _streched_ for this to work, this means adjacent borders must
// be separated by empty spaces.
//
// > ######
// > #.##.#
// > #....#
// > ######

// is an invalid shape, but
//
// > ###.###
// > #.###.#
// > #.....#
// > #######
//
// is a correct one.
//
// This is due to the fact that if not stretched, multiple shape can be reprsented
// the same way.
//
// > ┌────┐ ┌─┐┌─┐
// > │┌──┐│ └┐└┘┌┘
// > ││ X││ ┌┘ X└┐
// > │└┐┌┘│ │┌┐┌┐│
// > └─┘└─┘ └┘└┘└┘
//
// Are both represented by:
//
// > ######
// > ######
// > ## X##
// > ######
// > ######
//
// But in the first case X is outside the shape, will in the second case it is inside.
//
// If shapes are stretched, they have a distinct representations:
//
// > #########   #####..####
// > #.......#   #...#..#..#
// > #.#####.#   ##..####.##
// > #.#...#.#   .#.......#.
// > #.##.##.#   ##.......##
// > #..#.#..#   #.........#
// > ####.####   #.###.###.#
// >             ###.###.###
//
// There are multiple was to stretch a shape.
func (s Shape) Contains(p Point) bool {
	if s.points.Contains(p) {
		return true
	}

	pointsToPX := func(yield func(Point) bool) {
		for x := s.xMin; x <= p.X; x++ {
			if !yield(XY(x, p.Y)) {
				return
			}
		}
	}
	pointsFromPX := func(yield func(Point) bool) {
		for x := p.X + 1; x <= s.xMax+1; x++ {
			if !yield(XY(x, p.Y)) {
				return
			}
		}
	}
	pointsToPY := func(yield func(Point) bool) {
		for y := s.yMin; y <= p.Y; y++ {
			if !yield(XY(p.X, y)) {
				return
			}
		}
	}
	pointsFromPY := func(yield func(Point) bool) {
		for y := p.Y + 1; y <= s.yMax+1; y++ {
			if !yield(XY(p.X, y)) {
				return
			}
		}
	}

	bordersXTo := s.countBorders(pointsToPX)
	if bordersXTo == 0 {
		return false
	}
	bordersXFrom := s.countBorders(pointsFromPX)
	if bordersXFrom == 0 {
		return false
	}
	bordersYTo := s.countBorders(pointsToPY)
	if bordersYTo == 0 {
		return false
	}
	bordersYFrom := s.countBorders(pointsFromPY)
	if bordersYFrom == 0 {
		return false
	}

	return (bordersXTo%2 == 1 || bordersXFrom%2 == 1) && (bordersYTo%2 == 1 || bordersYFrom%2 == 1)
}

func (s Shape) countBorders(points iter.Seq[Point]) int {
	onBorder := false
	count := 0

	for p := range points {
		if !onBorder {
			if s.points.Contains(p) {
				onBorder = true
			}
		} else if !s.points.Contains(p) {
			onBorder = false
			count++
		}
	}

	return count
}
