module Aoc2020
  struct SeatSystem
    getter map : Array(Array(Char))

    def initialize(input, @occupied_limit : Int32)
      @map = input.lines.map &.each_char.to_a
    end

    def count_occupied
      @map.sum { |l| l.count('#') }
    end

    def stabilize
      iterate do |i, j|
        occupied = 0
        {-1, 0, 1}.each do |dx|
          {-1, 0, 1}.each do |dy|
            x = i + dx
            y = j + dy

            if (x == i && y == j) || x < 0 || x == @map.size || y < 0 || y == @map[0].size
              next
            end

            occupied += 1 if @map[x][y] == '#'
          end
        end

        occupied
      end
    end

    def stabilize2
      iterate do |i, j|
        occupied = 0
        {-1, 0, 1}.each do |dx|
          {-1, 0, 1}.each do |dy|
            (1..).each do |m|
              x = i + m*dx
              y = j + m*dy

              if (x == i && y == j) || x < 0 || x == @map.size || y < 0 || y == @map[0].size
                break
              end

              if @map[x][y] == '#'
                occupied += 1
                break
              elsif @map[x][y] == 'L'
                break
              end
            end
          end
        end

        occupied
      end
    end

    private def iterate
      edits = Deque(Tuple(Int32, Int32, Char)).new
      loop do
        @map.each_index do |i|
          @map[i].each_index do |j|
            occupied = yield i, j

            if @map[i][j] == 'L' && occupied == 0
              edits << {i, j, '#'}
            elsif @map[i][j] == '#' && occupied >= @occupied_limit
              edits << {i, j, 'L'}
            end
          end
        end

        if edits.size == 0
          return
        else
          while edits.size > 0
            i, j, c = edits.shift
            @map[i][j] = c
          end
        end
      end
    end
  end

  INPUT_DAY11 = File.read("./inputs/11.txt")

  def self.day11p1
    s = SeatSystem.new(INPUT_DAY11, 4)
    s.stabilize
    s.count_occupied
  end

  def self.day11p2
    s = SeatSystem.new(INPUT_DAY11, 5)
    s.stabilize2
    s.count_occupied
  end
end
