module Aoc2020
  def self.count_jolt_diff(input)
    counter = [0].each.chain(input.each).cons(2, reuse: true).map do |x|
      x[1] - x[0]
    end.tally
    counter.select { |x| {1, 3}.includes? x }
    counter[1] * (counter[3] + 1)
  end

  def self.spinner
    loop do
      print "\b\\"
      sleep 0.5
      print "\b|"
      sleep 0.5
      print "\b/"
      sleep 0.5
      print "\b-"
      sleep 0.5
    end
  end

  def self.count_valid_combinations(input)
    spawn spinner
    chargers = (0..input.max + 3).map { |x| input.includes? x }
    chargers[0] = true  # outlet
    chargers[-1] = true # device
    get_valid_combinations(0, 1, chargers) +
      get_valid_combinations(0, 2, chargers) +
      get_valid_combinations(0, 3, chargers)
  end

  private def self.get_valid_combinations(from, gap, chargers)
    if from + gap == chargers.size
      Fiber.yield
      1_u64
    elsif chargers[from + gap]?
      get_valid_combinations(from + gap, 1, chargers) +
        get_valid_combinations(from + gap, 2, chargers) +
        get_valid_combinations(from + gap, 3, chargers)
    else
      0_u64
    end
  end

  def self.count_valid_combinations2(input)
    multiplicators = Hash(Int32, UInt64).new { |h, k| h[k] = 0 }
    multiplicators[0] = 1

    [0].each.chain(input.each).each do |x|
      multiplicators[x + 1] += multiplicators[x]
      multiplicators[x + 2] += multiplicators[x]
      multiplicators[x + 3] += multiplicators[x]
    end

    multiplicators[input.max + 3]
  end

  INPUT_DAY10 = File.read("./inputs/10.txt").each_line.map(&.to_i).to_a.sort!

  def self.day10p1
    count_jolt_diff(INPUT_DAY10)
  end

  def self.day10p2
    self.count_valid_combinations2(INPUT_DAY10)
  end
end
