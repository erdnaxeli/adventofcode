package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func TestDay01p1(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"R2, L3",
			"5",
		},
		{
			"R2, R2, R2",
			"2",
		},
		{
			"R5, L5, R5, R3",
			"12",
		},
	}

	for _, test := range tests {
		input := aoc.NewInput(test.input)

		result := solver{}.Day1p1(input)

		if result != test.expected {
			t.Logf("Expected %s but got %s", test.expected, result)
			t.Fail()
		}
	}
}

func TestDay01p2(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"R8, R4, R4, R8",
			"4",
		},
	}

	for _, test := range tests {
		input := aoc.NewInput(test.input)

		result := solver{}.Day1p2(input)

		if result != test.expected {
			t.Logf("Expected %s but got %s", test.expected, result)
			t.Fail()
		}
	}
}
