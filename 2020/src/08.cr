class Computer
  enum InstructionType
    Acc
    Jmp
    Nop
  end

  record Instruction, type : InstructionType, arg : Int32 do
    def self.from_code(code : String) : Instruction
      result = /\A(?<inst>[a-z]{3}) (?<arg>(?:\+|-)\d+)\z/.match(code).not_nil!
      Instruction.new(
        type: InstructionType.parse(result["inst"]),
        arg: result["arg"].to_i,
      )
    end
  end

  @idx = 0

  getter acc = 0
  getter code : Array(Instruction)

  def initialize(input : String)
    @code = read_code(input)
  end

  def initialize(@code : Array(Instruction))
    @idx = 0
  end

  def read_code(input : String) : Array(Instruction)
    input.each_line.map { |c| Instruction.from_code(c) }.to_a
  end

  # Runs the program and returns the last instruction's index.
  def run(can_loop = true)
    seen = Set{@idx}

    while instruction = @code[@idx]?
      case instruction.type
      in .acc?
        @acc += instruction.arg
        @idx += 1
      in .jmp?
        @idx += instruction.arg
      in .nop?
        @idx += 1
      end

      if !can_loop && !seen.add?(@idx)
        break
      end
    end
  end

  def success?
    @idx < 0 || @idx >= @code.size
  end

  def self.fix_code(input)
    code = Computer.new(input).code
    jmp_or_nop = code.each.with_index.select do |x|
      x[0].type.jmp? || x[0].type.nop?
    end
    jmp_or_nop.each do |x, i|
      new_type = case x.type
                 when .jmp?
                   InstructionType::Nop
                 when .nop?
                   InstructionType::Jmp
                 else
                   raise "BUG!: unreachable"
                 end
      new_code = code.dup
      new_code[i] = Instruction.new(
        type: new_type,
        arg: code[i].arg,
      )
      c = Computer.new(new_code)
      c.run(can_loop: false)
      if c.success?
        return c.acc
      end
    end
  end
end

module Aoc2020
  INPUT_D08 = File.read("./inputs/08.txt")

  def self.day8p1
    c = Computer.new(INPUT_D08)
    c.run(can_loop: false)
    c.acc
  end

  def self.day8p2
    Computer.fix_code(INPUT_D08)
  end
end
