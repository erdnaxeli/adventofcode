package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
	"golang.org/x/exp/maps"
)

type Direction int

const (
	NORTH = iota
	WEST
	SOUTH
	EAST
)

type LightMove struct {
	Point aoc.Point
	From  Direction
}

func (s solver) Day16p1(input aoc.Input) string {
	grid := ToGrid(input)
	return aoc.ResultI(
		getEnergizedTiles(
			grid,
			LightMove{Point: aoc.Point{X: 0, Y: 0}, From: WEST},
		),
	)
}

func (s solver) Day16p2(input aoc.Input) string {
	grid := ToGrid(input)
	maxEnergized := 0

	for x := range grid {
		maxEnergized = max(
			maxEnergized,
			getEnergizedTiles(
				grid,
				LightMove{Point: aoc.Point{X: x, Y: 0}, From: WEST},
			),
			getEnergizedTiles(
				grid,
				LightMove{Point: aoc.Point{X: x, Y: len(grid[0]) - 1}, From: EAST},
			),
		)
	}

	for y := range grid[0] {
		maxEnergized = max(
			maxEnergized,
			getEnergizedTiles(
				grid,
				LightMove{Point: aoc.Point{X: 0, Y: y}, From: NORTH},
			),
			getEnergizedTiles(
				grid,
				LightMove{Point: aoc.Point{X: len(grid) - 1, Y: y}, From: SOUTH},
			),
		)
	}

	return aoc.ResultI(maxEnergized)
}

func getEnergizedTiles(grid [][]byte, move LightMove) int {
	queue := []LightMove{move}
	visited := make(map[LightMove]bool)

	for len(queue) > 0 {
		m := queue[0]
		queue = queue[1:]

		if m.Point.X < 0 || m.Point.Y < 0 || m.Point.X >= len(grid) || m.Point.Y >= len(grid[0]) {
			continue
		}

		if _, ok := visited[m]; ok {
			continue
		}

		visited[m] = true

		switch grid[m.Point.X][m.Point.Y] {
		case '|':
			if m.From == NORTH {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X + 1, Y: m.Point.Y},
						From:  NORTH,
					},
				)
			} else if m.From == SOUTH {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X - 1, Y: m.Point.Y},
						From:  SOUTH,
					},
				)
			} else {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X - 1, Y: m.Point.Y},
						From:  SOUTH,
					},
					LightMove{
						Point: aoc.Point{X: m.Point.X + 1, Y: m.Point.Y},
						From:  NORTH,
					},
				)
			}
		case '-':
			if m.From == WEST {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X, Y: m.Point.Y + 1},
						From:  WEST,
					},
				)
			} else if m.From == EAST {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X, Y: m.Point.Y - 1},
						From:  EAST,
					},
				)
			} else {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X, Y: m.Point.Y + 1},
						From:  WEST,
					},
					LightMove{
						Point: aoc.Point{X: m.Point.X, Y: m.Point.Y - 1},
						From:  EAST,
					},
				)
			}
		case '\\':
			if m.From == NORTH {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X, Y: m.Point.Y + 1},
						From:  WEST,
					},
				)
			} else if m.From == WEST {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X + 1, Y: m.Point.Y},
						From:  NORTH,
					},
				)
			} else if m.From == SOUTH {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X, Y: m.Point.Y - 1},
						From:  EAST,
					},
				)
			} else if m.From == EAST {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X - 1, Y: m.Point.Y},
						From:  SOUTH,
					},
				)
			}
		case '/':
			if m.From == NORTH {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X, Y: m.Point.Y - 1},
						From:  EAST,
					},
				)
			} else if m.From == WEST {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X - 1, Y: m.Point.Y},
						From:  SOUTH,
					},
				)
			} else if m.From == SOUTH {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X, Y: m.Point.Y + 1},
						From:  WEST,
					},
				)
			} else if m.From == EAST {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X + 1, Y: m.Point.Y},
						From:  NORTH,
					},
				)
			}
		default:
			if m.From == NORTH {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X + 1, Y: m.Point.Y},
						From:  NORTH,
					},
				)
			} else if m.From == WEST {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X, Y: m.Point.Y + 1},
						From:  WEST,
					},
				)
			} else if m.From == SOUTH {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X - 1, Y: m.Point.Y},
						From:  SOUTH,
					},
				)
			} else if m.From == EAST {
				queue = append(
					queue,
					LightMove{
						Point: aoc.Point{X: m.Point.X, Y: m.Point.Y - 1},
						From:  EAST,
					},
				)
			}
		}
	}

	energized := make(map[aoc.Point]bool)
	for lm := range visited {
		energized[lm.Point] = true
	}

	return len(maps.Keys(energized))
}
