require "../src/11"

describe Aoc2020 do
    it "works p1 1" do
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
        s = Aoc2020::SeatSystem.new(input)
        s.stabilize
        s.count_occupied.should eq(37)
    end
end