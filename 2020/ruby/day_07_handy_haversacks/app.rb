load 'bag.rb'
require 'set'
@bags = {}

def input
  File.open('input.txt').readlines.map(&:chomp)
end

def bag(name)
  if @bags.key?(name)
    bag = @bags[name]
  else
    bag = Bag.new(name)
    @bags[name] = bag
  end
  bag
end

def process_lines(line)
  matches = line.scan(/(\w+ \w+) bags contain|(\d) (\w+ \w+)/)

  parent_bag = bag(matches[0][0])
  matches[1..matches.count].map { |c| [bag(c[2]), c[1].to_i] }.each do |child_bag, count|
    child_bag.add_parent(parent_bag)
    count.times do
      parent_bag.add_child(child_bag)
    end
  end
end

def generate_graph
  input.each do |line|
    process_lines(line)
  end
end

def all_parents(node, parent_set = Set.new)
  node.parents.each do |p|
    parent_set.add(p)
    all_parents(p, parent_set)
  end

  parent_set
end

def count_children(node)
  count = 1
  node.children.each do |c|
    count += count_children(c)
  end

  count
end

generate_graph
shiny_gold = @bags['shiny gold']

puts all_parents(shiny_gold).count
puts count_children(shiny_gold) - 1
