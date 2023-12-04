package main

import (
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day4p1(input aoc.Input) string {
	sum := 0
	for _, line := range input.ToStringSlice() {
		parts := line.Split()
		var winningNumbers []aoc.String
		score := 0
		isFoundNumbers := false

		for _, part := range parts[2:] {
			if part == "|" {
				isFoundNumbers = true
				continue
			}

			if isFoundNumbers {
				if slices.Contains(winningNumbers, part) {
					if score == 0 {
						score = 1
					} else {
						score = score << 1
					}
				}
			} else {
				winningNumbers = append(winningNumbers, part)
			}
		}

		sum += score
	}

	return aoc.ResultI(sum)
}

func (s solver) Day4p2(input aoc.Input) string {
	cards := input.ToStringSlice()
	cardsCount := make(map[int]int)

	for i := 1; i < len(cards)+1; i++ {
		cardsCount[i] = 1
	}

	cardNumber := 0
	for _, line := range cards {
		cardNumber++

		parts := line.Split()
		var winningNumbers []aoc.String
		foundNumbers := 0
		isFoundNumbers := false

		for _, part := range parts[2:] {
			if part == "|" {
				isFoundNumbers = true
				continue
			}

			if isFoundNumbers {
				if slices.Contains(winningNumbers, part) {
					foundNumbers++
				}
			} else {
				winningNumbers = append(winningNumbers, part)
			}
		}

		for i := 1; i <= foundNumbers; i++ {
			cardsCount[cardNumber+i] += cardsCount[cardNumber]
		}
	}

	sum := 0
	for _, count := range cardsCount {
		sum += count
	}

	return aoc.ResultI(sum)
}
