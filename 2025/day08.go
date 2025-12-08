package main

import (
	"cmp"
	"maps"
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type positionsCouple struct {
	p1 aoc.Point
	p2 aoc.Point
	d  float64
}

func (s solver) Day8p1(input aoc.Input) string {
	positions := readInputDay08(input)
	couples := makePositionsCouples(positions)

	return computeD8P1(positions, couples)
}

func computeD8P1(positions []aoc.Point, couples []positionsCouple) string {
	circuits := make(map[aoc.Point]*aoc.Set[aoc.Point], len(positions))

	// Make one circuit for each position
	for _, p := range positions {
		s := aoc.NewSet(p)
		circuits[p] = &s
	}

	for _, couple := range couples[:1000] {
		circuitP1 := circuits[couple.p1]
		circuitP2 := circuits[couple.p2]

		for p := range *circuitP2 {
			circuitP1.Add(p)
			// To be sure they all share a pointer to the same circuit.
			circuits[p] = circuitP1
		}

		circuits[couple.p2] = circuitP1
	}

	uniqueCircuits := aoc.NewSetFromIter(maps.Values(circuits))
	values := slices.SortedFunc(
		uniqueCircuits.Values(),
		func(c1 *aoc.Set[aoc.Point], c2 *aoc.Set[aoc.Point]) int {
			return -cmp.Compare(c1.Len(), c2.Len())
		},
	)

	return aoc.ResultI(values[0].Len() * values[1].Len() * values[2].Len())
}

func (s solver) Day8p2(input aoc.Input) string {
	positions := readInputDay08(input)
	couples := makePositionsCouples(positions)
	circuitsCount := len(positions)
	circuits := make(map[aoc.Point]*aoc.Set[aoc.Point], circuitsCount)

	// Make one circuit for each position
	for _, p := range positions {
		s := aoc.NewSet(p)
		circuits[p] = &s
	}

	for _, couple := range couples {
		circuitP1 := circuits[couple.p1]
		circuitP2 := circuits[couple.p2]

		if circuitP1 == circuitP2 {
			// they are already in the same circuit
			continue
		}

		for p := range *circuitP2 {
			circuitP1.Add(p)
			// To be sure they all share a pointer to the same circuit.
			circuits[p] = circuitP1
		}

		circuits[couple.p2] = circuitP1

		circuitsCount--
		if circuitsCount == 1 {
			return aoc.ResultI(couple.p1.X * couple.p2.X)
		}
	}

	// error
	return aoc.ResultI(0)
}

func readInputDay08(input aoc.Input) []aoc.Point {
	lines := input.ToStringSlice()
	var positions []aoc.Point

	for _, line := range lines {
		coordinates := line.SplitOnAtoi(",")
		positions = append(positions, aoc.NewPointXYZ(coordinates[0], coordinates[1], coordinates[2]))
	}

	return positions
}

// makePositionsCouples returns a list of tuple (p1, p2, d) where d is the distance
// between p1 and p2.
//
// The list is ordered by the shortest distance.
func makePositionsCouples(positions []aoc.Point) []positionsCouple {
	// Both slices points to the same underlying array, but as we only read it it
	// is ok.
	work := positions
	var result []positionsCouple

	for len(work) > 0 {
		p1 := work[0]
		work = work[1:]

		if len(work) == 0 {
			continue
		}

		for _, p2 := range work {
			d := p1.Distance(p2)
			result = append(result, positionsCouple{p1: p1, p2: p2, d: d})
		}

	}

	slices.SortFunc(
		result,
		func(pc1 positionsCouple, pc2 positionsCouple) int { return cmp.Compare(pc1.d, pc2.d) },
	)
	return result
}
