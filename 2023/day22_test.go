package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay22p1(t *testing.T) {
	input := aoc.NewInput(`1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`)

	result := solver{}.Day22p1(input)

	assert.Equal(t, "5", result)
}

func TestDay22p2(t *testing.T) {
	input := aoc.NewInput(`1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`)

	result := solver{}.Day22p2(input)

	assert.Equal(t, "7", result)
}
