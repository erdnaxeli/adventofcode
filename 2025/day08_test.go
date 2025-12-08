package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay08p1(t *testing.T) {
	t.Skip("You must edit the code to change the number of couples joined from 1000 to 10.")

	input := aoc.NewInput(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`)

	result := solver{}.Day8p1(input)

	assert.Equal(t, "40", result)
}

func BenchmarkMakePositionsCouples(b *testing.B) {
	input := aoc.NewDefaultRunner(2025, solver{}).GetInput(8, 1)
	positions := readInputDay08(input)

	for b.Loop() {
		makePositionsCouples(positions)
	}
}

func BenchmarkCompteD8P1(b *testing.B) {
	input := aoc.NewDefaultRunner(2025, solver{}).GetInput(8, 1)
	positions := readInputDay08(input)
	couples := makePositionsCouples(positions)

	for b.Loop() {
		computeD8P1(positions, couples)
	}
}
