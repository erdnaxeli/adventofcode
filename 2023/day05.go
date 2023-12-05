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

type Range struct {
	Start int
	End   int
}

func (s solver) Day5p2(input aoc.Input) string {
	lines := input.ToStringSlice()
	seeds := lines[0][len("seeds: "):].Split()

	seedsMap := make(map[Range]Range)
	for i := 0; i < len(seeds); i += 2 {
		// seeds are actually range of seeds
		start := seeds[i].Atoi()
		length := seeds[i+1].Atoi()
		seed := Range{
			Start: start,
			End:   start + length - 1,
		}
		seedsMap[seed] = seed
	}

	processed := make(map[Range]bool)
	for _, line := range lines[2:] {
		// map separator
		if line == "" {
			processed = make(map[Range]bool)
			continue
		}

		// map header
		if line[len(line)-1] == ':' {
			continue
		}

		parts := line.Split()
		destinationStart, sourceStart, count := parts[0].Atoi(), parts[1].Atoi(), parts[2].Atoi()
		shift := destinationStart - sourceStart
		sourceEnd := sourceStart + count - 1
		destinationEnd := destinationStart + count - 1

		for seedRange, valueRange := range seedsMap {
			if processed[seedRange] {
				continue
			}
			delete(seedsMap, seedRange)

			if sourceStart <= valueRange.Start && valueRange.End <= sourceEnd {
				seedsMap[seedRange] = Range{
					Start: valueRange.Start + shift,
					End:   valueRange.End + shift,
				}
				processed[seedRange] = true
				continue
			}

			if valueRange.Start < sourceStart && sourceEnd < valueRange.End {
				seedRange1 := Range{
					Start: seedRange.Start,
					End:   seedRange.Start + (sourceStart - valueRange.Start) - 1,
				}
				seedRange2 := Range{
					Start: seedRange1.End + 1,
					End:   seedRange.End - (valueRange.End - sourceEnd),
				}
				seedRange3 := Range{
					Start: seedRange2.End + 1,
					End:   seedRange.End,
				}

				// not shifted
				seedsMap[seedRange1] = Range{
					Start: valueRange.Start,
					End:   valueRange.Start + (seedRange1.End - seedRange1.Start),
				}
				// shifted
				seedsMap[seedRange2] = Range{
					Start: destinationStart,
					End:   destinationEnd,
				}
				// not shifted
				seedsMap[seedRange3] = Range{
					Start: valueRange.End - (seedRange3.End - seedRange3.Start),
					End:   valueRange.End,
				}

				processed[seedRange2] = true

				continue
			}

			if sourceStart <= valueRange.Start && valueRange.Start <= sourceEnd {
				overlap := sourceEnd - valueRange.Start + 1

				seedRange1 := Range{
					Start: seedRange.Start,
					End:   seedRange.Start + overlap - 1,
				}
				seedRange2 := Range{
					Start: seedRange1.End + 1,
					End:   seedRange.End,
				}

				// shifted
				valueRange1 := Range{
					Start: valueRange.Start + shift,
					End:   destinationEnd,
				}
				// not shifted
				valueRange2 := Range{
					Start: valueRange.Start + overlap,
					End:   valueRange.End,
				}

				seedsMap[seedRange1] = valueRange1
				seedsMap[seedRange2] = valueRange2

				processed[seedRange1] = true

				continue
			}

			if sourceStart <= valueRange.End && valueRange.End < sourceStart+count {
				overlap := valueRange.End - sourceStart + 1

				seedRange1 := Range{
					Start: seedRange.Start,
					End:   seedRange.End - overlap,
				}
				seedRange2 := Range{
					Start: seedRange1.End + 1,
					End:   seedRange.End,
				}

				// not shifted
				valueRange1 := Range{
					Start: valueRange.Start,
					End:   valueRange.End - overlap,
				}
				// shifted
				valueRange2 := Range{
					Start: destinationStart,
					End:   valueRange.End + shift,
				}

				seedsMap[seedRange1] = valueRange1
				seedsMap[seedRange2] = valueRange2

				processed[seedRange2] = true

				continue
			}

			seedsMap[seedRange] = valueRange
		}
	}

	// we know location is the last map
	lowestLocation := -1
	for _, valueRange := range seedsMap {
		if lowestLocation == -1 || valueRange.Start < lowestLocation {
			lowestLocation = valueRange.Start
		}
	}

	return aoc.ResultI(lowestLocation)
}
