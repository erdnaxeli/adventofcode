package aoc.twentyone

case class Line(x1: Int, y1: Int, x2: Int, y2: Int)
    extends Iterable[(Int, Int)]:
  def iterator =
    val rx =
      if x1 < x2 then x1 to x2
      else (x2 to x1).reverse
    val ry =
      if y1 < y2 then y1 to y2
      else (y2 to y1).reverse

    // This works because lines are either horizontal, vertical or 45°.
    rx.iterator.zipAll(ry, x1, y1)

class Day5 extends SimplePuzzle[List[Line], Int]:
  override protected def input(raw: Iterator[String]): List[Line] =
    val regex = raw"(\d+),(\d+) -> (\d+),(\d+)".r
    raw.collect { case regex(x1, y1, x2, y2) =>
      Line(x1.toInt, y1.toInt, x2.toInt, y2.toInt)
    }.toList

  override protected def part1(input: List[Line]): Option[Int] =
    Some(
      input
        .filter(l => l.x1 == l.x2 || l.y1 == l.y2)
        .flatten
        .foldLeft(Map[(Int, Int), Int]())((acc, point) =>
          acc.updated(point, acc.getOrElse(point, 0) + 1)
        )
        .filter((k, v) => v > 1)
        .size
    )

  override protected def part2(input: List[Line]): Option[Int] =
    Some(
      input.flatten
        .foldLeft(Map[(Int, Int), Int]())((acc, point) =>
          acc.updated(point, acc.getOrElse(point, 0) + 1)
        )
        .filter((k, v) => v > 1)
        .size
    )
