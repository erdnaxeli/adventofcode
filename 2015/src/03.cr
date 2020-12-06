require "iterator"

struct Santa
  def initialize(@directions : String)
  end

  def distinct_houses
    x, y = 0, 0
    houses = Set{ {0, 0} }
    @directions.each_char do |c|
      case c
      when '^'
        y += 1
      when '>'
        x += 1
      when 'v'
        y -= 1
      when '<'
        x -= 1
      else
        next
      end

      houses << {x, y}
    end

    houses.size
  end

  def distinct_houses_iterator
    houses = Set{ {0, 0} }
    HousesIterator.new(@directions.each_char).each { |h| houses << h }
    houses.size
  end

  def distinct_houses_with_robot
    iter = @directions.each_char
    santa = HousesIterator.new(iter)
    robot = HousesIterator.new(iter)

    houses = Set{ {0, 0} }
    santa.zip(robot).each do |santa_house, robot_house|
      houses << santa_house
      houses << robot_house
    end

    houses.size
  end

  def distinct_houses_with_robot_2
    house_santa = {0, 0}
    house_robot = {0, 0}
    houses = Set{ {0, 0} }

    iter = @directions.each_char.reject('\n')
    iter.zip?(iter) do |s, r|
      house_santa = next_house(house_santa, s)
      houses << house_santa

      if !r.nil?
        house_robot = next_house(house_robot, r)
        houses << house_robot
      end
    end

    houses.size
  end

  private def next_house(house, c)
    x, y = house
    case c
    when '^'
      y += 1
    when '>'
      x += 1
    when 'v'
      y -= 1
    when '<'
      x -= 1
    else
      raise "BUG: unreachable"
    end

    {x, y}
  end

  private struct HousesIterator
    include Iterator(self)

    def initialize(@iterator : Iterator(Char))
      @x, @y = 0, 0
    end

    def next
      case c = @iterator.next
      when '^'
        @y += 1
      when '>'
        @x += 1
      when 'v'
        @y -= 1
      when '<'
        @x -= 1
      when Iterator::Stop
        return stop
      else
        return self.next
      end

      {@x, @y}
    end
  end
end
