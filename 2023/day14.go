package main

import (
	"log"
	"time"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day14p1(input aoc.Input) string {
	grid := ToGrid(input)
	load := tiltGridNorth(grid)
	return aoc.ResultI(load)
}

func (s solver) Day14p2(input aoc.Input) string {
	grid := ToGrid(input)
	var load int

	t := time.Now()
	for i := 0; i < 1_000_000_000; i++ {
		if i%10_000 == 0 {
			log.Print(i, time.Since(t))
		}
		tiltGridNorth(grid)
		RotateGridClockwise(grid)
		tiltGridNorth(grid)
		RotateGridClockwise(grid)
		tiltGridNorth(grid)
		RotateGridClockwise(grid)
		load = tiltGridNorth(grid)
		RotateGridClockwise(grid)
	}

	return aoc.ResultI(load)
}

func tiltGridNorth(grid [][]byte) int {
	// for each column (y), list of available row (x)
	available := make(map[int][]int)
	load := 0

	for x := range grid {
		for y := range grid[0] {
			switch grid[x][y] {
			case '.':
				available[y] = append(available[y], x)
			case '#':
				available[y] = available[y][:0]
			case 'O':
				if len(available[y]) > 0 {
					newX := available[y][0]
					available[y] = available[y][1:]

					grid[newX][y] = 'O'
					grid[x][y] = '.'

					available[y] = append(available[y], x)

					load += len(grid) - newX
				} else {
					load += len(grid) - x
				}
			}
		}
	}

	return load
}

func RotateGridClockwise(grid [][]byte) {
	maxCoordinate := len(grid) - 1

	for x := 0; x < len(grid)/2; x++ {
		for y := x; y < maxCoordinate-x; y++ {
			xy := grid[x][y]
			grid[x][y] = grid[maxCoordinate-y][x]
			grid[maxCoordinate-y][x] = grid[maxCoordinate-x][maxCoordinate-y]
			grid[maxCoordinate-x][maxCoordinate-y] = grid[y][maxCoordinate-x]
			grid[y][maxCoordinate-x] = xy
		}
	}
}
