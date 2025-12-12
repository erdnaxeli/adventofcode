package aoc_test

import (
	"slices"
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestAllCombinations(t *testing.T) {
	s := []int{1, 2, 3, 4}

	result := aoc.AllCombinations(s)

	assert.ElementsMatch(
		t,
		[][]int{
			{1},
			{1, 2},
			{1, 3},
			{1, 4},
			{1, 2, 3},
			{1, 2, 4},
			{1, 2, 3, 4},
			{1, 3, 4},
			{2},
			{2, 3},
			{2, 4},
			{2, 3, 4},
			{3},
			{3, 4},
			{4},
		},
		slices.Collect(result),
	)
}
