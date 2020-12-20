load 'commands.rb'
load 'ship.rb'
load 'waypoint.rb'

@command_dictionary = { 
  'F' => Forward,
  'N' => North, 
  'S' => South,
  'E' => East,
  'W' => West,
  'R' => Right,
  'L' => Left
}

def input
  File.readlines('input.txt').map(&:chomp)
end

@commands = input.map { |line| p line; @command_dictionary[line[0]].new(line[1..].to_i) }

@initial_waypoint = Waypoint.new(10,1)
@ship = Ship.new(Position.new(0,0),@initial_waypoint)

@commands.each { |c| @ship.execute(c) }
puts @ship.position.x.abs + @ship.position.y.abs