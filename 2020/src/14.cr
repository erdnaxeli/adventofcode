require "log"

Log.setup_from_env

class DockingComputer
  enum Operator
    Mask
    Mem
  end

  enum Version
    V1
    V2
  end

  record Instruction, operator : Operator, arg1 : UInt64 | String, arg2 : Int32? = nil do
    def self.from_input(input : String)
      if match = /mask = (?<arg1>.*)/.match(input)
        self.new(operator: Operator::Mask, arg1: match["arg1"])
      elsif match = /mem\[(?<arg1>\d+)\] = (?<arg2>\d+)/.match(input)
        self.new(operator: Operator::Mem, arg1: match["arg1"].to_u64, arg2: match["arg2"].to_i)
      else
        raise "Unknown instruction"
      end
    end
  end

  @mask : String = ""
  @program : Iterator(Instruction)

  getter mem = Hash(UInt64, UInt64).new { 0_u64 }

  def initialize(input, @version = Version::V1)
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
      case @version
      in Version::V1
        @mem[instruction.arg1.as(UInt64)] = DockingComputer.apply_mask(instruction.arg2.as(Int32), @mask)
      in Version::V2
        store_with_mask(instruction.arg1.as(UInt64), instruction.arg2.as(Int32))
      end
    end
  end

  def self.apply_mask(value : Int32, mask : String) : UInt64
    value = value.to_s(2).reverse
    String.build do |str|
      mask.reverse.each_char_with_index do |c, i|
        case c
        when 'X'
          str << (value[i]? || '0')
        else
          str << c
        end
      end
    end.reverse.to_u64(2)
  end

  private def apply_mask_v2(value, mask) : String
    value = value.to_s(2).reverse
    String.build do |str|
      mask.reverse.each_char_with_index do |c, i|
        case c
        when '0'
          str << (value[i]? || '0')
        else
          str << c
        end
      end
    end.reverse
  end

  private def store_with_mask(pointer, value)
    each_mask_value(apply_mask_v2(pointer, @mask)) do |p|
      @mem[p] = value.to_u64
    end
  end

  private def each_mask_value(mask)
    each_mask_value_s(mask).each do |v|
      yield v.to_u64(2)
    end
  end

  private def each_mask_value_s(mask) : Array(String)
    before, separator, after = mask.partition('X')
    if after.size == 0
      if separator.size > 0
        [before + "0", before + "1"]
      else
        [before]
      end
    else
      each_mask_value_s(after).flat_map do |v|
        [before + "0" + v, before + "1" + v]
      end
    end
  end
end

INPUT_DAY14 = File.read("./inputs/14.txt")

module Aoc2020
  def self.day14p1
    computer = DockingComputer.new(INPUT_DAY14)
    computer.run
    computer.mem.values.sum
  end

  def self.day14p2
    computer = DockingComputer.new(INPUT_DAY14, version: DockingComputer::Version::V2)
    computer.run
    computer.mem.values.sum
  end
end

puts Aoc2020.day14p2
