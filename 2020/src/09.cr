module Aoc2020
    struct XMAS
        getter last_item : Int32?

        def initialize(@input : Iterator(Int32))
            @last_25 = Deque(Int32).new()
            # https://github.com/crystal-lang/crystal/issues/9924
            25.times do
                item = @input.next
                break if item.is_a?(Iterator::Stop)
                @last_25 << item
            end
            @last_item = nil
        end

        # Consummes the next item
        def next_valid? : Bool
            puts "valid"
            item = @input.next
            if item.is_a?(Iterator::Stop)
                return false
            end

            valid = @last_25.any? do |y|
                item - y != y && @last_25.includes? item - y
            end
            pp! item, valid, @last_25

            @last_item = item
            @last_25.shift
            @last_25 << item

            valid
        end
    end

    def self.day9p1
        input = File.read("./inputs/09.txt").each_line.map &.to_i
        xmas = XMAS.new(input)
        while xmas.next_valid?
        end

        xmas.last_item
    end

    def self.day9p2
    end
end

pp Aoc2020.day9p1