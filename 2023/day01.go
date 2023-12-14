package main

import (
	"fmt"
	"regexp"

	"github.com/erdnaxeli/adventofcode/aoc"
)

var (
	CALIBRATION_VALUE_RGX       = regexp.MustCompile(`\d`)
	CALIBRATION_VALUE_RGX_2     = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	CALIBRATION_VALUE_RGX_2_rev = regexp.MustCompile(`\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`)
	CALIBRATION_VALUE_TO_INT    = map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
		"eno": 1, "owt": 2, "eerht": 3, "ruof": 4, "evif": 5, "xis": 6, "neves": 7, "thgie": 8, "enin": 9,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	}
)

func (s solver) Day1p1(input aoc.Input) string {
	sum := 0
	for _, line := range input.ToStringSlice() {
		first := CALIBRATION_VALUE_RGX.FindString(string(line))
		last := CALIBRATION_VALUE_RGX.FindString(reverse(string(line)))
		sum += aoc.Atoi(fmt.Sprintf("%s%s", first, last))
		// log.Printf("%s: %s%s = %d", line, first, last, sum)
	}

	return aoc.ResultI(sum)
}

func (s solver) Day1p2(input aoc.Input) string {
	sum := 0
	for _, line := range input.ToStringSlice() {
		firstS := CALIBRATION_VALUE_RGX_2.FindString(string(line))
		lastS := CALIBRATION_VALUE_RGX_2_rev.FindString(reverse(string(line)))

		first := CALIBRATION_VALUE_TO_INT[firstS]
		last := CALIBRATION_VALUE_TO_INT[lastS]
		sum += aoc.Atoi(fmt.Sprintf("%d%d", first, last))
		// log.Printf("%s: %d%d = %d", line, first, last, sum)
	}

	return aoc.ResultI(sum)
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
