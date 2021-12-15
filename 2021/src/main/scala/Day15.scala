package aoc.twentyone

import scala.annotation.tailrec

class Day15 extends SimplePuzzle[Map[(Int, Int), Int], Int]:
  def input(raw: Iterator[String]): Map[(Int, Int), Int] =
    raw.zipWithIndex
      .flatMap((line, y) =>
        line.zipWithIndex.map((risk, x) => ((x, y), risk.toString.toInt))
      )
      .toMap

  override protected def part1(input: Map[(Int, Int), Int]): Option[Int] =
    val end = input.keys.max
    Some(
      findPaths(
        end,
        input,
        Map(
          (0, 0) -> (
            0, List((0, 0))
          )
        )
      )
    )

  override protected def part2(input: Map[(Int, Int), Int]): Option[Int] =
    @tailrec
    def incrRisk(risk: Int, incr: Int): Int =
      if incr <= 0 then risk
      else
        val newRisk = risk + 1
        if newRisk > 9 then incrRisk(1, incr - 1)
        else incrRisk(newRisk, incr - 1)

    val endSmall = input.keys.max
    val bigInput = input
      .flatMap((point, risk) =>
        Map(
          point -> risk,
          (point(0) + (1 + endSmall(0)), point(1)) -> incrRisk(risk, 1),
          (point(0) + (1 + endSmall(0)) * 2, point(1)) -> incrRisk(risk, 2),
          (point(0) + (1 + endSmall(0)) * 3, point(1)) -> incrRisk(risk, 3),
          (point(0) + (1 + endSmall(0)) * 4, point(1)) -> incrRisk(risk, 4)
        )
      )
      .flatMap((point, risk) =>
        Map(
          point -> risk,
          (point(0), point(1) + (1 + endSmall(1))) -> incrRisk(risk, 1),
          (point(0), point(1) + (1 + endSmall(1)) * 2) -> incrRisk(risk, 2),
          (point(0), point(1) + (1 + endSmall(1)) * 3) -> incrRisk(risk, 3),
          (point(0), point(1) + (1 + endSmall(1)) * 4) -> incrRisk(risk, 4)
        )
      )
      .map((point, risk) => (point, if risk == 0 then 1 else risk))

    val end = bigInput.keys.max

    Some(
      findPaths(
        end,
        bigInput,
        Map(
          (0, 0) -> (
            0, List((0, 0))
          )
        )
      )
    )

  @tailrec
  private def findPaths(
      end: (Int, Int),
      map: Map[(Int, Int), Int],
      costs: Map[(Int, Int), (Int, List[(Int, Int)])],
      done: List[(Int, Int)] = List.empty[(Int, Int)]
  ): Int =
    // Select the reachable point with the lowest risk.
    val (point, v) = costs.minBy((_, v) => v(0))
    val (currentRisk, currentPath) = v
    if point == end then currentRisk
    else
      // Update points around it.
      // for each point defined
      val newCosts = for
        x <- (point(0) - 1) to (point(0) + 1)
        y <- (point(1) - 1) to (point(1) + 1)
        if (x == point(0)) != (y == point(1)) && !done.contains((x, y))
        neighbor = map.get((x, y))
        if neighbor.isDefined
      yield
        val Some(neighborsCost) = neighbor
        // compute the risk to reach this point
        val risk = currentRisk + neighborsCost
        // is this risk better than a risk using another path?
        val newRisk = costs.get((x, y)) match
          case Some((knownRisk, _)) if knownRisk < risk => knownRisk
          case _                                        => risk
        // return the point with the updated risk and path
        ((x, y), (newRisk, (x, y) :: currentPath))

      findPaths(
        end,
        map,
        // We remove the current point from the available points
        costs.filter(_._1 != point) ++ newCosts,
        // but we save it to not go to it again.
        point :: done
      )
