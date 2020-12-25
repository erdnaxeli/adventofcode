require "./spec_helper"

require "../src/25"

describe Aoc2020 do
  describe ".compute_key" do
    it "works" do
      key = 1
      11.times { key = Aoc2020.compute_key(7, key) }
      key.should eq(17807724)
    end
  end

  describe ".find_loop_size" do
    it "works" do
      Aoc2020.find_loop_size(17807724).should eq(11)
    end
  end
end
