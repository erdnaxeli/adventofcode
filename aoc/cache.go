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
		return errors.New("no env var $HOME defined")
	}

	directory := fmt.Sprintf("%s/.cache/adventofcode/%d", home, year)
	err := os.MkdirAll(directory, 0700)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s/%d", directory, day)
	err = os.WriteFile(filename, []byte(input.content), 0600)
	if err != nil {
		return err
	}

	return nil
}
