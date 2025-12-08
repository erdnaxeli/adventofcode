package aoc

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

// An object implementing the solutions for all days.
//
// A day method receives an Input object, and must return a string.
// If a day method returns an empty string, it will consider the solution is not
// implemented yet and it will not send it.
type Solver interface {
	Day1p1(Input) string
	Day1p2(Input) string
	Day2p1(Input) string
	Day2p2(Input) string
	Day3p1(Input) string
	Day3p2(Input) string
	Day4p1(Input) string
	Day4p2(Input) string
	Day5p1(Input) string
	Day5p2(Input) string
	Day6p1(Input) string
	Day6p2(Input) string
	Day7p1(Input) string
	Day7p2(Input) string
	Day8p1(Input) string
	Day8p2(Input) string
	Day9p1(Input) string
	Day9p2(Input) string
	Day10p1(Input) string
	Day10p2(Input) string
	Day11p1(Input) string
	Day11p2(Input) string
	Day12p1(Input) string
	Day12p2(Input) string
	Day13p1(Input) string
	Day13p2(Input) string
	Day14p1(Input) string
	Day14p2(Input) string
	Day15p1(Input) string
	Day15p2(Input) string
	Day16p1(Input) string
	Day16p2(Input) string
	Day17p1(Input) string
	Day17p2(Input) string
	Day18p1(Input) string
	Day18p2(Input) string
	Day19p1(Input) string
	Day19p2(Input) string
	Day20p1(Input) string
	Day20p2(Input) string
	Day21p1(Input) string
	Day21p2(Input) string
	Day22p1(Input) string
	Day22p2(Input) string
	Day23p1(Input) string
	Day23p2(Input) string
	Day24p1(Input) string
	Day24p2(Input) string
	Day25p1(Input) string
	Day25p2(Input) string
}

type Runner struct {
	cache     Cache
	client    Client
	solver    Solver
	daysParts [][]func(Input) string

	year int
}

// Return a new Runner object with default cache and client.
//
// This is the main constructor you should use.
func NewDefaultRunner(year int, solver Solver) Runner {
	session, ok := os.LookupEnv("AOC_SESSION")
	if !ok {
		log.Fatal("An environment variable AOC_SESSION must be defined.")
	}

	client, err := NewDefaultClient(session)
	if err != nil {
		log.Fatal(err)
	}

	cache := NewDefaultCache()

	return NewRunner(
		year,
		cache,
		client,
		solver,
	)
}

// Return a new Runner.
//
// Unless you want to inject specific implementation for the cache and the client
// you should use NewDefaultRunner.
func NewRunner(year int, cache Cache, client Client, solver Solver) Runner {
	return Runner{
		cache:  cache,
		client: client,
		solver: solver,
		daysParts: [][]func(Input) string{
			{solver.Day1p1, solver.Day1p2},
			{solver.Day2p1, solver.Day2p2},
			{solver.Day3p1, solver.Day3p2},
			{solver.Day4p1, solver.Day4p2},
			{solver.Day5p1, solver.Day5p2},
			{solver.Day6p1, solver.Day6p2},
			{solver.Day7p1, solver.Day7p2},
			{solver.Day8p1, solver.Day8p2},
			{solver.Day9p1, solver.Day9p2},
			{solver.Day10p1, solver.Day10p2},
			{solver.Day11p1, solver.Day11p2},
			{solver.Day12p1, solver.Day12p2},
			{solver.Day13p1, solver.Day13p2},
			{solver.Day14p1, solver.Day14p2},
			{solver.Day15p1, solver.Day15p2},
			{solver.Day16p1, solver.Day16p2},
			{solver.Day17p1, solver.Day17p2},
			{solver.Day18p1, solver.Day18p2},
			{solver.Day19p1, solver.Day19p2},
			{solver.Day20p1, solver.Day20p2},
			{solver.Day21p1, solver.Day21p2},
			{solver.Day22p1, solver.Day22p2},
			{solver.Day23p1, solver.Day23p2},
			{solver.Day24p1, solver.Day24p2},
			{solver.Day25p1, solver.Day25p2},
		},

		year: year,
	}
}

// RunCli runs the solution for a given day and part.
//
// It reads the day and the part for the executable arguments. The first argument
// is the day, the second is the part.
func (r Runner) RunCli() {
	if len(os.Args) > 1 && os.Args[1] == "bench" {
		r.Benchmark()
		return
	}

	if len(os.Args) != 3 {
		log.Fatal("It expects two arguments: day and part.")
	}

	day := Atoi(os.Args[1])
	part := Atoi(os.Args[2])
	r.Run(day, part)
}

// Run runs the solution for a given day and part
func (r Runner) Run(day int, part int) {
	if day < 1 || day > 25 {
		log.Fatal("Invalid day")
	}

	if part < 1 || part > 2 {
		log.Fatal("Invalid part")
	}

	log.Printf("Running day %d, part %d", day, part)

	input := r.GetInput(day, part)

	t := time.Now()
	solution := r.daysParts[day-1][part-1](input)
	if solution == "" {
		log.Fatalf("Day %d part %d is not implemented", day, part)
	}

	log.Printf("Got solution in %s: %s", time.Since(t), solution)
	err := r.client.SendSolution(r.year, day, part, solution)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Success!")
}

// Benchmark runs a benchmark over all implemented solutions, and print the result.
func (r Runner) Benchmark() {
	stdout := os.Stdout
	defer func() { os.Stdout = stdout }()
	os.Stdout, _ = os.Open(os.DevNull)

	for day := range 25 {
		for part := range 2 {
			input := r.GetInput(day+1, part+1)
			result := r.daysParts[day][part](input)
			if result == "" {
				continue
			}

			bench := testing.Benchmark(func(b *testing.B) {
				for b.Loop() {
					r.daysParts[day][part](input)
				}
			})

			ns := float64(bench.NsPerOp())
			var timePerOp string
			if ns > 1_000_000 {
				timePerOp = fmt.Sprintf("%.2f ms/op", ns/1_000_000)
			} else if ns > 1000 {
				timePerOp = fmt.Sprintf("%.2f µs/op", ns/1000)
			} else {
				timePerOp = fmt.Sprintf("%.0f ns/op", ns)
			}

			bytes := float64(bench.AllocedBytesPerOp())
			var memPerOp string
			if bytes > 1_000_000 {
				memPerOp = fmt.Sprintf("%.2f Mb/op", bytes/1_000_000)
			} else if bytes > 1000 {
				memPerOp = fmt.Sprintf("%.2f kb/op", bytes/1000)
			} else {
				memPerOp = fmt.Sprintf("%.0f b/op", bytes)
			}

			fmt.Fprintf(stdout, "day %d, part %d: %s\t%s\n", day+1, part+1, timePerOp, memPerOp)
		}
	}
}

func (r Runner) GetInput(day int, part int) Input {
	input := r.cache.GetInput(r.year, day, part)
	if input.content != "" {
		return input
	}

	input, err := r.client.GetInput(r.year, day, part)
	if err != nil {
		log.Print("Error while getting input.")
		log.Fatal(err)
	}

	err = r.cache.StoreInput(r.year, day, part, input)
	if err != nil {
		log.Printf("Unable to save input to cache: %s", err)
	}

	return input
}
