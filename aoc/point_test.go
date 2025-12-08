package aoc_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	for i, testCase := range []struct {
		p1       aoc.Point
		p2       aoc.Point
		expected float64
	}{
		{aoc.Point{0, 0, 0}, aoc.Point{0, 0, 0}, 0},
		{aoc.Point{0, 0, 0}, aoc.Point{1, 0, 0}, 1},
		{aoc.Point{0, 0, 0}, aoc.Point{0, 1, 0}, 1},
		{aoc.Point{0, 0, 0}, aoc.Point{0, 0, 1}, 1},
		{aoc.Point{0, 0, 0}, aoc.Point{1, 1, 0}, math.Sqrt(2)},
		{aoc.Point{3, 2, 1}, aoc.Point{7, 8, 9}, math.Sqrt(116)},
	} {
		t.Run(fmt.Sprintf("TestDistance/%d", i), func(t *testing.T) {
			d := testCase.p1.Distance(testCase.p2)

			assert.Equal(t, testCase.expected, d)
		})
	}
}
