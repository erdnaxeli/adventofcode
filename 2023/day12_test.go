package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay12p1_1(t *testing.T) {
	input := aoc.NewInput(`?###???????? 3,2,1`)

	result := solver{}.Day12p1(input)

	assert.Equal(t, "10", result)
}

func TestDay12p1_2(t *testing.T) {
	input := aoc.NewInput(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`)

	result := solver{}.Day12p1(input)

	assert.Equal(t, "21", result)
}

func TestDay12p1_3(t *testing.T) {
	input := aoc.NewInput(`.???#??????#???????? 1,1,11`)

	result := solver{}.Day12p1(input)

	assert.Equal(t, "11", result)
}

func TestDay12p2(t *testing.T) {
	input := aoc.NewInput(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`)

	result := solver{}.Day12p2(input)

	assert.Equal(t, "525152", result)
}
