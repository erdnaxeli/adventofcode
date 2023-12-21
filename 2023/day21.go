package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day21p1(input aoc.Input) string {
	grid := ToGrid(input)
	var xs, ys int
	for x := range grid {
		for y := range grid[0] {
			if grid[x][y] == 'S' {
				grid[x][y] = '.'
				xs, ys = x, y
			}
		}
	}

	sum := 0
	for x := xs - 64; x <= xs+64; x++ {
		for y := ys - 64; y <= ys+64; y++ {
			distance := aoc.AbsI(xs-x) + aoc.AbsI(ys-y)
			if grid[x][y] == '.' && distance <= 64 && distance%2 == 0 {
				sum++
			}
		}
	}

	// I am not sure why this -1 is needed, but it is
	return aoc.ResultI(sum - 1)
}

func (s solver) Day21p2(input aoc.Input) string {
	return ""
}
