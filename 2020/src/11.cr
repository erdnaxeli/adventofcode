module Aoc2020
  struct SeatSystem
    getter map : Array(Array(Char))

    def initialize(input)
      @map = input.lines.map &.each_char.to_a
    end

    def count_occupied
      @map.sum { |l| l.count('#') }
    end

    def stabilize
      edits = Deque(Tuple(Int32, Int32, Char)).new
      loop do
        
        map = @map.map { |l| l.join }.join("\n")
        # puts map + "\n\n"

        @map.each_index do |i|
          @map[i].each_index do |j|
            # p! i, j
            occupied = 0
            {-1, 0, 1}.each do |dx|
              {-1, 0, 1}.each do |dy|
                x = i + dx
                y = j + dy

                if (x == i && y == j) || x < 0 || x == @map.size || y < 0 || y == @map[0].size
                  next
                end

                occupied += 1 if @map[x][y] == '#'
                # p! x, y, occupied
              end
            end

            if @map[i][j] == 'L' && occupied == 0
              edits << {i, j, '#'}
            elsif @map[i][j] == '#' && occupied >= 4
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
    s = SeatSystem.new(INPUT_DAY11)
    s.stabilize
    s.count_occupied
  end
end

puts Aoc2020.day11p1