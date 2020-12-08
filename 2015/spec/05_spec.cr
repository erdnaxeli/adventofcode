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

  it "nice2 string 1" do
    NiceString.nice2?("qjhvhtzxzqqjkmpb").should eq(true)
  end

  it "nice2 string 2" do
    NiceString.nice2?("xxyxx").should eq(true)
  end

  it "not nice2 string 1" do
    NiceString.nice2?("uurcxstgmygtbstg").should eq(false)
  end

  it "not nice2 string 2" do
    NiceString.nice2?("ieodomkazucvgmuy").should eq(false)
  end

  it "not nice2 string 3" do
    NiceString.nice2?("xaxxx").should eq(false)
  end
end
