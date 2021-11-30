class Position
  attr_accessor :x, :y

  def initialize(x_coordinate, y_coordinate)
    @x = x_coordinate
    @y = y_coordinate
  end

  def to_s
    "#{x} | #{y}"
  end
end
