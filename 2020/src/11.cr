require "stumpy_png"

module Aoc2020
  struct SeatSystem
    getter map : Array(Array(Char))

    def initialize(input, @occupied_limit : Int32, @folder : String)
      @map = input.lines.map &.each_char.to_a
    end

    def count_occupied
      @map.sum { |l| l.count('#') }
    end

    def plot_map(i)
      zoom = 4
      canvas = StumpyPNG::Canvas.new(@map.size * zoom, @map[0].size * zoom)

      colors = {
        'L' => StumpyPNG::RGBA.from_rgb_n(0, 186, 227, 8), # some light blue
        '#' => StumpyPNG::RGBA.from_rgb_n(255, 0, 0, 8),   # red
        '.' => StumpyPNG::RGBA.from_rgb_n(0, 0, 0, 8),     # black
      }

      @map.each_index do |x|
        @map[x].each_index do |y|
          (x*zoom...(x*zoom + zoom)).each do |m|
            (y*zoom...(y*zoom + zoom)).each do |n|
              canvas[m, n] = colors[@map[x][y]]
            end
          end
        end
      end

      file_name = sprintf "%02d.png", i
      StumpyPNG.write(canvas, "#{@folder}/#{file_name}")
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
      generation = 0
      edits = Deque(Tuple(Int32, Int32, Char)).new
      loop do
        generation += 1
        plot_map(generation)

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
    s = SeatSystem.new(INPUT_DAY11, 4, "graphs/11/1")
    s.stabilize
    s.count_occupied
  end

  def self.day11p2
    s = SeatSystem.new(INPUT_DAY11, 5, "graphs/11/2")
    s.stabilize2
    s.count_occupied
  end
end
