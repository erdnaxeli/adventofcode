package aoc_test

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestToIntSlice(t *testing.T) {
	input := aoc.NewInput(
		`1
45
6
83
2
0
+9
-1
`,
	)

	result := input.ToIntSlice()

	expected := []int{1, 45, 6, 83, 2, 0, 9, -1}
	assert.Equal(t, expected, result)
}

func TestToStringSlice(t *testing.T) {
	input := aoc.NewInput(
		`some
string
a
b
c
asrnuite
dd
bonjour
`,
	)

	result := input.ToStringSlice()

	expected := []aoc.String{"some", "string", "a", "b", "c", "asrnuite", "dd", "bonjour"}
	assert.Equal(t, expected, result)
}
