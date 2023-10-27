package aoc

import (
	"strconv"
	"strings"
)

type Input struct {
	content string
}

func NewInput(content string) Input {
	return Input{strings.TrimRight(content, "\n")}
}

func (i Input) Content() string {
	return i.content
}

func (i Input) ToIntSlice() []int {
	var result []int
	for _, e := range strings.Split(i.content, "\n") {
		n, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}

		result = append(result, n)
	}

	return result
}

func (i Input) ToStringSlice() []string {
	return strings.Split(i.content, "\n")
}
