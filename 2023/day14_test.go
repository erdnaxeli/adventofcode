package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay14p1(t *testing.T) {
	input := aoc.NewInput(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`)

	result := solver{}.Day14p1(input)

	assert.Equal(t, "136", result)
}

func TestDay14p2(t *testing.T) {
	input := aoc.NewInput(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`)

	result := solver{}.Day14p2(input)

	assert.Equal(t, "64", result)
}
