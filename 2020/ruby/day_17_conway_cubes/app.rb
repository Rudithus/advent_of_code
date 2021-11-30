load 'rules.rb'
load 'cube.rb'
load 'game.rb'

def input
  lines = File.readlines('input.txt').map(&:chomp)
  lines.map.with_index { |line, row| skip_inactive_cubes(line).map { |_, column| active_cube(row, column, 0) } }
end

def active_cube(row, column, depth)
  Cube.new(@space, row, column, depth)
end

def skip_inactive_cubes(line)
  line.chars.map.with_index.reject { |c, _| c == '.' }
end

@rules = [ActivateIfThreeNeighbours.new, DeactivateIfCrowded.new, DeactivateIfLonely.new]
@space = Array.new { Array.new { Array.new } }

input.map.with_index {}

@game = Game.new(@rules, @cubes)
