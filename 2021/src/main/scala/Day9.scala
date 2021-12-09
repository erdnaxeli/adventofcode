package aoc.twentyone

class Day9 extends SimplePuzzle[Array[Array[Int]], Int]:
  override protected def input(raw: Iterator[String]): Array[Array[Int]] =
    raw.map(_.split("").map(_.toInt)).toArray

  override protected def part1(input: Array[Array[Int]]): Option[Int] =
    val comparisons = for
      i <- 0 until input.size
      j <- 0 until input(0).size
    yield
      val point = input(i)(j)
      (
        point,
        for
          ii <- (i - 1) to (i + 1)
          jj <- (j - 1) to (j + 1)
          if (ii == i) != (jj == j)
        yield input.lift(ii) match
          case None => true
          case Some(v) =>
            v.lift(jj) match
              case None                         => true
              case Some(other) if other > point => true
              case _                            => false
      )

    Some(
      comparisons
        .filter((_, v) => v.forall(identity))
        .map((point, _) => point + 1)
        .sum
    )

  override protected def part2(input: Array[Array[Int]]): Option[Int] =
    // Get a Map((x,y) => height)
    val floor = input.zipWithIndex
      .flatMap((line, x) =>
        line.zipWithIndex.map((height, y) => ((x, y), height))
      )
      .toMap

    // Given a floor, a map of point -> basin's min point, and a point (x, y),
    // return None if the point is not part of a basin, or Option((a, b)) if
    // the point is part of the basin of min (a, b).
    def getBasin(
        floor: Map[(Int, Int), Int],
        basins: Map[(Int, Int), (Int, Int)],
        x: Int,
        y: Int
    ): Option[(Int, Int)] =
      // A height of 9 is never part of a basin.
      if floor((x, y)) == 9 then None
      else
        val around = for
          i <- (x - 1) to (x + 1)
          j <- (y - 1) to (y + 1)
          if (i == x) != (j == y)
        yield floor.get((i, j)) match
          case None                          => Some((x, y))
          case Some(v) if v >= floor((x, y)) => Some((x, y))
          case Some(v) =>
            basins.get((i, j)) match
              case None         => Some((i, j))
              case Some((a, b)) => Some((a, b))

        val isSelfBasin = around.forall(_ match
          case Some((a, b)) if (a, b) == (x, y) => true
          case _                                => false
        )
        if isSelfBasin then Some((x, y))
        else
          around
            .collect { case Some((a, b)) if (a, b) != (x, y) => (a, b) }
            .toSet
            .toList match
            case v :: Nil => Some(v)
            case _        => None

    def getBasins(
        floor: Map[(Int, Int), Int],
        basins: Map[(Int, Int), (Int, Int)] = Map.empty[(Int, Int), (Int, Int)]
    ): Map[(Int, Int), (Int, Int)] =
      val newBasins =
        floor.foldLeft(basins)((basins, v) =>
          val ((x, y), height) = v
          getBasin(floor, basins, x, y) match
            case None         => basins
            case Some((a, b)) => basins.updated((x, y), (a, b))
        )
      if newBasins == basins then basins
      else getBasins(floor, newBasins)

    Some(
      getBasins(floor)
        .groupBy((k, v) => v)
        .map((k, v) => v.size)
        .toList
        .sorted
        .reverse
        .take(3)
        .product
    )
