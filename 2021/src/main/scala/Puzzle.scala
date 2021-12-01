package aoc.twentyone

import scala.io.Source

trait Puzzle[I1, I2, P1, P2]:
  protected def input1(raw: Iterator[String]): I1
  protected def input2(raw: Iterator[String]): I2
  protected def part1(input: I1): P1
  protected def part2(input: I2): P2

  def solve(day: Int): (P1, P2) =
    val r1 = Source.fromResource(f"$day%02d/p1.txt").getLines
    val r2 = Source.fromResource(f"$day%02d/p2.txt").getLines

    (part1(input1(r1)), part2(input2(r2)))

trait SimplePuzzle[I, P] extends Puzzle[I, I, P, P]:
  protected def input(raw: Iterator[String]): I

  override protected def input1(raw: Iterator[String]): I = input(raw)
  override protected def input2(raw: Iterator[String]): I = input(raw)
