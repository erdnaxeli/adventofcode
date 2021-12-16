package aoc.twentyone

import scala.annotation.tailrec

sealed class Packet(val version: Int, val typeId: Int, val value: Long)

class LiteralValue(
    override val version: Int,
    override val typeId: Int,
    override val value: Long
) extends Packet(version, typeId, value)

class Operator(
    override val version: Int,
    override val typeId: Int,
    override val value: Long,
    val subpackets: List[Packet]
) extends Packet(version, typeId, value)

class Day16 extends SimplePuzzle[List[Packet], Int]:
  override protected def input(raw: Iterator[String]): List[Packet] =
    if raw.isEmpty then ???
    else
      val binary = raw.next
        .collect {
          case '0' => "0000"
          case '1' => "0001"
          case '2' => "0010"
          case '3' => "0011"
          case '4' => "0100"
          case '5' => "0101"
          case '6' => "0110"
          case '7' => "0111"
          case '8' => "1000"
          case '9' => "1001"
          case 'A' => "1010"
          case 'B' => "1011"
          case 'C' => "1100"
          case 'D' => "1101"
          case 'E' => "1110"
          case 'F' => "1111"
        }
        .flatMap(_.map(_.toString.toShort))
        .toList

      readPackets(binary)

  override def part1(input: List[Packet]): Option[Int] =
    def sumVersion(packets: List[Packet]): Int =
      packets match
        case Nil                       => 0
        case (x: LiteralValue) :: tail => x.version + sumVersion(tail)
        case (x: Operator) :: tail =>
          x.version + sumVersion(x.subpackets) + sumVersion(tail)

    Some(sumVersion(input))

  override def part2(input: List[Packet]): Option[Int] =
    input match
      case packet :: Nil => println(packet.value)
      case _             => ???

    Some(0)

  private def readPackets(stream: List[Short]): List[Packet] =
    readPacketsCount(stream)(0)

  @tailrec
  private def readPacketsCount(
      stream: List[Short],
      count: Option[Int] = None,
      acc: List[Packet] = List.empty[Packet]
  ): (List[Packet], List[Short]) =
    count match
      case Some(0) => (acc.reverse, stream)
      case Some(v) =>
        readNextPacket(stream) match
          case None => (acc.reverse, stream)
          case Some(streamTail, packet) =>
            readPacketsCount(streamTail, Some(v - 1), packet :: acc)
      case None =>
        readNextPacket(stream) match
          case None => (acc.reverse, stream)
          case Some(streamTail, packet) =>
            readPacketsCount(streamTail, None, packet :: acc)

  private def readNextPacket(
      stream: List[Short]
  ): Option[(List[Short], Packet)] =
    stream match
      case Nil => None
      case a :: b :: c :: 1 :: 0 :: 0 :: tail =>
        val version = Integer.parseInt(s"$a$b$c", 2)
        val typeId = 4
        val (value, streamTail) = readLiteralValue(tail)
        Some(streamTail, LiteralValue(version, typeId, value))
      case a :: b :: c :: d :: e :: f :: tail =>
        val version = Integer.parseInt(s"$a$b$c", 2)
        val typeId = Integer.parseInt(s"$d$e$f", 2)
        val (subpackets, streamTail) = readSubpackets(tail)
        val value = typeId match
          case 0 => subpackets.map(_.value).sum
          case 1 => subpackets.map(_.value).product
          case 2 => subpackets.map(_.value).min
          case 3 => subpackets.map(_.value).max
          case 5 =>
            subpackets match
              case a :: b :: Nil => if a.value > b.value then 1L else 0L
              case _             => ???
          case 6 =>
            subpackets match
              case a :: b :: Nil => if a.value < b.value then 1L else 0L
              case _             => ???
          case 7 =>
            subpackets match
              case a :: b :: Nil => if a.value == b.value then 1L else 0L
              case _             => ???
        Some(streamTail, Operator(version, typeId, value, subpackets))
      case _ => None

  private def readLiteralValue(stream: List[Short]): (Long, List[Short]) =
    @tailrec
    def read(
        stream: List[Short],
        acc: List[Short] = List.empty[Short],
        bitsRead: Int = 6
    ): (Long, List[Short]) =
      stream match
        case 1 :: a :: b :: c :: d :: tail =>
          read(tail, d :: c :: b :: a :: acc, bitsRead + 5)
        case 0 :: a :: b :: c :: d :: tail =>
          val value =
            java.lang.Long.parseLong(
              (d :: c :: b :: a :: acc).reverse.map(_.toString).mkString,
              2
            )
          (value, tail)
        case _ => ???

    read(stream)

  private def readSubpackets(stream: List[Short]): (List[Packet], List[Short]) =
    stream match
      case 0 :: tail =>
        val (lengthBin, streamTail) = tail.splitAt(15)
        val length = Integer.parseInt(lengthBin.map(_.toString).mkString, 2)
        val (packetsStream, streamTailTail) = streamTail.splitAt(length)
        (readPackets(packetsStream), streamTailTail)
      case 1 :: tail =>
        val (countBin, streamTail) = tail.splitAt(11)
        val count = Integer.parseInt(countBin.map(_.toString).mkString, 2)
        readPacketsCount(streamTail, Some(count))
      case _ => ???
