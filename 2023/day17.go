package main

import (
	"errors"
	"log"
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type Path struct {
	cost   int
	points []aoc.Point
}

func (s solver) Day17p1(input aoc.Input) string {
	grid := ToGrid(input)
	bestPath, _ := FindCruciblePath(
		grid,
		Path{},
		aoc.Point{X: 0, Y: 0},
		0,
		make([]Direction, 0, 4),
	)

	return aoc.ResultI(bestPath.cost)
}

func (s solver) Day17p2(input aoc.Input) string {
	return ""
}

func FindCruciblePath(grid [][]byte, path Path, point aoc.Point, maxCost int, direction []Direction) (Path, error) {
	log.Printf("%+v", path)
	if len(direction) >= 2 && direction[len(direction)-1] != direction[len(direction)-2] {
		direction = direction[:0]
	}

	if len(direction) == 4 {
		return Path{}, errors.New("too many steps in the same direction")
	} else if point.X < 0 || point.Y < 0 || point.X >= len(grid) || point.Y >= len(grid[0]) {
		return Path{}, errors.New("invalid point")
	} else if slices.Contains(path.points, point) {
		return Path{}, errors.New("looping path")
	}

	path.points = append(path.points, point)
	path.cost += aoc.Atoi(string(grid[point.X][point.Y]))

	if point.X == len(grid)-1 && point.Y == len(grid[0])-1 {
		return path, nil
	}

	p1 := aoc.Point{X: point.X + 1, Y: point.Y}
	p2 := aoc.Point{X: point.X - 1, Y: point.Y}
	p3 := aoc.Point{X: point.X, Y: point.Y + 1}
	p4 := aoc.Point{X: point.X, Y: point.Y - 1}

	var bestPath Path

	path1, err := FindCruciblePath(grid, path, p1, maxCost, append(direction, SOUTH))
	if err != nil {
		if maxCost == 0 || path1.cost < maxCost {
			maxCost = path1.cost
			bestPath = path1
		}
	}

	path2, err := FindCruciblePath(grid, path, p2, maxCost, append(direction, NORTH))
	if err != nil {
		if path2.cost < maxCost {
			maxCost = path2.cost
			bestPath = path2
		}
	}

	path3, err := FindCruciblePath(grid, path, p3, maxCost, append(direction, EAST))
	if err != nil {
		if path3.cost < maxCost {
			maxCost = path3.cost
			bestPath = path3
		}
	}

	path4, err := FindCruciblePath(grid, path, p4, maxCost, append(direction, WEST))
	if err != nil {
		if path4.cost < maxCost {
			bestPath = path4
		}
	}

	return bestPath, nil
}
