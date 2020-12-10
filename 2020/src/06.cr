module Aoc2020
  def self.count_anyone_questions(input)
    Set.new(input.each_char.select('a'..'z')).size
  end

  def self.count_everyone_questions(input)
    people_count = input.lines.size
    input.each_char.select('a'..'z').tally.count { |_, v| v == people_count }
  end

  def self.day6p1
    puts File.read("./inputs/06.txt").split("\n\n").sum { |group| count_anyone_questions(group) }
  end

  def self.day6p2
    puts File.read("./inputs/06.txt").split("\n\n").sum { |group| count_everyone_questions(group) }
  end
end
