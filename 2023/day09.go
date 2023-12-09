package main

import "github.com/erdnaxeli/adventofcode/aoc"

func (s solver) Day9p1(input aoc.Input) string {
	sum := 0
	for _, line := range input.ToStringSlice() {
		sum += getNextOasisValue(line.SplitAtoi())
	}

	return aoc.ResultI(sum)
}

func (s solver) Day9p2(input aoc.Input) string {
	sum := 0
	for _, line := range input.ToStringSlice() {
		sum += getPrevOasisValue(line.SplitAtoi())
	}

	return aoc.ResultI(sum)
}

func getNextOasisValue(sequence []int) int {
	sequences := [][]int{sequence}

	for {
		var newSequence []int
		allZeroes := true

		for i := 1; i < len(sequence); i++ {
			diff := sequence[i] - sequence[i-1]
			if diff != 0 {
				allZeroes = false
			}

			newSequence = append(newSequence, diff)
		}

		if allZeroes {
			break
		}

		sequences = append(sequences, newSequence)
		sequence = newSequence
	}

	nextValue := 0
	for i := len(sequences) - 1; i >= 0; i-- {
		sequence := sequences[i]
		nextValue += sequence[len(sequence)-1]
	}

	return nextValue
}

func getPrevOasisValue(sequence []int) int {
	sequences := [][]int{sequence}

	for {
		var newSequence []int
		allZeroes := true

		for i := 1; i < len(sequence); i++ {
			diff := sequence[i] - sequence[i-1]
			if diff != 0 {
				allZeroes = false
			}

			newSequence = append(newSequence, diff)
		}

		if allZeroes {
			break
		}

		sequences = append(sequences, newSequence)
		sequence = newSequence
	}

	prevValue := 0
	for i := len(sequences) - 1; i >= 0; i-- {
		sequence := sequences[i]
		prevValue = sequence[0] - prevValue
	}

	return prevValue
}
