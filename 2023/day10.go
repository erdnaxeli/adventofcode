package main

import (
	"fmt"
	"log"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day10p1(input aoc.Input) string {
	grid := input.ToStringSlice()
	xs, ys := getPipeStart(grid)

	var x, y int
	var move byte
	if xs > 0 && (grid[xs-1][ys] == '|' || grid[xs-1][ys] == 'F' || grid[xs-1][ys] == '7') {
		x, y, move = xs-1, ys, 'N'
	} else if ys > 0 && (grid[xs][ys-1] == '-' || grid[xs][ys-1] == 'F' || grid[xs][ys-1] == 'L') {
		x, y, move = xs, ys-1, 'W'
	} else if xs < len(grid)-1 && (grid[xs+1][ys] == '|' || grid[xs+1][ys] == 'L' || grid[xs+1][ys] == 'J') {
		x, y, move = xs+1, ys, 'S'
	} else if ys < len(grid[0])-1 && (grid[xs][ys+1] == '-' || grid[xs][ys+1] == '7' || grid[xs][ys+1] == 'J') {
		x, y, move = xs, ys+1, 'E'
	}

	steps := 1
	for {
		switch grid[x][y] {
		case '|':
			if move == 'N' {
				x--
			} else {
				x++
			}
		case 'L':
			if move == 'S' {
				move = 'E'
				y++
			} else {
				move = 'N'
				x--
			}
		case '-':
			if move == 'E' {
				y++
			} else {
				y--
			}
		case 'J':
			if move == 'E' {
				move = 'N'
				x--
			} else {
				move = 'W'
				y--
			}
		case '7':
			if move == 'N' {
				move = 'W'
				y--
			} else {
				move = 'S'
				x++
			}
		case 'F':
			if move == 'W' {
				move = 'S'
				x++
			} else {
				move = 'E'
				y++
			}
		}

		steps++
		if grid[x][y] == 'S' {
			break
		}
	}

	return aoc.ResultI(steps / 2)
}

func (s solver) Day10p2(input aoc.Input) string {
	var grid [][]byte
	for _, line := range input.ToStringSlice() {
		grid = append(grid, []byte(line))
	}

	// remove pipes other than the loop
	xs, ys := getPipeStartB(grid)
	loopPipes := make(map[aoc.Point]bool)
	loopPipes[aoc.Point{X: xs, Y: ys}] = true
	// This function is too big, I don't know which variables are used where, so as this code is almost copied from p1
	// in doubt I put it in its own scope to not alter the rest of the function.
	{
		var x, y int
		var move byte
		if xs > 0 && (grid[xs-1][ys] == '|' || grid[xs-1][ys] == 'F' || grid[xs-1][ys] == '7') {
			x, y, move = xs-1, ys, 'N'
		} else if ys > 0 && (grid[xs][ys-1] == '-' || grid[xs][ys-1] == 'F' || grid[xs][ys-1] == 'L') {
			x, y, move = xs, ys-1, 'W'
		} else if xs < len(grid)-1 && (grid[xs+1][ys] == '|' || grid[xs+1][ys] == 'L' || grid[xs+1][ys] == 'J') {
			x, y, move = xs+1, ys, 'S'
		} else if ys < len(grid[0])-1 && (grid[xs][ys+1] == '-' || grid[xs][ys+1] == '7' || grid[xs][ys+1] == 'J') {
			x, y, move = xs, ys+1, 'E'
		}

		for {
			loopPipes[aoc.Point{X: x, Y: y}] = true

			switch grid[x][y] {
			case '|':
				if move == 'N' {
					x--
				} else {
					x++
				}
			case 'L':
				if move == 'S' {
					move = 'E'
					y++
				} else {
					move = 'N'
					x--
				}
			case '-':
				if move == 'E' {
					y++
				} else {
					y--
				}
			case 'J':
				if move == 'E' {
					move = 'N'
					x--
				} else {
					move = 'W'
					y--
				}
			case '7':
				if move == 'N' {
					move = 'W'
					y--
				} else {
					move = 'S'
					x++
				}
			case 'F':
				if move == 'W' {
					move = 'S'
					x++
				} else {
					move = 'E'
					y++
				}
			}

			if grid[x][y] == 'S' {
				break
			}
		}
	}

	for x := range grid {
		for y := range grid[0] {
			if !loopPipes[aoc.Point{X: x, Y: y}] {
				grid[x][y] = '.'
			}
		}
	}

	// replace start by the correct pipe
	if xs > 0 && (grid[xs-1][ys] == '|' || grid[xs-1][ys] == 'F' || grid[xs-1][ys] == '7') {
		if ys > 0 && (grid[xs][ys-1] == '-' || grid[xs][ys-1] == 'F' || grid[xs][ys-1] == 'L') {
			grid[xs][ys] = 'J'
		} else if xs < len(grid)-1 && (grid[xs+1][ys] == '|' || grid[xs+1][ys] == 'L' || grid[xs+1][ys] == 'J') {
			grid[xs][ys] = '|'
		} else if ys < len(grid[0])-1 && (grid[xs][ys+1] == '-' || grid[xs][ys+1] == '7' || grid[xs][ys+1] == 'J') {
			grid[xs][ys] = 'L'
		}
	} else if ys > 0 && (grid[xs][ys-1] == '-' || grid[xs][ys-1] == 'F' || grid[xs][ys-1] == 'L') {
		if xs < len(grid)-1 && (grid[xs+1][ys] == '|' || grid[xs+1][ys] == 'L' || grid[xs+1][ys] == 'J') {
			grid[xs][ys] = '7'
		} else if ys < len(grid[0])-1 && (grid[xs][ys+1] == '-' || grid[xs][ys+1] == '7' || grid[xs][ys+1] == 'J') {
			grid[xs][ys] = '-'
		}
	} else {
		grid[xs][ys] = 'F'
	}

	// fix corridors
	//
	// ..----..|.
	// ..----....
	// becomes
	// ..----..|.
	//         |
	// ..----....
	//
	// Same for vertical corridors.
	//
	// We mark new void by spaces to not count them.
	mX := len(grid) - 2
	mY := len(grid[0]) - 2
	for x := 0; x <= mX; x++ {
		for y := 0; y <= mY; y++ {
			if (grid[x][y] == '-' || grid[x][y] == 'L' || grid[x][y] == 'J') && (grid[x+1][y] == '-' || grid[x+1][y] == '7' || grid[x+1][y] == 'F') {
				// add an empty row
				var row []byte
				for yy := 0; yy < len(grid[0]); yy++ {
					switch grid[x][yy] {
					case '|', '7', 'F':
						row = append(row, '|')
					default:
						row = append(row, ' ')
					}
				}

				grid = append(grid[:x+2], grid[x+1:]...)
				grid[x+1] = row
				mX++

			}

			if (grid[x][y] == '|' || grid[x][y] == 'J' || grid[x][y] == '7') && (grid[x][y+1] == '|' || grid[x][y+1] == 'L' || grid[x][y+1] == 'F') {
				// add empty column
				for xx := 0; xx < len(grid); xx++ {
					switch grid[xx][y] {
					case '-', 'F', 'L':
						grid[xx] = append(grid[xx][:y+2], grid[xx][y+1:]...)
						grid[xx][y+1] = '-'
					default:
						grid[xx] = append(grid[xx][:y+2], grid[xx][y+1:]...)
						grid[xx][y+1] = ' '
					}
				}

				mY++
			}
		}
	}

	// find groups of continuous void
	var surrounded int
	var surroundedGroud rune
	markerChar := 'A'
	for x := range grid {
		for y := range grid[x] {
			// log.Print(">>>>", x, y)
			isInternal := true
			if grid[x][y] != '.' && grid[x][y] != ' ' {
				continue
			}

			marked := 0
			// we use toMark like a set
			toMark := make(map[aoc.Point]bool)
			toMark[aoc.Point{X: x, Y: y}] = true
			// log.Printf("Found void!")
			for len(toMark) > 0 {
				// get a random key
				var p aoc.Point
				for p = range toMark {
					break
				}

				delete(toMark, p)

				if p.X == 0 || p.Y == 0 || p.X == len(grid)-1 || p.Y == len(grid[0])-1 {
					isInternal = false
				}

				if grid[p.X][p.Y] == '.' {
					marked++
				}
				grid[p.X][p.Y] = byte(markerChar)

				for xx := max(0, p.X-1); xx <= min(len(grid)-1, p.X+1); xx++ {
					for yy := max(0, p.Y-1); yy <= min(len(grid[0])-1, p.Y+1); yy++ {
						if grid[xx][yy] == '.' || grid[xx][yy] == ' ' {
							toMark[aoc.Point{X: xx, Y: yy}] = true
						}
					}
				}
			}

			if isInternal {
				surrounded = marked
				surroundedGroud = markerChar
			}

			// to ease debug
			markerChar++
			if markerChar == 'F' {
				markerChar++
			}
		}
	}

	// debug
	for _, line := range grid {
		for _, char := range line {
			if char == byte(surroundedGroud) {
				fmt.Printf("\033[1;31m%c\033[0m", char)
			} else if char == '|' || char == '-' || char == 'L' || char == 'F' || char == 'J' || char == '7' {
				fmt.Printf("\033[1;32m%c\033[0m", char)
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
	fmt.Println()
	log.Print("'", string(markerChar), "'")
	return aoc.ResultI(surrounded)
}

func getPipeStart(grid []aoc.String) (int, int) {
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'S' {
				return x, y
			}
		}
	}

	panic("start not found")
}

func getPipeStartB(grid [][]byte) (int, int) {
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'S' {
				return x, y
			}
		}
	}

	panic("start not found")
}
