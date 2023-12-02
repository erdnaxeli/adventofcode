package main

import (
	"regexp"

	"github.com/erdnaxeli/adventofcode/aoc"
)

var (
	CUBE_GAME_RGX       = regexp.MustCompile(`Game (\d+): (.+)`)
	CUBE_GAME_CUBES_RGX = regexp.MustCompile(`(\d+) (red|green|blue)`)
)

func (s solver) Day2p1(input aoc.Input) string {
	sum := 0

	for _, line := range input.ToStringSlice() {
		match := CUBE_GAME_RGX.FindStringSubmatch(string(line))
		game_id := aoc.Atoi(match[1])
		sets := match[2]
		matches := CUBE_GAME_CUBES_RGX.FindAllStringSubmatch(sets, -1)

		possible := true
		for _, cubes := range matches {
			if (cubes[2] == "red" && aoc.Atoi(cubes[1]) > 12) || (cubes[2] == "green" && aoc.Atoi(cubes[1]) > 13) || (cubes[2] == "blue" && aoc.Atoi(cubes[1]) > 14) {
				possible = false
				break
			}
		}

		if possible {
			sum += game_id
		}
	}

	return aoc.ResultI(sum)
}

func (s solver) Day2p2(input aoc.Input) string {
	sum := 0

	for _, line := range input.ToStringSlice() {
		match := CUBE_GAME_RGX.FindStringSubmatch(string(line))
		sets := match[2]
		matches := CUBE_GAME_CUBES_RGX.FindAllStringSubmatch(sets, -1)

		cubes_min_count := make(map[string]int)
		for _, cubes := range matches {
			cubes_min_count[cubes[2]] = max(cubes_min_count[cubes[2]], aoc.Atoi(cubes[1]))
		}

		power := 1
		for _, value := range cubes_min_count {
			power *= value
		}

		sum += power
	}

	return aoc.ResultI(sum)
}
