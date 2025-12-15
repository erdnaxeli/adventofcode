package aoc

import (
	"iter"
)

// AllCombinations returns an iterator over all combinations of elements in s.
func AllCombinations[S ~[]E, E any](s S) iter.Seq[[]E] {
	return func(yield func([]E) bool) {
		for subset := 1; subset < (1 << len(s)); subset++ {
			var c []E
			for i := range s {
				if (subset>>i)&1 == 1 {
					c = append(c, s[i])
				}
			}

			if !yield(c) {
				return
			}
		}
	}
}

// Combinations returns an iterator over all two sized combinations of elements in s.
func Combinations[S ~[]E, E any](s S) iter.Seq[[]E] {
	return func(yield func([]E) bool) {
		for i, x := range s {
			for _, y := range s[i+1:] {
				if !yield([]E{x, y}) {
					return
				}
			}
		}
	}
}

// FirstMaxIndex returns the index of the first max value found, and the value.
func FirstMaxIndex(s []int) (int, int) {
	maxV := s[0]
	maxI := 0
	for i, e := range s {
		if e > maxV {
			maxV = e
			maxI = i
		}
	}

	return maxI, maxV
}

// FirstMaxIndex returns the index of the first min value found, and the value.
func FirstMinIndex(s []int) (int, int) {
	max := s[0]
	maxI := 0
	for i, e := range s {
		if e < max {
			max = e
			maxI = i
		}
	}

	return maxI, max
}

func SumSlice(s []int) int {
	result := 0

	for _, v := range s {
		result += v
	}

	return result
}

func MultSlice(s []int) int {
	result := 1

	for _, v := range s {
		result *= v
	}

	return result
}
