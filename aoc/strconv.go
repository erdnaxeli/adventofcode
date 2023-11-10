package aoc

import (
	"fmt"
	"strconv"
)

// Atoi is exactly like strconv.Atoi but does not return an error.
//
// Instead it panics if anything goes wrong.
func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func ResultI(i int) string {
	return fmt.Sprint(i)
}
