package aoc

import "math"

// AbsI is exactly like math.Abs but for integers.
func AbsI(x int) int {
	return int(math.Abs(float64(x)))
}
