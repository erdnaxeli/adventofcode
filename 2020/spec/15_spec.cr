require "./spec_helper"

require "../src/15"

describe Aoc2020 do
  describe "#day15p1" do
    [
      {[0, 3, 6], 436},
      {[1, 3, 2], 1},
      {[2, 1, 3], 10},
    ].each do |x, y|
      it "works" do
        g = MemoryGame.new(x)
        2020.times { g.play }
        g.last_value.should eq(y)
      end
    end
  end
end
