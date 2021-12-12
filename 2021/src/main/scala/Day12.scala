package aoc.twentyone

import scala.annotation.tailrec

class Day12 extends SimplePuzzle[Map[String, Set[String]], Int]:
  override protected def input(
      raw: Iterator[String]
  ): Map[String, Set[String]] =
    val segment = raw"([A-Za-z]+)-([A-Za-z]+)".r
    val segments = raw.collect { case segment(start, end) =>
      (start, end)
    }.toList

    @tailrec
    def buildNodes(
        segments: List[(String, String)],
        nodes: Map[String, Set[String]] = Map.empty[String, Set[String]]
    ): Map[String, Set[String]] =
      segments match
        case Nil => nodes
        case head :: tail =>
          val (start, end) = head
          buildNodes(
            tail,
            nodes
              .updatedWith(start) {
                case None            => Some(Set(end))
                case Some(neighbors) => Some(neighbors + end)
              }
              .updatedWith(end) {
                case None            => Some(Set(start))
                case Some(neighbors) => Some(neighbors + start)
              }
          )

    buildNodes(segments)

  override protected def part1(input: Map[String, Set[String]]): Option[Int] =
    def findPaths(
        map: Map[String, Set[String]],
        start: String = "start",
        end: String = "end",
        currentPath: List[String] = List.empty[String]
    ): Set[List[String]] =
      if start == end then Set(start :: currentPath)
      else if map(start).isEmpty then Set.empty[List[String]]
      else
        map(start)
          .flatMap(node =>
            if node.toLowerCase != node || !currentPath.contains(node) then
              Some(findPaths(map, node, end, start :: currentPath))
            else None
          )
          .flatten

    Some(findPaths(input).size)

  override protected def part2(input: Map[String, Set[String]]): Option[Int] =
    def findPaths(
        map: Map[String, Set[String]],
        start: String = "start",
        end: String = "end",
        currentPath: List[String] = List.empty[String]
    ): Set[List[String]] =
      currentPath match
        case Nil =>
          map(start).flatMap(n =>
            findPaths(map, start, end, n :: start :: currentPath)
          )
        case node :: _ =>
          if node == end then Set(currentPath)
          else if node == start then Set.empty[List[String]]
          else if map(node).isEmpty then Set.empty[List[String]]
          else if currentPath
              .filter(n => n.toLowerCase == n)
              .groupBy(identity)
              .find(_._2.size > 2)
              .isDefined
          then Set.empty[List[String]]
          else if currentPath
              .filter(n => n.toLowerCase == n)
              .groupBy(identity)
              .count(_._2.size == 2) > 1
          then Set.empty[List[String]]
          else
            map(node).flatMap(n => findPaths(map, start, end, n :: currentPath))

    Some(findPaths(input).size)
