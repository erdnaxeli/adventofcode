package main

import (
	"strings"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type Machine struct {
	lightsOn []int
	buttons  [][]int
	joltage  []int
}

func (s solver) Day10p1(input aoc.Input) string {
	machines := readMachines(input)
	return ""
}

func (s solver) Day10p2(input aoc.Input) string {
	return ""
}

func readMachines(input aoc.Input) []Machine {
	for line := range input.Lines() {
		var machine Machine
		parts := line.SplitS()

		lights := parts[0]
		for i, l := range lights[1 : len(lights)-1] {
			if l == '#' {
				machine.lightsOn = append(machine.lightsOn, i)
			}
		}

		joltage := parts[len(parts)-1]
		for j := range strings.Split(joltage[1:len(joltage)-1], ",") {
			machine.joltage = append(machine.joltage, aoc.Atoi(j))
		}

		parts := parts[1 : len(parts)-1]

	}
}
