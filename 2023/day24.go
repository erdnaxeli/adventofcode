package main

import (
	"regexp"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type Hailstone struct {
	Position aoc.Point
	Velocity aoc.Point
}

func (h Hailstone) IsInPast(x float64, y float64) bool {
	return ((h.Velocity.X > 0 && x < float64(h.Position.X)) || (float64(h.Velocity.X) < 0 && x > float64(h.Position.X)) ||
		h.Velocity.Y > 0 && y < float64(h.Position.Y)) || (float64(h.Velocity.Y) < 0 && y > float64(h.Position.Y))
}

var HAILSTONE_RGX = regexp.MustCompile(`(.*), (.*), (.*) @ (.*), (.*), (.*)`)

func (s solver) Day24p1(input aoc.Input) string {
	hailstones := ReadHailstones(input)
	sum := 0
	for i := range hailstones {
		for j := i + 1; j < len(hailstones); j++ {
			h1 := hailstones[i]
			h2 := hailstones[j]

			if AreColinear2D(h1, h2) {
				continue
			}

			a1, b1 := GetLinearEquationFactors2D(h1)
			a2, b2 := GetLinearEquationFactors2D(h2)
			x, y := Solve2D(a1, b1, a2, b2)

			// the intersection must be in the future
			if h1.IsInPast(x, y) || h2.IsInPast(x, y) {
				continue
			}

			if 200_000_000_000_000 <= x && x <= 400_000_000_000_000 && 200_000_000_000_000 <= y && y <= 400_000_000_000_000 {
				sum++
			}
		}
	}

	return aoc.ResultI(sum)
}

func (s solver) Day24p2(input aoc.Input) string {
	return ""
}

func ReadHailstones(input aoc.Input) []Hailstone {
	var hailstones []Hailstone
	for _, line := range input.ToStringSlice() {
		match := HAILSTONE_RGX.FindStringSubmatch(string(line))
		hailstones = append(
			hailstones,
			Hailstone{
				Position: aoc.Point{
					X: aoc.Atoi(match[1]),
					Y: aoc.Atoi(match[2]),
					Z: aoc.Atoi(match[3]),
				},
				Velocity: aoc.Point{
					X: aoc.Atoi(match[4]),
					Y: aoc.Atoi(match[5]),
					Z: aoc.Atoi(match[6]),
				},
			},
		)
	}

	return hailstones
}

func AreColinear2D(h1 Hailstone, h2 Hailstone) bool {
	return h1.Velocity.X*h2.Velocity.Y == h1.Velocity.Y*h2.Velocity.X
}

func GetLinearEquationFactors2D(h Hailstone) (float64, float64) {
	// position: (a b)
	// vector director (velocity): (c, d)
	//
	// If (x, y) is a point of the line, (x - a, y - b) is another vector director.
	// Then (x - a)*d == (y - b)*c
	//      dx - ad == cy - bc
	//      cy = dx - ad + bc
	//       y = (d/c)x - (ad - bc)/c
	a := float64(h.Position.X)
	b := float64(h.Position.Y)
	c := float64(h.Velocity.X)
	d := float64(h.Velocity.Y)

	// y = ax + b
	// return a, b
	return d / c, -(a*d - b*c) / c
}

func Solve2D(a1, b1, a2, b2 float64) (float64, float64) {
	// y = a1 * x + b1
	// y = a2 * x + b2
	// a1 * x + b1 = a2 * x + b2
	// x * (a1 - a2) = b2 - b1
	// x = (b2 - b1) / (a1 - a2)

	x := (b2 - b1) / (a1 - a2)
	y := a1*x + b1
	return x, y
}
