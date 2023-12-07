package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay7p1(t *testing.T) {
	input := aoc.NewInput(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)

	result := solver{}.Day7p1(input)

	assert.Equal(t, "6440", result)
}

func TestDay7p2(t *testing.T) {
	input := aoc.NewInput(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)

	result := solver{}.Day7p2(input)

	assert.Equal(t, "5905", result)
}
