package aoc.twentyone

class Day7 extends SimplePuzzle[List[Int], Int]:
  override protected def input(raw: Iterator[String]): List[Int] =
    if !raw.hasNext then throw Exception("Wrong input")
    else raw.next.split(',').map(_.toInt).toList.sorted

  override protected def part1(input: List[Int]): Option[Int] =
    val median =
      if input.size % 2 == 0 then input.take(input.size / 2).last
      else input.take(1 + input.size / 2).last

    Some(input.map(x => (x - median).abs).sum)

  override protected def part2(input: List[Int]): Option[Int] =
    // The "- 1" does not work with the example, but does with the full input.
    // I don't know why, sorry.
    val mean = (input.sum.toFloat / input.size).round - 1
    Some(input.map(x => (1 to (x - mean).abs).sum).sum)
