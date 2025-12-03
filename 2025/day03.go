package main

import (
	"math"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day3p1(input aoc.Input) string {
	banks := input.ToStringSlice()
	joltage := 0

	for _, bank := range banks {
		batteries := bank.ToIntSlice()
		i, x := aoc.FirstMaxIndex(batteries[0 : len(batteries)-1])
		_, y := aoc.FirstMaxIndex(batteries[i+1:])

		joltage += x*10 + y
	}

	return aoc.ResultI(joltage)
}

func (s solver) Day3p2(input aoc.Input) string {
	banks := input.ToStringSlice()
	joltage := 0.

	for _, bank := range banks {
		batteries := bank.ToIntSlice()

		lastIdx := -1
		for batteryIdx := range 12 {
			i, x := aoc.FirstMaxIndex(batteries[lastIdx+1 : len(batteries)-(11-batteryIdx)])
			lastIdx += 1 + i
			joltage += float64(x) * math.Pow10(11-batteryIdx)
		}
	}

	return aoc.ResultF64(joltage)
}
