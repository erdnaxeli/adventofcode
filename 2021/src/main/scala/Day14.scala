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

  override protected def part2(input: Input14): Option[Int] =
    val (template, instructions) = input
    val counts = instructions.map((k, v) =>
      val x = iterate(k.toList, instructions, 40).groupBy(identity).map(_.size)
      println(x)
      x
    )
    counts.tapEach(println)
    Some(0)

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
