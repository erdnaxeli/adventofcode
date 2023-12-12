package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
)

type Spring int

const (
	OperationalSpring Spring = iota
	DamagedSpring
	UnknownSpring
)

func (s Spring) String() string {
	switch s {
	case OperationalSpring:
		return "OperationalSpring"
	case DamagedSpring:
		return "DamagedSpring"
	case UnknownSpring:
		return "UnknownSpring"
	default:
		panic("")
	}
}

type springRow struct {
	springs []Spring
	groups  []int
}

func (s solver) Day12p1(input aoc.Input) string {
	rows := readSprings(input)
	sum := 0
	for _, row := range rows {
		sum += countPossibleSpringArrangements(row)
	}

	return aoc.ResultI(sum)
}

func (s solver) Day12p2(input aoc.Input) string {
	return ""
}

func readSprings(input aoc.Input) []springRow {
	var result []springRow
	for _, line := range input.ToStringSlice() {
		var row springRow
		parts := line.Split()

		for _, c := range parts[0] {
			switch c {
			case '.':
				row.springs = append(row.springs, OperationalSpring)
			case '#':
				row.springs = append(row.springs, DamagedSpring)
			case '?':
				row.springs = append(row.springs, UnknownSpring)
			default:
				panic("unknown char")
			}
		}

		for _, c := range parts[1].SplitOn(",") {
			row.groups = append(row.groups, aoc.Atoi(string(c)))
		}

		result = append(result, row)
	}

	return result
}

func countPossibleSpringArrangements(row springRow) int {
	// log.Printf("new row %+v", row)
	prevDamagedSpringsCount := 0
	groupIdx := 0
	prevSpring := OperationalSpring
	for i, spring := range row.springs {
		// log.Printf("Spring %s", spring)

		switch spring {
		case OperationalSpring:
			if prevSpring == DamagedSpring {
				if prevDamagedSpringsCount != row.groups[groupIdx] {
					// log.Print("invalid row: group incorrect size")
					return 0
				}

				groupIdx++
				prevDamagedSpringsCount = 0
			}
		case DamagedSpring:
			if groupIdx >= len(row.groups) {
				// log.Print("invalid row: too many groups")
				return 0
			}

			prevDamagedSpringsCount++
			if prevDamagedSpringsCount > row.groups[groupIdx] {
				// log.Print("invalid row: group too big")
				return 0
			}
		case UnknownSpring:
			if groupIdx == len(row.groups) || prevDamagedSpringsCount == row.groups[groupIdx] {
				// all groups have been found or current group has correct size,
				// next spring can only be operational
				rowTry := springRow{springs: make([]Spring, len(row.springs)), groups: row.groups}
				copy(rowTry.springs, row.springs)
				rowTry.springs[i] = OperationalSpring

				return countPossibleSpringArrangements(rowTry)
			} else if prevSpring == DamagedSpring && prevDamagedSpringsCount < row.groups[groupIdx] {
				// current group is not big enough, next spring can only be damaged
				rowTry := springRow{springs: make([]Spring, len(row.springs)), groups: row.groups}
				copy(rowTry.springs, row.springs)
				rowTry.springs[i] = DamagedSpring

				return countPossibleSpringArrangements(rowTry)
			} else {
				// we try both
				rowTry1 := springRow{springs: make([]Spring, len(row.springs)), groups: row.groups}
				copy(rowTry1.springs, row.springs)
				rowTry1.springs[i] = DamagedSpring

				rowTry2 := springRow{springs: make([]Spring, len(row.springs)), groups: row.groups}
				copy(rowTry2.springs, row.springs)
				rowTry2.springs[i] = OperationalSpring

				return countPossibleSpringArrangements(rowTry1) + countPossibleSpringArrangements(rowTry2)
			}
		}

		prevSpring = spring
	}

	if prevSpring == DamagedSpring {
		if prevDamagedSpringsCount != row.groups[groupIdx] {
			// last group too small
			return 0
		}

		// final count of groups
		groupIdx++
	}

	// is there enough groups?
	if groupIdx != len(row.groups) {
		return 0
	}

	// log.Printf("Result found: %+v", row)
	return 1
}
