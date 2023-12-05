package main

import (
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day5p1(input aoc.Input) string {
	lines := input.ToStringSlice()
	seeds := lines[0][len("seeds: "):].Split()

	seedsMap := make(map[int]int)
	for _, seedS := range seeds {
		seed := seedS.Atoi()
		seedsMap[seed] = seed
	}

	var seedsProcessed []int
	for _, line := range lines[2:] {
		// map separator
		if line == "" {
			// empty the slice keeping the underlying memory
			seedsProcessed = seedsProcessed[:0]
			continue
		}

		// map header
		if line[len(line)-1] == ':' {
			continue
		}

		parts := line.Split()
		destination, source, count := parts[0].Atoi(), parts[1].Atoi(), parts[2].Atoi()

		for seed, value := range seedsMap {
			if slices.Contains(seedsProcessed, seed) {
				continue
			}

			if source <= value && value < source+count {
				seedsMap[seed] = value + (destination - source)
				seedsProcessed = append(seedsProcessed, seed)
			}
		}
	}

	// we know location is the last map
	lowestLocation := -1
	for _, value := range seedsMap {
		if lowestLocation == -1 {
			lowestLocation = value
			continue
		}

		if value < lowestLocation {
			lowestLocation = value
		}
	}

	return aoc.ResultI(lowestLocation)
}

func (s solver) Day5p2(input aoc.Input) string {
	lines := input.ToStringSlice()
	seeds := lines[0][len("seeds: "):].Split()

	seedsMap := make(map[int]int)
	for i := 0; i < len(seeds); i += 2 {
		// seeds are actually range of seeds
		seed := seeds[i].Atoi()
		for j := 0; j < seeds[i+1].Atoi(); j++ {
			seedsMap[seed+j] = seed + j
		}
	}

	var seedsProcessed []int
	for _, line := range lines[2:] {
		// map separator
		if line == "" {
			// empty the slice keeping the underlying memory
			seedsProcessed = seedsProcessed[:0]
			continue
		}

		// map header
		if line[len(line)-1] == ':' {
			continue
		}

		parts := line.Split()
		destination, source, count := parts[0].Atoi(), parts[1].Atoi(), parts[2].Atoi()

		for seed, value := range seedsMap {
			if slices.Contains(seedsProcessed, seed) {
				continue
			}

			if source <= value && value < source+count {
				seedsMap[seed] = value + (destination - source)
				seedsProcessed = append(seedsProcessed, seed)
			}
		}
	}

	// we know location is the last map
	lowestLocation := -1
	for _, value := range seedsMap {
		if lowestLocation == -1 {
			lowestLocation = value
			continue
		}

		if value < lowestLocation {
			lowestLocation = value
		}
	}

	return aoc.ResultI(lowestLocation)
}
