package main

import (
	"errors"
	"fmt"
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type Path struct {
	cost   int
	points []aoc.Point
}

func (s solver) Day17p1(input aoc.Input) string {
	grid := ToGrid(input)
	grid[0][0] = '0'
	path, _, _ := FindCruciblePath(
		grid,
		Path{},
		aoc.Point{X: 0, Y: 0},
		0,
		make([]Direction, 0, 4),
		make(map[aoc.Point]int),
	)

	for x := range grid {
		for y := range grid[0] {
			if slices.Contains(path.points, aoc.Point{X: x, Y: y}) {
				fmt.Print("\033[0;31m")
				fmt.Print(string(grid[x][y]))
				fmt.Print("\033[0m")
			} else {
				fmt.Print(string(grid[x][y]))
			}
		}
		fmt.Println()
	}
	return aoc.ResultI(path.cost)
}

func (s solver) Day17p2(input aoc.Input) string {
	return ""
}

// grid is the grid
// path is the current path to get to point
// point is the new point being testing
// direction is the list of same directions we took to get to point
// cache is the list of minimal cost to reach a point
func FindCruciblePath(
	grid [][]byte,
	path Path,
	point aoc.Point,
	maxCost int,
	direction []Direction,
	cache map[aoc.Point]int,
) (Path, map[aoc.Point]int, error) {
	// direction = append([]Direction{}, direction...)

	/*
		if i >= j {
			for x := range grid {
				for y := range grid[0] {
					if x == point.X && y == point.Y {
						fmt.Print("\033[0;33m")
						fmt.Print(string(grid[x][y]))
						fmt.Print("\033[0m")
					} else if slices.Contains(path.points, aoc.Point{X: x, Y: y}) {
						fmt.Print("\033[0;31m")
						fmt.Print(string(grid[x][y]))
						fmt.Print("\033[0m")
					} else if _, ok := cache[aoc.Point{X: x, Y: y}]; ok {
						fmt.Print("\033[0;32m")
						fmt.Print(string(grid[x][y]))
						fmt.Print("\033[0m")
					} else {
						fmt.Print(string(grid[x][y]))
					}
				}
				fmt.Println()
			}
			fmt.Printf("max score: %d; current score: %d\n", maxCost, path.cost)
		}
	*/

	if len(direction) >= 2 && direction[len(direction)-1] != direction[len(direction)-2] {
		direction = direction[len(direction)-1:]
	} else if len(direction) > 3 {
		return Path{}, cache, errors.New("too many steps in the same direction")
	}

	if point.X < 0 || point.Y < 0 || point.X >= len(grid) || point.Y >= len(grid[0]) {
		return Path{}, cache, errors.New("invalid point")
	} else if slices.Contains(path.points, point) {
		return Path{}, cache, errors.New("looping path")
	}

	// path.points = append([]aoc.Point{}, path.points...)
	path.points = append(path.points, point)
	path.cost += aoc.Atoi(string(grid[point.X][point.Y]))
	if maxCost > 0 && path.cost >= maxCost {
		return Path{}, cache, errors.New("path not worth it")
	}

	// we use Z to store the number of consecutive forward movements
	cachePoint := aoc.Point{X: point.X, Y: point.Y, Z: len(direction)}
	if minCost, ok := cache[cachePoint]; ok && path.cost >= minCost {
		return Path{}, cache, errors.New("we already went to this point with a better path")
	} else {
		cache[cachePoint] = path.cost
	}

	if point.X == len(grid)-1 && point.Y == len(grid[0])-1 {
		return path, cache, nil
	}

	var bestPath Path

	p1 := aoc.Point{X: point.X + 1, Y: point.Y}
	path1, cache, err := FindCruciblePath(grid, path, p1, maxCost, append(direction, SOUTH), cache)
	if err == nil {
		if maxCost == 0 || path1.cost < maxCost {
			maxCost = path1.cost
			bestPath = Path{
				cost:   path1.cost,
				points: append([]aoc.Point{}, path1.points...),
			}
		}
	}

	p2 := aoc.Point{X: point.X - 1, Y: point.Y}
	path2, cache, err := FindCruciblePath(grid, path, p2, maxCost, append(direction, NORTH), cache)
	if err == nil {
		if maxCost == 0 || path2.cost < maxCost {
			maxCost = path2.cost
			bestPath = Path{
				cost:   path2.cost,
				points: append([]aoc.Point{}, path2.points...),
			}
		}
	}

	p3 := aoc.Point{X: point.X, Y: point.Y + 1}
	path3, cache, err := FindCruciblePath(grid, path, p3, maxCost, append(direction, EAST), cache)
	if err == nil {
		if maxCost == 0 || path3.cost < maxCost {
			maxCost = path3.cost
			bestPath = Path{
				cost:   path3.cost,
				points: append([]aoc.Point{}, path3.points...),
			}
		}
	}

	p4 := aoc.Point{X: point.X, Y: point.Y - 1}
	path4, cache, err := FindCruciblePath(grid, path, p4, maxCost, append(direction, WEST), cache)
	if err == nil {
		if maxCost == 0 || path4.cost < maxCost {
			bestPath = path4
		}
	}

	if len(bestPath.points) == 0 {
		return Path{}, cache, errors.New("no path found")
	}

	// log.Print(bestPath.cost, bestI)
	return bestPath, cache, nil
}
