package aoc.twentyone

import scala.annotation.tailrec

case class Point(val x: Int, val y: Int, val z: Int)

case class Cube(val points: List[Point] = List.empty[Point]):
  def +(other: Cube): Cube =
    Cube((points.toSet ++ other.points.toSet).toList)

  def +:(point: Point): Cube =
    Cube(point :: points)

  def size: Int = points.size

  def positions: List[Cube] =
    val xp1 = this
    val xp2 = xp1.rotateX
    val xp3 = xp2.rotateX
    val xp4 = xp3.rotateX
    val yp1 = xp1.rotateZ
    val yp2 = yp1.rotateX
    val yp3 = yp2.rotateX
    val yp4 = yp3.rotateX
    val xm1 = yp1.rotateZ
    val xm2 = xm1.rotateX
    val xm3 = xm2.rotateX
    val xm4 = xm3.rotateX
    val ym1 = xm1.rotateZ
    val ym2 = ym1.rotateX
    val ym3 = ym2.rotateX
    val ym4 = ym3.rotateX
    val zp1 = xm1.rotateY
    val zp2 = zp1.rotateX
    val zp3 = zp2.rotateX
    val zp4 = zp3.rotateX
    val zm1 = xp1.rotateY
    val zm2 = zm1.rotateX
    val zm3 = zm2.rotateX
    val zm4 = zm3.rotateX
    List(
      xp1,
      xp2,
      xp3,
      xp4,
      xm1,
      xm2,
      xm3,
      xm4,
      yp1,
      yp2,
      yp3,
      yp4,
      ym1,
      ym2,
      ym3,
      ym4,
      zp1,
      zp2,
      zp3,
      zp4,
      zm1,
      zm2,
      zm3,
      zm4
    )

  def rotateX: Cube = Cube(points.map(p => Point(p.x, -p.z, p.y)))
  def rotateY: Cube = Cube(points.map(p => Point(p.z, p.y, -p.x)))
  def rotateZ: Cube = Cube(points.map(p => Point(-p.y, p.x, p.z)))

  def translate(vector: Point): Cube =
    Cube(points.map(p => Point(p.x + vector.x, p.y + vector.y, p.z + vector.z)))

class Day19 extends SimplePuzzle[List[Cube], Int]:
  override protected def input(raw: Iterator[String]): List[Cube] =
    @tailrec
    def rec(
        input: Iterator[String],
        acc: List[Cube] = List.empty[Cube],
        cube: Cube = Cube()
    ): List[Cube] =
      if input.isEmpty then cube :: acc
      else
        val regex = """(-?\d+),(-?\d+),(-?\d+)""".r
        input.next match
          case "" => rec(input, cube :: acc)
          case regex(x, y, z) =>
            rec(input, acc, Point(x.toInt, y.toInt, z.toInt) +: cube)
          case _ => rec(input, acc, cube)

    rec(raw)

  override protected def part1(input: List[Cube]): Option[Int] =
    Some(combine(input).size)

  override protected def part2(input: List[Cube]): Option[Int] =
    None

  private def combine(cubes: List[Cube]): Cube =
    @tailrec
    def combineRec(cubes: List[Cube], result: Cube): Cube =
      println("test")
      cubes match
        case Nil => result
        case c :: tail =>
          combineAllPositions(c, result) match
            case None => combineRec(tail :+ c, result)
            case Some(r) =>
              println(s"found a match, to find: ${tail.size}")
              combineRec(tail, r)

    cubes match
      case Nil    => Cube()
      case h :: t => combineRec(t, h)

  private def combineAllPositions(c1: Cube, c2: Cube): Option[Cube] =
    // val combinations = c1.positions.map((_, c2))
    val combinations = for
      cp1 <- c1.positions
      cp2 <- c2.positions
    yield (cp1, cp2)

    recTest(combinations)((c1, c2) => combine(c1, c2))

  private def combine(c1: Cube, c2: Cube): Option[Cube] =
    val combinations = for
      p1 <- c1.points
      p2 <- c2.points
    yield (p1, p2)

    recTest(combinations)((p1, p2) =>
      val vector = Point(p1.x - p2.x, p1.y - p2.y, p1.z - p2.z)
      val sumCube = c1 + c2.translate(vector)
      if (c1.points.size + c2.points.size) - sumCube.points.size >= 12 then
        // if sumCube.points.size == (c1.points.size + c2.points.size) - 12 then
        Some(sumCube)
      else None
    )

  private def recTest[A, B](input: List[A])(f: A => Option[B]): Option[B] =
    input.iterator.map(f).find(_.isDefined).flatten
