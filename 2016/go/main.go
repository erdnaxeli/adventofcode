package main

import "github.com/erdnaxeli/adventofcode/aoc"

type solver struct{}

func main() {
	var solver solver
	runner := aoc.NewDefaultRunner(2016, solver)
	runner.RunCli()
}
