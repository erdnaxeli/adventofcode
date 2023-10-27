# AoC golang utilities

This package provides both a CLI tool to generate a default AoC project,
and a library to ease the solution writing.

## Generating an AoC project

### Installation

```shell
$ go install github.com/erdnaxeli/adventofcode/aoc/cmd/aoc
```

### Creating a new project

```
$ mkdir adventofcode2023
$ aoc
```

This command will generate the boilerplate for a project in the current directory.
It generates the following files:

* `main.go`: a `main()` function with all the wiring to run the project
* `day01.go`: the solutions for the day 1:
  * `func day1p1(input Input) string`: the solution for the part 1 of the day 1
  * `func day1p2(input Input) string`: the solution for the part 2 of the day 1
* and similar files for all the other days

### Usage

To use your project, you have to implement the solution in the corresponding file,
then you run it like this:

```shell
$ # firt export your adventofcode.com session cookie
$ export AOC_SESSION=...
$ # then run your solution
$ go build
$ ./adventofcode2023 1 1  # run the solution for the part 1 of the day 1
```

## The `aoc` library

This library provides many types, the two most important are:
* `Runner`: the code that will get your puzzle input, run your solution, and submit
  the result. You should not have to deal with it as the CLI utility create a `main.go`
  file with everything set up.
* `Input`: represent your puzzle input.

The `Runner` object wants a type implementing the `Solver` interface. This interface
defines methods `dayXpY()` for each day and each part.

Every method `dayXpY()` takes a single parameter `input Input` which is the puzzle input.
The type `Input` provides many method to parse the input. See the [type documentation]
for more details.

Every method must return a single string. If the string is empty, the `Runner` assumes
the solution is not implemented.

Have fun!
