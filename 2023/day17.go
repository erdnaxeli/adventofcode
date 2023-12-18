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
	grid := ToGrid(input)
	grid[0][0] = '0'
	path, _, _ := FindCruciblePath2(
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

	path.points = append(path.points, point)
	path.cost += aoc.Atoi(string(grid[point.X][point.Y]))
	if maxCost > 0 && path.cost >= maxCost {
		return Path{}, cache, errors.New("path not worth it")
	}

	// we use Z to store consecutive forward movements
	var z int
	if len(direction) == 0 {
		z = 0
	} else {
		z = len(direction)*10 + int(direction[0])
	}
	cachePoint := aoc.Point{X: point.X, Y: point.Y, Z: z}
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
			bestPath = Path{
				cost:   path4.cost,
				points: append([]aoc.Point{}, path4.points...),
			}
		}
	}

	if len(bestPath.points) == 0 {
		return Path{}, cache, errors.New("no path found")
	}

	return bestPath, cache, nil
}

// grid is the grid
// path is the current path to get to point
// point is the new point being testing
// direction is the list of same directions we took to get to point
// cache is the list of minimal cost to reach a point
func FindCruciblePath2(
	grid [][]byte,
	path Path,
	point aoc.Point,
	maxCost int,
	direction []Direction,
	cache map[aoc.Point]int,
) (Path, map[aoc.Point]int, error) {
	if len(direction) <= 4 {
		for i := 1; i < len(direction); i++ {
			if direction[i] != direction[0] {
				return Path{}, cache, errors.New("must move in the same direction 4 times before turning")
			}
		}
	} else if len(direction) >= 2 && direction[len(direction)-1] != direction[len(direction)-2] {
		direction = direction[len(direction)-1:]
	} else if len(direction) > 10 {
		return Path{}, cache, errors.New("too many steps in the same direction")
	}

	if point.X < 0 || point.Y < 0 || point.X >= len(grid) || point.Y >= len(grid[0]) {
		return Path{}, cache, errors.New("invalid point")
	} else if slices.Contains(path.points, point) {
		return Path{}, cache, errors.New("looping path")
	}

	path.points = append(path.points, point)
	path.cost += aoc.Atoi(string(grid[point.X][point.Y]))
	if maxCost > 0 && path.cost >= maxCost {
		return Path{}, cache, errors.New("path not worth it")
	}

	// we use Z to store consecutive forward movements
	var z int
	if len(direction) == 0 {
		z = 0
	} else {
		z = len(direction)*10 + int(direction[0])
	}
	cachePoint := aoc.Point{X: point.X, Y: point.Y, Z: z}
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
	path1, cache, err := FindCruciblePath2(grid, path, p1, maxCost, append(direction, SOUTH), cache)
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
	path2, cache, err := FindCruciblePath2(grid, path, p2, maxCost, append(direction, NORTH), cache)
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
	path3, cache, err := FindCruciblePath2(grid, path, p3, maxCost, append(direction, EAST), cache)
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
	path4, cache, err := FindCruciblePath2(grid, path, p4, maxCost, append(direction, WEST), cache)
	if err == nil {
		if maxCost == 0 || path4.cost < maxCost {
			bestPath = Path{
				cost:   path4.cost,
				points: append([]aoc.Point{}, path4.points...),
			}
		}
	}

	if len(bestPath.points) == 0 {
		return Path{}, cache, errors.New("no path found")
	}

	return bestPath, cache, nil
}
