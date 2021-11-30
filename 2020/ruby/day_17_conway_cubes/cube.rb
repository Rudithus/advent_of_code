class Cube
  attr_reader :neighbours

  def initialize(space, x, y, z)
    @neighbours = []
    @space = space
    @active = true
    @x = x
    @y = y
    @z = z
  end

  def active_neighbours_count
    @neighbours.count
  end

  def activate; end

  def deactivate; end
end
