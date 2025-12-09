package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay09p1(t *testing.T) {
	input := aoc.NewInput(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)

	result := solver{}.Day9p1(input)

	assert.Equal(t, "50", result)
}

func TestDay09p2(t *testing.T) {
	input := aoc.NewInput(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)

	result := solver{}.Day9p2(input)

	assert.Equal(t, "24", result)
}
