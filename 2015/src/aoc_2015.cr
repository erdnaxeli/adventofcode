require "./01"
require "./02"
require "./03"

module Aoc2015
  VERSION = "0.1.0"

  module Day1
    INPUT = File.read("./inputs/01.txt")

    def self.p1
      Elevator.new(INPUT).floor
    end

    def self.p2
      Elevator.new(INPUT).index(-1)
    end
  end

  module Day2
    INPUT = File.read("./inputs/02.txt")

    def self.p1
      INPUT.each_line.map { |l| Present.new(l) }.sum &.paper
    end

    def self.p2
      INPUT.each_line.map { |l| Present.new(l) }.sum &.ribbon
    end
  end

  module Day3
    extend self

    INPUT = File.read("./inputs/03.txt")

    def p1
      Santa.new(INPUT).distinct_houses
    end

    def p2
      Santa.new(INPUT).distinct_houses_with_robot_2
    end
  end
end

pp Aoc2015::Day3.p2
