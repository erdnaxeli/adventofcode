package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay11p2(t *testing.T) {
	t.Skip("You need to change added value to 9 in the p2 code.")
	input := aoc.NewInput(`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`)

	result := solver{}.Day11p2(input)

	assert.Equal(t, "1030", result)
}
