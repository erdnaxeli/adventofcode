package main

import "github.com/erdnaxeli/adventofcode/aoc"

func (s solver) Day1p1(input aoc.Input) string {
	// Direct copy of the python implementation.
	zeros := 0
	position := 50

	for _, line := range input.ToStringSlice() {
		direction := line.At(0)
		count := line.From(1).Atoi()

		if direction == 'L' {
			position -= count
		} else {
			position += count
		}

		if position%100 == 0 {
			zeros++
		}
	}

	return aoc.ResultI(zeros)
}

func (s solver) Day1p2(input aoc.Input) string {
	return ""
}
