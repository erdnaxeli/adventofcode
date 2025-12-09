package aoc

import (
	"iter"
	"strings"
)

var aroundRange = [...]int{-1, 0, 1}

// Grid represent a grid.
//
// It expects to be manipulated using Point objects to address positions in the grid.
type Grid[E any] struct {
	grid [][]E
}

func NewGridFromPoints[E any](points []Point, empty E, notEmpty E) Grid[E] {
	xMax, yMax := 0, 0
	for _, p := range points {
		if p.X > xMax {
			xMax = p.X
		}

		if p.Y > yMax {
			yMax = p.Y
		}
	}

	var grid [][]E
	for range xMax + 1 {
		var line []E
		for range yMax + 1 {
			line = append(line, empty)
		}

		grid = append(grid, line)
	}

	for _, p := range points {
		grid[p.X][p.Y] = notEmpty
	}

	return Grid[E]{grid}
}

func (g Grid[E]) At(p Point) E {
	return g.grid[p.X][p.Y]
}

func (g Grid[E]) AtXY(x int, y int) E {
	return g.grid[x][y]
}

func (g Grid[E]) IterColumn(y int, xMin int, xMax int) iter.Seq2[Point, E] {
	return func(yield func(Point, E) bool) {
		for x := xMin; x <= xMax; x++ {
			if !yield(Point{X: x, Y: y}, g.grid[x][y]) {
				return
			}
		}
	}
}

func (g Grid[E]) IterAllColumn(y int) iter.Seq2[Point, E] {
	return g.IterColumn(y, 0, g.MaxX())
}

func (g Grid[E]) IterLine(x int, yMin int, yMax int) iter.Seq2[Point, E] {
	return func(yield func(Point, E) bool) {
		for y, e := range g.grid[x][yMin : yMax+1] {
			if !yield(Point{X: x, Y: y}, e) {
				return
			}
		}
	}
}

func (g Grid[E]) IterAllLine(x int) iter.Seq2[Point, E] {
	return g.IterLine(x, 0, g.MaxX())
}

// Set sets a value on the grid
//
// It actually mutates the grid even if the receiver is not on a pointer because
// it involves slices on the inside, which use pointers.
func (g Grid[E]) Set(p Point, v E) {
	g.grid[p.X][p.Y] = v
}

func (g Grid[E]) LenX() int {
	return len(g.grid)
}

func (g Grid[E]) LenY() int {
	if g.LenX() == 0 {
		return 0
	}

	return len(g.grid[0])
}

func (g Grid[E]) MaxX() int {
	return g.LenX() - 1
}

func (g Grid[E]) MaxY() int {
	return g.LenY() - 1
}

// Around return an iterator over all the 8 points around a position.
//
// If the position is on an edge of the grid, fewer than 8 points can be returned.
func (g Grid[E]) Around(p Point) iter.Seq2[Point, E] {
	return func(yield func(Point, E) bool) {
		for _, x := range aroundRange {
			for _, y := range aroundRange {
				pp := Point{X: p.X + x, Y: p.Y + y}
				if pp.X < 0 || pp.Y < 0 || (pp.X == p.X && pp.Y == p.Y) || pp.X > g.MaxX() || pp.Y > g.MaxY() {
					continue
				}

				if !yield(pp, g.At(pp)) {
					return
				}
			}
		}
	}
}

// Points return an iterator over all points on the grid.
func (g Grid[E]) Points() iter.Seq2[Point, E] {
	return func(yield func(Point, E) bool) {
		for x := range g.LenX() {
			for y := range g.LenY() {
				p := Point{X: x, Y: y}

				if !yield(p, g.At(p)) {
					return
				}
			}
		}
	}
}

func (g Grid[E]) String() string {
	var b strings.Builder

	if gB, ok := any(g).(Grid[byte]); ok {
		for x := range gB.grid {
			for y := range gB.grid[x] {
				b.WriteByte(gB.grid[x][y])
				b.WriteByte(' ')
			}

			b.WriteByte('\n')
		}

		return b.String()
	}

	return ""
}
