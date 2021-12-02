package aoc.twentyone

import javax.print.attribute.standard.PresentationDirection

enum Direction:
  case Forward, Up, Down

case class MoveInstruction(direction: Direction, count: Int)

class Day2 extends aoc.twentyone.SimplePuzzle[List[MoveInstruction], Int]:
  override protected def input(raw: Iterator[String]): List[MoveInstruction] =
    raw
      .map(s =>
        s.split(' ')
          .toList
          .match
            case "forward" :: c :: Nil =>
              MoveInstruction(Direction.Forward, c.toInt)
            case "up" :: c :: Nil =>
              MoveInstruction(Direction.Up, c.toInt)
            case "down" :: c :: Nil =>
              MoveInstruction(Direction.Down, c.toInt)
            case x => throw Exception(s"Unknown instruction $x")
      )
      .toList

  override protected def part1(input: List[MoveInstruction]): Int =
    input
      .foldLeft((0, 0))((acc, mi) =>
        mi match
          case MoveInstruction(Direction.Forward, c) => (acc(0) + c, acc(1))
          case MoveInstruction(Direction.Up, c)      => (acc(0), acc(1) - c)
          case MoveInstruction(Direction.Down, c)    => (acc(0), acc(1) + c)
      )
      .toList
      .product

  override protected def part2(input: List[MoveInstruction]): Int =
    input
      .foldLeft((0, 0, 0))((acc, mi) =>
        mi match
          case MoveInstruction(Direction.Forward, c) =>
            (acc(0), acc(1) + c, acc(2) + c * acc(0))
          case MoveInstruction(Direction.Up, c) => (acc(0) - c, acc(1), acc(2))
          case MoveInstruction(Direction.Down, c) =>
            (acc(0) + c, acc(1), acc(2))
      )
      .tail
      .toList
      .product
