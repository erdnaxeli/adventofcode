require "./spec_helper"

require "../src/14.cr"

describe DockingComputer do
  [
    {"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 11, 73},
    {"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 101, 101},
    {"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 0, 64},
  ].each do |mask, value, expected|
    it "apply mask" do  
      DockingComputer.apply_mask(value, mask).should eq(expected)
    end
  end
end
