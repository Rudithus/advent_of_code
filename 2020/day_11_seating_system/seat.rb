load 'position.rb'

class Seat
  attr_reader :position, :neighbours

  def initialize(position)
    @position = position
    @occupied = false
    @neighbours = []
  end

  def occupied?
    @occupied
  end

  def occupy
    @occupied = true
  end

  def vacate
    @occupied = false
  end

  def add_neighbour(seat)
    add_adjacent(seat)
    seat.add_adjacent(self)
  end

  def to_s
    occupied? ? '#' : 'L'
  end

  protected

  def add_adjacent(seat)
    @neighbours.push(seat)
  end
end