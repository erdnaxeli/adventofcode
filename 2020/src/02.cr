require "./inputs/02"

module Aoc2020
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

    def initialize(input)
      @channel = Channel(PasswordEntry).new(0)
      spawn parse_password_db(input)
    end

    def next
      if v = @channel.receive?
        v
      else
        stop
      end
    end

    private def parse_password_db(input)
      input.each_line do |l|
        if result = l.match(/(?<min>\d+)-(?<max>\d+)\s(?<letter>[a-z]):\s(?<password>[a-z0-0]+)/)
          password_entry = PasswordEntry.new(
            password: result["password"],
            policy: Policy.new(
              min: result["min"].to_i,
              max: result["max"].to_i,
              letter: result["letter"][0],
            )
          )

          @channel.send(password_entry)
        end
      end

      @channel.close
    end
  end

  def self.day2p1
    count = 0
    puts PasswordDb.new(INPUT_DAY2).count { |p| p.valid? }
  end

  def self.day2p2
    count = 0
    puts PasswordDb.new(INPUT_DAY2).count { |p| p.positions_valid? }
  end
end

Aoc2020.day2p1
Aoc2020.day2p2
