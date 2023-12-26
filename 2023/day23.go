package main

import (
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type Move LightMove

func (s solver) Day23p1(input aoc.Input) string {
	grid := ToGrid(input)
	start := aoc.Point{X: 0, Y: 1}
	path := GetLongestHike(
		Path{cost: 0, points: []aoc.Point{start}},
		aoc.Point{X: len(grid) - 1, Y: len(grid[0]) - 2},
		grid,
	)
	return aoc.ResultI(path.cost)
}

func (s solver) Day23p2(input aoc.Input) string {
	grid := ToGrid(input)
	for x := range grid {
		for y := range grid[0] {
			if grid[x][y] == '#' {
				continue
			}

			grid[x][y] = '.'
		}
	}

	start := aoc.Point{X: 0, Y: 1}
	// does not end
	path := GetLongestHike(
		Path{cost: 0, points: []aoc.Point{start}},
		aoc.Point{X: len(grid) - 1, Y: len(grid[0]) - 2},
		grid,
	)
	return aoc.ResultI(path.cost)
}

func GetLongestHike(path Path, end aoc.Point, grid [][]byte) Path {
	point := path.points[len(path.points)-1]
	if point.X == end.X && point.Y == end.Y {
		return path
	}

	var nextPoints []aoc.Point
	switch grid[point.X][point.Y] {
	case '#':
		// we can't walk on trees
		return Path{}
	case '>':
		p := point
		p.Y++
		nextPoints = append(nextPoints, p)
	case '<':
		p := point
		p.Y--
		nextPoints = append(nextPoints, p)
	case '^':
		p := point
		p.X--
		nextPoints = append(nextPoints, p)
	case 'v':
		p := point
		p.X++
		nextPoints = append(nextPoints, p)
	default:
		p1, p2, p3, p4 := point, point, point, point
		p1.X--
		p2.X++
		p3.Y--
		p4.Y++
		nextPoints = append(nextPoints, p1, p2, p3, p4)
	}

	bestPath := Path{cost: path.cost}
	for _, np := range nextPoints {
		if np.X < 0 || np.Y < 0 || np.X >= len(grid) || np.Y >= len(grid[0]) {
			continue
		} else if grid[np.X][np.Y] == '#' {
			continue
		}

		if slices.Contains(path.points, np) {
			continue
		}

		nextPath := Path{
			cost:   path.cost + 1,
			points: append([]aoc.Point{}, path.points...),
		}
		nextPath.points = append(nextPath.points, np)
		resultPath := GetLongestHike(nextPath, end, grid)
		if len(resultPath.points) > 0 && resultPath.cost > bestPath.cost {
			bestPath = resultPath
		}
	}

	if len(bestPath.points) == 0 {
		bestPath.cost = 0
	}
	return bestPath
}
