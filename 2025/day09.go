package main

import (
	"fmt"
	"slices"

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
	borderPoints := aoc.NewSet(redTiles[0])
	minX, minY, maxX, maxY := redTiles[0].X, redTiles[0].Y, redTiles[0].X, redTiles[0].Y

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

	ppp := slices.Collect(borderPoints.Values())
	for i := range ppp {
		ppp[i].X /= 300
		ppp[i].Y /= 300
	}

	g := aoc.NewGridFromPoints[byte](ppp, '.', '#')
	fmt.Println(g)

	/*
		img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{maxX-minX, maxY-minY}})
		red := color.RGBA{255, 0, 0, 0xff}
		for p := range borderPoints {
			fmt.Println(p)
			img.Set(p.X-minX, p.Y-minY, red)
		}

		f, _ := os.Create("day08.png")
		png.Encode(f, img)
	*/

	combinations := aoc.Combinations(redTiles)
	/*
		combinationsSorted := slices.SortedFunc(combinations, func(c1 []aoc.Point, c2 []aoc.Point) int {
			return -cmp.Compare(
				aoc.AbsI(meanX-(c1[0].X+c1[1].X)/2)+aoc.AbsI(meanY-(c1[0].Y+c1[1].Y)/2),
				aoc.AbsI(meanX-(c2[0].X+c2[1].X)/2)+aoc.AbsI(meanY-(c2[0].Y+c2[1].Y)/2),
			)
		})

		fmt.Println(combinationsSorted[0])
		return aoc.ResultI(area(combinationsSorted[0][0], combinationsSorted[0][1]))
	*/

	maxArea := 0

	for c := range combinations {
		p1, p2 := c[0], c[1]
		var p3 aoc.Point
		if p1.X < p2.X {
			p3 = aoc.NewPointXY(p1.X, p2.Y)
		} else {
			p3 = aoc.NewPointXY(p2.X, p1.Y)
		}

		if !borderPoints.Contains(p3) {
			// check if p3 is outside the circle
			dLeft := p3.X - minX
			dRight := maxX - p3.X
			borderCrossCount := 0

			if dRight < dLeft {
				for x := p3.X + 1; x < maxX; x++ {
					if borderPoints.Contains(aoc.NewPointXY(x, p3.Y)) {
						borderCrossCount++
					}
				}
			} else {
				for x := minX; x < p3.X; x++ {
					if borderPoints.Contains(aoc.NewPointXY(x, p3.Y)) {
						borderCrossCount++
					}
				}
			}

			/*
				if borderCrossCount > 2 {
					fmt.Println(p1, p2, p3, borderCrossCount, borderCrossCount%2)
				}
			*/
			if borderCrossCount%2 == 0 {
				// we cross two times the border, we are outside the shape
				continue
			}
		}

		area := area(c[0], c[1])
		if area > maxArea {
			maxArea = area
			//fmt.Println(c, area)
		}
	}

	return aoc.ResultI(maxArea)
}

func area(p1 aoc.Point, p2 aoc.Point) int {
	return (aoc.AbsI(p1.X-p2.X) + 1) * (aoc.AbsI(p1.Y-p2.Y) + 1)
}
