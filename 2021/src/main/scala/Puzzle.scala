package aoc.twentyone

import scala.io.Source
import scala.util.Failure
import scala.util.Success
import scala.util.Try

trait Puzzle[I1, I2, P1, P2]:
  protected def input1(raw: Iterator[String]): I1
  protected def input2(raw: Iterator[String]): I2
  protected def part1(input: I1): Option[P1] = None
  protected def part2(input: I2): Option[P2] = None

  def solve(day: Int): (Option[P1], Option[P2]) =
    val (r1, r2) =
      val it1 = Source.fromResource(f"$day%02d/p1.txt").getLines
      Try { Source.fromResource(f"$day%02d/p2.txt").getLines } match
        case Success(it2) => (it1, it2)
        case Failure(_)   => it1.duplicate

    (part1(input1(r1)), part2(input2(r2)))

trait SimplePuzzle[I, P] extends Puzzle[I, I, P, P]:
  protected def input(raw: Iterator[String]): I

  override protected def input1(raw: Iterator[String]): I = input(raw)
  override protected def input2(raw: Iterator[String]): I = input(raw)
