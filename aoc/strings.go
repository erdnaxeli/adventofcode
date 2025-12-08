package aoc

import (
	"regexp"
	"strings"
)

type String string

var SPACE_REGEX = regexp.MustCompile(`\s+`)

// At return the char at the given position in the String.
func (s String) At(i int) byte {
	return string(s)[i]
}

// From return the String from the given char position to the end.
func (s String) From(i int) String {
	return String(string(s)[i:])
}

// Atoi convert the String to an it.
//
// It panics on error.
func (s String) Atoi() int {
	return Atoi(string(s))
}

// Split splits the String on spaces.
func (s String) Split() []String {
	var result []String
	for _, p := range SPACE_REGEX.Split(string(s), -1) {
		result = append(result, String(p))
	}

	return result
}

// SplitS splits the String on spaces.
func (s String) SplitS() []string {
	return SPACE_REGEX.Split(string(s), -1)
}

// SplitAtoi splits the String on spaces and convert the results to integers.
func (s String) SplitAtoi() []int {
	var result []int
	for _, p := range s.Split() {
		result = append(result, p.Atoi())
	}

	return result
}

// SplitOn splits the String using a given delimiter.
func (s String) SplitOn(d string) []String {
	var result []String
	for _, p := range strings.Split(string(s), d) {
		result = append(result, String(p))
	}

	return result
}

// SplitOn splits the String using a given delimiter and convert the results to integers.
func (s String) SplitOnAtoi(d string) []int {
	var result []int
	for _, p := range strings.Split(string(s), d) {
		result = append(result, Atoi(p))
	}

	return result
}

// SplitOn splits the String using a given delimiter and convert the results to string.
func (s String) SplitOnS(d string) []string {
	return strings.Split(string(s), d)
}

// ToIntSlice convert the String to a slice of int.
//
// It considers each char as an element of the result slice.
func (s String) ToIntSlice() []int {
	return s.SplitOnAtoi("")
}
