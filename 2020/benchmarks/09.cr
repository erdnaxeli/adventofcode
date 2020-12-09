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
  bench.report("better oneliner") do
    input = File.read("./inputs/09.txt").each_line.map &.to_i64
    input.cons(26).find { |x| x.each { |y| x[25] - y != y && x.includes? x[25] - y } }.try &.last
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

# XMAS   8.91k (112.24µs) (± 7.80%)  26.7kB/op        fastest
# oneliner 105.69  (  9.46ms) (± 4.51%)  15.9MB/op  84.30× slower
# better oneliner   4.41k (226.94µs) (± 7.21%)   496kB/op   2.02× slower
#
# array  33.86M ( 29.54ns) (±16.63%)  0.0B/op        fastest
# deque  28.21M ( 35.44ns) (±14.98%)  0.0B/op   1.20× slower
# set   9.21M (108.61ns) (±14.45%)  0.0B/op   3.68× slower
