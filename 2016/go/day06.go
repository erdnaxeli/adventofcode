package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day6p1(input aoc.Input) string {
	var message []rune
	for _, line := range input.Rotate() {
		counter := aoc.NewCounter(line)
		keys := counter.GetKeys()
		message = append(message, keys[0].Key)
	}

	return string(message)
}

func (s solver) Day6p2(input aoc.Input) string {
	var message []rune
	for _, line := range input.Rotate() {
		counter := aoc.NewCounter(line)
		keys := counter.GetKeys()
		message = append(message, keys[len(keys)-1].Key)
	}

	return string(message)
}
