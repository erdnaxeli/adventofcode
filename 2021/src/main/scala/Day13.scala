package aoc.twentyone

import scala.annotation.tailrec

class Day13 extends SimplePuzzle[(Set[(Int, Int)], List[(Int, Int)]), Int]:
  def input(raw: Iterator[String]): (Set[(Int, Int)], List[(Int, Int)]) =
    @tailrec
    def readInput(
        input: Iterator[String],
        acc: (Set[(Int, Int)], List[(Int, Int)])
    ): (Set[(Int, Int)], List[(Int, Int)]) =
      if input.isEmpty then acc
      else
        val rPoint = raw"(\d+),(\d+)".r
        val rFold = raw"fold along (x|y)=(\d+)".r
        input.next match
          case rPoint(x, y) =>
            readInput(input, (acc(0) + ((x.toInt, y.toInt)), acc(1)))
          case rFold(axe, value) if axe == "x" =>
            readInput(input, (acc(0), acc(1) :+ (value.toInt, 0)))
          case rFold(axe, value) if axe == "y" =>
            readInput(input, (acc(0), acc(1) :+ (0, value.toInt)))
          case "" => readInput(input, acc)
          case _  => ???

    readInput(raw, (Set.empty[(Int, Int)], List.empty[(Int, Int)]))

  override protected def part1(
      input: (Set[(Int, Int)], List[(Int, Int)])
  ): Option[Int] =
    val (points, folds) = input
    val firstFold = folds.head
    Some(
      applyFold(points, firstFold).size
    )

  override protected def part2(
      input: (Set[(Int, Int)], List[(Int, Int)])
  ): Option[Int] =
    val (points, folds) = input
    val result =
      folds.foldLeft(points)((points, fold) => applyFold(points, fold))

    val maxX =
      result.foldLeft(0)((max, p) => if p(0) > max then p(0) else max)
    val maxY =
      result.foldLeft(0)((max, p) => if p(1) > max then p(1) else max)

    for y <- 0 to maxY do
      for x <- 0 to maxX do print(if result.contains(x, y) then "#" else ".")
      println

    Some(0)

  private def applyFold(
      points: Set[(Int, Int)],
      fold: (Int, Int)
  ): Set[(Int, Int)] =
    val (fx, fy) = fold
    points
      .foldLeft(Set.empty[(Int, Int)])((acc, point) =>
        val (x, y) = point
        val newX =
          if fx == 0 then x
          else fx - (x - fx).abs
        val newY =
          if fy == 0 then y
          else fy - (y - fy).abs
        acc + ((newX, newY))
      )
