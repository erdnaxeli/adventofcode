module Aoc2020
  enum Op
    Addition
    Multiplication
    Nop
  end

  class Expression
    @op : Op
    @a : Expression | UInt64
    @b : Expression | UInt64 | Nil

    def initialize(@op, @a, @b = nil)
    end

    def compute
      if a = @a.as?(Expression)
        x = a.compute
      else
        x = @a.as(UInt64)
      end

      if b = @b.as?(Expression)
        y = b.compute
      elsif b = @b.as?(UInt64)
        y = b
      else
        y = nil
      end

      case @op
      in .addition?
        x + y.not_nil!
      in .multiplication?
        x * y.not_nil!
      in .nop?
        x
      end
    end
  end

  def self.compute(input) : UInt64
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

  def self.read_expression(input)
    if input.size == 1
      return Expression.new(op: Op::Nop, a: input[0].to_u64)
    end

    parentheses = 0
    op = Op::Multiplication
    idx = find_index('*', input)
    if idx.nil?
      op = Op::Addition
      idx = find_index('+', input)
    end

    if !idx.nil?
      Expression.new(
        op: op,
        a: read_expression(fix_parentheses(0, idx - 1, input)),
        b: read_expression(fix_parentheses(idx + 1, input.size - 1, input)),
      )
    else
      raise "BUG: unreachable"
    end
  end

  def self.find_index(op, input)
    acc = 0
    min_parentheses = input[...-1].map do |c|
      if c == '('
        acc += 1
        acc
      elsif c == ')'
        acc -= 1
        acc
      else
        acc
      end
    end.min
    parentheses = 0
    index = nil
    input.each_with_index do |c, i|
      if c == '('
        parentheses += 1
      elsif c == ')'
        parentheses -= 1
      elsif c == op
        if parentheses == min_parentheses
          index = i
        end
      end
    end
    index
  end

  def self.fix_parentheses(x, y, input)
    tally = input[x..y].tally
    lp = tally.fetch('(', 0)
    rp = tally.fetch(')', 0)
    if lp > rp
      input[x + 1..y]
    elsif lp < rp
      input[x..y -1]
    else
      input[x..y]
    end
  end

  INPUT_DAY18 = File.read("./inputs/18.txt")

  def self.day18p1
    INPUT_DAY18.each_line.map { |l| compute(l) }.sum
  end

  def self.day18p2
    INPUT_DAY18.each_line.map { |l| read_expression(l.chars.reject!(' ')).compute.as(UInt64) }.sum
  end
end

puts Aoc2020.day18p2
