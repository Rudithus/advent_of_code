class Waypoint
  attr_accessor :x, :y

  def initialize(x, y)
    @x = x
    @y = y
  end

  def to_s
    "#{x} | #{y}"
  end
end
