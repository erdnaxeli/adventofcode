package aoc.twentyone

import scala.concurrent.duration.fromNow

enum Amphipod(cost: Int):
  case A extends Amphipod(1)
  case B extends Amphipod(10)
  case C extends Amphipod(100)
  case D extends Amphipod(1000)

  def getEnergy(steps: Int): Int = cost * steps

case class Burrow(
    rooms: Map[Amphipod, List[(Int, Int)]],
    hallway: List[(Int, Int)],
    moves: Map[(Int, Int), List[(Int, Int, Int)]],
    positions: Map[(Int, Int), Amphipod] = Map.empty[(Int, Int), Amphipod]
):
  def printMap(): Unit =
    println("#############")
    print("#")
    for x <- 1 to 11 do
      print(
        positions.get((x, 1)) match
          case None           => "."
          case Some(amphipod) => amphipod
      )
    println("#")
    print("###")
    print(
      positions.get((3, 2)) match
        case None    => "."
        case Some(a) => a
    )
    print("#")
    print(
      positions.get((5, 2)) match
        case None    => "."
        case Some(a) => a
    )
    print("#")
    print(
      positions.get((7, 2)) match
        case None    => "."
        case Some(a) => a
    )
    print("#")
    print(
      positions.get((9, 2)) match
        case None    => "."
        case Some(a) => a
    )
    println("###")
    print("  #")
    print(
      positions.get((3, 3)) match
        case None    => "."
        case Some(a) => a
    )
    print("#")
    print(
      positions.get((5, 3)) match
        case None    => "."
        case Some(a) => a
    )
    print("#")
    print(
      positions.get((7, 3)) match
        case None    => "."
        case Some(a) => a
    )
    print("#")
    print(
      positions.get((9, 3)) match
        case None    => "."
        case Some(a) => a
    )
    println("#")
    println("  #########")

  def +:(position: Map[(Int, Int), Amphipod]): Burrow =
    copy(positions = position ++ positions)

  def isFinalState: Boolean =
    positions.forall((k, amphipod) => rooms(amphipod).contains(k))

  def nextStates: List[(Burrow, Int)] =
    // For all amphipod current position, get the next positions availble.
    val newPositions = positions.toList.flatMap((k, amphipod) =>
      val (x, y) = k
      // If the amphipod is in the hallway it can only move to its final room
      // if there is a space and it does not contains a amphipod who does not
      // belong to this room.
      if hallway.contains((x, y)) then
        if roomIsFreeOrContains(amphipod) then
          getRoomPlaceFor(amphipod) match
            case None => None
            case Some((xx, yy)) =>
              if isHallwayReachable((xx, yy), (x, y)) then
                List(
                  (amphipod, x, y, xx, yy, Burrow.getSteps((x, y), (xx, yy)))
                )
              else None
        else None
      // Else the amphipod must be in a room and can move to any hallway place
      // reachable.
      else if shouldStayInRoom(amphipod, x, y) then None
      else
        findHallwayReachablePositionsFrom(x, y).map((xx, yy, steps) =>
          (amphipod, x, y, xx, yy, steps)
        )
    )

    // Return new states
    newPositions.map((amphipod, x, y, xx, yy, steps) =>
      (
        copy(positions = (positions - ((x, y))) ++ Map((xx, yy) -> amphipod)),
        amphipod.getEnergy(steps)
      )
    )

  private def roomIsFreeOrContains(amphipod: Amphipod): Boolean =
    rooms(amphipod).forall((x, y) =>
      positions.get((x, y)) match
        case None                             => true
        case Some(other) if other == amphipod => true
        case _                                => false
    )

  private def getRoomPlaceFor(amphipod: Amphipod): Option[(Int, Int)] =
    rooms(amphipod).reverse.collectFirst {
      case (x, y) if !positions.contains((x, y)) => (x, y)
    }

  private def findHallwayReachablePositionsFrom(
      x: Int,
      y: Int
  ): List[(Int, Int, Int)] =
    moves((x, y))
      .map((xx, yy, e) =>
        if !positions.contains((xx, yy)) &&
          isHallwayReachable((x, y), (xx, yy))
        then Some(xx, yy, e)
        else None
      )
      .flatten

  private def isHallwayReachable(
      room: (Int, Int),
      hallway: (Int, Int)
  ): Boolean =
    // Check the way from the room point to reach the hallway. We skip the point
    // where we are.
    lazy val roomEntranceEmpty = (hallway(1) to (room(1) - 1)).forall(y =>
      !positions.contains((room(0), y))
    )
    // Check the way from the hallway in front of our room to the hallawy point.
    // We skip the point where we want to go.
    lazy val hallwayEmpty =
      if hallway(0) > room(0) then
        (room(0) to (hallway(0) - 1)).forall(x =>
          !positions.contains((x, hallway(1)))
        )
      else
        ((hallway(0) + 1) to room(0)).forall(x =>
          !positions.contains((x, hallway(1)))
        )

    roomEntranceEmpty && hallwayEmpty

  private def shouldStayInRoom(amphipod: Amphipod, x: Int, y: Int): Boolean =
    // If the room contains only one correct amphipod, it must be myself.
    rooms(amphipod).contains((x, y)) && roomIsFreeOrContains(amphipod)

object Burrow:
  def apply(
      rooms: Map[Amphipod, List[(Int, Int)]],
      hallway: List[(Int, Int)]
  ): Burrow =
    Burrow(rooms, hallway, getMoves(hallway, rooms))

  /*
   * Returns a map of all positions reachable from a given point, with the steps
   * needed to.
   */
  def getMoves(
      hallway: List[(Int, Int)],
      rooms: Map[Amphipod, List[(Int, Int)]]
  ): Map[(Int, Int), List[(Int, Int, Int)]] =
    val fromHallway = hallway
      .map((x, y) =>
        val points = rooms.values
          .flatMap(roomPoints =>
            roomPoints.map((xx, yy) => (xx, yy, getSteps((x, y), (xx, yy))))
          )
          .toList
        (x, y) -> points
      )
      .toMap
    // Reverse fromHallway
    val toHallway = fromHallway
      .foldLeft(Map.empty[(Int, Int), List[(Int, Int, Int)]])((acc, e) =>
        val (k, points) = e
        val (x, y) = k
        acc ++ points
          .map((xx, yy, steps) =>
            val oldPoints = acc.getOrElse((xx, yy), List.empty[(Int, Int, Int)])
            val newPoints = (x, y, steps) :: oldPoints
            (xx, yy) -> newPoints
          )
          .toMap
      )

    fromHallway ++ toHallway

  private def getSteps(from: (Int, Int), to: (Int, Int)): Int =
    (from(0) - to(0)).abs + (from(1) - to(1)).abs

class Day23 extends Puzzle[Burrow, Burrow, Int, Int]:
  override protected def input1(raw: Iterator[String]): Burrow =
    val burrow = Burrow(
      // rooms must be ordered from the entrance to the back
      rooms = Map(
        Amphipod.A -> List((3, 2), (3, 3)),
        Amphipod.B -> List((5, 2), (5, 3)),
        Amphipod.C -> List((7, 2), (7, 3)),
        Amphipod.D -> List((9, 2), (9, 3))
      ),
      hallway = List(
        (1, 1),
        (2, 1),
        (4, 1),
        (6, 1),
        (8, 1),
        (10, 1),
        (11, 1)
      )
    )

    raw.zipWithIndex
      .flatMap((line, y) =>
        line.zipWithIndex.collect {
          case ('A', x) => Map((x, y) -> Amphipod.A)
          case ('B', x) => Map((x, y) -> Amphipod.B)
          case ('C', x) => Map((x, y) -> Amphipod.C)
          case ('D', x) => Map((x, y) -> Amphipod.D)
        }
      )
      .foldLeft(burrow)((burrow, position) => position +: burrow)

  override protected def input2(raw: Iterator[String]): Burrow =
    val burrow = Burrow(
      // rooms must be ordered from the entrance to the back
      rooms = Map(
        Amphipod.A -> List((3, 2), (3, 3), (3, 4), (3, 5)),
        Amphipod.B -> List((5, 2), (5, 3), (5, 4), (5, 5)),
        Amphipod.C -> List((7, 2), (7, 3), (7, 4), (7, 5)),
        Amphipod.D -> List((9, 2), (9, 3), (9, 4), (9, 5))
      ),
      hallway = List(
        (1, 1),
        (2, 1),
        (4, 1),
        (6, 1),
        (8, 1),
        (10, 1),
        (11, 1)
      )
    )

    raw.zipWithIndex
      .flatMap((line, y) =>
        line.zipWithIndex.collect {
          case ('A', x) => Map((x, y) -> Amphipod.A)
          case ('B', x) => Map((x, y) -> Amphipod.B)
          case ('C', x) => Map((x, y) -> Amphipod.C)
          case ('D', x) => Map((x, y) -> Amphipod.D)
        }
      )
      .foldLeft(burrow)((burrow, position) => position +: burrow)

  override protected def part1(input: Burrow): Option[Int] =
    val energy = getFinalState(input)
    Some(energy)

  override protected def part2(input: Burrow): Option[Int] =
    val energy = getFinalState(input)
    Some(energy)

  private def getFinalState(burrow: Burrow): Int =
    def rec(
        best: Option[(Burrow, Int)],
        nextSolutions: List[(Burrow, Int)],
        cache: Map[Burrow, Int]
    ): Option[(Burrow, Int)] =
      nextSolutions match
        case Nil => best
        case (burrow, energy) :: tail =>
          if cache.contains(burrow) && cache(burrow) <= energy then
            rec(best, tail, cache)
          else if burrow.isFinalState then
            best match
              case None =>
                println(s"Find best $energy, still ${nextSolutions.size}")
                rec(Some(burrow, energy), tail, cache + (burrow -> energy))
              case Some(bestBurrow, bestEnergy) =>
                if energy < bestEnergy then
                  println(s"Find best $energy, still ${nextSolutions.size}")
                  rec(Some(burrow, energy), tail, cache)
                else rec(best, tail, cache + (burrow -> energy))
          else
            best match
              case None =>
                rec(
                  best,
                  tail ++ burrow.nextStates.map((b, e) => (b, e + energy)),
                  cache + (burrow -> energy)
                )
              case Some(bestBurrow, bestEnergy) =>
                rec(
                  best,
                  tail ++ burrow.nextStates
                    .map((b, e) => (b, e + energy))
                    .filter((_, e) => e < bestEnergy),
                  cache + (burrow -> energy)
                )

    rec(None, List((burrow, 0)), Map.empty[Burrow, Int]) match
      case None            => ???
      case Some(_, energy) => energy
