module Aoc2020
  class Boat
    enum Action
      North
      South
      East
      West
      Left
      Right
      Forward

      def self.parse(c)
        case c
        when 'N'
          Action::North
        when 'S'
          Action::South
        when 'E'
          Action::East
        when 'W'
          Action::West
        when 'L'
          Action::Left
        when 'R'
          Action::Right
        when 'F'
          Action::Forward
        else
          raise "BUG: unreachable"
        end
      end
    end

    record(
      Instruction,
      action : Action,
      value : Int32,
    )

    struct Circuit
      include Iterator(Instruction)

      @input : Iterator(String)

      def initialize(input)
        @input = input.each_line
      end

      def next
        line = @input.next
        return stop if line.is_a?(Iterator::Stop)

        Instruction.new(
          action: Action.parse(line[0]),
          value: line[1..].to_i
        )
      end
    end

    getter x : Int32
    getter y : Int32

    def initialize(input, @with_waypoint = false)
      @circuit = Circuit.new(input)

      @x, @y = 0, 0
      @waypoint = [10, 1]
      @direction = {1, 0}
    end

    def travel
      if @with_waypoint
        @circuit.each do |instruction|
          move_with_waypoint(instruction)
        end
      else
        @circuit.each do |instruction|
          move(instruction)
        end
      end
    end

    def move(instruction)
      case instruction.action
      in .north?
        @y += instruction.value
      in .south?
        @y -= instruction.value
      in .east?
        @x += instruction.value
      in .west?
        @x -= instruction.value
      in .left?
        (instruction.value // 90).times do
          @direction = {-@direction[1], @direction[0]}
        end
      in .right?
        (instruction.value // 90).times do
          @direction = {@direction[1], -@direction[0]}
        end
      in .forward?
        @x += @direction[0] * instruction.value
        @y += @direction[1] * instruction.value
      end
    end

    def move_with_waypoint(instruction)
      case instruction.action
      in .north?
        @waypoint[1] += instruction.value
      in .south?
        @waypoint[1] -= instruction.value
      in .east?
        @waypoint[0] += instruction.value
      in .west?
        @waypoint[0] -= instruction.value
      in .left?
        (instruction.value // 90).times do
          @waypoint = [-@waypoint[1], @waypoint[0]]
        end
      in .right?
        (instruction.value // 90).times do
          @waypoint = [@waypoint[1], -@waypoint[0]]
        end
      in .forward?
        @x += @waypoint[0] * instruction.value
        @y += @waypoint[1] * instruction.value
      end
    end
  end

  INPUT_DAY12 = File.read("./inputs/12.txt")

  def self.day12p1
    boat = Boat.new(INPUT_DAY12)
    boat.travel
    boat.x.abs + boat.y.abs
  end

  def self.day12p2
    boat = Boat.new(INPUT_DAY12, with_waypoint: true)
    boat.travel
    boat.x.abs + boat.y.abs
  end
end

