package aoc.twentyone

import scala.compiletime.ops.int

enum CubeState:
  case On, Off

object CubeState:
  def apply(state: String) =
    state match
      case "on"  => CubeState.On
      case "off" => CubeState.Off
      case _     => ???

type Cuboid = (Range, Range, Range)

object Cuboid:
  def substract(c1: Cuboid, c2: Cuboid): List[Cuboid] =
    intersect(c1, c2) match
      case None => List(c1)
      case Some(intersection) =>
        val x1 = c1(0).start to (intersection(0).start - 1)
        val x2 = (intersection(0).end + 1) to c1(0).end
        val y1 = c1(1).start to (intersection(1).start - 1)
        val y2 = (intersection(1).end + 1) to c1(1).end
        val z1 = c1(2).start to (intersection(2).start - 1)
        val z2 = (intersection(2).end + 1) to c1(2).end

        List(
          (
            (x1.start to x2.end),
            (y1.start to y2.end),
            z1
          ),
          (
            x1,
            (y1.start to y2.end),
            intersection(2)
          ),
          (
            intersection(0),
            y1,
            intersection(2)
          ),
          (
            intersection(0),
            y2,
            intersection(2)
          ),
          (
            x2,
            (y1.start to y2.end),
            intersection(2)
          ),
          (
            (x1.start to x2.end),
            (y1.start to y2.end),
            z2
          )
        ).filter((a, b, c) => a.size > 0 && b.size > 0 && c.size > 0)

  def size(c: Cuboid): Long =
    c(0).size.toLong * c(1).size.toLong * c(2).size.toLong

  private def overlap(c1: Cuboid, c2: Cuboid): Boolean =
    overlapAxe(c1(0), c2(0)) &&
      overlapAxe(c1(1), c2(1)) &&
      overlapAxe(c1(2), c2(2))

  private def overlapAxe(c1: Range, c2: Range): Boolean =
    (c1.start <= c2.end && c2.end <= c1.end) || (c2.start <= c1.end && c1.end <= c2.end)

  private def intersect(c1: Cuboid, c2: Cuboid): Option[Cuboid] =
    if overlap(c1, c2) then
      Some(
        c1(0).start.max(c2(0).start) to c1(0).end.min(c2(0).end),
        c1(1).start.max(c2(1).start) to c1(1).end.min(c2(1).end),
        c1(2).start.max(c2(2).start) to c1(2).end.min(c2(2).end)
      )
    else None

case class RebootStep(
    state: CubeState,
    cuboid: Cuboid
)

class Day22 extends SimplePuzzle[List[RebootStep], Long]:
  override protected def input(raw: Iterator[String]): List[RebootStep] =
    val regex =
      """(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)""".r
    raw
      .map(_ match
        case regex(state, x1, x2, y1, y2, z1, z2) =>
          RebootStep(
            CubeState(state),
            (
              x1.toInt to x2.toInt,
              y1.toInt to y2.toInt,
              z1.toInt to z2.toInt
            )
          )
        case _ => ???
      )
      .toList

  override protected def part1(input: List[RebootStep]): Option[Long] =
    Some(
      input
        .filter(step =>
          step.cuboid(0).start >= -50 &&
            step.cuboid(0).end <= 50 &&
            step.cuboid(1).start >= -50 &&
            step.cuboid(1).end <= 50 &&
            step.cuboid(2).start >= -50 &&
            step.cuboid(2).end <= 50
        )
        .foldLeft(Map.empty[(Int, Int, Int), CubeState])((acc, step) =>
          acc ++
            (
              for
                x <- step.cuboid(0)
                y <- step.cuboid(1)
                z <- step.cuboid(2)
              yield (x, y, z) -> step.state
            ).toMap
        )
        .count((_, v) => v == CubeState.On)
    )

  override protected def part2(input: List[RebootStep]): Option[Long] =
    Some(
      input
        .foldLeft(List.empty[Cuboid])((onCuboids, step) =>
          step.state match
            case CubeState.Off =>
              onCuboids.flatMap(cuboid => Cuboid.substract(cuboid, step.cuboid))
            case CubeState.On =>
              onCuboids ++ onCuboids.foldLeft(List(step.cuboid))(
                (newOnCuboids, onCuboid) =>
                  newOnCuboids.flatMap(Cuboid.substract(_, onCuboid))
              )
        )
        .map(Cuboid.size(_))
        .sum
    )
