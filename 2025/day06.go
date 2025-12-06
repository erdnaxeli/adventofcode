package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day6p1(input aoc.Input) string {
	grid := input.ToGridS()
	total := 0

	for y := range grid.LenY() {
		op := string(grid.AtXY(grid.MaxX(), y))
		problemResult := grid.AtXY(0, y).Atoi()

		for _, v := range grid.IterColumn(y, 1, grid.MaxX()-1) {
			switch op {
			case "+":
				problemResult += v.Atoi()
			case "*":
				problemResult *= v.Atoi()
			}
		}

		total += problemResult
	}

	return aoc.ResultI(total)
}

func (s solver) Day6p2(input aoc.Input) string {
	grid := input.ToGrid()
	total := 0
	var numbers []int
	var op byte

	for y := grid.MaxY(); y >= 0; y-- {
		number := 0

		for _, v := range grid.IterAllColumn(y) {
			if v == ' ' {
				continue
			}

			if v == '+' || v == '*' {
				op = v
				continue
			}

			number = number*10 + aoc.Atoi(string(v))
		}

		if number == 0 {
			// It equals zero if the whole column was empty, meaning it is the end
			// of the problem.

			switch op {
			case '+':
				total += aoc.SumSlice(numbers)
			case '*':
				total += aoc.MultSlice(numbers)
			}

			numbers = make([]int, 0)
		} else {
			numbers = append(numbers, number)
		}
	}

	// We have to do it once again for the last problem.
	switch op {
	case '+':
		total += aoc.SumSlice(numbers)
	case '*':
		total += aoc.MultSlice(numbers)
	}

	return aoc.ResultI(total)
}
