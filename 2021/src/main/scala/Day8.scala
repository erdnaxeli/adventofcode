package aoc.twentyone

type Digit = Set[Char]
case class Screen(
    digits: (
        Digit,
        Digit,
        Digit,
        Digit,
        Digit,
        Digit,
        Digit,
        Digit,
        Digit,
        Digit
    ),
    number: (Digit, Digit, Digit, Digit)
)

class Day8 extends SimplePuzzle[List[Screen], Int]:
  override protected def input(raw: Iterator[String]): List[Screen] =
    val regex =
      raw"([a-g]+) ([a-g]+) ([a-g]+) ([a-g]+) ([a-g]+) ([a-g]+) ([a-g]+) ([a-g]+) ([a-g]+) ([a-g]+) \| ([a-g]+) ([a-g]+) ([a-g]+) ([a-g]+)".r
    raw
      .map(_ match
        case regex(d0, d1, d2, d3, d4, d5, d6, d7, d8, d9, n0, n1, n2, n3) =>
          Screen(
            (
              d0.toSet,
              d1.toSet,
              d2.toSet,
              d3.toSet,
              d4.toSet,
              d5.toSet,
              d6.toSet,
              d7.toSet,
              d8.toSet,
              d9.toSet
            ),
            (n0.toSet, n1.toSet, n2.toSet, n3.toSet)
          )
      )
      .toList

  override protected def part1(input: List[Screen]): Option[Int] =
    Some(
      input
        .flatMap(screen =>
          val searched =
            screen.digits.toList.filter(d => List(2, 3, 4, 7).contains(d.size))
          screen.number.toList.filter(d => searched.contains(d))
        )
        .size
    )

  override protected def part2(input: List[Screen]): Option[Int] =
    Some(
      input
        .map(screen =>
          val digitBySize = screen.digits.toList
            .groupBy(_.size)
            .toMap
          val one = digitBySize.get(2).get.head
          val four = digitBySize.get(4).get.head
          val seven = digitBySize.get(3).get.head
          val eight = digitBySize.get(7).get.head
          val nine =
            digitBySize.get(6).get.find(d => (d & four) == four).get
          val zero =
            digitBySize
              .get(6)
              .get
              .filter(_ != nine)
              .find(d => (d & seven) == seven)
              .get
          val six =
            digitBySize.get(6).get.filter(d => d != nine && d != zero).head
          val three = digitBySize.get(5).get.find(d => (d & seven) == seven).get
          val five =
            digitBySize
              .get(5)
              .get
              .filter(_ != three)
              .find(d => (nine diff d).size == 1)
              .get
          var two =
            digitBySize.get(5).get.filter(d => d != three && d != five).head

          val digitsToDecimal = Map(
            zero -> '0',
            one -> '1',
            two -> '2',
            three -> '3',
            four -> '4',
            five -> '5',
            six -> '6',
            seven -> '7',
            eight -> '8',
            nine -> '9'
          )
          screen.number.toList
            .map(d => digitsToDecimal.get(d).get)
            .mkString
            .toInt
        )
        .sum
    )
