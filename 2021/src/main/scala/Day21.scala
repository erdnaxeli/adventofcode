package aoc.twentyone

import scala.annotation.tailrec

enum Player:
  case P1, P2

case class DiracDiceState(
    dice: Int,
    p1: Int,
    p2: Int,
    p1score: Int = 0,
    p2score: Int = 0,
    toPlay: Player = Player.P1,
    runs: Int = 0
):
  def nextState: DiracDiceState =
    val d1 = runDice(dice)
    val d2 = runDice(d1)
    val d3 = runDice(d2)

    toPlay match
      case Player.P1 =>
        val np1 = incrPlayer(p1, d1 + d2 + d3)
        DiracDiceState(
          d3,
          np1,
          p2,
          p1score + np1,
          p2score,
          Player.P2,
          runs + 3
        )
      case Player.P2 =>
        val np2 = incrPlayer(p2, d1 + d2 + d3)
        DiracDiceState(
          d3,
          p1,
          np2,
          p1score,
          p2score + np2,
          Player.P1,
          runs + 3
        )

  def isWin: Boolean = p1score >= 1000 || p2score >= 1000

  private def runDice(dice: Int) =
    val n = dice + 1
    if n > 100 then 1 else n

  @tailrec
  private def incrPlayer(player: Int, incr: Int): Int =
    if incr == 0 then player
    else
      val n = player + 1
      if n > 10 then incrPlayer(1, incr - 1)
      else incrPlayer(n, incr - 1)

class Day21 extends SimplePuzzle[DiracDiceState, Long]:
  override protected def input(raw: Iterator[String]): DiracDiceState =
    def pos(line: String): Int =
      val regex = """Player \d starting position: (\d+)""".r
      line match
        case regex(position) => position.toInt
        case _               => ???

    DiracDiceState(0, pos(raw.next), pos(raw.next))

  override protected def part1(input: DiracDiceState): Option[Long] =
    @tailrec
    def rec(state: DiracDiceState): DiracDiceState =
      if state.isWin then state
      else rec(state.nextState)

    val finalState = rec(input)
    if finalState.p1 > finalState.p2 then
      Some(finalState.p2score.toLong * finalState.runs.toLong)
    else Some(finalState.p1score.toLong * finalState.runs.toLong)

  override protected def part2(input: DiracDiceState): Option[Long] =
    val moves = (1 to 10).map(x => x -> getMovesFrom(x)).toMap
    val (p1, p2) = (input.p1, input.p2)
    val (p1Win, p2Win) = getWinsCount(input.p1, input.p2, moves)
    println((p1Win, p2Win))
    if p1Win > p2Win then Some(p1Win)
    else Some(p2Win)

  private def getMovesFrom(position: Int): Map[Int, Int] =
    def rec(position: Int, incr: Int): Int =
      if incr == 0 then position
      else
        val n = position + 1
        if n > 10 then rec(1, incr - 1) else rec(n, incr - 1)
    (
      for
        d1 <- 1 to 3
        d2 <- 1 to 3
        d3 <- 1 to 3
      yield rec(position, d1 + d2 + d3)
    ).groupBy(identity).map((k, v) => k -> v.size)

  private def getWinsCount(
      p1: Int,
      p2: Int,
      moves: Map[Int, Map[Int, Int]],
      p1Score: Int = 0,
      p2Score: Int = 0,
      universesCount: Long = 1L
  ): (Long, Long) =
    if p2Score >= 21 then (0L, universesCount)
    else
      moves(p1).toList
        .map((newPosition, count) =>
          getWinsCount(
            p2,
            newPosition,
            moves,
            p2Score,
            p1Score + newPosition,
            count * universesCount
          )
        )
        .reduce((a, b) => (a(0) + b(0), a(1) + b(1)))
        .swap
