require "../src/11"

describe Aoc2020 do
  it "works p1" do
    input = %(L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL)
    s = Aoc2020::SeatSystem.new(input, 4)
    s.stabilize
    s.count_occupied.should eq(37)
  end

  it "works p2" do
    input = %(L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL)
    s = Aoc2020::SeatSystem.new(input, 5)
    s.stabilize2
    s.count_occupied.should eq(26)
  end
end
