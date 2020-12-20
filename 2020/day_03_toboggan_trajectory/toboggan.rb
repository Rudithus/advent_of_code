class Toboggan
  attr_writer :tree_count
  attr_reader :pos_x, :pos_y

  def initialize(angle_x, angle_y)
    @angle_x = angle_x
    @angle_y = angle_y
    @pos_x = 0
    @pos_y = 0
  end

  def slide
    @pos_x += @angle_x
    @pos_y += @angle_y
  end
end
