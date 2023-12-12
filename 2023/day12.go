package main

import (
	"fmt"
	"strings"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type Spring int

const (
	OperationalSpring Spring = iota
	DamagedSpring
	UnknownSpring
)

/*func (s Spring) String() string {
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
}*/

type springRow struct {
	springs []Spring
	groups  []int
}

func (s springRow) Hash() string {
	return strings.Join(strings.Fields(fmt.Sprintf("%v-%v", s.springs, s.groups)), "-")
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
	rows := readSprings(input)
	sum := 0
	for _, row := range rows {
		s, g := row.springs, row.groups
		for i := 0; i < 4; i++ {
			row.springs = append(row.springs, UnknownSpring)
			row.springs = append(row.springs, s...)
			row.groups = append(row.groups, g...)
		}

		sum += countPossibleSpringArrangements(row)
	}

	return aoc.ResultI(sum)
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

var cache = make(map[string]int)

func countPossibleSpringArrangements(row springRow) int {
	rowHash := row.Hash()
	if count, ok := cache[rowHash]; ok {
		return count
	}

	// size of the current group of damaged springs
	currentGroupCount := 0
	// index of the current group in row.groups
	groupIdx := 0
	prevSpring := OperationalSpring
	for i, spring := range row.springs {
		switch spring {
		case OperationalSpring:
			if prevSpring == DamagedSpring {
				if currentGroupCount != row.groups[groupIdx] {
					cache[rowHash] = 0
					return 0
				}

				// preparing next group
				groupIdx++
				currentGroupCount = 0
			}
		case DamagedSpring:
			if groupIdx >= len(row.groups) {
				cache[rowHash] = 0
				return 0
			}

			currentGroupCount++
			if currentGroupCount > row.groups[groupIdx] {
				cache[rowHash] = 0
				return 0
			}
		case UnknownSpring:
			if groupIdx == len(row.groups) || (groupIdx == len(row.groups)-1 && currentGroupCount == row.groups[groupIdx]) {
				// all groups have been found
				for _, spring := range row.springs[i+1:] {
					if spring == DamagedSpring {
						cache[rowHash] = 0
						return 0
					}
				}

				cache[rowHash] = 1
				return 1
			} else if currentGroupCount == row.groups[groupIdx] {
				// current group has correct size, next spring can only be operational
				rowTry := springRow{springs: make([]Spring, len(row.springs)-i), groups: row.groups[groupIdx+1:]}
				copy(rowTry.springs, row.springs[i:])
				rowTry.springs[0] = OperationalSpring

				count := countPossibleSpringArrangements(rowTry)
				cache[rowHash] = count
				return count
			} else if prevSpring == DamagedSpring {
				// current group is not big enough, next spring can only be damaged
				rowTry := springRow{
					springs: make([]Spring, len(row.springs)-i),
					groups:  make([]int, len(row.groups)-groupIdx),
				}
				copy(rowTry.groups, row.groups[groupIdx:])
				rowTry.groups[0] -= currentGroupCount
				copy(rowTry.springs, row.springs[i:])
				rowTry.springs[0] = DamagedSpring

				count := countPossibleSpringArrangements(rowTry)
				cache[rowHash] = count
				return count
			} else {
				// we try both
				rowTry1 := springRow{
					springs: make([]Spring, len(row.springs)-i),
					groups:  row.groups[groupIdx:],
				}
				copy(rowTry1.springs, row.springs[i:])
				rowTry1.springs[0] = DamagedSpring

				rowTry2 := springRow{
					springs: make([]Spring, len(row.springs)-i),
					groups:  row.groups[groupIdx:],
				}
				copy(rowTry2.springs, row.springs[i:])
				rowTry2.springs[0] = OperationalSpring

				count1 := countPossibleSpringArrangements(rowTry1)
				count2 := countPossibleSpringArrangements(rowTry2)
				cache[rowHash] = count1 + count2
				return count1 + count2
			}
		}

		prevSpring = spring
	}

	// check if we need to count last group
	if prevSpring == DamagedSpring {
		if currentGroupCount != row.groups[groupIdx] {
			// last group too small
			cache[rowHash] = 0
			return 0
		}

		// final count of groups
		groupIdx++
	}

	// is there enough groups?
	if groupIdx != len(row.groups) {
		cache[rowHash] = 0
		return 0
	}

	cache[rowHash] = 1
	return 1
}
