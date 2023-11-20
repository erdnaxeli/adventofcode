package main

import (
	"sort"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day3p1(input aoc.Input) string {
	count := 0
	for _, line := range input.ToStringSlice() {
		parts := line.SplitAtoi()
		sort.Ints(parts)
		if isTriangleValid(parts) {
			count++
		}
	}

	return aoc.ResultI(count)
}

func (s solver) Day3p2(input aoc.Input) string {
	count := 0
	var t1, t2, t3 []int
	for _, line := range input.ToStringSlice() {
		parts := line.SplitAtoi()
		t1 = append(t1, parts[0])
		t2 = append(t2, parts[1])
		t3 = append(t3, parts[2])

		if len(t1) == 3 {
			for _, t := range [][]int{t1, t2, t3} {
				if isTriangleValid(t) {
					count++
				}
			}

			t1, t2, t3 = nil, nil, nil
		}
	}

	return aoc.ResultI(count)
}

func isTriangleValid(t []int) bool {
	sort.Ints(t)
	return t[0]+t[1] > t[2]
}
