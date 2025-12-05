package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay05p2(t *testing.T) {
	input := aoc.NewInput(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

	result := solver{}.Day5p2(input)

	assert.Equal(t, "14", result)
}
