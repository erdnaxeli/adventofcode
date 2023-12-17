package main

import (
	"log"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day14p1(input aoc.Input) string {
	grid := ToGrid(input)
	load := tiltGridNorth(grid)
	return aoc.ResultI(load)
}

func (s solver) Day14p2(input aoc.Input) string {
	grid := ToGrid(input)
	var cycle int
	var cache [][][]byte
	for step := 1; step <= 1_000_000_000; step++ {
		var g [][]byte
		if cycle == 0 {
			g = make([][]byte, len(grid))
			for x := range grid {
				g[x] = make([]byte, len(grid[0]))
				for y := range grid[0] {
					g[x][y] = grid[x][y]
				}
			}
		} else {
			g = grid
		}

		tiltGridNorth(g)
		RotateGridClockwise(g)
		tiltGridNorth(g)
		RotateGridClockwise(g)
		tiltGridNorth(g)
		RotateGridClockwise(g)
		tiltGridNorth(g)
		RotateGridClockwise(g)

		if cycle == 0 {
			cache = append(cache, grid)
			grid = g

			for i, g := range cache[:len(cache)-1] {
				equal := true
				for x := range grid {
					for y := range grid {
						if grid[x][y] != g[x][y] {
							equal = false
							break
						}
					}

					if !equal {
						break
					}
				}

				if equal {
					cycle = step - i
					log.Printf("Cycle of size %d found after %d steps", i, step)
					step += ((1_000_000_000-step)/cycle)*cycle - cycle
					log.Printf("Jumping to step %d", step)
					break
				}
			}
		}
	}

	load := getGridLoad(grid)
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

func getGridLoad(grid [][]byte) int {
	load := 0
	for x := range grid {
		for y := range grid {
			if grid[x][y] == 'O' {
				load += len(grid) - x
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
