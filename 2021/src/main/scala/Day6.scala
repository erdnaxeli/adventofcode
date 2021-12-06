package aoc.twentyone

class Lanternfish(val timer: Int):
  def nextTick: List[Lanternfish] =
    if timer == 0 then List(Lanternfish(6), Lanternfish(8))
    else List(Lanternfish(timer - 1))

class Lanternfishes(override val timer: Int, val count: Long)
    extends Lanternfish(timer):
  override def nextTick: List[Lanternfishes] =
    super.nextTick match
      case l :: Nil => List(Lanternfishes(l.timer, count))
      case l1 :: l2 :: Nil =>
        List(Lanternfishes(l1.timer, count), Lanternfishes(l2.timer, count))
      case _ => ???

class Day6 extends SimplePuzzle[List[Lanternfish], Long]:
  override protected def input(raw: Iterator[String]): List[Lanternfish] =
    if !raw.hasNext then throw Exception("Invalid input")
    else raw.next.split(',').map(t => Lanternfish(t.toInt)).toList

  override protected def part1(input: List[Lanternfish]): Option[Long] =
    Some(
      (0 until 80).foldLeft(input)((input, _) => input.flatMap(_.nextTick)).size
    )

  override protected def part2(input: List[Lanternfish]): Option[Long] =
    val groupedInput =
      input.groupBy(_.timer).map((k, v) => Lanternfishes(k, v.size))
    Some(
      (0 until 256)
        .foldLeft(groupedInput)((input, _) =>
          input
            .flatMap(_.nextTick)
            .groupBy(_.timer)
            .map((k, v) => Lanternfishes(k, v.map(_.count).sum))
        )
        .map(_.count)
        .sum
    )
