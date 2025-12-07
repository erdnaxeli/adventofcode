package main

import (
	"maps"
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day7p1(input aoc.Input) string {
	grid := input.ToGrid()
	beams := make(map[aoc.Point]struct{})

	// find first position
	for p, e := range grid.IterAllLine(0) {
		if e == 'S' {
			beams[p] = struct{}{}
			break
		}
	}

	splitsCount := 0

	for x := 0; x < grid.MaxX(); x++ {
		newBeams := make(map[aoc.Point]struct{})

		for p := range beams {
			if grid.AtXY(p.X+1, p.Y) == '^' {
				if p.Y > 0 {
					newBeams[aoc.Point{X: p.X + 1, Y: p.Y - 1}] = struct{}{}
				}

				if p.Y < grid.MaxY() {
					newBeams[aoc.Point{X: p.X + 1, Y: p.Y + 1}] = struct{}{}
				}

				splitsCount++
			} else {
				newBeams[p.MoveSouth()] = struct{}{}
			}
		}

		beams = newBeams
	}

	return aoc.ResultI(splitsCount)
}

func (s solver) Day7p2(input aoc.Input) string {
	grid := input.ToGrid()
	beams := make(map[aoc.Point]int)

	// find first position
	for p, e := range grid.IterAllLine(0) {
		if e == 'S' {
			beams[p] = 1
			break
		}
	}

	for x := 0; x < grid.MaxX(); x++ {
		newBeams := make(map[aoc.Point]int)

		for p := range beams {
			if grid.AtXY(p.X+1, p.Y) == '^' {
				if p.Y > 0 {
					pp := aoc.Point{X: p.X + 1, Y: p.Y - 1}
					newBeams[pp] = beams[p] + newBeams[pp]
				}

				if p.Y < grid.MaxY() {
					pp := aoc.Point{X: p.X + 1, Y: p.Y + 1}
					newBeams[pp] = beams[p] + newBeams[pp]
				}
			} else {
				pp := p.MoveSouth()
				newBeams[pp] = beams[p] + newBeams[pp]
			}
		}

		beams = newBeams
	}

	return aoc.ResultI(aoc.SumSlice(slices.Collect(maps.Values(beams))))
}
