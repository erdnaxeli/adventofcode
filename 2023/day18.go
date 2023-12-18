package main

import (
	"fmt"
	"strconv"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day18p1(input aoc.Input) string {
	point := aoc.Point{}
	grid := make(map[aoc.Point]bool)
	grid[point] = true

	for _, line := range input.ToStringSlice() {
		parts := line.Split()
		switch parts[0] {
		case "U":
			for i := 0; i < parts[1].Atoi(); i++ {
				point = point.MoveNorth()
				grid[point] = true
			}
		case "D":
			for i := 0; i < parts[1].Atoi(); i++ {
				point = point.MoveSouth()
				grid[point] = true
			}
		case "L":
			for i := 0; i < parts[1].Atoi(); i++ {
				point = point.MoveWest()
				grid[point] = true
			}
		case "R":
			for i := 0; i < parts[1].Atoi(); i++ {
				point = point.MoveEast()
				grid[point] = true
			}
		}

		grid[point] = true
	}

	// I cheated: I printed the map then looked for a point inside.
	queue := []aoc.Point{{X: -1, Y: 0}}
	processed := make(map[aoc.Point]bool)
	sum := 0
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		if processed[point] || grid[point] {
			continue
		}

		processed[point] = true
		sum++
		pn := point.MoveNorth()
		pw := point.MoveWest()
		ps := point.MoveSouth()
		pe := point.MoveEast()
		queue = append(queue, pn, pw, ps, pe)
	}

	return aoc.ResultI(sum + len(grid))
}

func (s solver) Day18p2(input aoc.Input) string {
	point := aoc.Point{}
	grid := make(map[aoc.Point]bool)
	grid[point] = true

	for _, line := range input.ToStringSlice() {
		parts := line.Split()
		distance, _ := strconv.ParseInt(string(parts[2][2:7]), 16, 64)
		direction := parts[2][7]

		switch direction {
		case '0':
			for i := 0; i < int(distance); i++ {
				point = point.MoveNorth()
				grid[point] = true
			}
		case '1':
			for i := 0; i < int(distance); i++ {
				point = point.MoveSouth()
				grid[point] = true
			}
		case '2':
			for i := 0; i < int(distance); i++ {
				point = point.MoveWest()
				grid[point] = true
			}
		case '3':
			for i := 0; i < int(distance); i++ {
				point = point.MoveEast()
				grid[point] = true
			}
		}

		grid[point] = true
	}

	var xmin, xmax, ymin, ymax int
	for point := range grid {
		xmin = min(xmin, point.X)
		xmax = max(xmax, point.X)
		ymin = min(ymin, point.Y)
		ymax = max(ymax, point.Y)
	}
	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			if x == 0 && y == 0 {
				fmt.Print("0000000")
			} else if grid[aoc.Point{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return ""

	// I cheated: I printed the map then looked for a point inside.
	queue := []aoc.Point{{X: -1, Y: 0}}
	processed := make(map[aoc.Point]bool)
	sum := 0
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		if processed[point] || grid[point] {
			continue
		}

		processed[point] = true
		sum++
		pn := point.MoveNorth()
		pw := point.MoveWest()
		ps := point.MoveSouth()
		pe := point.MoveEast()
		queue = append(queue, pn, pw, ps, pe)
	}

	return aoc.ResultI(sum + len(grid))
}
