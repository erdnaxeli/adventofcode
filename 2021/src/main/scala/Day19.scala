package aoc.twentyone

import scala.annotation.tailrec

case class Point(val x: Int, val y: Int, val z: Int)

case class Cube(name: String = "", points: List[Point] = List.empty[Point]):
  def +(other: Cube): Cube =
    Cube(name, (points.toSet ++ other.points.toSet).toList)

  def +:(point: Point): Cube =
    Cube(name, point :: points)

  def size: Int = points.size

  def positions: List[Cube] =
    val xp1 = this
    val xp2 = xp1.rotateX
    val xp3 = xp2.rotateX
    val xp4 = xp3.rotateX
    val ym1 = xp1.rotateZ
    val ym2 = ym1.rotateY
    val ym3 = ym2.rotateY
    val ym4 = ym3.rotateY
    val xm1 = ym1.rotateZ
    val xm2 = xm1.rotateX
    val xm3 = xm2.rotateX
    val xm4 = xm3.rotateX
    val yp1 = xm1.rotateZ
    val yp2 = yp1.rotateY
    val yp3 = yp2.rotateY
    val yp4 = yp3.rotateY
    val zp1 = xp1.rotateY
    val zp2 = zp1.rotateZ
    val zp3 = zp2.rotateZ
    val zp4 = zp3.rotateZ
    val zm1 = xm1.rotateY
    val zm2 = zm1.rotateZ
    val zm3 = zm2.rotateZ
    val zm4 = zm3.rotateZ
    List(
      xp1,
      xp2,
      xp3,
      xp4,
      ym1,
      ym2,
      ym3,
      ym4,
      xm1,
      xm2,
      xm3,
      xm4,
      yp1,
      yp2,
      yp3,
      yp4,
      zp1,
      zp2,
      zp3,
      zp4,
      zm1,
      zm2,
      zm3,
      zm4
    )

  def rotateX: Cube = Cube(name, points.map(p => Point(p.x, p.z, -p.y)))
  def rotateY: Cube = Cube(name, points.map(p => Point(-p.z, p.y, p.x)))
  def rotateZ: Cube = Cube(name, points.map(p => Point(p.y, -p.x, p.z)))

  def translate(vector: Point): Cube =
    Cube(
      name,
      points.map(p => Point(p.x + vector.x, p.y + vector.y, p.z + vector.z))
    )

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
          case name => rec(input, acc, Cube(name))

    rec(raw)

  override protected def part1(input: List[Cube]): Option[Int] =
    Some(combine(input).size)

  override protected def part2(input: List[Cube]): Option[Int] =
    val scanners = findAssociations(input).map(_._2)

    Some(
      (
        for
          s1 <- scanners
          s2 <- scanners
        yield (s1.x - s2.x).abs + (s1.y - s2.y).abs + (s1.z - s2.z).abs
      ).max
    )

  private def combine(cubes: List[Cube]): Cube =
    // findAssociations(cubes).tapEach((c, v) => println(s"${c.name} $v"))
    findAssociations(cubes).map(_._1).reduce(_ + _)

  /* Given a list of cubes, return a list of tuples with the cube rotated if
   * needed and a translation vector.
   */
  def findAssociations(cubes: List[Cube]): List[(Cube, Point)] =
    @tailrec
    def rec(
        cubes: List[Cube],
        result: List[(Cube, Point)]
    ): List[(Cube, Point)] =
      println(s"Test ${cubes.size} ${result.size}")
      cubes match
        case Nil => result
        case c2 :: tail =>
          getFirstResult(result)((c1, _) =>
            // println(s"${c1.name}, ${c2.name}")
            findMatching(c1, c2)
          ) match
            case None       => rec(tail :+ c2, result)
            case Some(c, v) => rec(tail, (c, v) :: result)

    cubes match
      case Nil => Nil
      case c :: tail =>
        rec(tail, List((c, Point(0, 0, 0))))

  /* For a given cube c1, return c2 with the correct rotation and translation
   * and the center position.
   */
  private def findMatching(c1: Cube, c2: Cube): Option[(Cube, Point)] =
    getFirstResult(c2.positions)(cp2 =>
      findTranslation(c1, cp2).map(vector => (cp2.translate(vector), vector))
    )

  private def findTranslation(c1: Cube, c2: Cube): Option[Point] =
    val translations = for
      p1 <- c1.points
      p2 <- c2.points
    yield Point(p1.x - p2.x, p1.y - p2.y, p1.z - p2.z)
    // println(translations.toSet.size)

    getFirstResult(translations.toSet)(vector =>
      if (c1.points.toSet & c2.translate(vector).points.toSet).size >= 12 then
        Some(vector)
      else None
    )

  private def getFirstResult[A, B](input: IterableOnce[A])(
      f: A => Option[B]
  ): Option[B] =
    input.iterator.map(f).find(_.isDefined).flatten
