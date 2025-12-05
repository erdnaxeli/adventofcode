package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day5p1(input aoc.Input) string {
	ranges, ids := aoc.MultiInput(
		input,
		func(i aoc.Input) []aoc.Range {
			return i.ToRanges(true)
		},
		func(i aoc.Input) []int {
			return i.ToIntSlice()
		},
	)

	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if r.Contains(id) {
				count += 1
				break
			}
		}
	}

	return aoc.ResultI(count)
}

func (s solver) Day5p2(input aoc.Input) string {
	ranges, _ := aoc.MultiInput(
		input,
		func(i aoc.Input) []aoc.Range {
			return i.ToRanges(true)
		},
		func(i aoc.Input) int { return 0 },
	)

	ranges = aoc.CombineRanges(ranges)
	count := 0
	for _, r := range ranges {
		count += r.Len()
	}

	return aoc.ResultI(count)
}
