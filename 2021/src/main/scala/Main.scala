package aoc.twentyone

@main def run(day: Int): Unit =
  val days = Map(
    1 -> Day1(),
    2 -> Day2(),
    3 -> Day3(),
    4 -> Day4(),
    5 -> Day5(),
    6 -> Day6()
  )
  println(days(day).solve(day))
