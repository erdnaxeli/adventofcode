module Aoc2020
  def self.read_map(input = "./inputs/03.txt")
    File.read_lines(input)
  end

  def self.count_trees(map, slop_x, slop_y)
    x_steps = (0...(map.size)).step(slop_x)
    y_steps = 0.step(by: slop_y)
    x_steps.zip(y_steps).count do |x, y|
      map[x][y % map[0].size] == '#'
    end
  end

  def self.day3p1
    map = read_map
    puts count_trees(map, 1, 3)
  end

  def self.day3p2
    map = read_map
    puts [{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}].product { |x, y| count_trees(map, x, y).to_i64 }
  end
end

Aoc2020.day3p1
Aoc2020.day3p2
