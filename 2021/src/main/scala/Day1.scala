package aoc.twentyone

class Day1 extends SimplePuzzle[List[Int], Int]:
  override def input(raw: Iterator[String]): List[Int] =
    raw.map(_.toInt).toList

  override def part1(input: List[Int]): Int =
    input
      .foldLeft((0, None): Tuple2[Int, Option[Int]])((acc, next) =>
        acc match
          case (increase, Some(value)) if next > value =>
            (increase + 1, Some(next))
          case (increase, _) => (increase, Some(next))
      )(0)

  override def part2(input: List[Int]): Int =
    input
      .sliding(3)
      .foldLeft((0, None): Tuple2[Int, Option[Int]])((acc, next) =>
        acc match
          case (increase, Some(value)) if (next.sum) > value =>
            (increase + 1, Some(next.sum))
          case (increase, _) =>
            (increase, Some(next.sum))
      )(0)
