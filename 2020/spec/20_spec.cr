require "./spec_helper"

require "../src/20"

input = "Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...
"

describe Aoc2020 do
  describe ".read_tiles" do
    it "works with input" do
      Aoc2020.read_tiles(input).size.should eq(9)
    end

    it "works with one tile" do
      tile = Aoc2020.read_tiles("Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...
")[0]

      tile.id.should eq(3079)
      tile.up.should eq("#.#.#####.".chars)
      tile.down.should eq("..#.###...".chars)
      tile.left.should eq("#..##.#...".chars)
      tile.right.should eq(".#....#...".chars)
    end
  end

  describe Aoc2020::Tile do
    describe "#each_position" do
      it "works" do
        tile = Aoc2020.read_tiles("Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...
")[0]
        result = Array(Aoc2020::Tile).new
        tile.each_position { |t| result << t }

        result.size.should eq(12)
      end
    end
  end

  describe Aoc2020::ImageTiles do
    describe "#four_corners" do
      it "works", focus: true do
        image = Aoc2020::ImageTiles.new(Aoc2020.read_tiles(input))
        image.resolve
        image.four_corners.product.should eq(20899048083289)
      end
    end
  end

  describe "p1" do
    it "works" do
      image = Aoc2020::ImageTiles.new(Aoc2020.read_tiles(input))
      image.resolve
      image.map.size.should eq(9)
    end
  end
end
