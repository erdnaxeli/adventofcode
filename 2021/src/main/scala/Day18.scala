package aoc.twentyone

import scala.annotation.tailrec
import scala.util.parsing.combinator.*

enum SnailfishNumber:
  case Regular(value: Int)
  case Pair(Left: SnailfishNumber, Right: SnailfishNumber)

  override def toString: String =
    this match
      case Regular(value)    => value.toString
      case Pair(left, right) => s"[$left, $right]"

  def +(other: SnailfishNumber): SnailfishNumber =
    Pair(this, other).reduce

  def magnitude: Long =
    this match
      case Regular(value)    => value
      case Pair(left, right) => 3 * left.magnitude + 2 * right.magnitude

  def split: SnailfishNumber =
    this match
      case Regular(value) if value >= 10 =>
        Pair(Regular(value / 2), Regular((value + 1) / 2))
      case Pair(left, right) =>
        val newLeft = left.split
        if newLeft != left then Pair(newLeft, right)
        else Pair(left, right.split)
      case x => x

  def reduce: SnailfishNumber =
    def reduceRec(
        number: SnailfishNumber,
        depth: Int = 0
    ): (SnailfishNumber, Int, Int) =
      val (newNumber, toAddLeft, toAddRight) = number match
        case Regular(value) => (Regular(value), 0, 0)
        case Pair(Regular(left), Regular(right)) if depth >= 4 =>
          (Regular(0), left, right)
        case Pair(left, right) =>
          val (newLeft, toAddLeft, toAddRight) = reduceRec(left, depth + 1)
          if newLeft != left then
            val newNumber = Pair(newLeft, addToLeft(right, toAddRight))
            (newNumber, toAddLeft, 0)
          else
            val (newRight, toAddLeft, toAddRight) = reduceRec(right, depth + 1)
            if newRight != right then
              val newNumber = Pair(addToRight(left, toAddLeft), newRight)
              (newNumber, 0, toAddRight)
            else (Pair(left, right), 0, 0)

      if newNumber == number then
        if depth > 0 then (number, 0, 0)
        else
          var newNumber = number.split
          if newNumber == number then (newNumber, 0, 0)
          else reduceRec(newNumber)
      else if depth > 0 then (newNumber, toAddLeft, toAddRight)
      else reduceRec(newNumber)

    reduceRec(this)(0)

  private def addToLeft(number: SnailfishNumber, toAdd: Int): SnailfishNumber =
    if toAdd == 0 then number
    else
      number match
        case Regular(value)    => Regular(value + toAdd)
        case Pair(left, right) => Pair(addToLeft(left, toAdd), right)

  private def addToRight(number: SnailfishNumber, toAdd: Int): SnailfishNumber =
    if toAdd == 0 then number
    else
      number match
        case Regular(value)    => Regular(value + toAdd)
        case Pair(left, right) => Pair(left, addToRight(right, toAdd))

object SnailfishNumber:
  def apply(str: String): SnailfishNumber =
    SnailfishNumberParser(str)

  object SnailfishNumberParser extends RegexParsers:
    def regular = """\d+""".r ^^ { t => Regular(t.toInt) }
    def pair = "[" ~ number ~ "," ~ number ~ "]" ^^ {
      case _ ~ left ~ _ ~ right ~ _ => Pair(left, right)
    }
    def number: Parser[SnailfishNumber] = regular | pair

    def apply(str: String): SnailfishNumber =
      parse(number, str) match
        case Success(result, _) => result
        case _                  => ???

class Day18 extends SimplePuzzle[List[SnailfishNumber], Long]:
  override protected def input(raw: Iterator[String]): List[SnailfishNumber] =
    raw.map(SnailfishNumber(_)).toList

  override protected def part1(input: List[SnailfishNumber]): Option[Long] =
    val result = input.reduceLeft(_ + _)
    Some(result.magnitude)

  override protected def part2(input: List[SnailfishNumber]): Option[Long] =
    Some(
      (
        for
          x <- input
          y <- input
        yield (x + y).magnitude
      ).max
    )
