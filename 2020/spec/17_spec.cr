require "./spec_helper"

require "../src/17"

describe ConwaySpace do
  describe "step 0" do
    # it "has a correct map" do
    #   space.map = [
    #     [
    #       ['.'],
    #       []
    #     ]
    #   ]
    # end

    it "has 5 cubes" do
      space = ConwaySpace.new(".#.\n..#\n###")
      space.count_cubes.should eq(5)
    end
  end

  describe "step 1", focus: true do
    it "has 11 cubes" do
      space = ConwaySpace.new(".#.\n..#\n###")
      space.next
      space.count_cubes.should eq(11)
    end
  end
end
