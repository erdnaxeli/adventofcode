package aoc.twentyone

import scala.annotation.tailrec

enum SeaCucumber:
  case East, South

class Day25 extends SimplePuzzle[(Map[(Int, Int), SeaCucumber], Int, Int), Int]:
  override protected def input(
      raw: Iterator[String]
  ): (Map[(Int, Int), SeaCucumber], Int, Int) =
    val rawL = raw.toList
    val maxY = rawL.size - 1
    val maxX = rawL.map(_.size - 1).max
    val map = rawL.zipWithIndex
      .flatMap((line, y) =>
        line.zipWithIndex
          .flatMap((p, x) =>
            p match
              case '.' => None
              case '>' => Some((x, y) -> SeaCucumber.East)
              case 'v' => Some((x, y) -> SeaCucumber.South)
              case _   => ???
          )
      )
      .toMap

    (map, maxX, maxY)

  override protected def part1(
      input: (Map[(Int, Int), SeaCucumber], Int, Int)
  ): Option[Int] =
    @tailrec
    def rec(
        input: Map[(Int, Int), SeaCucumber],
        maxX: Int,
        maxY: Int,
        steps: Int = 0
    ): Int =
      val nextState = move(input, maxX, maxY)
      if nextState == input then steps + 1
      else rec(nextState, maxX, maxY, steps + 1)

    Some(rec(input(0), input(1), input(2)))

  private def move(
      input: Map[(Int, Int), SeaCucumber],
      maxX: Int,
      maxY: Int
  ): Map[(Int, Int), SeaCucumber] =
    moveHerd(
      moveHerd(input, SeaCucumber.East, maxX).map((k, v) => (k(1), k(0)) -> v),
      SeaCucumber.South,
      maxY
    ).map((k, v) => (k(1), k(0)) -> v)

  private def moveHerd(
      input: Map[(Int, Int), SeaCucumber],
      herd: SeaCucumber,
      maxX: Int
  ): Map[(Int, Int), SeaCucumber] =
    input
      .map((k, seaCucumber) =>
        if seaCucumber == herd then
          val (x, y) = k
          val toTest =
            if x + 1 > maxX then (0, y) else (x + 1, y)
          input.get(toTest) match
            case None => toTest -> seaCucumber
            case _    => (x, y) -> seaCucumber
        else k -> seaCucumber
      )
