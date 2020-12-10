require "../src/10"

describe Aoc2020 do
  it "works p1" do
    input = %(28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3).each_line.map(&.to_i).to_a.sort!

    Aoc2020.count_jolt_diff(input).should eq(220)
  end

  it "works p2 1" do
    input = %w(16
      10
      15
      5
      1
      11
      7
      19
      6
      12
      4).map(&.to_i).sort!

    Aoc2020.count_valid_combinations2(input).should eq(8)
  end

  it "works p2 2" do
    input = %(28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3).each_line.map(&.to_i).to_a.sort!

    Aoc2020.count_valid_combinations2(input).should eq(19208)
  end
end
