package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay24p1_Solve2D(t *testing.T) {
	tests := []struct {
		input string
		x     float64
		y     float64
	}{
		{
			input: "19, 13, 30 @ -2, 1, -2\n18, 19, 22 @ -1, -1, -2",
			x:     14.333,
			y:     15.333,
		},
		{
			input: "19, 13, 30 @ -2, 1, -2\n20, 25, 34 @ -2, -2, -4",
			x:     11.667,
			y:     16.667,
		},
		{
			input: "19, 13, 30 @ -2, 1, -2\n12, 31, 28 @ -1, -2, -1",
			x:     6.2,
			y:     19.4,
		},
		{
			input: "20, 25, 34 @ -2, -2, -4\n12, 31, 28 @ -1, -2, -1",
			x:     -2,
			y:     3,
		},
	}

	for _, test := range tests {
		input := aoc.NewInput(test.input)
		hailstones := ReadHailstones(input)

		require.Len(t, hailstones, 2)

		a1, b1 := GetLinearEquationFactors2D(hailstones[0])
		a2, b2 := GetLinearEquationFactors2D(hailstones[1])

		x, y := Solve2D(a1, b1, a2, b2)

		assert.InDelta(t, test.x, x, 0.001)
		assert.InDelta(t, test.y, y, 0.001)
	}
}

func TestDay24p1_IntersectInThePast(t *testing.T) {
	tests := []string{
		"20, 25, 34 @ -2, -2, -4\n20, 19, 15 @ 1, -5, -3",
		"12, 31, 28 @ -1, -2, -1\n20, 19, 15 @ 1, -5, -3",
		"18, 19, 22 @ -1, -1, -2\n20, 19, 15 @ 1, -5, -3",
		"18, 19, 22 @ -1, -1, -2\n20, 25, 34 @ -2, -2, -4", // parallel
	}

	for _, test := range tests {
		input := aoc.NewInput(test)

		result := solver{}.Day24p1(input)

		assert.Equal(t, "0", result)
	}
}
