require "log"

Log.setup_from_env

class DockingComputer
  enum Operator
    Mask
    Mem
  end

  record Instruction, operator : Operator, arg1 : Int32 | String, arg2 : Int32? = nil do
    def self.from_input(input : String)
      if match = /mask = (?<arg1>.*)/.match(input)
        self.new(operator: Operator::Mask, arg1: match["arg1"])
      elsif match = /mem\[(?<arg1>\d+)\] = (?<arg2>\d+)/.match(input)
        self.new(operator: Operator::Mem, arg1: match["arg1"].to_i, arg2: match["arg2"].to_i)
      else
        raise "Unknown instruction"
      end
    end
  end

  @mask : String = ""
  @program : Iterator(Instruction)

  getter mem = Hash(Int32, UInt64).new { 0_u64 }

  def initialize(input)
    @program = input.each_line.map { |l| Instruction.from_input(l) }
  end

  def run
    @program.each do |instruction|
      exec(instruction)
    end
  end

  private def exec(instruction)
    case instruction.operator
    in .mask?
      @mask = instruction.arg1.as(String)
      Log.debug { "New mask #{@mask}" }
    in .mem?
      @mem[instruction.arg1.as(Int32)] = DockingComputer.apply_mask(instruction.arg2.as(Int32), @mask)
    end
  end

  def self.apply_mask(value : Int32, mask : String)
    Log.debug { "applying mask #{mask} to value #{value}"}
    value = value.to_s(2).reverse
    Log.debug { "value in binary reversed is #{value} "}
    String.build do |str|
      mask.reverse.each_char_with_index do |c, i|
        Log.debug { "applying #{c} to #{value[i]?} "}
        case c
        when 'X'
          Log.debug { "-> #{value[i]? || '0'}" }
          str << (value[i]? || '0')
        else
          Log.debug { "-> #{c}" }
          str << c
        end
      end
    end.reverse.tap { |x| Log.debug { "result is #{x}" } }.to_u64(2)
  end
end

INPUT_DAY14 = File.read("./inputs/14.txt")

module Aoc2020
  def self.day14p1
    computer = DockingComputer.new(INPUT_DAY14)
    computer.run
    computer.mem.values.sum
  end
end

puts Aoc2020.day14p1
