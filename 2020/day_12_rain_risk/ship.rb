load 'position.rb'

class Ship
  attr_reader :position

  def initialize(position, waypoint)
    @position = position
    @waypoint = waypoint
  end

  def execute(command)
    command.move(@waypoint, @position)
  end
end
