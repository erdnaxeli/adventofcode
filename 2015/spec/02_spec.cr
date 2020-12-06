require "../src/02"

describe Present do
  it "computes a box paper need" do
    Present.new("2x3x4").paper.should eq(58)
  end

  it "computes a box ribbon need" do
    Present.new("4x3x2").ribbon.should eq(34)
  end
end
