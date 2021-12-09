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
