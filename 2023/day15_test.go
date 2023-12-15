package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay15p2(t *testing.T) {
	input := aoc.NewInput(`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`)

	result := solver{}.Day15p2(input)

	assert.Equal(t, "145", result)
}
