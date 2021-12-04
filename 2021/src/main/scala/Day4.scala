package aoc.twentyone

type RandomNumbers = List[Int]
type Board = List[List[Int]]

class Day4 extends SimplePuzzle[(RandomNumbers, List[Board]), Int]:
  override protected def input(
      raw: Iterator[String]
  ): (RandomNumbers, List[Board]) =
    raw.toList match
      case numbers :: "" :: boards =>
        (readNumbers(numbers, ','), readBoards(boards))
      case _ => throw Exception("invalid input")

  override protected def part1(
      input: (RandomNumbers, List[Board])
  ): Option[Int] =
    val (numbers, boards) = input
    val (numbersCalled, winningBoard) = solveBoards(boards, numbers)
      // find the board that wins first
      .sorted.headOption.flatten match
      case None    => throw Exception("No wining board found")
      case Some(x) => x

    Some(computeScore(winningBoard, numbersCalled))

  override protected def part2(
      input: (RandomNumbers, List[Board])
  ): Option[Int] =
    val (numbers, boards) = input
    val (numbersCalled, winningBoard) = solveBoards(boards, numbers)
      // find the board that wins last
      .sorted.lastOption.flatten match
      case None    => throw Exception("No wining board found")
      case Some(x) => x

    Some(computeScore(winningBoard, numbersCalled))

  private def computeScore(board: Board, numbersCalled: RandomNumbers): Int =
    val unmarkedNumbers = board.flatten
      .filter(x => !numbersCalled.contains(x))
      .sum
    val justCalledNumber = numbersCalled.lastOption.getOrElse(0)
    unmarkedNumbers * justCalledNumber

  private def solveBoards(
      boards: List[Board],
      numbers: RandomNumbers
  ): List[Option[(RandomNumbers, Board)]] =
    boards
      // get Some(solution) for each board or None
      .map(board =>
        val solutionRow = solveBoard(board, numbers)
        val solutionColumn = solveBoard(board.transpose, numbers)
        (solutionRow, solutionColumn) match
          case (None, None)    => None
          case (Some(x), None) => Some(x, board)
          case (None, Some(y)) => Some(y, board)
          case (Some(x), Some(y)) =>
            if x.size < y.size then Some(x, board)
            else Some(y, board)
      )

  private def solveBoard(
      board: Board,
      numbers: RandomNumbers
  ): Option[RandomNumbers] =
    board
      .collect {
        // filter winning lines
        case line if line.forall(n => numbers.contains(n)) =>
          // find all numbers called until the line wines
          numbers.take(line.map(n => numbers.indexOf(n)).max + 1)
      }
      // get the line that wins first
      .sorted
      .headOption

  private def readNumbers(line: String, sep: Char = ' '): List[Int] =
    line.split(sep).filterNot(_ == "").map(_.toInt).toList

  private def readBoards(
      boards: List[String],
      accBoards: List[Board] = Nil,
      accBoard: Board = Nil
  ): List[Board] =
    boards match
      case Nil => accBoard :: accBoards
      case line :: "" :: tail =>
        readBoards(tail, (readNumbers(line) :: accBoard) :: accBoards, Nil)
      case line :: tail =>
        readBoards(tail, accBoards, readNumbers(line) :: accBoard)
