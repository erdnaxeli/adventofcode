require "benchmark"

require "../src/03"

INPUT = File.read("./inputs/03.txt")

Benchmark.ips do |x|
  x.report("distinct_houses") { Santa.new(INPUT).distinct_houses }
  x.report("distinct_houses_iterator") { Santa.new(INPUT).distinct_houses_iterator }
end

puts

Benchmark.ips do |x|
  x.report("distinct_houses_with_robot") { Santa.new(INPUT).distinct_houses_with_robot }
  x.report("distinct_houses_with_robot_2") { Santa.new(INPUT).distinct_houses_with_robot_2 }
end
