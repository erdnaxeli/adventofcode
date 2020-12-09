module Aoc2020
  struct XMAS(T)
    getter last_item : T?

    def self.find_invalid(input)
      xmas = XMAS(T).new(input)
      while xmas.next_valid?
      end

      xmas.last_item
    end

    def initialize(@input : Iterator(T))
      @last_25 = Deque(T).new
      input.first(25).each { |x| @last_25 << x }
      @last_item = nil
    end

    # Consumes the next item
    def next_valid? : Bool
      item = @input.next
      if item.is_a?(Iterator::Stop)
        return false
      end

      valid = @last_25.any? do |y|
        item - y != y && @last_25.includes? item - y
      end

      @last_item = item
      @last_25.shift
      @last_25 << item

      valid
    end

    def weakness(invalid : T) : T?
      buffer = @input.to_a
      (0...buffer.size - 2).each do |i|
        (i + 1...buffer.size).each do |j|
          total = buffer[i..j].sum
          if total == invalid
            return buffer[i..j].min + buffer[i..j].max
          elsif total > invalid
            break
          end
        end
      end
    end
  end

  def self.day9p1
    input = File.read("./inputs/09.txt").each_line.map &.to_i64
    XMAS(Int64).find_invalid(input)
  end

  def self.day9p2
    input = File.read("./inputs/09.txt").each_line.map &.to_i64
    invalid = XMAS(Int64).find_invalid(input).not_nil!
    input = File.read("./inputs/09.txt").each_line.map &.to_i64
    XMAS(Int64).new(input).weakness(invalid)
  end
end

pp Aoc2020.day9p1
