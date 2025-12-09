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
	//fmt.Println("shape contains", aoc.NewPointXY(7, 2), shape.Contains(aoc.NewPointXY(7, 2)))

	areaChan := make(chan int)
	maxArea := 0

	go func() {
		for area := range areaChan {
			if area > maxArea {
				fmt.Println(area)
				maxArea = area
			}
		}
	}()

	for c := range combinations {
		go func(c []aoc.Point) {
			p1, p2 := c[0], c[1]
			p3 := aoc.NewPointXY(p1.X, p2.Y)
			p4 := aoc.NewPointXY(p2.X, p1.Y)

			if !shape.Contains(p3) || !shape.Contains(p4) {
				return
			}

			area := area(c[0], c[1])
			areaChan <- area
		}(c)
	}

	return aoc.ResultI(maxArea)
}

func area(p1 aoc.Point, p2 aoc.Point) int {
	return (aoc.AbsI(p1.X-p2.X) + 1) * (aoc.AbsI(p1.Y-p2.Y) + 1)
}

func shapeContains(shapeBorder aoc.Set[aoc.Point], p aoc.Point, minX int, maxX int) bool {
	if shapeBorder.Contains(p) {
		return true
	}

	dLeft := p.X - minX
	dRight := maxX - p.X
	borderCrossCount := 0
	onBorderPreviously := false

	if dRight < dLeft {
		for x := maxX; x > p.X; x-- {
			onBorder := shapeBorder.Contains(aoc.NewPointXY(x, p.Y))
			if onBorder && !onBorderPreviously {
				borderCrossCount++
			}

			onBorderPreviously = onBorder
		}
	} else {
		for x := minX; x < p.X; x++ {
			onBorder := shapeBorder.Contains(aoc.NewPointXY(x, p.Y))
			if onBorder && !onBorderPreviously {
				borderCrossCount++
			}

			onBorderPreviously = onBorder
		}
	}

	if borderCrossCount%2 == 0 {
		return false
	}

	return true
}
