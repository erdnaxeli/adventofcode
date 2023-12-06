package main

import (
	"strings"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type Race struct {
	Time     int
	Distance int
}

func (s solver) Day6p1(input aoc.Input) string {
	lines := input.ToStringSlice()
	var races []Race

	for _, time := range lines[0].Split()[1:] {
		races = append(races, Race{
			Time:     time.Atoi(),
			Distance: 0,
		})
	}

	for i, distance := range lines[1].Split()[1:] {
		races[i].Distance = distance.Atoi()
	}

	var waysToWin []int
	for _, race := range races {
		ways := 0
		for i := 1; i < race.Time; i++ {
			if i*(race.Time-i) > race.Distance {
				ways += 1
			}
		}

		waysToWin = append(waysToWin, ways)
	}

	result := 1
	for _, ways := range waysToWin {
		result *= ways
	}

	return aoc.ResultI(result)
}

func (s solver) Day6p2(input aoc.Input) string {
	lines := input.ToStringSlice()
	time := aoc.Atoi(strings.Join(lines[0].SplitS()[1:], ""))
	distance := aoc.Atoi(strings.Join(lines[1].SplitS()[1:], ""))

	ways := 0
	for i := 1; i < time; i++ {
		if i*(time-i) > distance {
			ways += 1
		}
	}

	return aoc.ResultI(ways)
}
