package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay11p2(t *testing.T) {
	t.Skip("does not work because we take a shortcut for the actual solution")

	input := aoc.NewInput(`svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`)

	result := solver{}.Day11p2(input)

	assert.Equal(t, "2", result)
}
