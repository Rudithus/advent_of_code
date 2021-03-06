load 'node.rb'

def input
  File.open('input.txt').readlines.map(&:to_i).sort
end
jolts = input.to_a
adapters = jolts.unshift(0).push(jolts.last + 3).map { |jolt| Node.new(jolt) }

puts adapters.each_cons(2).with_object(Hash.new(0)) { |j, counts| counts[j[1].value - j[0].value] += 1 }

def create_tree(adapters)
  adapters.each_with_index do |adapter, i|
    adapters
      .drop(i + 1)
      .take_while { |connection| adapter.can_connect?(connection) }
      .each { |connection| adapter.add_connection(connection) }
  end
end

create_tree(adapters)
puts adapters.first.branch_count
