require "./spec_helper"

require "../src/13"

describe Aoc2020 do
  it "p2 1" do
    input = "7,13,x,x,59,x,31,19"
    planning = Aoc2020.read_bus_planning_with_gap(input)
    Aoc2020.find_t_for_gap(planning).should eq(1068781)
  end

  it "p2 2" do
    input = "17,x,13,19"
    planning = Aoc2020.read_bus_planning_with_gap(input)
    Aoc2020.find_t_for_gap(planning).should eq(3417)
  end
end
