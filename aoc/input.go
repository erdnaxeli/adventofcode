package aoc

import (
	"strings"
)

type Input struct {
	content   string
	delimiter string
	multiline bool
}

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

func (i Input) Delimiter(d string) Input {
	i.delimiter = d
	return i
}

func (i Input) SingleLine() Input {
	i.multiline = false
	return i
}

func (i Input) ToIntSlice() []int {
	var result []int
	for _, elt := range strings.Split(i.content, i.getDelimiter()) {
		result = append(result, Atoi(elt))
	}

	return result
}

func (i Input) ToStringSlice() []String {
	lines := strings.Split(i.content, i.getDelimiter())
	var trimmedLines []String
	for _, line := range lines {
		trimmedLines = append(trimmedLines, String(strings.TrimSpace(line)))
	}

	return trimmedLines
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
