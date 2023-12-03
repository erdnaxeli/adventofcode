package main

import (
	"regexp"

	"github.com/erdnaxeli/adventofcode/aoc"
)

var ENGINE_SCHEMATIC_NUMBER_RGX = regexp.MustCompile(`\d+`)

func (s solver) Day3p1(input aoc.Input) string {
	sum := 0
	lines := input.ToStringSlice()

	for x, line := range lines {
		matches := ENGINE_SCHEMATIC_NUMBER_RGX.FindAllStringIndex(string(line), -1)

		for _, match := range matches {
			yMin, yMax := match[0], match[1]-1
			symbolFound := false

			for yy := max(0, yMin-1); yy <= min(len(lines)-1, yMax+1); yy++ {
				for xx := max(0, x-1); xx <= min(len(line)-1, x+1); xx++ {
					if xx == x && yMin <= yy && yy <= yMax {
						continue
					}

					if lines[xx][yy] != '.' {
						symbolFound = true
						break
					}
				}

				if symbolFound {
					break
				}
			}

			if symbolFound {
				sum += aoc.Atoi(string(line[yMin : yMax+1]))
			}
		}
	}

	return aoc.ResultI(sum)
}

func (s solver) Day3p2(input aoc.Input) string {
	gears := make(map[aoc.Point][]int)
	lines := input.ToStringSlice()

	for x, line := range lines {
		matches := ENGINE_SCHEMATIC_NUMBER_RGX.FindAllStringIndex(string(line), -1)

		for _, match := range matches {
			yMin, yMax := match[0], match[1]-1

			for yy := max(0, yMin-1); yy <= min(len(lines)-1, yMax+1); yy++ {
				for xx := max(0, x-1); xx <= min(len(line)-1, x+1); xx++ {
					if xx == x && yMin <= yy && yy <= yMax {
						continue
					}

					if lines[xx][yy] == '*' {
						gear := aoc.Point{X: xx, Y: yy, Z: 0}
						gears[gear] = append(gears[gear], (aoc.Atoi(string(line[yMin : yMax+1]))))
					}
				}
			}
		}
	}

	sum := 0
	for _, numbers := range gears {
		if len(numbers) != 2 {
			continue
		}

		sum += numbers[0] * numbers[1]
	}

	return aoc.ResultI(sum)
}
