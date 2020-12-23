module Aoc2020
  INPUT_DAY23 = "784235916".chars.map &.to_i

  def self.move(cups : Array(Int32), moves : Int32)
    circle = Deque.new(cups)
    min = circle.min
    max = circle.max

    i = 0
    moves.times do
      i += 1
      puts i if i % 1_000 == 0
      circle.rotate!
      a = circle.shift
      b = circle.shift
      c = circle.shift

      current = circle[-1]
      destination = current - 1
      if destination < min
        destination = max - (min - destination) + 1
      end

      while {a, b, c}.includes?(destination)
        destination -= 1

        if destination < min
          destination = max - (min - destination) + 1
        end
      end

      destination_position = circle.index(destination).not_nil!
      circle.insert(destination_position + 1, c)
      circle.insert(destination_position + 1, b)
      circle.insert(destination_position + 1, a)
    end

    circle
    while circle[0] != 1
      circle.rotate!
    end

    circle.shift
    circle
  end

  def self.day23p1
    move(INPUT_DAY23, 100).join
  end

  def self.day23p2
    move(INPUT_DAY23 + (INPUT_DAY23.max..1_000_000).to_a, 10_000_000).to_a.[0..1].product
  end
end

puts Aoc2020.day23p2
