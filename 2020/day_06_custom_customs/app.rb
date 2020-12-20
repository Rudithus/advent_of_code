require 'set'

def input
  File.open('input.txt').read.split(/\n\n/).map { |group| group.split(/\n/) }
end
total_count = 0

input.each do |group|
  group_items = Hash.new(0)

  group.each do |person|
    person.each_char { |c| items[c] += 1 }
  end

  total_count += group_items.select { |_k, v| v == group_count }.count
end

puts total_count
