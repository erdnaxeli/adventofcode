package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day1p1(input aoc.Input) string {
	position := aoc.Point{}
	direction := aoc.Point{0, 1, 0}
	for _, elt := range input.SingleLine().ToStringSlice() {
		d, n := elt.At(0), elt.From(1).Atoi()
		switch d {
		case 'R':
			direction = direction.RotateRightZ()
		case 'L':
			direction = direction.RotateLeftZ()
		}

		position = position.Add(direction.MultI(n))
	}

	return aoc.ResultI(aoc.AbsI(position.X) + aoc.AbsI(position.Y))
}

func (s solver) Day1p2(input aoc.Input) string {
	positions := make(map[aoc.Point]bool)
	position := aoc.Point{}
	positions[position] = true
	direction := aoc.Point{0, 1, 0}
	for _, elt := range input.SingleLine().ToStringSlice() {
		d, n := elt.At(0), elt.From(1).Atoi()
		switch d {
		case 'R':
			direction = direction.RotateRightZ()
		case 'L':
			direction = direction.RotateLeftZ()
		}

		for i := 0; i < n; i++ {
			position = position.Add(direction)
			if positions[position] {
				return aoc.ResultI(aoc.AbsI(position.X) + aoc.AbsI(position.Y))
			}

			positions[position] = true
		}

	}

	return "fail"
}
