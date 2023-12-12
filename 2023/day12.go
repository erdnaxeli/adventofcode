package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/erdnaxeli/adventofcode/aoc"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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

func (s solver) Day12p1(input aoc.Input) string {
	rows := readSprings(input)
	sum := 0
	for _, row := range rows {
		sum += countPossibleSpringArrangements(row, 0, -1, 0, 0, OperationalSpring)
	}

	log.Print(count)
	return aoc.ResultI(sum)
}

func (s solver) Day12p2(input aoc.Input) string {
	rows := readSprings(input)
	sum := 0
	for i, row := range rows {
		log.Print(i, len(rows))
		s, g := row.springs, row.groups
		for i := 0; i < 4; i++ {
			row.springs = append(row.springs, UnknownSpring)
			row.springs = append(row.springs, s...)
			row.groups = append(row.groups, g...)
		}

		sum += countPossibleSpringArrangements(row, 0, -1, 0, 0, OperationalSpring)
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

var count = 0

func countPossibleSpringArrangements(row springRow, idx int, possibleDamaged int, prevDamagedSpringsCount int, groupIdx int, prevSpring Spring) int {
	if possibleDamaged < 0 {
		for _, spring := range row.springs {
			switch spring {
			case DamagedSpring, UnknownSpring:
				possibleDamaged++
			}
		}
	}

	if count%50_000_000 == 0 {
		log.Printf("new row %v", strings.Join(strings.Fields(fmt.Sprint(row.springs)), ""))
		log.Print(row.groups)
		p := message.NewPrinter(language.English)
		p.Printf("%d\n", count)
	}
	count++
	for i, spring := range row.springs[idx:] {
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
			possibleDamaged--
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
			if groupIdx < len(row.groups)-1 {
				remainingDamaged := 0
				for _, count := range row.groups[groupIdx+1:] {
					remainingDamaged += count
				}

				if possibleDamaged < remainingDamaged {
					return 0
				}
			}

			if groupIdx == len(row.groups) || prevDamagedSpringsCount == row.groups[groupIdx] {
				// all groups have been found or the current group has correct size,
				// the next spring can only be operational
				/*rowTry := springRow{springs: make([]Spring, len(row.springs)), groups: row.groups}
				copy(rowTry.springs, row.springs)*/
				row.springs[idx+i] = OperationalSpring

				return countPossibleSpringArrangements(row, idx+i, possibleDamaged-1, prevDamagedSpringsCount, groupIdx, prevSpring)
			} else if prevSpring == DamagedSpring && prevDamagedSpringsCount < row.groups[groupIdx] {
				// current group is not big enough, next spring can only be damaged
				/*rowTry := springRow{springs: make([]Spring, len(row.springs)), groups: row.groups}
				copy(rowTry.springs, row.springs)*/
				row.springs[idx+i] = DamagedSpring

				return countPossibleSpringArrangements(row, idx+i, possibleDamaged, prevDamagedSpringsCount, groupIdx, prevSpring)
			} else {
				// we try both
				rowTry1 := springRow{springs: make([]Spring, len(row.springs)), groups: row.groups}
				copy(rowTry1.springs, row.springs)
				rowTry1.springs[idx+i] = DamagedSpring

				/*rowTry2 := springRow{springs: make([]Spring, len(row.springs)), groups: row.groups}
				copy(rowTry2.springs, row.springs)*/
				row.springs[idx+i] = OperationalSpring

				return (countPossibleSpringArrangements(rowTry1, idx+i, possibleDamaged, prevDamagedSpringsCount, groupIdx, prevSpring) +
					countPossibleSpringArrangements(row, idx+i, possibleDamaged-1, prevDamagedSpringsCount, groupIdx, prevSpring))
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

	// log.Printf("Result found: %s", strings.Join(strings.Fields(fmt.Sprint(row.springs)), ""))
	return 1
}
