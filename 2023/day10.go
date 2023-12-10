package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day10p1(input aoc.Input) string {
	grid := input.ToStringSlice()
	xs, ys := getPipeStart(grid)

	var x, y int
	var move byte
	if xs > 0 && (grid[xs-1][ys] == '|' || grid[xs-1][ys] == 'F' || grid[xs-1][ys] == '7') {
		x, y, move = xs-1, ys, 'N'
	} else if ys > 0 && (grid[xs][ys-1] == '-' || grid[xs][ys-1] == 'F' || grid[xs][ys-1] == 'L') {
		x, y, move = xs, ys-1, 'W'
	} else if xs < len(grid)-1 && (grid[xs+1][ys] == '|' || grid[xs+1][ys] == 'L' || grid[xs+1][ys] == 'J') {
		x, y, move = xs+1, ys, 'S'
	} else if ys < len(grid[0]) && (grid[xs][ys+1] == '-' || grid[xs][ys+1] == '7' || grid[xs][ys+1] == 'J') {
		x, y, move = xs, ys+1, 'E'
	}

	steps := 1
	for {
		switch grid[x][y] {
		case '|':
			if move == 'N' {
				x--
			} else {
				x++
			}
		case 'L':
			if move == 'S' {
				move = 'E'
				y++
			} else {
				move = 'N'
				x--
			}
		case '-':
			if move == 'E' {
				y++
			} else {
				y--
			}
		case 'J':
			if move == 'E' {
				move = 'N'
				x--
			} else {
				move = 'W'
				y--
			}
		case '7':
			if move == 'N' {
				move = 'W'
				y--
			} else {
				move = 'S'
				x++
			}
		case 'F':
			if move == 'W' {
				move = 'S'
				x++
			} else {
				move = 'E'
				y++
			}
		}

		steps++
		if grid[x][y] == 'S' {
			break
		}
	}

	return aoc.ResultI(steps / 2)
}

func (s solver) Day10p2(input aoc.Input) string {
	return ""
}

func getPipeStart(grid []aoc.String) (int, int) {
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'S' {
				return x, y
			}
		}
	}

	panic("start not found")
}
