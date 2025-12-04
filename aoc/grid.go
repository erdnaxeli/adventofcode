package aoc

import (
	"iter"
)

var aroundRange = [...]int{-1, 0, 1}

// Grid represent a grid.
//
// It expects to be manipulated using Point objects to address positions in the grid.
type Grid struct {
	grid [][]byte
}

func (g Grid) At(p Point) byte {
	return g.grid[p.X][p.Y]
}

// Set sets a value on the grid
//
// It actually mutates the grid even if the receiver is not on a pointer because
// it involves slices on the inside, which use pointers.
func (g Grid) Set(p Point, v byte) {
	g.grid[p.X][p.Y] = v
}

func (g Grid) LenX() int {
	return len(g.grid)
}

func (g Grid) LenY() int {
	if g.LenX() == 0 {
		return 0
	}

	return len(g.grid[0])
}

func (g Grid) MaxX() int {
	return g.LenX() - 1
}

func (g Grid) MaxY() int {
	return g.LenY() - 1
}

// Around return an iterator over all the 8 points around a position.
//
// If the position is on an edge of the grid, fewer than 8 points can be returned.
func (g Grid) Around(p Point) iter.Seq2[Point, byte] {
	return func(yield func(Point, byte) bool) {
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
func (g Grid) Points() iter.Seq2[Point, byte] {
	return func(yield func(Point, byte) bool) {
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
