package main

import (
	"log"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day2p1(input aoc.Input) string {
	var buttons []int
	button := 5
	for _, line := range input.ToStringSlice() {
		for _, char := range line {
			switch char {
			case 'R':
				if button%3 == 0 {
					continue
				}

				button += 1
			case 'L':
				if button%3 == 1 {
					continue
				}

				button -= 1
			case 'U':
				if button <= 3 {
					continue
				}

				button -= 3
			case 'D':
				if button >= 7 {
					continue
				}

				button += 3
			}
		}

		log.Print(line, button)

		buttons = append(buttons, button)
	}

	return aoc.ResultSI(buttons)
}

func (s solver) Day2p2(input aoc.Input) string {
	var positions []string
	position := aoc.Point{-2, 0, 0}
	grid := map[aoc.Point]string{
		{0, 2, 0}:   "1",
		{-1, 1, 0}:  "2",
		{0, 1, 0}:   "3",
		{1, 1, 0}:   "4",
		{-2, 0, 0}:  "5",
		{-1, 0, 0}:  "6",
		{0, 0, 0}:   "7",
		{1, 0, 0}:   "8",
		{2, 0, 0}:   "9",
		{-1, -1, 0}: "A",
		{0, -1, 0}:  "B",
		{1, -1, 0}:  "C",
		{0, -2, 0}:  "D",
	}
	for _, line := range input.ToStringSlice() {
		for _, char := range line {
			var newPosition aoc.Point
			switch char {
			case 'R':
				newPosition = position.Add(aoc.Point{1, 0, 0})
			case 'L':
				newPosition = position.Add(aoc.Point{-1, 0, 0})
			case 'U':
				newPosition = position.Add(aoc.Point{0, 1, 0})
			case 'D':
				newPosition = position.Add(aoc.Point{0, -1, 0})
			}

			if grid[newPosition] == "" {
				continue
			}

			position = newPosition
		}

		positions = append(positions, grid[position])
	}

	return aoc.ResultS(positions)
}
