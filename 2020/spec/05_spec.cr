require "../src/05"

describe Aoc2020 do
  it "computes correct row" do
    pass = Aoc2020::BoardingPass.new("FBFBBFFRLR")
    pass.seat_row.should eq(44)
  end

  it "computes correct column" do
    pass = Aoc2020::BoardingPass.new("FBFBBFFRLR")
    pass.seat_column.should eq(5)
  end
end
