load 'repeating_array.rb'
load 'toboggan.rb'

def import_input
  File.open('input.txt').readlines.map { |l| RepeatingArray.new(l.chomp.chars) }
end

map = import_input

toboggans = [
  Toboggan.new(1, 1),
  Toboggan.new(1, 3),
  Toboggan.new(1, 5),
  Toboggan.new(1, 7),
  Toboggan.new(2, 1)
]
tree_counts = []
toboggans.each do |t|
  tree_count = 0
  while t.pos_x < map.count
    tree_count += 1 if map[t.pos_x][t.pos_y] == '#'
    t.slide
  end
  tree_counts.push(tree_count)
end

puts tree_counts.reduce(&:*)
