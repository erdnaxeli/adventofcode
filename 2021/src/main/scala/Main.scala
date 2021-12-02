package aoc.twentyone

@main def run(day: Int): Unit =
  val days = Map(
    1 -> Day1(),
    2 -> Day2()
  )
  println(days(day).solve(day))
