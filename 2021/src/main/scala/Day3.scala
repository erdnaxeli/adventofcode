package aoc.twentyone

import scala.math.pow

class Day3 extends SimplePuzzle[List[List[Int]], Double]:
  override protected def input(raw: Iterator[String]): List[List[Int]] =
    raw.map(_.toList.map(_.toString.toInt)).toList

  override protected def part1(input: List[List[Int]]): Option[Double] =
    val mostCommonBytes =
      input.transpose
        .map(_.groupBy(identity).mapValues(_.size).maxBy(_._2)._1)
        .reverse
        .zipWithIndex
    val leastCommonBytes = mostCommonBytes.map((v, i) => ((v + 1) % 2, i))
    val gamma = mostCommonBytes.map((v, i) => pow(2, i) * v).sum
    val epsilon = leastCommonBytes.map((v, i) => pow(2, i) * v).sum
    Some(gamma * epsilon)

  override protected def part2(input: List[List[Int]]): Option[Double] =
    def findMatching(
        input: List[List[Int]],
        cmp: (Int, Int) => Boolean,
        acc: List[Int] = List[Int]()
    ): List[Int] =
      input match
        case Nil      => ???
        case x :: Nil => acc ++ x
        case _ =>
          val mostCommon = input.transpose.head
            .groupBy(identity)
            .mapValues(_.size)
            .maxBy(_._2)
            ._1
          var newInput = input.filter(x => cmp(x.head, mostCommon))
          findMatching(
            newInput.map(_.tail),
            cmp,
            acc :+ newInput.head.head
          )

    val oxygen = findMatching(input, (x, y) => x == y).reverse.zipWithIndex
      .map((v, i) => pow(2, i) * v)
      .sum
    val co2Scrubber = findMatching(input, (x, y) => x != y).reverse.zipWithIndex
      .map((v, i) => pow(2, i) * v)
      .sum

    Some(oxygen * co2Scrubber)
