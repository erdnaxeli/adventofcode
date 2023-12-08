package main

import (
	"math"
	"regexp"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type DesertNode struct {
	Left  string
	Right string
}

var DESERT_NODE_RGX = regexp.MustCompile(`(...) = \((...), (...)\)`)

func (s solver) Day8p1(input aoc.Input) string {
	lines := input.ToStringSlice()
	instructions := lines[0]
	nodes := make(map[string]DesertNode)

	for _, node := range lines[2:] {
		match := DESERT_NODE_RGX.FindStringSubmatch(string(node))
		nodes[match[1]] = DesertNode{
			Left:  match[2],
			Right: match[3],
		}
	}

	steps := 0
	currentNode := "AAA"
	for {
		for _, instruction := range instructions {
			steps += 1
			switch instruction {
			case 'L':
				currentNode = nodes[currentNode].Left
			case 'R':
				currentNode = nodes[currentNode].Right
			}

			if currentNode == "ZZZ" {
				return aoc.ResultI(steps)
			}
		}
	}
}

func (s solver) Day8p2(input aoc.Input) string {
	lines := input.ToStringSlice()
	instructions := lines[0]
	nodes := make(map[string]DesertNode)
	var currentNodes []string

	for _, node := range lines[2:] {
		match := DESERT_NODE_RGX.FindStringSubmatch(string(node))
		nodes[match[1]] = DesertNode{
			Left:  match[2],
			Right: match[3],
		}

		if match[1][2] == 'A' {
			currentNodes = append(currentNodes, match[1])
		}
	}

	var cycles []int
	for i := range currentNodes {
		steps := 0
		for {
			stop := false

			for _, instruction := range instructions {
				steps += 1

				switch instruction {
				case 'L':
					currentNodes[i] = nodes[currentNodes[i]].Left
				case 'R':
					currentNodes[i] = nodes[currentNodes[i]].Right
				}

				if currentNodes[i][2] == 'Z' {
					cycles = append(cycles, steps)
					stop = true
					break
				}
			}

			if stop {
				break
			}
		}
	}

	result := getLeastCommonMultiple(cycles)
	return aoc.ResultF64(result)
}

func getPrimeFactors(n int) []int {
	var factors []int

	for i := 2; i < n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	factors = append(factors, n)
	return factors
}

func getLeastCommonMultiple(numbers []int) float64 {
	factors := make(map[int]int)

	for _, n := range numbers {
		nFactors := getPrimeFactors(n)
		nFactorsCount := make(map[int]int)

		for _, factor := range nFactors {
			nFactorsCount[factor]++
		}

		for factor, count := range nFactorsCount {
			factors[factor] = max(factors[factor], count)
		}
	}

	result := 1.
	for factor, count := range factors {
		result *= math.Pow(float64(factor), float64(count))
	}

	return result
}
