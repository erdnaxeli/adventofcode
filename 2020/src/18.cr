require "deque"

module Aoc2020
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

  enum Operator
    Add
    Mult
  end

  abstract class Node
    abstract def compute
  end

  class Operation < Node
    property left : Node
    property right : Node

    getter op : Operator

    def initialize(@op, @left, @right)
    end

    def compute
      case @op
      in .add?
        @left.compute + @right.compute
      in .mult?
        @left.compute * @right.compute
      end
    end
  end

  class Digit < Node
    def initialize(@value : UInt64)
    end

    def compute
      @value
    end
  end

  def self.read_tree(input)
    stack = Deque(Node).new

    token = input.next.as(Char)
    case token
    when .number?
      stack << Digit.new(token.to_u64)
    when '('
      stack << read_tree(input)
    else
      raise "Unexpected token #{token} at the beginning of the input"
    end

    first = true
    loop do
      pp! stack[0]
      token = input.next
      if token.is_a?(Iterator::Stop)
        break
      end

      case token
      when '+'
        current_node = stack.pop
        if current_node.is_a?(Operation)
          if first || current_node.op.add?
            node = new_operation(Operator::Add, current_node, input)

            if stack.size > 0
              previous_node = stack.pop.as(Operation)
              previous_node.right = node
              stack << previous_node
            end

            stack << node
          else
            current_node.right = new_operation(Operator::Add, current_node.right, input)
            stack << current_node << current_node.right
          end
        else
          stack << new_operation(Operator::Add, current_node, input)
        end
      when '*'
        current_node = stack.pop
        if current_node.is_a?(Operation) && current_node.op.mult?
          current_node.right = new_operation(Operator::Mult, current_node.right, input)
          stack << current_node << current_node.right
        else
          node = new_operation(Operator::Mult, current_node, input)

          if stack.size > 0
            previous_node = stack.pop.as(Operation)
            previous_node.right = node
            stack << previous_node
          end

          stack << node
        end
      when ')'
        return stack[0]
      else
        raise "Unexpected token #{token}, expecting an operator or right parenthesis"
      end

      first = false
    end

    pp stack[0]
    stack[0]
  end

  def self.new_operation(op, left, input)
    right = case token = input.next.as(Char)
            when .number?
              Digit.new(token.to_u64)
            when '('
              read_tree(input)
            else
              raise "Unexpected token #{token}, expecting a number or a left parenthesis"
            end
    Operation.new(op, left, right)
  end

  INPUT_DAY18 = File.read("./inputs/18.txt")

  def self.day18p1
    INPUT_DAY18.each_line.map { |l| compute(l.chars.reject(' ')) }.sum
  end

  def self.day18p2
    INPUT_DAY18.each_line.map do |line|
      read_tree(line.each_char.reject(' ')).compute.tap { |x| puts x }
    end.sum
  end
end

puts Aoc2020.day18p2
