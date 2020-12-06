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
#
# input = %(ahynbmqljzpwxokcfrtsgeud
# xwzcmdhkrjnupegqlyoaft
# fjnurhzoqmgwacxdlypkte
# qwrjxahtlnzcfdouepmkgy
# ezpqxfcmgrnhylukwajotd)
#
# Benchmark.ips do |x|
#   x.report("tally") { input.each_char.select('a'..'z').tally.size }
#   x.report("set") { Set(Char).new(input.each_char.select('a'..'z')).size }
#   x.report("set array") { Set(Char).new(input.chars.select('a'..'z')).size }
#   x.report("uniq") { input.each_char.select('a'..'z').uniq.size }
#   x.report("uniq array") { input.chars.select('a'..'z').uniq.size }
# end
#
# puts
#
# Benchmark.ips do |x|
#   x.report("new") do
#     people_count = input.lines.size
#     input.each_char.select('a'..'z').tally.select { |k, v| v == people_count }.size
#   end
#   x.report("inplace") do
#     people_count = input.lines.size
#     input.each_char.select('a'..'z').tally.tap(&.select! { |k, v| v == people_count }).size
#   end
#   x.report("reduce") { input.each_line.map(&.chars).reduce { |a, b| a & b }.size }
# end

# tally 231.25k (  4.32µs) (±10.15%)  1.04kB/op   1.52× slower
# set 351.81k (  2.84µs) (±10.67%)    784B/op        fastest
# set array 303.15k (  3.30µs) (± 5.71%)  3.85kB/op   1.16× slower
# uniq 305.19k (  3.28µs) (±10.34%)  1.04kB/op   1.15× slower
# uniq array 287.75k (  3.48µs) (± 8.10%)   3.3kB/op   1.22× slower
#
# new 171.70k (  5.82µs) (±10.37%)  2.36kB/op   1.21× slower
# inplace 208.55k (  4.80µs) (± 9.08%)  1.41kB/op        fastest
# reduce 143.02k (  6.99µs) (± 9.04%)  5.23kB/op   1.46× slower
