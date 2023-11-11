package aoc

import (
	"regexp"
	"strconv"
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
	n, err := strconv.Atoi(string(s))
	if err != nil {
		panic(err)
	}

	return n
}

func (s String) Split() []String {
	var result []String
	for _, p := range SPACE_REGEX.Split(string(s), -1) {
		result = append(result, String(p))
	}

	return result
}
func (s String) SplitAtoi() []int {
	var result []int
	for _, p := range s.Split() {
		result = append(result, p.Atoi())
	}

	return result
}
