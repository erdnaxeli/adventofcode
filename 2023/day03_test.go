package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay3p1(t *testing.T) {
	input := aoc.NewInput(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

	result := solver{}.Day3p1(input)

	assert.Equal(t, "4361", result)
}
