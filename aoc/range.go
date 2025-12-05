package aoc

import (
	"cmp"
	"slices"
)

type Range struct {
	Start int
	End   int

	inclusive bool
}

// NewRange creates a new Range object.
//
// If the inclusive parameter is true, the start and the end are part of the range,
// else they are not.
func NewRange(start int, end int, inclusive bool) Range {
	return Range{
		Start:     start,
		End:       end,
		inclusive: inclusive,
	}
}

// CombineRanges takes a list of ranges and return an equivalent list of non overlapping ranges.
func CombineRanges(ranges []Range) []Range {
	if len(ranges) <= 1 {
		return ranges
	}

	ranges = slices.Clone(ranges)
	slices.SortFunc(ranges, func(r1 Range, r2 Range) int { return cmp.Compare(r1.Start, r2.Start) })
	result := make([]Range, 0, len(ranges))
	result = append(result, ranges[0])

	for _, r := range ranges[1:] {
		if r.End <= result[len(result)-1].End {
			// the range is fully included in the previous range
			continue
		} else if r.Start <= result[len(result)-1].End {
			// the range is partially included in the previous range
			result[len(result)-1].End = r.End
		} else {
			result = append(result, r)
		}

	}

	return result
}

func (r Range) Contains(v int) bool {
	if r.inclusive {
		return r.Start <= v && v <= r.End
	} else {
		return r.Start < v && v < r.End
	}
}

func (r Range) Len() int {
	if r.inclusive {
		return r.End - r.Start + 1
	} else {
		return r.End - r.Start - 1
	}
}
