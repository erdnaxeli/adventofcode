module Aoc2020
  INPUT_DAY2 = File.read_lines("./inputs/02.txt")

  record Policy, min : Int32, max : Int32, letter : Char
  record PasswordEntry, password : String, policy : Policy do
    def valid?
      count = password.each_char.count(policy.letter)
      policy.min <= count <= policy.max
    end

    def positions_valid?
      (password[policy.min - 1]? == policy.letter) != (password[policy.max - 1]? == policy.letter)
    end
  end

  class PasswordDb
    include Iterator(PasswordEntry)

    @lines : Iterator(String)

    def initialize(input)
      @lines = input.each
    end

    def next
      parse_next_line
    end

    private def parse_next_line
      l = @lines.next
      if l.is_a?(Iterator::Stop)
        return l
      end

      result = l.match(/(?<min>\d+)-(?<max>\d+)\s(?<letter>[a-z]):\s(?<password>[a-z0-0]+)/).not_nil!
      password_entry = PasswordEntry.new(
        password: result["password"],
        policy: Policy.new(
          min: result["min"].to_i,
          max: result["max"].to_i,
          letter: result["letter"][0],
        )
      )
    end
  end

  def self.day2p1
    puts PasswordDb.new(INPUT_DAY2).count { |p| p.valid? }
  end

  def self.day2p2
    puts PasswordDb.new(INPUT_DAY2).count { |p| p.positions_valid? }
  end
end

Aoc2020.day2p1
Aoc2020.day2p2
