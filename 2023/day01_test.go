package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay01p2(t *testing.T) {
	input := aoc.NewInput(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)

	result := solver{}.Day1p2(input)

	assert.Equal(t, "281", result)
}
