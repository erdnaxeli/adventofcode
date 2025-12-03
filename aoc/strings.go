package aoc

import (
	"regexp"
	"strings"
)

type String string

var SPACE_REGEX = regexp.MustCompile(`\s+`)

func (s String) At(i int) byte {
	return string(s)[i]
}

func (s String) From(i int) String {
	return String(string(s)[i:])
}

func (s String) Atoi() int {
	return Atoi(string(s))
}

func (s String) Split() []String {
	var result []String
	for _, p := range SPACE_REGEX.Split(string(s), -1) {
		result = append(result, String(p))
	}

	return result
}

func (s String) SplitS() []string {
	return SPACE_REGEX.Split(string(s), -1)
}

func (s String) SplitAtoi() []int {
	var result []int
	for _, p := range s.Split() {
		result = append(result, p.Atoi())
	}

	return result
}

func (s String) SplitOn(d string) []String {
	var result []String
	for _, p := range strings.Split(string(s), d) {
		result = append(result, String(p))
	}

	return result
}

func (s String) SplitOnAtoi(d string) []int {
	var result []int
	for _, p := range strings.Split(string(s), d) {
		result = append(result, Atoi(p))
	}

	return result
}

func (s String) SplitOnS(d string) []string {
	return strings.Split(string(s), d)
}

func (s String) ToIntSlice() []int {
	return s.SplitOnAtoi("")
}
