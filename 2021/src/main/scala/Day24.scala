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

class ALU(
    val instructions: List[Instruction],
    val input: List[Int],
    val memory: Map[Variable, Int] =
      Map.empty[Variable, Int].withDefaultValue(0)
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
            ALU(
              tail,
              t,
              memory + (a -> h)
            )
      case Add(a, b) =>
        ALU(tail, input, memory + (a -> (memory(a) + getVar(b))))
      case Mul(a, b) =>
        ALU(tail, input, memory + (a -> (memory(a) * getVar(b))))
      case Div(a, b) =>
        ALU(tail, input, memory + (a -> (memory(a) / getVar(b))))
      case Mod(a, b) =>
        ALU(tail, input, memory + (a -> (memory(a) % getVar(b))))
      case Eql(a, b) =>
        ALU(
          tail,
          input,
          memory + (a -> (if getVar(a) == getVar(b) then 1 else 0))
        )

  def getVar(x: Variable | Int): Int =
    x match
      case v: Variable => memory(v)
      case i: Int      => i

class Day24 extends SimplePuzzle[List[List[Instruction]], Long]:
  override protected def input(raw: Iterator[String]): List[List[Instruction]] =
    def rec(
        result: List[List[Instruction]],
        current: List[Instruction],
        instructions: List[Instruction]
    ): List[List[Instruction]] =
      instructions match
        case Nil => ((current :: result).reverse).map(_.reverse)
        case (x @ Instruction.Inp(_)) :: tail =>
          rec(current :: result, List(x), tail)
        case x :: tail => rec(result, x :: current, tail)

    raw.map(Instruction(_)).toList match
      case h :: t => rec(List.empty[List[Instruction]], List(h), t)
      case _      => ???

  override protected def part1(input: List[List[Instruction]]): Option[Long] =
    // def rec(
    //     result: List[Int],
    //     z: Int,
    //     input: List[List[Instruction]],
    //     finalZ: Int
    // ): Option[Long] =
    //   // println(result.reverse.mkString)
    //   input match
    //     case Nil =>
    //       if z == finalZ then
    //         // println(s"ouh yeah ${result.reverse.mkString}")
    //         Some(result.reverse.mkString.toInt)
    //       else None
    //     case toTest :: tail =>
    //       val inputOutput = testPart(toTest, z)
    //       // println(inputOutput)
    //       inputOutput
    //         .flatMap((w, z) => rec(w :: result, z, tail, finalZ))
    //         .maxOption

    // println(testPart(input(13), 20).filter(_._2 == 0))
    // println(testPart(input(12), 0).filter(_._2 == 20))
    // println(testPart(input(11), 20).filter(_._2 == 0))
    // println(testPart(input(10), 537).filter(_._2 == 20))
    // println(testPart(input(9), 13981).filter(_._2 == 537))
    // println(testPart(input(8), 537).filter(_._2 == 13981))
    // println(testPart(input(7), 20).filter(_._2 == 537))
    // println(testPart(input(6), 535).filter(_._2 == 20))

    // println(testPart(input(5), 13930).filter(_._2 == 535))
    // println(testPart(input(4), 362205).filter(_._2 == 13930))
    // println(testPart(input(3), 13930).filter(_._2 == 362205))
    // println(testPart(input(2), 535).filter(_._2 == 13930))
    // println(testPart(input(1), 20).filter(_._2 == 535))

    // println(testPart(input(5), 13929).filter(_._2 == 535))
    // println(testPart(input(4), 362179).filter(_._2 == 13929))
    // println(testPart(input(3), 13929).filter(_._2 == 362179))

    // println(testPart(input(13), 24).filter(_._2 == 0))
    // println(testPart(input(12), 0).filter(_._2 == 20))
    // println(testPart(input(11), 20).filter(_._2 == 0))
    // println(testPart(input(10), 537).filter(_._2 == 20))
    // println(testPart(input(9), 13981).filter(_._2 == 537))
    // println(testPart(input(8), 537).filter(_._2 == 13981))
    // println(testPart(input(7), 20).filter(_._2 == 537))
    // println(testPart(input(6), 535).filter(_._2 == 20))
    // val result = (0 to 100000)
    //   .map(z => z -> testPart(input(12), z).filter(_._2 == 24))
    //   .filterNot((k, v) => v.isEmpty)
    // println(result)
    // println(testPart(input(0), 0))
    // println(testPart(input(1), 20))
    // println(rec(List.empty[Int], 0, input.take(6), 535))

    def rec(
        input: List[List[Instruction]],
        targetZ: Int,
        ws: List[Int],
        cache: Map[List[Instruction], Map[Int, Map[Int, Int]]] =
          Map.empty[List[Instruction], Map[Int, Map[Int, Int]]],
        n: Int = 0
    ): (Option[List[Int]], Map[List[Instruction], Map[Int, Map[Int, Int]]]) =
      input match
        case Nil => (Some(ws), cache)
        case head :: tail =>
          val (candidates, newCache) = findPotentialZ(head, targetZ, cache)

          if candidates.isEmpty then (None, newCache)
          else
            candidates.toList.sorted.reverse
              .foldLeft(
                (None: Option[List[Int]], newCache)
              )((acc, elt) =>
                val (result, cache) = acc
                val (w, z) = elt

                result match
                  case None =>
                    print(" " * n)
                    println(
                      s"Target Z : $targetZ ; Candidate W : $w ; Candidate z(n-1) : $z"
                    )
                    rec(tail, z, w :: ws, cache, n + 1)
                  case _ => acc
              )
    // .to(LazyList)
    // .map((w, z) =>
    //   print(" " * n)
    //   println(
    //     s"Target Z : $targetZ ; Candidate W : $w ; Candidate z(n-1) : $z"
    //   )
    //   rec(tail, z, w :: ws, newCache, n + 1)
    // )
    // .collectFirst {
    //   case (Some(result), cache) => (result, cache)
    //   case
    // }

    val result = rec(input.reverse, 0, List.empty[Int])(0)
    println(result)
    None

  private def findPotentialZ(
      input: List[Instruction],
      targetZ: Int = 0,
      cache: Map[List[Instruction], Map[Int, Map[Int, Int]]]
  ): (Map[Int, Int], Map[List[Instruction], Map[Int, Map[Int, Int]]]) =
    val forAllZ = cache.get(input) match
      case None =>
        println("No cache yet")
        (0 to 1000000)
          .map(z =>
            val ws = testPart(input, z)
              .groupBy(_._2)
              .map((zFinal, ws) => (zFinal -> ws.keys.max))
            (z -> ws)
          )
          .toMap
      case Some(cachedResult) =>
        println("Using cache")
        cachedResult
    val newCache = cache + (input -> forAllZ)
    val forTargetZ = forAllZ
      .map((z, result) =>
        result.get(targetZ) match
          case None    => None
          case Some(w) => Some(w -> z)
      )
      .flatten
      .toMap

    println(s"For z=$targetZ I found $forTargetZ")
    (forTargetZ, newCache)

  // def rec(
  //     z: Int,
  //     result: Map[Int, Int],
  //     input: List[Instruction]
  // ): Map[Int, Int] =
  //   if result.size == 9 || z >= 1000000 then result
  //   else
  //     val wCandidate = testPart(input, z)
  //       .filter(_._2 == targetZ)
  //       .keys
  //       .maxOption
  //     wCandidate match
  //       case None    => rec(z + 1, result, input)
  //       case Some(w) => rec(z + 1, result + (w -> z), input)

  // rec(0, Map.empty[Int, Int], input)

  private def testPart(
      instructions: List[Instruction],
      z: Int
  ): Map[Int, Int] =
    (1 to 9)
      .map(w =>
        val alu =
          ALU(instructions, List(w), Map(Variable.Z -> z).withDefaultValue(0))
        w -> exec(alu).memory(Variable.Z)
      )
      .toMap

  private def exec(alu: ALU): ALU =
    @tailrec
    def rec(alu: ALU): ALU =
      alu.next match
        case None          => alu
        case Some(nextAlu) => rec(nextAlu)

    rec(alu)
