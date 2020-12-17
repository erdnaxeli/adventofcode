require "./spec_helper"

require "../src/17"

describe ConwaySpace do
  describe "step 0" do
    it "has 5 cubes" do
      space = ConwaySpace::Resolver(Tuple(Int32, Int32, Int32), ConwaySpace::Cube3D).new(".#.\n..#\n###")
      space.cubes.size.should eq(5)
    end
  end

  describe "step 1" do
    it "has 11 cubes" do
      space = ConwaySpace::Resolver(Tuple(Int32, Int32, Int32), ConwaySpace::Cube3D).new(".#.\n..#\n###")
      space.next
      space.cubes.size.should eq(11)
    end
  end
end

describe ConwaySpace do
  describe "step 0" do
    it "has 5 cubes" do
      space = ConwaySpace::Resolver(Tuple(Int32, Int32, Int32, Int32), ConwaySpace::Cube4D).new(".#.\n..#\n###")
      space.cubes.size.should eq(5)
    end
  end

  describe "step 1" do
    it "has 11 cubes" do
      space = ConwaySpace::Resolver(Tuple(Int32, Int32, Int32, Int32), ConwaySpace::Cube4D).new(".#.\n..#\n###")
      space.next
      space.cubes.size.should eq(29)
    end
  end
end
