require "../src/05"

describe NiceString do
  it "nice string 1" do
    NiceString.nice?("ugknbfddgicrmopn").should eq(true)
  end

  it "nice string 2" do
    NiceString.nice?("aaa").should eq(true)
  end

  it "not nice string 1" do
    NiceString.nice?("jchzalrnumimnmhp").should eq(false)
  end

  it "not nice string 2" do
    NiceString.nice?("haegwjzuvuyypxyu").should eq(false)
  end

  it "not nice string 3" do
    NiceString.nice?("dvszwmarrgswjxmb").should eq(false)
  end

  it "not nice string 4" do
    NiceString.nice?("aabbeiou")
  end
end
