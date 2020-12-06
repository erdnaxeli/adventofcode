struct Present
  @x : Int32
  @y : Int32
  @z : Int32

  def initialize(input : String)
    @x, @y, @z = input.split('x').map(&.to_i).sort!
  end

  def paper
    3*@x*@y + 2*@x*@z + 2*@y*@z
  end

  def ribbon
    2*@x + 2*@y + @x*@y*@z
  end
end
