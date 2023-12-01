package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day1p1(input aoc.Input) string {
	sum := 0
	for _, line := range input.ToStringSlice() {
		var first, last int
		var firstFound bool
		for _, c := range line {
			var err error
			if !firstFound {
				first, err = strconv.Atoi(string(c))
				if err == nil {
					firstFound = true
				}
			}

			l, err := strconv.Atoi(string(c))
			if err == nil {
				last = l
			}
		}

		sum += aoc.Atoi(fmt.Sprintf("%d%d", first, last))
		log.Printf("%s: %d%d = %d", line, first, last, sum)
	}

	return aoc.ResultI(sum)
}

func (s solver) Day1p2(input aoc.Input) string {
	sum := 0

	for _, line := range input.ToStringSlice() {
		var first, last int
		var firstFound bool
		for i := 0; i < len(line); i++ {
			if !firstFound {
				r := startsWithNumber(string(line[i:]))
				if r > -1 {
					first = r
					firstFound = true
				}
			}

			r := startsWithNumber(string(line[i:]))
			if r > -1 {
				last = r
			}
		}

		sum += aoc.Atoi(fmt.Sprintf("%d%d", first, last))
	}

	return aoc.ResultI(sum)
}

func startsWithNumber(s string) int {
	i, err := strconv.Atoi(string(s[0]))
	if err == nil {
		return i
	}

	if len(s) >= 3 && s[:3] == "one" {
		return 1
	} else if len(s) >= 3 && s[:3] == "two" {
		return 2
	} else if len(s) >= 5 && s[:5] == "three" {
		return 3
	} else if len(s) >= 4 && s[:4] == "four" {
		return 4
	} else if len(s) >= 4 && s[:4] == "five" {
		return 5
	} else if len(s) >= 3 && s[:3] == "six" {
		return 6
	} else if len(s) >= 5 && s[:5] == "seven" {
		return 7
	} else if len(s) >= 5 && s[:5] == "eight" {
		return 8
	} else if len(s) >= 4 && s[:4] == "nine" {
		return 9
	}

	return -1
}
