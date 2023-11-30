package main

import "github.com/erdnaxeli/adventofcode/aoc"

type solver struct{}

func main() {
	var solver solver
	runner := aoc.NewDefaultRunner(2023, solver)
	runner.RunCli()
}
