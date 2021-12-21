package aoc.twentyone

import scala.sys.process.processInternal

enum Pixel:
  case Light, Dark

  override def toString: String =
    this match
      case Light => "#"
      case Dark  => "."

object Pixel:
  def apply(pixel: Char) =
    pixel match
      case '#' => Pixel.Light
      case '.' => Pixel.Dark
      case _   => ???

case class ScannersResult(
    val algorithm: Vector[Pixel],
    val image: Map[(Int, Int), Pixel]
)

class Day20 extends SimplePuzzle[ScannersResult, Int]:
  override protected def input(raw: Iterator[String]): ScannersResult =
    val algorithm = raw.next.map(Pixel(_)).toVector

    raw.next
    val image = raw.zipWithIndex
      .flatMap((line, y) =>
        line.zipWithIndex.map((pixel, x) => ((x, y), Pixel(pixel)))
      )
      .toMap

    ScannersResult(algorithm, image)

  override protected def part1(input: ScannersResult): Option[Int] =
    Some(
      enhance(
        enhance(input.image, input.algorithm, Pixel.Dark),
        input.algorithm,
        input.algorithm(0)
      ).values
        .count(_ == Pixel.Light)
    )

  override protected def part2(input: ScannersResult): Option[Int] =
    def rec(
        image: Map[(Int, Int), Pixel],
        algorithm: Vector[Pixel],
        count: Int,
        default: Pixel
    ): Map[(Int, Int), Pixel] =
      if count == 0 then image
      else
        rec(
          enhance(image, algorithm, default),
          algorithm,
          count - 1,
          default match
            case Pixel.Dark  => algorithm(0)
            case Pixel.Light => algorithm.last
        )

    Some(
      rec(input.image, input.algorithm, 50, Pixel.Dark).values
        .count(_ == Pixel.Light)
    )

  private def enhance(
      image: Map[(Int, Int), Pixel],
      algorithm: Vector[Pixel],
      default: Pixel
  ): Map[(Int, Int), Pixel] =
    // printImage(image)
    val (minX, minY) = image.keys.min
    val (maxX, maxY) = image.keys.max
    (
      for
        x <- (minX - 1) to (maxX + 1)
        y <- (minY - 1) to (maxY + 1)
      yield
        val binary = (
          for
            yy <- (y - 1) to (y + 1)
            xx <- (x - 1) to (x + 1)
          yield image.getOrElse((xx, yy), default) match
            case Pixel.Dark  => '0'
            case Pixel.Light => '1'
        ).mkString
        val idx = Integer.parseInt(binary, 2)
        ((x, y), algorithm(idx))
    ).toMap

  private def printImage(image: Map[(Int, Int), Pixel]): Unit =
    val (minX, minY) = image.keys.min
    val (maxX, maxY) = image.keys.max
    for y <- minY to maxY do
      for x <- minX to maxX do print(image((x, y)))
      println
