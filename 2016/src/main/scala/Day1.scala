package aoc.twentysixteen

enum MoveDirection:
  case Right, Left

case class MoveInstruction(val direction: MoveDirection, val count: Int)

class Day1 extends SimplePuzzle[List[MoveInstruction], Int]:
  import MoveDirection.*

  override protected def input(raw: Iterator[String]): List[MoveInstruction] =
    val regex = raw"(R|L)(\d+)".r
    if !raw.hasNext then ???
    else
      raw.next
        .split(", ")
        .collect {
          case regex(d, c) if d == "R" => MoveInstruction(Right, c.toInt)
          case regex(d, c) if d == "L" => MoveInstruction(Left, c.toInt)
        }
        .toList

  override protected def part1(input: List[MoveInstruction]): Option[Int] =
    def solve(
        input: List[MoveInstruction],
        position: (Int, Int) = (0, 0),
        currentDirection: (Int, Int) = (0, 1)
    ): (Int, Int) =
      input match
        case Nil => position
        case MoveInstruction(d, count) :: tail =>
          val newDirection = getNewDirection(d, currentDirection)
          solve(
            tail,
            (
              position(0) + newDirection(0) * count,
              position(1) + newDirection(1) * count
            ),
            newDirection
          )

    val finalPosition = solve(input)
    Some(finalPosition(0).abs + finalPosition(1).abs)

  private def getNewDirection(
      direction: MoveDirection,
      currentDirection: (Int, Int)
  ): (Int, Int) =
    direction match
      case Right => (currentDirection(1), currentDirection(0) * (-1))
      case Left  => (currentDirection(1) * (-1), currentDirection(0))
