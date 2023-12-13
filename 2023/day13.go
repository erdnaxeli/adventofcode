package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day13p1(input aoc.Input) string {
	grids := input.Delimiter("\n\n").ToStringSlice()
	sum := 0
	for _, g := range grids {
		grid := g.SplitOn("\n")
		x, y := findMirrorSymmetry(grid)
		sum += y + 100*x
	}

	return aoc.ResultI(sum)
}

func (s solver) Day13p2(input aoc.Input) string {
	grids := input.Delimiter("\n\n").ToStringSlice()
	sum := 0
	for _, g := range grids {
		grid := g.SplitOn("\n")
		x, y := findMirrorSymmetry2(grid)
		sum += y + 100*x
	}

	return aoc.ResultI(sum)
}

func findMirrorSymmetry(grid []aoc.String) (int, int) {
	var x, y int

	symmetry := true
	for x = range grid {
		for dx := 1; dx < len(grid); dx++ {
			xp := x + 1 - dx
			xpp := x + dx
			if xp < 0 || xpp >= len(grid) {
				break
			}

			if grid[xp] == grid[xpp] {
				symmetry = true
			} else {
				symmetry = false
				break
			}
		}

		if symmetry {
			break
		}
	}

	if symmetry {
		return x + 1, 0
	}

	symmetry = false
	for y = range grid[0] {
		for dy := 1; dy < len(grid); dy++ {
			yp := y + 1 - dy
			ypp := y + dy
			if yp < 0 || ypp >= len(grid[0]) {
				break
			}

			for x := range grid {
				if grid[x][yp] == grid[x][ypp] {
					symmetry = true
				} else {
					symmetry = false
					break
				}
			}

			if !symmetry {
				break
			}
		}

		if symmetry {
			break
		}
	}

	return 0, y + 1
}

func findMirrorSymmetry2(grid []aoc.String) (int, int) {
	var symmetry bool
	var x, y int

	for x = range grid {
		symmetry = false
		diff := 0
		for dx := 1; dx < len(grid); dx++ {
			xp := x + 1 - dx
			xpp := x + dx
			if xp < 0 || xpp >= len(grid) {
				break
			}

			for y := range grid[x] {
				if grid[xp][y] == grid[xpp][y] {
					symmetry = true
				} else {
					diff++

					if diff == 2 {
						symmetry = false
						break
					}
				}
			}

			if !symmetry {
				break
			}
		}

		if symmetry && diff == 1 {
			break
		}
	}

	if symmetry {
		return x + 1, 0
	}

	for y = range grid[0] {
		symmetry = false
		diff := 0
		for dy := 1; dy < len(grid); dy++ {
			yp := y + 1 - dy
			ypp := y + dy
			if yp < 0 || ypp >= len(grid[0]) {
				break
			}

			for x := range grid {
				if grid[x][yp] == grid[x][ypp] {
					symmetry = true
				} else {
					diff++

					if diff == 2 {
						symmetry = false
						break
					}
				}
			}

			if !symmetry {
				break
			}
		}

		if symmetry && diff == 1 {
			break
		}
	}

	return 0, y + 1
}
