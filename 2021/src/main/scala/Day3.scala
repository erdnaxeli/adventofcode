package aoc.twentyone

import scala.math.pow

class Day3 extends SimplePuzzle[List[List[Int]], Double]:
  override protected def input(raw: Iterator[String]): List[List[Int]] =
    raw.map(_.toList.map(_.toString.toInt)).toList

  override protected def part1(input: List[List[Int]]): Option[Double] =
    val mostCommonBytes =
      input
        .map(_.reverse)
        .foldLeft(Map[Int, Int]())((acc, x) =>
          acc ++ x.zipWithIndex.map((v, i) => (i, v + acc.getOrElse(i, 0)))
        )
        .map((i, v) => (i, (v.toFloat / input.size).round))
    val leasCommonBytes = mostCommonBytes.map((i, v) => (i, (v + 1) % 2))
    val gamma = mostCommonBytes.map((i, v) => pow(2, i) * v).sum
    val epsilon = leasCommonBytes.map((i, v) => pow(2, i) * v).sum
    Some(gamma * epsilon)
