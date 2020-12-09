require "benchmark"

input = %(ahynbmqljzpwxokcfrtsgeud
xwzcmdhkrjnupegqlyoaft
fjnurhzoqmgwacxdlypkte
qwrjxahtlnzcfdouepmkgy
ezpqxfcmgrnhylukwajotd)

Benchmark.ips do |x|
  x.report("tally") { input.each_char.select('a'..'z').tally.size }
  x.report("set") { Set(Char).new(input.each_char.select('a'..'z')).size }
  x.report("set array") { Set(Char).new(input.chars.select('a'..'z')).size }
  x.report("uniq") { input.each_char.select('a'..'z').uniq.size }
  x.report("uniq array") { input.chars.select('a'..'z').uniq.size }
end

puts

Benchmark.ips do |x|
  x.report("tally, new hash") do
    people_count = input.lines.size
    input.each_char.select('a'..'z').tally.count { |_, v| v == people_count }
  end
  x.report("tally, inplace") do
    people_count = input.lines.size
    input.each_char.select('a'..'z').tally.tap(&.select! { |_, v| v == people_count }).size
  end
  x.report("reduce") { input.each_line.map(&.chars).reduce { |a, b| a & b }.size }
end

# tally 232.48k (  4.30µs) (± 7.87%)  1.04kB/op   1.52× slower
# set 352.37k (  2.84µs) (± 9.69%)    784B/op        fastest
# set array 299.81k (  3.34µs) (± 6.63%)  3.85kB/op   1.18× slower
# uniq 319.09k (  3.13µs) (± 9.17%)  1.04kB/op   1.10× slower
# uniq array 279.03k (  3.58µs) (± 7.67%)   3.3kB/op   1.26× slower

# tally, new hash 169.54k (  5.90µs) (± 8.90%)  2.36kB/op   1.21× slower
# tally, inplace 205.60k (  4.86µs) (± 9.74%)  1.41kB/op        fastest
#   reduce 139.66k (  7.16µs) (± 9.51%)  5.23kB/op   1.47× slower
