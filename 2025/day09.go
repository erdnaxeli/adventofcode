package main

import (
	"fmt"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day9p1(input aoc.Input) string {
	points := input.ToPoints()

	maxArea := 0
	for c := range aoc.Combinations(points) {
		area := area(c[0], c[1])
		if area > maxArea {
			maxArea = area
		}
	}

	return aoc.ResultI(maxArea)
}

func (s solver) Day9p2(input aoc.Input) string {
	redTiles := input.ToPoints()
	// the exercice statement use a grid where X is horizontal and Y is vertical
	for i := range redTiles {
		redTiles[i].X, redTiles[i].Y = redTiles[i].Y, redTiles[i].X
		// We scale the grid by 2 to _streth_ the shape.
		redTiles[i].X *= 2
		redTiles[i].Y *= 2
	}

	// we add the first point at the end to complete the loop
	redTiles = append(redTiles, redTiles[0])
	borderPoints := aoc.NewSet[aoc.Point]()
	minX, minY, maxX, maxY := redTiles[0].Y, redTiles[0].X, redTiles[0].Y, redTiles[0].X

	for i := 0; i < len(redTiles)-1; i++ {
		a, b := redTiles[i], redTiles[i+1]
		borderPoints.Add(b)

		if a.X == b.X {
			if a.Y > b.Y {
				a, b = b, a
			}

			for y := a.Y + 1; y < b.Y; y++ {
				borderPoints.Add(aoc.NewPointXY(a.X, y))
			}
		} else {
			if a.X > b.X {
				a, b = b, a
			}

			for x := a.X + 1; x < b.X; x++ {
				borderPoints.Add(aoc.NewPointXY(x, a.Y))
			}
		}

		minX = min(a.X, minX)
		minY = min(a.Y, minY)
		maxX = max(a.X, maxX)
		maxY = max(a.Y, maxY)
	}

	/*
		fmt.Println(borderPoints)
		g := aoc.NewGridFromPoints[byte](slices.Collect(borderPoints.Values()), '.', '#')
		fmt.Println(g)
	*/

	combinations := aoc.Combinations(redTiles)
	shape := aoc.NewShape(borderPoints)
	// fmt.Println("shape contains", aoc.NewPointXY(7, 2), shape.Contains(aoc.NewPointXY(7, 2)))
	maxArea := 0
	cache := make(map[aoc.Point]bool)

	for c := range combinations {
		if !isSquareValid(c, shape, cache) {
			continue
		}

		area := area(c[0], c[1])
		if area > maxArea {
			fmt.Println(area)
			maxArea = area
		}

	}

	return aoc.ResultI(maxArea)
}

func area(p1 aoc.Point, p2 aoc.Point) int {
	return (aoc.AbsI(p1.X-p2.X) + 1) * (aoc.AbsI(p1.Y-p2.Y) + 1)
}

func isSquareValid(c []aoc.Point, shape aoc.Shape, cache map[aoc.Point]bool) bool {
	for x := min(c[0].X, c[1].X); x <= max(c[0].X, c[1].X); x++ {
		for y := min(c[0].Y, c[1].Y); y <= max(c[0].Y, c[1].Y); y++ {
			p := aoc.XY(x, y)
			if r, ok := cache[p]; ok {
				if !r {
					return false
				}

				continue
			}

			if !shape.Contains(p) {
				cache[p] = false
				return false
			}

			cache[p] = true
		}
	}

	return true
}
