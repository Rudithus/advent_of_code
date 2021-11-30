class Position
  attr_reader :x, :y

  def initialize(x_coordinate, y_coordinate)
    @x = x_coordinate
    @y = y_coordinate
  end

  def to_s
    "x: #{x}, y: #{y}"
  end
end
