require "../src/01"

describe Elevator do
  describe "index" do
    it "works 1" do
      Elevator.new(")").index(-1).should eq(1)
    end

    it "works 2" do
      Elevator.new("()())").index(-1).should eq(5)
    end
  end
end
