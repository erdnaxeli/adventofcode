package aoc

import (
	"fmt"
	"strconv"
	"strings"
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

func ResultF64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func ResultS(s []string) string {
	return strings.Join(s, "")
}

func ResultSI(s []int) string {
	result := strings.Builder{}

	for _, elt := range s {
		fmt.Fprint(&result, elt)
	}

	return result.String()
}
