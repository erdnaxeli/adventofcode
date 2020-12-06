struct Elevator
  def initialize(@input : String)
  end

  def floor
    chars = @input.each_char.tally
    (chars['(']? || 0) - (chars[')']? || 0)
  end

  def index(floor)
    current_floor = 0
    @input.each_char_with_index(offset: 1) do |c, i|
      case c
      when '('
        current_floor += 1
      when ')'
        current_floor -= 1
      end

      if current_floor == floor
        return i
      end
    end
  end
end
