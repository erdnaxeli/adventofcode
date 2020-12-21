module Aoc2020
  @[Flags]
  enum Side
    Up
    Down
    Left
    Right
  end

  class Tile
    getter id : Int32
    getter up : Array(Char)
    getter down : Array(Char)
    getter left : Array(Char)
    getter right : Array(Char)
    getter value : Array(Array(Char))

    property taken = Side::None

    def initialize(@id, @value)
      @up = @value[0]
      @down = @value[-1]
      @left = @value.map { |l| l[0] }
      @right = @value.map { |l| l[-1] }
    end

    def each_position
      yield self

      if @taken.none?
        yield h_sym
        yield v_sym
        yield rotate
        yield rotate.h_sym
        yield rotate.v_sym
        yield rotate.rotate
        yield rotate.rotate.h_sym
        yield rotate.rotate.v_sym
        yield rotate.rotate.rotate
        yield rotate.rotate.rotate.h_sym
        yield rotate.rotate.rotate.v_sym
      end
    end

    # Rotates 90° right
    protected def rotate
      tmp = @value.map_with_index do |line, x|
        line.map_with_index do |_, y|
          @value[y][@value[0].size - 1 - x]
        end
      end
      Tile.new(id: @id, value: tmp)
    end

    protected def h_sym
      Tile.new(id: @id, value: @value.reverse)
    end

    protected def v_sym
      Tile.new(id: @id, value: @value.map &.reverse)
    end
  end

  class ImageTiles
    getter map = Hash(Tuple(Int32, Int32), Tile).new

    def initialize(@tiles : Array(Tile))
    end

    def resolve
      to_check = Deque(Tuple(Int32, Int32)).new
      to_check << {0, 0}
      @map[{0, 0}] = @tiles.pop

      while to_check.size > 0
        k = to_check.shift
        x, y = k
        map_tile = @map[k]
        if @tiles.size == 0
          break
        end

        matches = Array(Tile).new
        @tiles.each do |tile|
          tile.each_position do |candidate|
            if !map_tile.taken.up? && map_tile.up == candidate.down
              @map[{x, y + 1}] = candidate
              to_check << {x, y + 1}
              map_tile.taken |= Side::Up
              candidate.taken |= Side::Down
            elsif !map_tile.taken.down? && map_tile.down == candidate.up
              @map[{x, y - 1}] = candidate
              to_check << {x, y - 1}
              map_tile.taken |= Side::Down
              candidate.taken |= Side::Up
            elsif !map_tile.taken.left? && map_tile.left == candidate.right
              @map[{x - 1, y}] = candidate
              to_check << {x - 1, y}
              map_tile.taken |= Side::Left
              candidate.taken |= Side::Right
            elsif !map_tile.taken.right? && map_tile.right == candidate.left
              @map[{x + 1, y}] = candidate
              to_check << {x + 1, y}
              map_tile.taken |= Side::Right
              candidate.taken |= Side::Left
            else
              next
            end

            matches << tile
            break
          end

          if map_tile.taken.is_a?(Side::All)
            break
          end
        end

        matches.each { |tile| @tiles.delete(tile) }
      end
    end

    def print_map : Nil
      line = nil
      @map.keys.sort.each do |k|
        if line.nil?
          line = k[0]
        elsif k[0] != line
          puts
          line = k[0]
        end

        print @map[k]
      end
    end

    def four_corners
      min = @map.keys.min
      max = @map.keys.max

      [
        @map[min].id,
        @map[{min[0], max[1]}].id,
        @map[{max[0], min[1]}].id,
        @map[max].id,
      ].map &.to_u64
    end
  end

  def self.read_tiles(input) : Array(Tile)
    input.split("\n\n").map do |tile|
      id = tile.lines[0][5...-1].to_i
      Tile.new(id, tile.lines[1..].map(&.chars))
    end
  end

  INPUT_DAY20 = File.read("./inputs/20.txt")

  def self.day20p1
    image = ImageTiles.new(read_tiles(INPUT_DAY20))
    image.resolve
    image.four_corners.product
  end
end

puts Aoc2020.day20p1
