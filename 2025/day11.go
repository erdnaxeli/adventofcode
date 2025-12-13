package main

import (
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type ServerGraph map[string][]string

func (g ServerGraph) pathsCount(start string, end string) int {
	queue := [][]string{{start}}
	cache := make(map[string]int)

	for len(queue) > 0 {
		// LIFO for depth-first search
		// much more memory efficient than a FIFO
		path := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		server := path[len(path)-1]

		if len(g[server]) == 0 {
			// dead end
			cache[server] = 0
		}

		for _, s := range g[server] {
			if count, ok := cache[s]; ok {
				for i := range path {
					cache[path[i]] += count
				}

				continue
			}

			if s == end {
				for i := range path {
					cache[path[i]]++
				}
			} else {
				newPath := slices.Clone(path)
				newPath = append(newPath, s)
				queue = append(queue, newPath)
			}
		}
	}

	return cache[start]
}

func (s solver) Day11p1(input aoc.Input) string {
	g := readInputD11(input)
	count := g.pathsCount("you", "out")

	return aoc.ResultI(count)
}

func (s solver) Day11p2(input aoc.Input) string {
	g := readInputD11(input)

	// count := g.pathsCount("dac", "fft") => 0
	// this means the path is svr -> ... -> fft -> ... -> dac -> ... -> out

	countToFft := g.pathsCount("svr", "fft")
	countToDac := g.pathsCount("fft", "dac")
	countToOut := g.pathsCount("dac", "out")

	return aoc.ResultI(countToFft * countToDac * countToOut)
}

func readInputD11(input aoc.Input) ServerGraph {
	g := make(ServerGraph)

	for line := range input.Lines() {
		parts := line.SplitS()
		server := parts[0][:len(parts[0])-1]
		g[server] = parts[1:]
	}

	return g
}
