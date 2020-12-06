module Aoc2020
  def self.count_anyone_questions(input)
    Set.new(input.each_char.select('a'..'z')).size
  end

  def self.count_everyone_questions(input)
    people_count = input.lines.size
    input.each_char.select('a'..'z').tally.select { |k, v| v == people_count }.size
  end

  def self.day6p1
    puts File.read("./inputs/06.txt").split("\n\n").sum { |group| count_anyone_questions(group) }
  end

  def self.day6p2
    puts File.read("./inputs/06.txt").split("\n\n").sum { |group| count_everyone_questions(group) }
  end
end

Aoc2020.day6p1
Aoc2020.day6p2

# require "benchmark"

# input = %(ahynbmqljzpwxokcfrtsgeud
# xwzcmdhkrjnupegqlyoaft
# fjnurhzoqmgwacxdlypkte
# qwrjxahtlnzcfdouepmkgy
# ezpqxfcmgrnhylukwajotd)

# Benchmark.ips do |x|
#     x.report("tally") { input.each_char.select('a'..'z').tally.size }
#     x.report("set") { Set(Char).new(input.each_char.select('a'..'z')).size }
# end

# people_count = input.lines.size
# Benchmark.ips do |x|
#     x.report("new") { input.each_char.select('a'..'z').tally.select { |k, v| v == people_count }.size }
#     x.report("inplace") { input.each_char.select('a'..'z').tally.tap(&.select! { |k, v| v == people_count }).size }
# end
#
#
# tally 242.41k (  4.13µs) (± 7.94%)  1.04kB/op   1.62× slower
#   set 392.87k (  2.55µs) (± 8.42%)    784B/op        fastest
#     new 191.95k (  5.21µs) (± 7.61%)   2.0kB/op   1.19× slower
# inplace 228.74k (  4.37µs) (± 9.27%)  1.04kB/op        fastest
