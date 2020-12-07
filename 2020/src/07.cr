module Aoc2020
  def self.read_bags_rules_reverse(input)
    rules = Hash(String, Array(String)).new { |h, k| h[k] = Array(String).new }
    input.each_line do |l|
      read_bag_rule(l) do |container, _, contained|
        rules[contained] << container
      end
    end

    rules
  end

  def self.read_bags_rules(input)
    rules = Hash(String, Array(Tuple(Int32, String))).new do |h, k|
      h[k] = Array(Tuple(Int32, String)).new
    end

    input.each_line do |l|
      read_bag_rule(l) do |container, count, contained|
        rules[container] << {count, contained}
      end
    end

    rules
  end

  def self.read_bag_rule(input)
    result = /^(?<container>[a-z\s]+) bags contain (?<contained>.*)$/.match(input).not_nil!
    container = result["container"]

    result["contained"].scan(/(?<count>\d+) (?<bag>[a-z\s]+) bags?/).each do |result|
      yield container, result["count"].to_i, result["bag"]
    end
  end

  # Counts distinct bags that can contain the given *bag*.
  def self.count_container_bags(rules, bag)
    bags = Set(String).new
    keys = [bag]
    while keys.size != 0
      key = keys.shift

      rules[key]?.try &.each do |k|
        keys << k
        bags << k
      end
    end

    bags.size
  end

  # Counts the total number of bags that *bag* must contains.
  def self.count_contained_bags(rules, bag)
    total_count = 0
    bags = [{1, bag}]
    while bags.size != 0
      count, bag = bags.shift
      # We count the current bag.
      total_count += count

      if rule = rules[bag]?
        rule.each do |contained_bag_count, contained_bag|
          # We add any contained bag. There count is the number of bags contained
          # in the current bag multiplied by the current bag count.
          bags << {count * contained_bag_count, contained_bag}
        end
      end
    end

    # We don't count the initial bag.
    total_count - 1
  end

  module Day7
    INPUT = File.read("./inputs/07.txt")

    def self.p1
      rules = Aoc2020.read_bags_rules_reverse(INPUT)
      Aoc2020.count_container_bags(rules, "shiny gold")
    end

    def self.p2
      rules = Aoc2020.read_bags_rules(INPUT)
      Aoc2020.count_contained_bags(rules, "shiny gold")
    end
  end
end

puts Aoc2020::Day7.p2
