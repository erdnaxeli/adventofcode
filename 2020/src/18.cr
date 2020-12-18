module Aoc2020
  INPUT_DAY18 = File.read("./inputs/18.txt")

  def self.compute(input)
    input.each_char.reject(' ').reduce(Array(UInt64 | Char | Nil).new) do |acc, x|
      case x.to_s
      when /\d/
        compute_digit(acc, x.to_u64)
      when "*"
        acc << x
      when "+"
        acc << x
      when "("
        acc << x
      when ")"
        x = acc.pop.as(UInt64)
        acc.pop
        compute_digit(acc, x)
      else
        acc
      end
    end.first.as(UInt64)
  end

  def self.compute_digit(acc, x)
    case acc.pop?
    when '*'
      acc << acc.pop.as(UInt64) * x
    when '+'
      acc << acc.pop.as(UInt64) + x
    when '('
      acc << '(' << x
    when nil
      acc << x
    else
      raise "BUG: invalid expression"
    end

    acc
  end

  def self.day18p1
    INPUT_DAY18.each_line.map { |l| compute(l) }.sum
  end
end

puts Aoc2020.day18p1
