package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay20p1(t *testing.T) {
	input := aoc.NewInput(`broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`)

	result := solver{}.Day20p1(input)

	assert.Equal(t, "32000000", result)
}
