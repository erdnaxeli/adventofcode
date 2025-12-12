package main

import (
	"strings"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type Machine struct {
	lights  []int
	buttons [][]int
	joltage []int
}

func (m Machine) minPresses() int {
	presses := len(m.buttons)

	for buttons := range aoc.AllCombinations(m.buttons) {
		lights := make([]int, len(m.lights))

		if len(buttons) >= presses {
			continue
		}

		for _, button := range buttons {
			for _, light := range button {
				lights[light]++
			}
		}

		valid := true
		for i := range m.lights {
			if lights[i]%2 != m.lights[i] {
				valid = false
				break
			}
		}

		if valid {
			presses = len(buttons)
		}
	}

	return presses
}

func (s solver) Day10p1(input aoc.Input) string {
	machines := readMachines(input)
	var total int

	for _, machine := range machines {
		p := machine.minPresses()
		total += p
	}

	return aoc.ResultI(total)
}

func (s solver) Day10p2(input aoc.Input) string {
	return ""
}

func readMachines(input aoc.Input) []Machine {
	var machines []Machine

	for line := range input.Lines() {
		var machine Machine
		parts := line.SplitS()

		lights := parts[0]
		for _, l := range lights[1 : len(lights)-1] {
			if l == '#' {
				machine.lights = append(machine.lights, 1)
			} else {
				machine.lights = append(machine.lights, 0)
			}
		}

		joltage := parts[len(parts)-1]
		for _, j := range strings.Split(joltage[1:len(joltage)-1], ",") {
			machine.joltage = append(machine.joltage, aoc.Atoi(j))
		}

		parts = parts[1 : len(parts)-1]
		for _, part := range parts {
			var button []int
			button_parts := strings.Split(part[1:len(part)-1], ",")
			for _, button_part := range button_parts {
				button = append(button, aoc.Atoi(button_part))
			}

			machine.buttons = append(machine.buttons, button)
		}

		machines = append(machines, machine)
	}

	return machines
}
