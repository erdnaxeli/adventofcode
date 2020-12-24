require "./17"

module Aoc2020
  def self.flip(instructions : Array(Array(String)))
    map = Hash(Tuple(Int32, Int32), Bool).new { false }

    instructions.each do |instruction|
      coordinates = read_coordinates(instruction)
      map[coordinates] = !map[coordinates]
    end

    map
  end

  def self.read_coordinates(instruction) : Tuple(Int32, Int32)
    x = y = 0
    instruction.each do |direction|
      case direction
      when "nw"
        x -= 1
        y += 1
      when "w"
        x -= 2
      when "sw"
        x -= 1
        y -= 1
      when "se"
        x += 1
        y -= 1
      when "e"
        x += 2
      when "ne"
        x += 1
        y += 1
      else
        raise "Unknown direction #{direction}"
      end
    end

    {x, y}
  end

  INPUT_DAY24 = File.read("./inputs/24.txt").lines.map do |line|
    chars = line.each_char
    r = Array(String).new
    while (c = chars.next).is_a?(Char)
      r << String.build do |str|
        if c == 'w' || c == 'e'
          str << c
        else
          str << c << chars.next
        end
      end
    end
    r
  end

  module Day24Rules
    def alive?(neighbors_alive)
      neighbors_alive == 2
    end

    def dead?(neighbors_alive)
      !(1..2).includes?(neighbors_alive)
    end
  end

  struct Tile
    extend Day24Rules

    def initialize(@x = 0, @y = 0)
    end

    def coordinates
      {@x, @y}
    end

    def each_neighbors
      yield Tile.new(@x - 1, @y + 1)
      yield Tile.new(@x - 2, @y)
      yield Tile.new(@x - 1, @y - 1)
      yield Tile.new(@x + 1, @y - 1)
      yield Tile.new(@x + 2, @y)
      yield Tile.new(@x + 1, @y + 1)
    end
  end

  def self.day24p1
    flip(INPUT_DAY24).count { |_, v| v }
  end

  def self.day24p2
    map = flip(INPUT_DAY24).compact_map do |k, v|
      if v
        {k, Tile.new(k[0], k[1])}
      end
    end.to_h
    resolver = ConwaySpace::Resolver(Tuple(Int32, Int32), Tile).new(map)
    100.times { resolver.next }
    resolver.cubes.size
  end
end

puts Aoc2020.day24p2
