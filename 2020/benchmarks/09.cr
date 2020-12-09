require "benchmark"

require "../src/09.cr"

Benchmark.ips do |bench|
  bench.report("XMAS") do
    input = File.read("./inputs/09.txt").each_line.map &.to_i64
    Aoc2020::XMAS(Int64).find_invalid(input)
  end
  bench.report("oneliner") do
    input = File.read("./inputs/09.txt").each_line.map &.to_i64
    input.cons(26).find { |x| !x[0...25].combinations(2).map(&.sum).includes? x[25] }.try &.last
  end
end

puts

Benchmark.ips do |x|
  a = Array(Int32).new(25, 0)
  d = Deque(Int32).new(a)
  s = Set(Int32).new(a)

  x.report("array") do
    a.includes? 1
    a.shift
    a << 0
  end
  x.report("deque") do
    d.includes? 1
    d.shift
    d << 0
  end
  x.report("set") do
    s.includes? 1
    s.delete(s.first)
    s << 0
  end
end
