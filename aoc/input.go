package aoc

import (
	"iter"
	"strings"
)

type Input struct {
	content   string
	delimiter string
	multiline bool
}

// NewInput return a new Input object.
//
// By default the input is configured in multiline mode, without any delimiter,
// which is equivalent to using a "\n" delimiter.
func NewInput(content string) Input {
	return Input{
		content:   strings.TrimRight(content, "\n"),
		delimiter: "",
		multiline: true,
	}
}

func (i Input) Content() string {
	return i.content
}

// Delimiter sets the delimiter used by methods to split the input.
//
// If the delimiter set is different than "" it overwrite the multiline
// configuration.
//
// It returns the same input, so methods can be chained.
func (i Input) Delimiter(d string) Input {
	i.delimiter = d
	return i
}

func (i Input) Rotate() []String {
	lines := i.ToStringSlice()
	tmp := make([][]rune, len(lines[0]))
	for i := range tmp {
		tmp[i] = make([]rune, len(lines))
	}

	for i, line := range lines {
		for j, c := range line {
			tmp[j][i] = c
		}
	}

	var result []String
	for _, line := range tmp {
		result = append(result, String(line))
	}

	return result
}

// SingleLine set the multiline mode to false.
//
// In multiline mode, the input is split on "\n".
// In singleline, the input is split on each char.
//
// If a delimiter is set, the multiline mode is ignored.
func (i Input) SingleLine() Input {
	i.multiline = false
	return i
}

func (i Input) Lines() iter.Seq[String] {
	return func(yield func(String) bool) {
		for line := range strings.SplitSeq(i.content, "\n") {
			yield(String(line))
		}
	}
}

// ToIntSlice parse the input as a slice of int.
//
// If a delimiter is set, it is used to split the input.
// Else in multiline mode the input is split on each line.
// Else the input is plit on each char.
//
// If any part cannot be parsed as an int, it panics.
func (i Input) ToIntSlice() []int {
	var result []int
	for elt := range strings.SplitSeq(i.content, i.getDelimiter()) {
		result = append(result, Atoi(elt))
	}

	return result
}

// ToStringSlice parse the input as a slice of String.
//
// If a delimiter is set, it is used to split the input.
// Else in multiline mode the input is split on each line.
// Else the input is plit on each char.
func (i Input) ToStringSlice() []String {
	var trimmedLines []String
	for line := range strings.SplitSeq(i.content, i.getDelimiter()) {
		trimmedLines = append(trimmedLines, String(strings.TrimSpace(line)))
	}

	return trimmedLines
}

// ToGrid parse the input as a grid.
//
// x is the axe from top to bottom, y is the axis from left to right.
// (0, 0) is the top left point.
func (i Input) ToGrid() Grid[byte] {
	var grid [][]byte
	for line := range strings.SplitSeq(i.content, i.getDelimiter()) {
		grid = append(grid, []byte(line))
	}

	return Grid[byte]{grid: grid}
}

func (i Input) ToGridS() Grid[String] {
	var grid [][]String
	for _, line := range i.ToStringSlice() {
		grid = append(grid, line.Split())
	}

	return Grid[String]{grid: grid}
}

func (i Input) ToRanges(inclusive bool) []Range {
	var ranges []Range

	for line := range i.Lines() {
		parts := line.SplitOnAtoi("-")
		ranges = append(ranges, NewRange(parts[0], parts[1], inclusive))
	}

	return ranges
}

// MultiInput split and parse the input as two differents input separated by an empty line.
func MultiInput[R1, R2 any](i Input, parseInput1 func(i Input) R1, parseInput2 func(i Input) R2) (R1, R2) {
	parts := strings.Split(i.content, "\n\n")
	return parseInput1(NewInput(parts[0])), parseInput2(NewInput(parts[1]))
}

func (i Input) getDelimiter() string {
	if i.delimiter != "" {
		return i.delimiter
	} else if i.multiline {
		return "\n"
	} else {
		return ", "
	}
}
