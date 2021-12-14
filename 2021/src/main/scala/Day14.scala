package aoc.twentyone

import scala.annotation.tailrec
import scala.compiletime.ops.int

type Input14 = (List[Char], Map[(Char, Char), Char])

class Day14 extends SimplePuzzle[Input14, Int]:
  def input(raw: Iterator[String]): Input14 =
    @tailrec
    def readInput(
        input: Iterator[String],
        template: List[Char],
        instructions: Map[(Char, Char), Char]
    ): Input14 =
      if !input.hasNext then (template, instructions)
      else
        val regex = raw"([A-Z])([A-Z]) -> ([A-Z])".r
        input.next match
          case "" => readInput(input, template, instructions)
          case regex(a, b, c) =>
            readInput(input, template, instructions.updated((a(0), b(0)), c(0)))
          case x => readInput(input, x.toCharArray.toList, instructions)

    readInput(raw, List.empty[Char], Map.empty[(Char, Char), Char])

  override protected def part1(input: Input14): Option[Int] =
    val result = iterate(input(0), input(1), 10)
      .groupBy(identity)
      .values
      .toList
      .sortBy(_.size)
    Some(result.last.size - result.head.size)

  private def nextStep(
      template: List[Char],
      instructions: Map[(Char, Char), Char]
  ): List[Char] =
    template.head :: template
      .sliding(2)
      .flatMap(v => List(instructions((v(0), v(1))), v(1)))
      .toList

  @tailrec
  private def iterate(
      template: List[Char],
      instructions: Map[(Char, Char), Char],
      count: Int
  ): List[Char] =
    if count == 0 then template
    else iterate(nextStep(template, instructions), instructions, count - 1)

  override protected def part2(input: Input14): Option[Int] =
    val (template, instructions) = input
    val countersInit =
      instructions.map((k, v) =>
        (
          (k(0), k(1), 1),
          List(v).groupBy(identity).mapValues(_.size.toLong).toMap
        )
      )
    val result =
      countCharsTemplate(template, countersInit, 40).values.toList.sorted

    // Print the result because I don't want to change everything in part 1 to
    // int.
    println(result.last - result.head)
    Some(-1)

  private def countCharsTemplate(
      template: List[Char],
      counters: Map[(Char, Char, Int), Map[Char, Long]],
      iterations: Int
  ): Map[Char, Long] =
    sumMap(
      template
        .sliding(2)
        .foldLeft((counters, Map.empty[Char, Long]))((acc, pair) =>
          val (counters, result) = acc
          val x :: y :: Nil = pair
          val nextCounters = countChars(x, y, counters, iterations)
          (
            nextCounters,
            sumMap(
              sumMap(result, nextCounters((x, y, iterations))),
              // We count the first char.
              Map(x -> 1)
            )
          )
        )(1),
      // Finally we count the last char which was never counted.
      Map(template.last -> 1)
    )

  /* Return a copy of `counters` with an entry `(x, y, iterations)`.
   *
   * `counters((x, y, iterations))` is a map associating, for each distinct char
   * that would be added between `x` and `y` after `iterations` iterations, the
   * count of this char.
   * If the entry `(x, y, iterations)` already exist, the copy is returned without
   * modification.
   * `counters` must have an entry for each pair of char and a count of iterations
   * of 1 for the algorithm to work.
   */
  private def countChars(
      x: Char,
      y: Char,
      counters: Map[(Char, Char, Int), Map[Char, Long]],
      iterations: Int
  ): Map[(Char, Char, Int), Map[Char, Long]] =
    if counters.contains((x, y, iterations)) then counters
    else
      // Get the char to insert between x and y.
      // We know if we do only one iteration there can be only one chart to
      // insert, so the counter wil be Map(char -> 1).
      val newChar = counters((x, y, 1)).keys.head
      // Get the updated counters for [x, newChar], iterated one iteration less.
      val counters2 =
        countChars(x, newChar, counters, iterations - 1)
      // Get the updated counters for [newChar, X], iterated one iteration less.
      val counters3 =
        countChars(newChar, y, counters2, iterations - 1)

      // The count of chars after iterating [x, y] n times is the sum of the chars
      // after iterating [x, newChar] (n-1) times and [newChar, y] (n-1) times.
      val counterXY =
        sumMap(
          sumMap(
            counters3((x, newChar, iterations - 1)),
            counters3((newChar, y, iterations - 1))
          ),
          // We need to add newChar, with is the junction between [x, newChar] and
          // [newChar, y].
          Map(newChar -> 1)
        )

      // Finally return the counters with the result for the given values.
      counters3 ++ Map((x, y, iterations) -> counterXY)

  // Sum two map, summing each value with the same key.
  private def sumMap[T](map1: Map[T, Long], map2: Map[T, Long]): Map[T, Long] =
    map1 ++ map2.map((k, v) => (k, v + map1.getOrElse(k, 0L)))
