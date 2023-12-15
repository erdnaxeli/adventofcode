package main

import (
	"github.com/erdnaxeli/adventofcode/aoc"
)

type Lens struct {
	Label aoc.String
	Focal int
}

func (s solver) Day15p1(input aoc.Input) string {
	sum := 0
	for _, step := range input.Delimiter(",").ToStringSlice() {
		sum += ComputeHash(step)
	}

	return aoc.ResultI(sum)
}

func (s solver) Day15p2(input aoc.Input) string {
	boxes := make(map[int][]Lens)
	for _, step := range input.Delimiter(",").ToStringSlice() {
		parts := step.SplitOn("=")
		if len(parts) == 2 {
			// add or replace lens
			lens := Lens{Label: parts[0], Focal: parts[1].Atoi()}
			i := ComputeHash(lens.Label)
			// fmt.Println(step)
			// fmt.Print(boxes[i])

			found := false
			for j, l := range boxes[i] {
				if l.Label == lens.Label {
					// replace
					boxes[i][j] = lens
					found = true
					break
				}
			}

			if !found {
				// add
				boxes[i] = append(boxes[i], lens)
			}
		} else {
			// remove lens
			label := step[:len(step)-1]
			i := ComputeHash(label)
			// fmt.Print(boxes[i])

			for j, l := range boxes[i] {
				if l.Label == label {
					// log.Print(len(boxes[i]), j)
					if j == len(boxes[i])-1 {
						boxes[i] = boxes[i][:j]
					} else {
						boxes[i] = append(boxes[i][:j], boxes[i][j+1:]...)
					}

					break
				}
			}
		}
	}

	sum := 0
	for i := range boxes {
		if len(boxes[i]) == 0 {
			continue
		}

		for j := range boxes[i] {
			sum += (1 + i) * (1 + j) * boxes[i][j].Focal
		}
	}

	return aoc.ResultI(sum)
}

func ComputeHash(s aoc.String) int {
	value := 0
	for _, c := range s {
		value += int(c)
		value *= 17
		value %= 256
	}

	return value
}
