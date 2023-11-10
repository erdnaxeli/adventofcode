package aoc

import (
	"errors"
	"fmt"
	"os"
)

type Cache interface {
	GetInput(year int, day int, part int) Input
	StoreInput(year int, day int, part int, input Input) error
}

type DefaultCache struct{}

func NewDefaultCache() DefaultCache {
	return DefaultCache{}
}

func (c DefaultCache) GetInput(year int, day int, part int) Input {
	home := os.Getenv("HOME")
	if home == "" {
		return Input{}
	}

	filename := fmt.Sprintf("%s/.cache/adventofcode/%d/%d", home, year, day)
	content, err := os.ReadFile(filename)
	if err != nil {
		return Input{}
	}

	return NewInput(string(content))
}

func (c DefaultCache) StoreInput(year int, day int, part int, input Input) error {
	home := os.Getenv("HOME")
	if home == "" {
		return errors.New("No env var $HOME defined.")
	}

	directory := fmt.Sprintf("%s/.cache/adventofcode/%d", home, year)
	os.MkdirAll(directory, 0700)

	filename := fmt.Sprintf("%s/%d", directory, day)
	os.WriteFile(filename, []byte(input.content), 0600)

	return nil
}
