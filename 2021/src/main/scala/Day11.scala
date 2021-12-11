package aoc.twentyone

import scala.annotation.tailrec

type Grid = Map[(Int, Int), Int]

class Day11 extends SimplePuzzle[Grid, Int]:
  override protected def input(raw: Iterator[String]): Grid =
    raw.zipWithIndex
      .flatMap((line, y) =>
        line.zipWithIndex.map((e, x) => ((x, y), e.toString.toInt))
      )
      .toMap

  override protected def part1(input: Grid): Option[Int] =
    var grid = input
    var flashes = 0
    for s <- 1 to 100 do
      grid = newStep(grid)
      flashes += grid.count(_._2 == 0)

    Some(flashes)

  override protected def part2(input: Grid): Option[Int] =
    @tailrec
    def findSync(grid: Grid, step: Int = 0): Int =
      if grid.forall(_._2 == 0) then step
      else findSync(newStep(grid), step + 1)

    Some(findSync(input))

  private def newStep(grid: Grid): Grid =
    @tailrec
    def stabilize(grid: Grid, flashes: Int = 0): Grid =
      val newGrid = grid
        .map((k, e) =>
          // has already flashed this step
          if e == 0 then (k, e)
          // will flash anyway
          else if e > 9 then (k, 0)
          else
            // add energy from flashes around
            val (x, y) = k
            val flashesAround = (
              for
                i <- (x - 1) to (x + 1)
                j <- (y - 1) to (y + 1)
                if (i, j) != (x, y)
              yield grid.get((i, j))
            ).flatten.count(e => e > 9)
            (k, e + flashesAround)
        )

      if newGrid == grid then grid
      else stabilize(newGrid)

    stabilize(grid.map((k, e) => (k, e + 1)))

  private def print(grid: Grid): Unit =
    for y <- 0 to 9 do
      println((0 to 9).map(x => f"${grid((x, y))}%2d").mkString)
