require "spec"

require "../src/03"

describe Aoc2020 do
  it "works with subject example" do
    map = %w(..##.......
    #...#...#..
    .#....#..#.
    ..#.#...#.#
    .#...##..#.
    ..#.##.....
    .#.#.#....#
    .#........#
    #.##...#...
    #...##....#
    .#..#...#.#)

    Aoc2020.count_trees(map, 1, 1).should eq(2)
    Aoc2020.count_trees(map, 1, 3).should eq(7)
    Aoc2020.count_trees(map, 1, 5).should eq(3)
    Aoc2020.count_trees(map, 1, 7).should eq(4)
    Aoc2020.count_trees(map, 2, 1).should eq(2)
  end
end
