module ConwaySpace
  struct Resolver(Coordinates, Cube)
    getter cubes = Hash(Coordinates, Cube).new

    def initialize(input)
      read_cubes_from_map(input).each do |cube|
        @cubes[cube.coordinates] = cube
      end
    end

    def initialize(@cubes)
    end

    def next
      to_delete = Array(Coordinates).new
      to_add = Array(Cube).new
      checked = Set(Coordinates).new

      @cubes.each_value do |cube|
        valid = count_valid_around(cube)

        if Cube.dead?(valid)
          to_delete << cube.coordinates
        end

        cube.each_neighbors do |neighbor|
          if checked.add?(neighbor.coordinates)
            valid = count_valid_around(neighbor)

            if Cube.alive?(valid)
              to_add << neighbor
            end
          end
        end
      end

      to_delete.each do |coordinates|
        @cubes.delete(coordinates)
      end

      to_add.each do |cube|
        @cubes[cube.coordinates] = cube
      end
    end

    private def read_cubes_from_map(input)
      input.each_line.with_index.flat_map do |line_index|
        line, x = line_index
        line.each_char.with_index.compact_map do |char_index|
          char, y = char_index
          if char == '#'
            Cube.new(x, y)
          end
        end
      end
    end

    private def count_valid_around(cube : Cube)
      count = 0
      cube.each_neighbors do |neighbor|
        if @cubes.has_key?(neighbor.coordinates)
          count += 1
        end
      end
      count
    end
  end

  module Day17Rules
    # From death to alive?
    def alive?(neighbors_alive)
      neighbors_alive == 3
    end

    # From alive to death?
    def dead?(neighbors_alive)
      !(2..3).includes?(neighbors_alive)
    end
  end

  struct Cube3D
    extend Day17Rules

    def initialize(@x = 0, @y = 0, @z = 0)
    end

    def coordinates
      {@x, @y, @z}
    end

    def each_neighbors
      (-1..1).each do |dx|
        (-1..1).each do |dy|
          (-1..1).each do |dz|
            if !(dx == dy == dz == 0)
              yield Cube3D.new(@x + dx, @y + dy, @z + dz)
            end
          end
        end
      end
    end
  end

  struct Cube4D
    extend Day17Rules

    def initialize(@w = 0, @x = 0, @y = 0, @z = 0)
    end

    def coordinates
      {@w, @x, @y, @z}
    end

    def each_neighbors
      (-1..1).each do |dw|
        (-1..1).each do |dx|
          (-1..1).each do |dy|
            (-1..1).each do |dz|
              if !(dw == dx == dy == dz == 0)
                yield Cube4D.new(@w + dw, @x + dx, @y + dy, @z + dz)
              end
            end
          end
        end
      end
    end
  end
end

module Aoc2020
  INPUT_DAY17 = File.read("./inputs/17.txt")

  def self.day17p1
    resolver = ConwaySpace::Resolver(Tuple(Int32, Int32, Int32), ConwaySpace::Cube3D).new(INPUT_DAY17)
    6.times { resolver.next }
    resolver.cubes.size
  end

  def self.day17p2
    resolver = ConwaySpace::Resolver(Tuple(Int32, Int32, Int32, Int32), ConwaySpace::Cube4D).new(INPUT_DAY17)
    6.times { resolver.next }
    resolver.cubes.size
  end
end
