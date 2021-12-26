package aoc.twentyone

import scala.annotation.tailrec
import scala.util.Try

enum Variable:
  case W, X, Y, Z

enum Instruction:
  case Inp(a: Variable)
  case Add(a: Variable, b: Variable | Int)
  case Mul(a: Variable, b: Variable | Int)
  case Div(a: Variable, b: Variable | Int)
  case Mod(a: Variable, b: Variable | Int)
  case Eql(a: Variable, b: Variable | Int)

object Instruction:
  def apply(input: String): Instruction =
    val inp = """inp ([w-z])""".r
    val add = """add ([w-z]) ([w-z]|-?\d+)""".r
    val mul = """mul ([w-z]) ([w-z]|-?\d+)""".r
    val div = """div ([w-z]) ([w-z]|-?\d+)""".r
    val mod = """mod ([w-z]) ([w-z]|-?\d+)""".r
    val eql = """eql ([w-z]) ([w-z]|-?\d+)""".r
    input match
      case inp(a)    => Inp(Variable.valueOf(a.toUpperCase))
      case add(a, b) => Add(Variable.valueOf(a.toUpperCase), tryInt(b))
      case mul(a, b) => Mul(Variable.valueOf(a.toUpperCase), tryInt(b))
      case div(a, b) => Div(Variable.valueOf(a.toUpperCase), tryInt(b))
      case mod(a, b) => Mod(Variable.valueOf(a.toUpperCase), tryInt(b))
      case eql(a, b) => Eql(Variable.valueOf(a.toUpperCase), tryInt(b))
      case x         => throw Exception(s"Unknown instruction '$x'")

  private def tryInt(c: String): Variable | Int =
    Try(c.toInt).getOrElse(Variable.valueOf(c.toUpperCase))

case class ALU(
    instructions: List[Instruction],
    input: List[Int],
    memory: Map[Variable, Int] = Map.empty[Variable, Int].withDefaultValue(0)
):
  import Instruction.*

  def next: Option[ALU] =
    instructions match
      case Nil          => None
      case inst :: tail => Some(exec(inst, tail))

  def exec(inst: Instruction, tail: List[Instruction]): ALU =
    inst match
      case Inp(a) =>
        input match
          case Nil => throw Exception("Excepted input but none given")
          case h :: t =>
            copy(
              tail,
              t,
              memory + (a -> h)
            )
      case Add(a, b) =>
        copy(tail, input, memory + (a -> (memory(a) + getVar(b))))
      case Mul(a, b) =>
        copy(tail, input, memory + (a -> (memory(a) * getVar(b))))
      case Div(a, b) =>
        copy(tail, input, memory + (a -> (memory(a) / getVar(b))))
      case Mod(a, b) =>
        copy(tail, input, memory + (a -> (memory(a) % getVar(b))))
      case Eql(a, b) =>
        copy(
          tail,
          input,
          memory + (a -> (if getVar(a) == getVar(b) then 1 else 0))
        )

  def getVar(x: Variable | Int): Int =
    x match
      case v: Variable => memory(v)
      case i: Int      => i

class Day24 extends SimplePuzzle[List[Instruction], Long]:
  override protected def input(raw: Iterator[String]): List[Instruction] =
    raw
      .map(Instruction(_))
      .filter {
        case Instruction.Div(_, 1) => false
        case _                     => true
      }
      .toList

  override protected def part1(input: List[Instruction]): Option[Long] =
    def rec(toTest: Long = 99999999999999L): Long =
      if toTest % 1000000 == 0 then println(s"$toTest")
      val alu =
        exec(input, "%014d".format(toTest).split("").map(_.toInt).toList)
      if alu.memory(Variable.Z) == 0 then toTest
      else rec(toTest - 1)

    Some(rec())

  private def exec(instructions: List[Instruction], input: List[Int]): ALU =
    @tailrec
    def rec(alu: ALU): ALU =
      alu.next match
        case None          => alu
        case Some(nextAlu) => rec(nextAlu)

    val alu = ALU(instructions, input)
    rec(alu)
