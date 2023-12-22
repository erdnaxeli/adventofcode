package main

import (
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
	"golang.org/x/exp/maps"
)

type SandPile struct {
	grid   map[aoc.Point]int
	bricks [][]aoc.Point
	maxX   int
	maxY   int
	maxZ   int
}

func readSandPile(input aoc.Input) SandPile {
	lines := input.ToStringSlice()

	// A 3D grid of points to brick index.
	grid := make(map[aoc.Point]int)
	// A slice of bricks. Each brick is a list of points.
	bricks := make([][]aoc.Point, len(lines))
	maxX, maxY, maxZ := 0, 0, 0

	for brickIdx, line := range lines {
		coordinates := line.SplitOn("~")
		start := coordinates[0]
		end := coordinates[1]

		startParts := start.SplitOn(",")
		endParts := end.SplitOn(",")
		xs, ys, zs := startParts[0].Atoi(), startParts[1].Atoi(), startParts[2].Atoi()
		xe, ye, ze := endParts[0].Atoi(), endParts[1].Atoi(), endParts[2].Atoi()

		for x := xs; x <= xe; x++ {
			for y := ys; y <= ye; y++ {
				for z := zs; z <= ze; z++ {
					p := aoc.Point{X: x, Y: y, Z: z}
					bricks[brickIdx] = append(bricks[brickIdx], p)
					grid[p] = brickIdx
				}
			}
		}

		maxX = max(maxX, xe)
		maxY = max(maxY, ye)
		maxZ = max(maxZ, ze)
	}

	return SandPile{grid, bricks, maxX, maxY, maxZ}
}

// Settle makes all the bricks fall downward as far as they can go.
func (pile *SandPile) Settle() {
	// z = 0 is the ground, first layer is at z = 1
	for z := 2; z <= pile.maxZ; z++ {
		for x := 0; x <= pile.maxX; x++ {
			for y := 0; y <= pile.maxY; y++ {
				brickIdx, ok := pile.grid[aoc.Point{X: x, Y: y, Z: z}]
				if !ok {
					continue
				}

				maxFall := pile.GetMaxFall(brickIdx)
				pile.MakeFall(brickIdx, maxFall)
			}
		}
	}
}

func (pile SandPile) GetMaxFall(brickIdx int) int {
	maxFall := pile.maxZ
	for _, point := range pile.bricks[brickIdx] {
		fall := 0
		for z := point.Z - 1; z > 0; z-- {
			bellowBrickIdx, ok := pile.grid[aoc.Point{X: point.X, Y: point.Y, Z: z}]
			if ok {
				if bellowBrickIdx == brickIdx {
					// this is a vertical brick
					// this point can fall as much as the maxFall
					fall = maxFall
				}

				break
			}

			fall++
		}

		maxFall = min(maxFall, fall)
	}

	return maxFall
}

func (pile *SandPile) MakeFall(brickIdx int, fall int) {
	var newPoints []aoc.Point
	for _, point := range pile.bricks[brickIdx] {
		newPoint := point
		newPoint.Z -= fall
		newPoints = append(newPoints, newPoint)

		delete(pile.grid, point)
		pile.grid[newPoint] = brickIdx
	}

	pile.bricks[brickIdx] = newPoints
}

func (pile SandPile) FindDisintegrableBrick() []int {
	// brick that cannot be disintegrated
	notSafe := make(map[int]bool)

	for brickIdx := range pile.bricks {
		bricksBellow := pile.BricksBellow(brickIdx)
		if len(bricksBellow) == 1 {
			notSafe[bricksBellow[0]] = true
		}
	}

	var bricks []int
	for idx := range pile.bricks {
		if !notSafe[idx] {
			bricks = append(bricks, idx)
		}
	}

	return bricks
}

func (pile SandPile) BricksBellow(brick int) []int {
	bricksBellow := make(map[int]bool)

	for _, point := range pile.bricks[brick] {
		brickBellow, ok := pile.grid[aoc.Point{X: point.X, Y: point.Y, Z: point.Z - 1}]
		if !ok || brickBellow == brick {
			continue
		}

		bricksBellow[brickBellow] = true
	}

	return maps.Keys(bricksBellow)
}

func (pile SandPile) CountFallen(brickDisintegrated int) int {
	// we count the disintegrated brick as it have fallen
	bricksFallen := []int{brickDisintegrated}
	processed := make(map[int]bool)

	for z := 2; z <= pile.maxZ; z++ {
		for x := 0; x <= pile.maxX; x++ {
			for y := 0; y <= pile.maxY; y++ {
				brickIdx, ok := pile.grid[aoc.Point{X: x, Y: y, Z: z}]
				if !ok || processed[brickIdx] {
					continue
				}

				bricksBellow := pile.BricksBellow(brickIdx)
				if len(bricksBellow) == 0 {
					continue
				}

				allFallen := true
				for _, brickBellow := range bricksBellow {
					if !slices.Contains(bricksFallen, brickBellow) {
						allFallen = false
						break
					}
				}

				if allFallen {
					// All the brick below the current brick have fallen, the current
					// brick falls too.
					bricksFallen = append(bricksFallen, brickIdx)
				}

				processed[brickIdx] = true
			}
		}
	}

	return len(bricksFallen) - 1
}

func (s solver) Day22p1(input aoc.Input) string {
	pile := readSandPile(input)
	pile.Settle()
	bricks := pile.FindDisintegrableBrick()
	return aoc.ResultI(len(bricks))
}

func (s solver) Day22p2(input aoc.Input) string {
	pile := readSandPile(input)
	pile.Settle()

	safeBricks := pile.FindDisintegrableBrick()
	sum := 0
	for brickIdx := range pile.bricks {
		if slices.Contains(safeBricks, brickIdx) {
			// disintegrating this brick would not make any other brick fall
			continue
		}

		sum += pile.CountFallen(brickIdx)
	}

	return aoc.ResultI(sum)
}
