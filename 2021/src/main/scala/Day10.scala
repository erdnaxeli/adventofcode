package aoc.twentyone

case class Chunk(
    start: Char,
    content: List[Chunk] = Nil,
    end: Option[Char] = None
):
  override def toString: String =
    s"${start}${content.map(_.toString).mkString}${end.getOrElse("")}"

class Day10 extends SimplePuzzle[List[List[Chunk]], Long]:
  override protected def input(raw: Iterator[String]): List[List[Chunk]] =
    def unpopTree(chunks: List[Chunk]): Option[Chunk] =
      chunks match
        case Nil          => None
        case chunk :: Nil => Some(chunk)
        case c1 :: c2 :: tail =>
          unpopTree(c2.copy(content = c2.content :+ c1) :: tail)

    def constructChunk(
        input: List[Char],
        currentChunk: Option[Chunk] = None,
        acc: List[Chunk] = List.empty[Chunk]
    ): List[Chunk] =
      input match
        case Nil =>
          val finalChunk = currentChunk match
            case None        => unpopTree(acc)
            case Some(chunk) => unpopTree(chunk :: acc)
          finalChunk match
            case None        => List.empty[Chunk]
            case Some(chunk) => List(chunk)
        case x :: tail if "([{<" contains x =>
          val nextChunk = Chunk(x)
          currentChunk match
            case None => constructChunk(tail, Some(nextChunk), acc)
            case Some(chunk) =>
              constructChunk(tail, Some(nextChunk), chunk :: acc)
        case x :: tail if ")]}>" contains x =>
          acc match
            case Nil =>
              currentChunk match
                case None => ???
                case Some(chunk) =>
                  chunk.copy(end = Some(x)) :: constructChunk(tail, None, acc)
            case previousChunk :: accTail =>
              currentChunk match
                case None => ???
                case Some(chunk) =>
                  constructChunk(
                    tail,
                    Some(
                      previousChunk.copy(content =
                        previousChunk.content :+ chunk.copy(end = Some(x))
                      )
                    ),
                    accTail
                  )
        case _ => ???

    raw.map(line => constructChunk(line.toList)).toList

  def foundInvalidChar(chunks: List[Chunk]): Option[Char] =
    chunks match
      case Nil => None
      case chunk :: tail =>
        foundInvalidChar(chunk.content) match
          case None =>
            (chunk.start, chunk.end) match
              case ('(', Some(c)) if c != ')' => Some(c)
              case ('[', Some(c)) if c != ']' => Some(c)
              case ('{', Some(c)) if c != '}' => Some(c)
              case ('<', Some(c)) if c != '>' => Some(c)
              case _                          => foundInvalidChar(tail)
          case Some(c) =>
            Some(c)

  def autocomplete(chunks: List[Chunk]): List[Char] =
    val startToEnd = Map(
      '(' -> ')',
      '[' -> ']',
      '{' -> '}',
      '<' -> '>'
    )
    chunks match
      case Nil => List.empty[Char]
      case chunk :: Nil if chunk.end.isEmpty =>
        val missing = startToEnd(chunk.start)
        autocomplete(chunk.content) :+ missing
      case _ :: tail => autocomplete(tail)

  override protected def part1(input: List[List[Chunk]]): Option[Long] =
    val points = Map(
      ')' -> 3,
      ']' -> 57,
      '}' -> 1197,
      '>' -> 25137
    )
    Some(
      input
        .flatMap(chunks =>
          foundInvalidChar(chunks) match
            case None    => None
            case Some(v) => Some(points(v))
        )
        .sum
        .toLong
    )

  override protected def part2(input: List[List[Chunk]]): Option[Long] =
    val points = Map(
      ')' -> 1,
      ']' -> 2,
      '}' -> 3,
      '>' -> 4
    )
    val incompleteChunks = input
      .filter(foundInvalidChar(_) match
        case None    => true
        case Some(_) => false
      )
    incompleteChunks
      .map(chunks =>
        autocomplete(chunks).foldLeft(0L)((acc, v) => acc * 5 + points(v))
      )
      .sorted
      .drop(incompleteChunks.size / 2)
      .headOption
