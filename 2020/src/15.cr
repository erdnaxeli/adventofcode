struct MemoryGame
  @input = Array(Int32).new
  @values = Hash(Int32, Int32).new { 0 }
  @turn = 0

  getter last_value : Int32? = nil

  def initialize(input)
    @input = input.dup
  end

  def play
    if @input.size > 0
      value = @input.shift
    else
      value = @values[@last_value]
      if value != 0
        value = @turn - value
      end
    end

    if last_value = @last_value
      @values[last_value] = @turn
    end

    @turn += 1
    @last_value = value
  end
end

module Aoc2020
  INPUT_DAY15 = "11,18,0,20,1,7,16".split(",").map &.to_i

  def self.day15p1
    g = MemoryGame.new(INPUT_DAY15)
    2020.times do
      g.play
    end

    g.last_value
  end

  def self.day15p2
    g = MemoryGame.new(INPUT_DAY15)
    30_000_000.times do
      g.play
    end

    g.last_value
  end
end

puts Aoc2020.day15p2
