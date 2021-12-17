package aoc.twentyone

import scala.annotation.tailrec

type Input17 = ((Int, Int), (Int, Int))

class Day17 extends SimplePuzzle[Input17, Int]:
  def input(raw: Iterator[String]): Input17 =
    if raw.isEmpty then ???
    else
      val regex = raw"target area: x=(\d+)..(\d+), y=(-?\d+)..(-?\d+)".r
      raw.next match
        case regex(x1, x2, y1, y2) =>
          ((x1.toInt, y1.toInt), (x2.toInt, y2.toInt))
        case _ => ???

  override protected def part1(input: Input17): Option[Int] =
    val ((x1, y1), (x2, y2)) = input

    // The move on the Y axis is independent of the one on the X axis.
    // We known the probe will raise then pass again to a point where y=0.
    // We need to arrive to the lowest possible point (y1) with the maximum
    // velocity (0 - y1). Then we sum all the velocity from the highest point
    // (velocity 0) to the lowest point (velocity -y1) to get the distance between
    // those two point. We add y1 to get the distance from 0 to the highest point.
    // Trust me it works.
    Some(y1 + (0 to (-y1)).sum)

  override protected def part2(input: Input17): Option[Int] =
    val ((x1, y1), (x2, y2)) = input
    val xVelocities = allX(x1, x2)
    val maxHeight = y1 + (0 to (-y1)).sum
    // Assume the max X velocity, find the initial Y velocity so Y reach
    // maxHeight when X is zero. This value is way bigger than the actual maxY,
    // but we just want an upper bound, not the exact max.
    val maxYV = (0 to xVelocities.max).sum

    val points = for
      yV <- xVelocities
      yV <- y1 to maxYV
    yield
      if goThrough(yV, yV, x1, y1, x2, y2) then Some(yV, yV)
      else None
    Some(points.count(_.isDefined))

  private def allX(x1: Int, x2: Int): List[Int] =
    @tailrec
    def allXRec(
        x1: Int,
        x2: Int,
        acc: List[Int] = List.empty[Int],
        x: Int = 0
    ): List[Int] =
      if x > x2 then acc
      else
        val distances = (0 to x).scan(0)(_ + _)
        if distances.find(d => x1 <= d && d <= x2).isDefined then
          allXRec(x1, x2, x :: acc, x + 1)
        else allXRec(x1, x2, acc, x + 1)

    allXRec(x1, x2)

  private def goThrough(
      xV: Int,
      yV: Int,
      x1: Int,
      y1: Int,
      x2: Int,
      y2: Int
  ): Boolean =
    @tailrec
    def goThroughRec(xV: Int, yV: Int, x: Int = 0, y: Int = 0): Boolean =
      if x > x2 || (yV < 0 && y < y1) then false
      else if x1 <= x && x <= x2 && y1 <= y && y <= y2 then true
      else goThroughRec(List(xV - 1, 0).max, yV - 1, x + xV, y + yV)

    goThroughRec(xV, yV)
