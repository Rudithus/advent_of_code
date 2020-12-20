expanses = File.open('input.txt').readlines.map(&:to_i)

puts expanses.combination(3).find { |group| group.sum == 2020 }.reduce(:*)
