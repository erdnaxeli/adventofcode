package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay06p1(t *testing.T) {
	input := aoc.NewInput(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `)

	result := solver{}.Day6p1(input)

	assert.Equal(t, "4277556", result)
}

func TestDay06p2(t *testing.T) {
	input := aoc.NewInput(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `)

	result := solver{}.Day6p2(input)

	assert.Equal(t, "3263827", result)
}
