struct ConwaySpace
  enum State
    Active
    Inactive
  end

  getter map = [[[] of ConwaySpace::State]]

  @x_size = 0
  @y_size = 0
  @z_size = 0

  def initialize(input : String)
    x, y = 0, 0
    input.each_line do |l|
      l.each_char do |c|
        case c
        when '#'
          @map[x][y] << State::Active
        when '.'
          @map[x][y] << State::Inactive
        else
          raise "BUG: unknown cube type"
        end

        y += 1
        @map[x] << [] of State
      end

      x += 1
      y = 0
      @map << [[] of State]
    end

    update_sizes
    print_map
  end

  def next
    diff = get_diff
    # diff += update_outer_cubes
    apply_diff(diff)

    puts "New map"
    print_map
  end

  def each_cube
    @map.each_with_index do |_, x|
      @map[x].each_with_index do |_, y|
        @map[x][y].each_with_index do |cube, z|
          if cube == State::Active
            yield({x, y, z})
          end
        end
      end
    end
  end

  def count_cubes
    count = 0
    each_cube { count += 1 }
    count
  end

  def print_map
    (0...@z_size).each do |z|
      puts "z=#{z}"
      (0...@x_size).each do |x|
        (0...@y_size).each do |y|
          case @map[x]?.try(&.[y]?.try &.[z]?)
          when State::Active
            print '#'
          else
            print '.'
          end
        end

        print '\n'
      end

      print "\n\n"
    end
  end

  private def update_sizes
    @x_size = @map.size
    @y_size = @map.map { |y| y.size }.max
    @z_size = @map.map { |y| y.map { |z| z.size }.max }.max
  end

  private def get_diff
    diff = Array(Tuple(Int32, Int32, Int32, State)).new

    (-1..@x_size).each do |x|
      (-1..@y_size).each do |y|
        (-1..@z_size).each do |z|
          actives = get_actives_around(x, y, z)

          if (0...@x_size).includes?(x) && (0...@map[x].size).includes?(y) && (0...@map[x][y].size).includes?(z)
            if @map[x][y][z] == State::Active && actives < 2 || actives > 3
              diff << {x, y, z, State::Inactive}
            elsif @map[x][y][z] == State::Inactive && actives == 3
              diff << {x, y, z, State::Active}
            end
          else
            if actives == 3
              diff << {x, y, z, State::Active}
            end
          end
        end
      end
    end

    diff
  end

  private def get_actives_around(x, y, z)
    (-1..1).sum do |dx|
      (-1..1).sum do |dy|
        (-1..1).count do |dz|
          if dx == dy == dz == 0
            false
          else
            nx = x + dx
            ny = y + dy
            nz = z + dz

            if nx < 0 || ny < 0 || nz < 0
              false
            elsif nx >= @map.size || ny >= @map[nx].size || nz >= @map[nx][ny].size
              false
            else
              @map[nx][ny][nz] == State::Active
            end
          end
        end
      end
    end
  end

  private def apply_diff(diff)
    dx, dy, dz = apply_unshifts(diff)

    diff.each do |x, y, z, state|
      x += dx
      y += dy
      z += dz

      while x >= @map.size
        @map << [[] of State]
      end

      while y >= @map[x].size
        @map[x] << [] of State
      end

      while z >= @map[x][y].size
        @map[x][y] << State::Inactive
      end

      @map[x][y][z] = state
    end

    update_sizes
  end

  private def apply_unshifts(diff)
    dx = dy = dz = 0

    diff.each do |x, y, z, _|
      x += dx
      y += dy
      z += dz

      if x < 0 || y < 0 || z < 0
        ddx, ddy, ddz = unshift(x, y, z)
        dx += ddx
        dy += ddy
        dz += ddz
      end
    end

    {dx, dy, dz}
  end

  private def unshift(x, y, z)
    dx = dy = dz = 0

    if x < 0
      dx = 1
      @map.unshift([[] of State])
    end

    if y < 0
      dy = 1
      @map.each do |x_row|
        x_row.unshift([] of State)
      end
    end

    if z < 0
      dz = 1
      @map.each do |x_row|
        x_row.each do |y_row|
          y_row.unshift(State::Inactive)
        end
      end
    end

    {dx, dy, dz}
  end
end

struct ConwaySpace4
  getter map = [[[[] of ConwaySpace::State]]]

  @x_size = 0
  @y_size = 0
  @z_size = 0
  @a_size = 0

  def initialize(input : String)
    x, y, z = 0, 0, 0
    input.each_line do |l|
      l.each_char do |c|
        case c
        when '#'
          @map[x][y][z] << ConwaySpace::State::Active
        when '.'
          @map[x][y][z] << ConwaySpace::State::Inactive
        else
          raise "BUG: unknown cube type"
        end

        y += 1
        @map[x] << [[] of ConwaySpace::State]
      end

      x += 1
      y = 0
      @map << [[[] of ConwaySpace::State]]
    end

    update_sizes
    print_map
  end

  def next
    diff = get_diff
    puts diff
    apply_diff(diff)

    puts "new map"
    print_map
  end

  def each_cube
    @map.each_with_index do |_, x|
      @map[x].each_with_index do |_, y|
        @map[x][y].each_with_index do |_, z|
          @map[x][y][z].each_with_index do |cube, a|
            if cube == ConwaySpace::State::Active
              yield({x, y, z, a})
            end
          end
        end
      end
    end
  end

  def count_cubes
    count = 0
    each_cube { count += 1 }
    count
  end

  def print_map
    (0...@a_size).each do |a|
      puts "a=#{a}"
      (0...@z_size).each do |z|
        puts "z=#{z}"
        (0...@x_size).each do |x|
          (0...@y_size).each do |y|
            case @map[x]?.try(&.[y]?.try(&.[z]?.try &.[a]?))
            when ConwaySpace::State::Active
              print '#'
            else
              print '.'
            end
          end

          print '\n'
        end

        print "\n\n"
      end

      print "\n\n"
    end
  end

  private def update_sizes
    @x_size = @map.size
    @y_size = @map.map { |y| y.size }.max
    @z_size = @map.map { |y| y.map { |z| z.size }.max }.max
    @a_size = @map.map { |y| y.map { |z| z.map { |a| a.size }.max }.max }.max
  end

  private def get_diff
    diff = Array(Tuple(Int32, Int32, Int32, Int32, ConwaySpace::State)).new

    (-1..@x_size).each do |x|
      (-1..@y_size).each do |y|
        (-1..@z_size).each do |z|
          (-1..@a_size).each do |a|
            actives = get_actives_around(x, y, z, a)

            if (0...@x_size).includes?(x) && (0...@map[x].size).includes?(y) && (0...@map[x][y].size).includes?(z) && (0...@map[x][y][z].size).includes?(a)
              if @map[x][y][z][a] == ConwaySpace::State::Active && actives < 2 || actives > 3
                diff << {x, y, z, a, ConwaySpace::State::Inactive}
              elsif @map[x][y][z][a] == ConwaySpace::State::Inactive && actives == 3
                diff << {x, y, z, a, ConwaySpace::State::Active}
              end
            else
              if actives == 3
                diff << {x, y, z, a, ConwaySpace::State::Active}
              end
            end
          end
        end
      end
    end

    diff
  end

  private def get_actives_around(x, y, z, a)
    (-1..1).sum do |dx|
      (-1..1).sum do |dy|
        (-1..1).sum do |dz|
          (-1..1).count do |da|
            if dx == dy == dz == da == 0
              false
            else
              nx = x + dx
              ny = y + dy
              nz = z + dz
              na = a + da

              if nx < 0 || ny < 0 || nz < 0 || na < 0
                false
              elsif nx >= @map.size || ny >= @map[nx].size || nz >= @map[nx][ny].size || na >= @map[nx][ny][nz].size
                false
              else
                @map[nx][ny][nz][na] == ConwaySpace::State::Active
              end
            end
          end
        end
      end
    end
  end

  private def apply_diff(diff)
    dx, dy, dz, da = apply_unshifts(diff)

    diff.each do |x, y, z, a, state|
      x += dx
      y += dy
      z += dz
      a += da

      while x >= @map.size
        @map << [[[] of ConwaySpace::State]]
      end

      while y >= @map[x].size
        @map[x] << [[] of ConwaySpace::State]
      end

      while z >= @map[x][y].size
        @map[x][y] << [] of ConwaySpace::State
      end

      while a >= @map[x][y][z].size
        @map[x][y][z] << ConwaySpace::State::Inactive
      end

      @map[x][y][z][a] = state
    end

    update_sizes
  end

  private def apply_unshifts(diff)
    dx = dy = dz = da = 0

    diff.each do |x, y, z, a, _|
      x += dx
      y += dy
      z += dz
      a += da

      if x < 0 || y < 0 || z < 0 || a < 0
        ddx, ddy, ddz, dda = unshift(x, y, z, a)
        dx += ddx
        dy += ddy
        dz += ddz
        da += dda
      end
    end

    {dx, dy, dz, da}
  end

  private def unshift(x, y, z, a)
    dx = dy = dz = da = 0

    if x < 0
      dx = 1
      @map.unshift([[[] of ConwaySpace::State]])
    end

    if y < 0
      dy = 1
      @map.each do |x_row|
        x_row.unshift([[] of ConwaySpace::State])
      end
    end

    if z < 0
      dz = 1
      @map.each do |x_row|
        x_row.each do |y_row|
          y_row.unshift([] of ConwaySpace::State)
        end
      end
    end

    if a < 0
      da = 1
      @map.each do |x_row|
        x_row.each do |y_row|
          y_row.each do |z_row|
            z_row.unshift(ConwaySpace::State::Inactive)
          end
        end
      end
    end

    {dx, dy, dz, da}
  end
end

module Aoc2020
  INPUT_DAY17 = File.read("./inputs/17.txt")

  def self.day17p1
    map = ConwaySpace.new(INPUT_DAY17)
    6.times { map.next }
    total = 0
    map.each_cube { total += 1 }
    total
  end

  def self.day17p2
    map = ConwaySpace4.new(INPUT_DAY17)
    6.times { map.next }
    total = 0
    map.each_cube { total += 1 }
    total
  end
end
