package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day4p1(input aoc.Input) string {
	grid := input.ToGrid()
	accessibleCount := len(findRollsAccessibles(grid))
	return aoc.ResultI(accessibleCount)
}

func (s solver) Day4p2(input aoc.Input) string {
	grid := input.ToGrid()
	accessibles := findRollsAccessibles(grid)
	accessiblesCount := len(accessibles)
	accessiblesTotalCount := accessiblesCount

	for accessiblesCount != 0 {
		for _, p := range accessibles {
			grid.Set(p, '.')
		}

		accessibles = findRollsAccessibles(grid)
		accessiblesCount = len(accessibles)
		accessiblesTotalCount += accessiblesCount
	}

	return aoc.ResultI(accessiblesTotalCount)
}

func findRollsAccessibles(grid aoc.Grid) []aoc.Point {
	var accessibles []aoc.Point

	for p, v := range grid.Points() {
		if v != '@' {
			continue
		}

		count := 0
		for _, vv := range grid.Around(p) {
			if vv == '@' {
				count++
			}

			if count >= 4 {
				break
			}
		}

		if count < 4 {
			accessibles = append(accessibles, p)
		}
	}

	return accessibles
}
