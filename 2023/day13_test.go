package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay13p1(t *testing.T) {
	input := aoc.NewInput(`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`)

	result := solver{}.Day13p1(input)

	assert.Equal(t, "405", result)
}

func TestDay13p2(t *testing.T) {
	input := aoc.NewInput(`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`)

	result := solver{}.Day13p2(input)

	assert.Equal(t, "400", result)
}

func TestDay13p2_2(t *testing.T) {
	input := aoc.NewInput(`######.#.#.#...
.#..#.#.##...#.
...#....#.#.#..
#.##.#.#.....#.
.......##.#.###
#######.##.#..#
......##...##..
#.##.#....#.#..
#.##.#....#.#..`)

	result := solver{}.Day13p2(input)

	assert.Equal(t, "3", result)
}
