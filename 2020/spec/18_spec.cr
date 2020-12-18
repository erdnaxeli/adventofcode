require "./spec_helper"

require "../src/18"

describe Aoc2020 do
  describe "#compute" do
    [
      # {"1 + 2 * 3 + 4 * 5 + 6", 71},
      {"1 + (2 * 3) + (4 * (5 + 6))", 51},
      {"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
    ].each do |test, expected|
      it "works" do
        Aoc2020.compute(test).should eq(expected)
      end
    end
  end
end
