package main

import (
	"math"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day11p1(input aoc.Input) string {
	grid := ToGrid(input)
	lenX := len(grid)
	for x := 0; x < lenX; x++ {
		isEmpty := true
		for y := range grid[0] {
			if grid[x][y] != '.' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			grid = AddGridRow(grid, x+1)
			lenX++
			x++
		}
	}

	lenY := len(grid[0])
	for y := 0; y < lenY; y++ {
		isEmpty := true
		for x := range grid {
			if grid[x][y] != '.' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			grid = AddGridColumn(grid, y+1)
			lenY++
			y++
		}
	}

	var coordinates []aoc.Point
	for x := range grid {
		for y := range grid[0] {
			if grid[x][y] == '#' {
				coordinates = append(coordinates, aoc.Point{X: x, Y: y})
			}
		}
	}

	sum := 0
	for i, p1 := range coordinates {
		for j := i; j < len(coordinates); j++ {
			p2 := coordinates[j]
			distance := math.Abs(float64(p2.X-p1.X)) + math.Abs(float64(p2.Y-p1.Y))
			sum += int(distance)
		}
	}

	return aoc.ResultI(sum)
}

func (s solver) Day11p2(input aoc.Input) string {
	return ""
}

func ToGrid(input aoc.Input) [][]byte {
	var grid [][]byte
	for _, line := range input.ToStringSlice() {
		grid = append(grid, []byte(line))
	}

	return grid
}

func AddGridRow(grid [][]byte, x int) [][]byte {
	grid = append(grid[:x+1], grid[x:]...)
	grid[x] = make([]byte, len(grid[0]))
	for y := range grid[x] {
		grid[x][y] = '.'
	}

	return grid
}

func AddGridColumn(grid [][]byte, y int) [][]byte {
	for x := range grid {
		grid[x] = append(grid[x][:y+1], grid[x][y:]...)
		grid[x][y] = '.'
	}

	return grid
}
