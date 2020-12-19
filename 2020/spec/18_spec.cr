require "./spec_helper"

require "../src/18"

describe Aoc2020 do
  describe "#compute" do
    [
      {"1 + 2 * 3 + 4 * 5 + 6", 231},
      {"1 + (2 * 3) + (4 * (5 + 6))", 51},
      {"2 * 3 + (4 * 5", 46},
      {"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
      {"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
      {"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
      {"((((3 * 5 + 1) * (1 + 5 * 3)) + 1 * 1 ) + 9 * (5 + (2 * 5 + 1)) + 9 ) + 1", 8685},
      {"1 * 2 * 3 * 4 * 5 * 6 * 7", 5040},
      {"3 * 8 * 9 + 3 + 2", 336},
    ].each do |test, expected|
      it "works", focus: true do
        Aoc2020.read_tree(test.each_char.reject(' ')).compute.should eq(expected)
      end
    end
  end
end
