package aoc.twentyone

class Day1 extends SimplePuzzle[List[Int], Int]:
  override def input(raw: Iterator[String]): List[Int] =
    raw.map(_.toInt).toList

  override def part1(input: List[Int]): Int =
    input.sliding(2).count(p => p(0) < p(1))

  override def part2(input: List[Int]): Int =
    input
      .sliding(3)
      .map(_.sum)
      .sliding(2)
      .count(p => p(0) < p(1))
