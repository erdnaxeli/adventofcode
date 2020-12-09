module Aoc2020
  struct BoardingPass
    include Comparable(self)

    getter seat : String
    @seat_row : String
    @seat_column : String

    def initialize(@seat)
      @seat_row = @seat[0..6]
      @seat_column = @seat[7..10]
    end

    def seat_row
      compute_position(@seat_row.tr("FB", "AB"), 0, 127)
    end

    def seat_column
      compute_position(@seat_column.tr("LR", "AB"), 0, 7)
    end

    def id
      seat_row * 8 + seat_column
    end

    def <=>(other)
      @seat.tr("FBLR", "ABAB") <=> other.seat.tr("FBLR", "ABAB")
    end

    private def compute_position(instructions, min, max)
      instruction = instructions[0]
      middle = (max - min) // 2

      case instruction
      when 'A'
        if middle == 0
          min
        else
          compute_position(instructions[1..], min, min + middle)
        end
      when 'B'
        if middle == 0
          max
        else
          compute_position(instructions[1..], max - middle, max)
        end
      else
        raise "BUG: unreachable"
      end
    end
  end

  def self.day5p1
    input = File.read("./inputs/05.txt").each_line.map { |l| BoardingPass.new(l) }.to_a.sort
    puts input[-1].id
  end

  def self.day5p2
    File.read("./inputs/05.txt").each_line.map { |l| BoardingPass.new(l) }.to_a.sort.reduce(0) do |acc, pass|
      if acc == 0
        pass.id
      else
        if pass.id - acc > 1
          puts acc + 1
          break
        else
          pass.id
        end
      end
    end
  end
end

Aoc2020.day5p1
Aoc2020.day5p2
